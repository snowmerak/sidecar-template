package main

import (
	"fmt"
	"log"
	// This import path is replaced by the local path in go.mod
	// "github.com/my-org/core-platform/gen/go/company/auth/v1"
)

func main() {
	fmt.Println("Sidecar Go Server Starting...")
	// Example usage of generated code:
	// req := &authv1.LoginRequest{Username: "admin"}
	// fmt.Println(req)
	log.Println("Ready to serve.")
}
