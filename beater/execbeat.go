// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package beater

import (
	"os"
	"os/exec"
	"os/signal"
	"fmt"
	"time"
	"bytes"
	// "log"
	"strings"
	"syscall"
	"sync"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/edgarm1964/execbeat/config"

	// "github.com/aikxeed/dump"
)

// Execbeat configuration.
type Execbeat struct {
	done   chan bool
	config config.Config
	client beat.Client
	waitGroup sync.WaitGroup
}

// New creates an instance of execbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	for i := range c.Commands {
		if len (c.Commands[i].DocumentType) == 0 {
			c.Commands[i].DocumentType = config.DefaultDocumentType
		}

		if c.Commands[i].Schedule == 0 {
			c.Commands[i].Schedule = config.DefaultSchedule
		}
	}

	bt := &Execbeat{
		done:   make(chan bool),
		config: c,
	}

	return bt, nil
}

// Run starts execbeat.
func (bt *Execbeat) Run(b *beat.Beat) error {
	logp.Info("execbeat is running! Hit CTRL-C to stop it.")

	c := bt.config

	// connect to publisher
	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// set up signals
	bt.SetupSignals ()

	// start workers for all commands
	for _, cmd := range c.Commands {
		// create and start worker
		go bt.CreateAndRunWorker (cmd)

		bt.waitGroup.Add(1)
	}

	// wait for done
	<-bt.done

	time.Sleep(200 * time.Millisecond)

	close (bt.done)

	bt.waitGroup.Wait ()

	return nil
}

// create worker
func (bt *Execbeat) CreateAndRunWorker (cfg config.ExecConfig) error {
	var outb, errb bytes.Buffer
	var err error
	var cmd *exec.Cmd
	var execCommand string

	// mark goroutine finshed when worker ends
	defer bt.waitGroup.Done ()

	logp.Info ("create worker for cmd: %s", strings.Trim (cfg.Command, " "))
	ticker := time.NewTicker(cfg.Schedule)

	for {
		// wait for either 'done' or ticker
		select {
		case <-bt.done:
			logp.Info ("worker for cmd: %s finished", execCommand)
			return nil
		case <-ticker.C:
		}

		logp.Info ("wake up goroutine for cmd: %s", cfg.Command)

		// set up command to run
		if len(cfg.Args) > 0 {
			execCommand = cfg.Command + " " + cfg.Args
			cmd = exec.Command (cfg.Command, cfg.Args)
		} else {
			execCommand = cfg.Command
			cmd = exec.Command (cfg.Command)
		}

		// attach buffers to stdout and stderr
		cmd.Stdout = &outb
		cmd.Stderr = &errb

		// Run command
		// As libbeat is now compiled with seccomp support, it is
		// possible that Run() returns with 'Operation not permitted'
		// In that case add 'seccomp.enabled: false' to your
		// configuration file
		err = cmd.Run ()

		if err != nil {
			logp.Err ("Couldn't execute %s: %v. Maybe add 'seccomp.enabled: false' to the configuration file?", execCommand, err)
			return nil
		}
		exitcode := cmd.ProcessState.ExitCode()
		logp.Info ("stdout: %s, stderr: %s, exitcode: %d\n",
			    strings.Trim (outb.String(), " \n"),
			    strings.Trim (errb.String(), " \n"),
			    exitcode)

		fields := common.MapStr {
			"command":	execCommand,
			"stdout":	strings.Trim (outb.String(), " \n"),
			"stderr":	strings.Trim (errb.String(), " \n"),
			"exitCode":	exitcode,
		}

		// If available, add custom fields
		if cfg.Fields != nil {
			if (cfg.FieldsUnderRoot) {
				// fields are to be placed at root
				for k, v := range cfg.Fields {
					fields[k] = v
				}
			} else {
				// fields are to be added separate
				fields["fields"] = cfg.Fields
			}
		}

		// create event
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: fields,
		}

		// send event to end points
		bt.client.Publish(event)
		logp.Info("Event sent")

		// reset buffers
		outb.Reset()
		errb.Reset()
	}

	return nil
}

// interrupt service routine
func (bt *Execbeat) SetupSignals() error {
	// Set up signal channel
	c := make (chan os.Signal)

	// Catch signals ^C, ^\ and 15 (TERM)
	signal.Notify (c, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	// create an anonymous blocking function that catches the
	// signals and stops all other goroutines
	go func() {
		// block and wait for signal
		s := <-c
		msg := "other"

		switch s {
			case os.Interrupt:
				msg = "^C"
			case syscall.SIGQUIT:
				msg = "^\\"
			case syscall.SIGTERM:
				msg = "terminate"
		}

		logp.Info ("Received signal %s, shutting down...\n", msg)

		// call Stop to signal other threads
		bt.Stop ()
	}()

	// all done, return
	return nil
}

// Stop stops execbeat.
func (bt *Execbeat) Stop() {
	bt.done <- true
	//close (bt.done)
}
