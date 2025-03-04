package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "img-converter",
	Short: "A CLI tool to convert images from HEIC and other formats",
	Long:  "img-converter is a command-line tool to convert images from HEIC to PNG, JPEG, and WebP.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	fmt.Println("Image Converter CLI")
}
