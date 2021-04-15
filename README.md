[![build](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml/badge.svg)](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/pgollangi/go-portfinder/branch/main/graph/badge.svg?token=MI1VM2O6AU)](https://codecov.io/gh/pgollangi/go-portfinder)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/go-portfinder)](https://goreportcard.com/report/github.com/pgollangi/go-portfinder)
[![Maintainability](https://api.codeclimate.com/v1/badges/032b766c28546267c545/maintainability)](https://codeclimate.com/github/pgollangi/go-portfinder/maintainability)
![License: MIT](https://img.shields.io/github/license/pgollangi/go-portfinder)
# go-portfinder
Go implementation of npm [portfinder](https://www.npmjs.com/package/portfinder). A simple tool to find an open port on the current machine.

Installation
--------------

```bash
$ go get github.com/pgollangi/go-portfinder
```
# Usage

```go
package main

import (
	"github.com/pgollangi/go-portfinder"
)

func main(){
     // scans and returns first open port on all network interfaces of current machine.
     openPort, err := portfinder.GetPort(PortFinderOptions {
         StartPort: 9090,
         StopPort: 9099
     })

    // Check if a port is open
    isOpen, err := portfinder.IsOpen(8080)
}


```
## Author
Built with ❤ by [Prasanna Kumar](https://pgollangi.com/tabs/about/)
