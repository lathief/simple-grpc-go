
# Project Title

A brief description of what this project does and who it's for

## Prerequisites

This is an example of how to list things you need to use the software and how to install them in linux.

* Protocol Buffer Compiler
    1. Manually download from github.com/google/protobuf/releases the zip file corresponding to your operating system and computer architecture (protoc-<version>-<os>-<arch>.zip), or fetch the file using commands such as the following:
  ```sh
  PB_REL="https://github.com/protocolbuffers/protobuf/releases"
  curl -LO $PB_REL/download/<version>/protoc-<version>-<os>-<arch>.zip
  ```
    2. Unzip the file under $HOME/.local or a directory of your choice. For example:
  ```sh
  unzip protoc-<version>-<os>-<arch>.zip -d $HOME/.local
  ```
    3. Update your environment’s path variable to include the path to the protoc executable. For example:
  ```sh
  export PATH="$PATH:$HOME/.local/bin"
  ```
* Go protocol buffers plugin:
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```
## Installation

Install my-project

```bash
go get -u github.com/jmoiron/sqlx
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf
```
    