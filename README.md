# requestid [![GoDoc](https://godoc.org/github.com/ascarter/uuid?status.svg)](http://godoc.org/github.com/ascarter/uuid)[![Go Report Card](https://goreportcard.com/badge/github.com/ascarter/uuid)](https://goreportcard.com/report/github.com/ascarter/uuid)

RequestID middleware for Go.

# Example

```go

package main

import (
	"fmt"

	"github.com/ascarter/uuid"
)

func main() {
	u := uuid.NewUUID()
	fmt.Println("UUID:", u.String())
}

```

# References

