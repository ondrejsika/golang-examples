package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {}

func main() {
	Cmd.Execute()
}

var Cmd = &cobra.Command{
	Use:   "go_expand_env_with_cobra <inputfile> <outputfile>",
	Short: "A simple CLI to expand environment variables in a file",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputFilePath := os.Args[1]
		outputFilePath := os.Args[2]
		expandEnvInFiles(inputFilePath, outputFilePath)
	},
}

func expandEnvInFiles(inputFilePath string, outputFilePath string) {
	// Read the content of the input file
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Expand environment variables in the content
	expandedContent := os.ExpandEnv(string(content))

	// Write the expanded content to the output file
	err = os.WriteFile(outputFilePath, []byte(expandedContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
