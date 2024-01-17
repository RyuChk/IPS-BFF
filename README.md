# IPS BFF (Backend for frontend)

Backend server that stay in between front end and services.

## Table of Content
- [Quick Start](#quick-start)
    - [Install Dependency](#install-bufbuild)
    - [Pull Git submodule](#pull-git-submodule)
    - [Setup Local Env]()
    - [Setup Debug Tools (vscode)](#install-protoc)
- [Protobuf Info](#protobuf-structure)
    - [Structure](#structure)
    - [Generate proto](#generate-proto)
- [Mock Generate for testing](#using-gomock-generate)

## Quick Start
---
### Install Dependency

#### Install buf/build

**Install using brew**
```
brew install bufbuild/buf/buf
```

#### Install protoc

**Install using go**
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

**Install path**
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Pull Git Submodule
using this command ```git submodule update --recursive --remote```

### Setup Local Env

#### Option 1 : using make command

command : `make initbeta`


#### Option 2 : Manual

1. Copy `.env.example` to `.env.beta`
2. Fill up unfill variable inside
3. Run `make initmongo`
> **Hint:** you can use .env.example as a references on what environment variable need to be set

### Setup Debug in Vscode
create file in `.vscode` directory named `launch.json`
filepath: `.vscode/launch.json`

copy and paste this code into the file.
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "rssi-grpc",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "envFile": "${workspaceFolder}/.env.beta",
            "program": "cmd/rssi-grpc/main.go",
        },
        {
            "name": "rssi-api",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "envFile": "${workspaceFolder}/.env.beta",
            "program": "cmd/rssi-api/main.go",
        }
    ]
}
```

now you can go to **Run and Debug** section to start debuging.

## Protobuf
---

### Structure
```
ips
|- bff (declare method using in bff service)
    |- v1
        |- bff.proto
|- shared (set enum or message that use in many project)
    |- bff
        |- v1
            |- bff.proto
```

### Generate proto

#### Option 1: use make
```
make generate
```

#### Option 2: directly generate
1. `cd` into proto directory
2. run `go generate`
```
$ cd proto
$ go generate
```

after run go generate file will be generated into `internal/gen` directory

## Using Gomock Generate
---
### Setup gomock command

argrument list:
- `-source` : filename to generate (current file name)
- `-destination` : generate mock destination path `mock_<service_name>/mock_<service_name>.go`
- `-package` : package to use `mock_<service_name>`

`example_service.go`
```
package example

//go:generate mockgen -source=example_service.go -destination=mock_example_service/mock_example.go -package=mock_example_service
type Service interface{

}

type ExampleService struct{
}
```

### Run generate
#### Option 1: use make
```
make generate
```

#### Option 2: self generate
cd into file directory and run `go generate` command