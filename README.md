# Flip-go

[![Go Report Card](https://goreportcard.com/badge/github.com/rl404/flip-go)](https://goreportcard.com/report/github.com/rl404/flip-go)
![License: MIT](https://img.shields.io/github/license/rl404/flip-go.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/rl404/flip-go.svg)](https://pkg.go.dev/github.com/rl404/flip-go)

_flip-go_ is unofficial golang API wrapper for [flip.id](https://flip.id).

For official documentation, go [here](https://docs.flip.id).

## Features

- Get flip account's balance
- Get bank list and info
- Get flip's maintenance status
- Get bank account info and status
- Get city code list
- Get country code list
- Create disbursement (v2)
- Get all disbursement list + filter (v2)
- Get disbursement by id (v2)
- Create special disbursement (v2)

## Installation

```
go get github.com/rl404/flip-go
```

## Quick Start

```go
package main

import (
	"log"

	"github.com/rl404/flip-go"
)

func main() {
	secretKey := "abc123"

	f := flip.NewDefault(secretKey, flip.Sandbox)

	balance, err := f.GetBalance()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(balance)
}
```

*For more detail config and usage, please go to the [documentation](https://pkg.go.dev/github.com/rl404/flip-go).*

## License

MIT License

Copyright (c) 2021 Axel
