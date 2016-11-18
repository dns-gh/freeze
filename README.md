# freeze package

[![Go Report Card](https://goreportcard.com/badge/github.com/dns-gh/freeze)](https://goreportcard.com/report/github.com/dns-gh/freeze)

[![GoDoc](https://godoc.org/github.com/dns-gh/freeze?status.png)]
(https://godoc.org/github.com/dns-gh/freeze)

Simple wrapper over the time and rand packages to make more complex behaviors around random sleep functionalities

## Motivation

Used in other projects such as the Twitter Bot or the Nasa Neo Client

## Installation

- It requires Go language of course. You can set it up by downloading it here: https://golang.org/dl/
- Install it here C:/Go.
- Set your GOPATH, GOROOT and PATH environment variables:

```
export GOROOT=C:/Go
export GOPATH=WORKING_DIR
export PATH=C:/Go/bin:${PATH}
```

- Download and install the package:

```
@working_dir $ go get github.com/dns-gh/freeze/...
@working_dir $ go install github.com/dns-gh/freeze
```

## Usage

```
package main

import "github.com/dns-gh/freeze"

func main() {
	freeze.Sleep(10) // sleeps randomly from 0 to 9 seconds
    freeze.MaybeSleepMinMax(1,10,20,30) // 1 chance over 10 to sleep from 10 to 30 seconds
}

```

## Tests

TODO

## LICENSE

See included LICENSE file.