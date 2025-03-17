package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("Starting all services...")

	// Start URL Service
	go runService("url-service/main.go")

	// Start Auth Service
	go runService("auth-service/main.go")

	// Start Rate Limiting Service
	go runService("rate-limiting-service/main.go")

	// Prevent exit
	select {}
}

func runService(servicePath string) {
	cmd := exec.Command("go", "run", servicePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting", servicePath, ":", err)
	}
}
