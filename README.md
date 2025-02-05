[![Build Status](https://travis-ci.org/edgarm1964/execbeat.svg?branch=master)](https://travis-ci.org/edgarm1964/execbeat)
[![codecov.io](http://codecov.io/github/christiangalsterer/execbeat/coverage.svg?branch=master)](http://codecov.io/github/christiangalsterer/execbeat?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/edgarm1964/execbeat)](https://goreportcard.com/report/github.com/edgarm1964/execbeat)
[![license](https://img.shields.io/github/license/edgarm1964/execbeat.svg)](https://github.com/edgarm1964/execbeat)
[![Github All Releases](https://img.shields.io/github/downloads/edgarm1964/execbeat/total.svg)](https://github.com/edgarm1964/execbeat)

![Elastic Beats 7.2.0](https://img.shields.io/badge/Elastic%20Beats-v7.2.0-blue.svg)
![Golang 1.12.4](https://img.shields.io/badge/Golang-v1.12.4-blue.svg)

# Overview

Execbeat is the [Beat](https://www.elastic.co/products/beats) used to execute any command.
Multiple commands can be configured which are executed in a regular interval and the standard output and standard error is shipped to the configured output channel.

Execbeat is inspired by the Logstash [exec](https://www.elastic.co/guide/en/logstash/current/plugins-inputs-exec.html) input filter but doesn't require that the endpoint is reachable by Logstash as Execbeat pushes the data to Logstash or Elasticsearch.
This is often necessary in security restricted network setups, where Logstash is not able to reach all servers. Instead the server to be monitored itself has Execbeat installed and can send the data or a collector server has Execbeat installed which is deployed in the secured network environment and can reach all servers to be monitored.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/edgarm1964`

## Build information

Execbeat is built against the following Beats versions and if execbeat was able to connect to a running Elastic Search instance or not.

| Build against | Connect to ES 6.5.1 | ES 6.8.0 | ES 7.0.0 | ES 7.1.1 | ES 7.2.0 |
| ---- | ---- | ---- | ---- | ---- | ---- |
| Beats 6.5.1 | OK | OK | Fails | N/T | N/T |
| Beats 6.8.0 | OK | OK | OK | OK | OK |
| Beats 6.8.1 | OK | OK | OK | OK | OK |
| Beats 7.0.0 | OK | OK | OK | OK | OK |
| Beats 7.1.1 | OK | OK | OK | OK | OK |
| Beats 7.2.0 | OK | OK | OK | OK | OK |

N/T: Not Tested

## Installation

### Download
Pre-compiled binaries for different operating systems are available for [download](#releases).

### Installation
Install the package for your operation system by running the respective package manager or unzipping the package.

### Configuration
Adjust the `execbeat.yml` configuration file to your needs. You may take `execbeat.reference.yml` as an example containing all possible configuration values. The output of the executed command is stored in the strings `stdout` and `stderr,` The exit code is stored in `extiCode.` The command itself is stored in `command.` All fields can be accessed using the `processors` Beats provides. See [Decode JSON fields example](#decode-json-fields-example)

#### Simple example
The list is a [YAML](http://yaml.org/) array, so each command begins with a dash (`-`). You can specify multiple commands, and you can specify the same command type more than once. For example:

```
execbeat.commands:
  - command: date
    period: 2m
    args: '+%Y%m%dT%H%M%S'
    fields:
      app: MyApplication
      env: test
    fields_under_root: true
```

#### Decode JSON fields example
If a command returns a [JSON](http://json.org) formatted string, it is possible to use processors to split the fields of such a string into separate fields. Example:

```
execbeat.commands:
  - command: /usr/local/bin/a-json-script.sh
    period: 5m
    # args:

processors:
  - decode_json_fields:
      fields: ["stdout"]
      process_array: true
      max_depth: 1
      target: ""
      overwrite_keys: false
```

Visit [processors](https://www.elastic.co/guide/en/beats/filebeat/current/defining-processors.html) for more information on processors and their use.

### Running
In order to start Execbeat please use the respective startup script, e.g. `/usr/bin/execbeat.sh`. For more information, run `execbeat --help`

### Starting Execbeat as Service
Where supported Execbeat can be started also using the respetive service scripts, e.g. `etc/init.d/execbeat`.

## Building and Releasing Execbeat

### Requirements

* [Golang](https://golang.org/dl/) = 1.12.4
* [Glide](https://github.com/Masterminds/glide) >= 0.13.0
* [Mage](https://magefile.org) >= 1.8.0

### Build

To build the binary for execbeat run the command below. This will generate a binary
in the same directory with the name execbeat.

```
make clean && make
```

### Run

To run execbeat with debugging output enabled, run:

```
./execbeat -c execbeat.yml -e -d '*'
```

### Test

To test execbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `_meta/fields.yml`.
To generate docs/execbeat.template.json and docs/execbeat.asciidoc

```
make update
```


### Cleanup

To clean execbeat source code, run the following commands:

```
make check
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```

### Clone

To clone execbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/edgarm1964
cd ${GOPATH}/github.com/edgarm1964
git clone https://github.com/edgarm1964/execbeat
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The complete process to finish can take several minutes.

# Releases

## 7.2.0 (2019-07-15) [Download](https://github.com/edgarm1964/execbeat/releases/tag/7.2.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/7.2.0/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/7.2.0)

Feature and Bugfix release containing the following changes:
* Update to beats v7.2.0

## 7.1.1 (2019-07-15) [Download](https://github.com/edgarm1964/execbeat/releases/tag/7.1.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/7.1.1/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/7.1.1)

Feature and Bugfix release containing the following changes:
* Update to beats v7.1.1

## 7.0.0 (2019-07-11) [Download](https://github.com/edgarm1964/execbeat/releases/tag/7.0.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/7.0.0/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/7.0.0)

Feature and Bugfix release containing the following changes:
* Update to beats v7.0.0

## 6.8.1 (2019-07-08) [Download](https://github.com/edgarm1964/execbeat/releases/tag/6.8.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/6.8.1/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/6.8.1)

Feature and Bugfix release containing the following changes:
* Update to beats v6.8.1

## 6.8.0 (2019-06-20) [Download](https://github.com/edgarm1964/execbeat/releases/tag/6.8.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/6.8.0/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/6.8.0)

Feature and Bugfix release containing the following changes:
* Update to beats v6.8.0

## 6.5.1 (2019-06-20) [Download](https://github.com/edgarm1964/execbeat/releases/tag/6.5.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/edgarm1964/execbeat/6.5.1/total.svg)](https://github.com/edgarm1094/execbeat/releases/tag/6.5.1)

Feature and Bugfix release containing the following changes:
* Update to beats v6.5.1
* Redesigned from the ground up following the [Creating a New Beat](https://www.elastic.co/guide/en/beats/devguide/current/new-beat.html) guide
* execbeat.yml is incompatible with the ones from previous versions: change 'schedule' into 'period'


## 3.3.0 (2017-10-06) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.3.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.3.0/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/3.3.0)

Feature and Bugfix release containing the following changes:
* Update to beats v5.6.2

## 3.2.0 (2017-06-05) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.2.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.2.0/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/3.2.0)

Feature and bugfix release containing the following changes:
* Fix: [Use exit code 127 when command to execute is not found](https://github.com/christiangalsterer/execbeat/issues/15)

## 3.1.1 (2017-02-24) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.1.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.1.1/total.svg)](https://g0thub.com/christiangalsterer/execbeat/releases/tag/3.1.1)

Bugfix release containing the following changes:
* [Set correct version in package names and package metadata](https://github.com/christiangalsterer/execbeat/issues/10)

## 3.1.0 (2017-02-23) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.1.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.1.0/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/3.1.0)

Feature and bugfix release containing the following changes:
* The exit code of the command executed is now exported in field `exitCode`.
* Fix: Examples were not fully updated with configuration changes introduced in 3.0.0.

## 3.0.1 (2017-02-21) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.0.1/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.1)

Bugfix release containing the following changes:
* [Multiple arguments are not properly passed](https://github.com/christiangalsterer/execbeat/issues/7)

## 3.0.0 (2017-02-19) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/3.0.0/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/3.0.0)

Feature and bugfix release containing the following **breaking** changes:
* Renamed configuration parameter `execs` to `commands`. Please update your configuration accordingly.
* Renamed configuration parameter `cron` to `schedule`. Please update your configuration accordingly.
* Update to beats v5.2.1
* Fix: [Default schedule not working](https://github.com/christiangalsterer/execbeat/issues/6)

## 2.2.0 (2017-02-04) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/2.2.0)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/2.2.0/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/2.2.0)

Feature release containing the following changes:
* Update to beats v5.2.0

## 2.1.1 (2017-01-14) [Download](https://github.com/christiangalsterer/execbeat/releases/tag/2.1.1)
[![Github Releases (by Release)](https://img.shields.io/github/downloads/christiangalsterer/execbeat/2.1.1/total.svg)](https://github.com/christiangalsterer/execbeat/releases/tag/2.1.1)

Starting with this release pre-compiled binaries for different operating systems are available under the respective tag in the github project.

Bugfix release containing the following changes:
* Move files into correct place to allow correct bulding with `make package`
* Move files into correct place to allow correct bulding with `make update`
* Cleanup of documentation
* Update to beats v5.1.2
* Update to Go 1.7.4

## 2.1.0 (2016-12-23)

Feature release containing the following changes:
* Update to beats v5.1.1

## 2.0.0 (2016-11-26)

Feature release containing the following changes:
* Update to beats v5.0.1

Please note that this release contains the following breaking changes introduced by beats 5.0.X, see also [Beats Changelog](https://github.com/elastic/beats/blob/v5.0.0-beta1/CHANGELOG.asciidoc)
* SSL Configuration
    * rename tls configurations section to ssl
    * rename certificate_key configuration to key.
    * replace tls.insecure with ssl.verification_mode setting.
    * replace tls.min/max_version with ssl.supported_protocols setting requiring full protocol name

## 1.1.0 (2016-07-19)

Feature release containing the following changes:
* Update to Go 1.6
* Update to libbeat 1.2.3
* Use [Glide](https://github.com/Masterminds/glide) for dependency management

## 1.0.1 (2016-02-15)

Bugfix release containing the following changes:
* Fix: [Hanging during shutdown](https://github.com/christiangalsterer/execbeat/issues/2)

## 1.0.0 (2015-12-26)
* Initial release

# Configuration

## Configuration Options

See [here](docs/configuration.asciidoc) for more information.

## Exported Document Types

There is exactly one document type exported:

- `type: execbeat` command execution information, e.g. standard output and standard error. The type can be changed by setting the document_type attribute.

## Exported Fields

See [here](docs/fields.asciidoc) for a detailed description of all exported fields.

### execbeat type

<pre>
{
  "_index": "execbeat-2015.12.26",
  "_type": "execbeat",
  "_source": {
    "@timestamp": "2015-12-26T02:18:53.001Z",
    "beat": {
      "hostname": "mbp.box",
      "name": "mbp.box"
    },
    "count": 1,
    "fields": {
      "host": "test"
    },
    "exec": {
      "command": "echo",
      "exitCode": 0,
      "stdout": "Hello World\n"
    },
    "fields": {
      "host": "test2"
    },
    "type": "execbeat"
    },
  "sort": [
    1449314173
  ]
}
</pre>


## Elasticsearch Template

To apply the Execbeat template:

    curl -XPUT 'http://localhost:9200/_template/execbeat' -d@etc/execbeat.template.json

# Contribution
All sorts of contributions are welcome. Please create a pull request and/or issue.
