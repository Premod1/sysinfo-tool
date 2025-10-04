package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// Take input keyword from command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <keyword>")
		return
	}
	keyword := os.Args[1]

	if keyword == "system" || keyword == "info" {
		fmt.Println("=== System Information ===")
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Architecture: %s\n", runtime.GOARCH)
		fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
		fmt.Printf("Go Version: %s\n", runtime.Version())
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		fmt.Printf("Allocated Memory: %v MB\n", mem.Alloc/1024/1024)
	} else {
		fmt.Println("Unknown keyword. Try 'system' or 'info'")
	}
}
