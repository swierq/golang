# Golang

Just some golang learning repository with examples for various go libraries I am using.

## Install

### build from source

```
make instal-dev-deps
make build
```

Will download dev tools needed to build the project and compile binaries into ./bin directory.

### docker

```
docker run -it --rm -p 8080:8080 swierq/golang loadek -cpumi 500 -memmb 100
```

# Apps in the repository

## Loadek - artificial load generator

Generates artificial load based on command line parameters. Exposes simple ui over http.

```
Usage: bin/loadek [flags]

Flags:

  -cpumi value
        Cpu milocores (default 100)
  -memmb value
        Memory mb (default 200)
  -port value
        Listen Port (default 8080)

bin/loadek: bad flags: flag: help requested
```
