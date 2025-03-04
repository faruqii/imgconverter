package main

import (
	"fmt"

	"github.com/faruqii/imgconverter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
