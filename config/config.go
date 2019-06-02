// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

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
