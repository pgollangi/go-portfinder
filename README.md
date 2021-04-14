[![build](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml/badge.svg)](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/pgollangi/go-portfinder/branch/main/graph/badge.svg?token=MI1VM2O6AU)](https://codecov.io/gh/pgollangi/go-portfinder)

[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/fastget)](https://goreportcard.com/report/github.com/pgollangi/fastget)
[![Maintainability](https://api.codeclimate.com/v1/badges/032b766c28546267c545/maintainability)](https://codeclimate.com/github/pgollangi/go-portfinder/maintainability)
# go-portfinder
Go implementation of [node-portfinder](https://www.npmjs.com/package/portfinder). A simple tool to find an open port on the current machine.

Installation
--------------

```bash
$ go get github.com/pgollang/go-portfinder
```
# Usage

```go
package main

import (
	"github.com/pgollang/go-portfinder"
)

func main(){
     // scan localhost with a 2 second timeout per port in 5 concurrent threads
     openPort, err := portfinder.GetPort(PortFinderOptions {
         StartPort: 9090,
         StopPort: 9099
     })
}


```
## Author
Built with ❤ by [Prasanna Kumar](https://pgollangi.com/tabs/about/)