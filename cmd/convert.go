package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/faruqii/imgconverter/internal/converter"
	"github.com/spf13/cobra"
)

var inputFile, outputFile, format string

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert an image to a different format",
	RunE: func(cmd *cobra.Command, args []string) error {
		if inputFile == "" || outputFile == "" || format == "" {
			return fmt.Errorf("input, output, and format are required")
		}

		ext := strings.ToLower(filepath.Ext(inputFile))
		if ext == ".heic" {
			return converter.ConvertHEIC(inputFile, outputFile, format)
		}
		return converter.ConvertImage(inputFile, outputFile, format)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input file path")
	convertCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path")
	convertCmd.Flags().StringVarP(&format, "format", "f", "", "Output format (png, jpg, webp)")
}
