// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

// Copyright (c) 2019 Edgar Matzinger
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"time"

	"github.com/elastic/beats/libbeat/common"
)

type Config struct {
	Commands []ExecConfig          `config:"commands"`
}

type ExecConfig struct {
	Schedule         time.Duration `config:"schedule"`
	Command          string        `config:"command"`
	Args             string        `config:"args"`
	DocumentType     string        `config:"document_type"`
	Fields           common.MapStr `config:"fields"`
	FieldsUnderRoot  bool          `config:"fields_under_root"`
}

var DefaultConfig = Config{
}

// Defaults for config variables which are not set
const (
	DefaultSchedule     time.Duration = 60 * time.Second
	DefaultDocumentType string = "execbeat"
	DefaultFieldsUnderRoot bool = false
)
