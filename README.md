# Go MSC SDK

[![GoDoc](https://godoc.org/github.com/buyco/go-msc-sdk?status.svg)](http://godoc.org/github.com/buyco/go-msc-sdk) [![Build Status](https://github.com/buyco/go-msc-sdk/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/buyco/go-msc-sdk/actions/workflows/test.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/buyco/go-msc-sdk)](https://goreportcard.com/report/github.com/buyco/go-msc-sdk) [![codecov](https://codecov.io/gh/buyco/go-msc-sdk/branch/master/graph/badge.svg?token=7BM0AEZAJ8)](https://codecov.io/gh/buyco/go-msc-sdk) [![license](https://img.shields.io/github/license/buyco/go-msc-sdk.svg)](https://github.com/buyco/go-msc-sdk/LICENSE)

SDK built to use MSC APIs

## What is done ?

- [x] Tracking
- [ ] Schedules
- [ ] Booking

## Usage recommandations

- To match _HttpClient_ interface use: https://github.com/imroc/req

## How to install Go ?

#### Debian / Ubuntu:

```bash
$ sudo apt update
$ sudo apt install golang
```

#### Arch:

```bash
$ sudo pacman -Sy go
```

#### Mac OS X:

```bash
$ brew update
$ brew install golang
```

#### Last release from script:

See: https://github.com/udhos/update-golang

#### From tarballs:

See: https://golang.org/doc/install

## Check golang version

```bash
$ go version
```

## Install from Makefile

### Commands:

To list available commands:

```bash
$ make help
```

To check (vet / lint / fmt)

```bash
$ make check
$ make lint # only lint
$ make vet # only vet
$ make fmt # only fmt
```

To run tests:

```bash
$ go install github.com/joho/godotenv/cmd/godotenv
$ godotenv -f <env_file> make test
```