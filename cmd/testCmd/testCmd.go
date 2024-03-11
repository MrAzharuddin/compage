/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package testcmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type testCmd struct {
	accessToken string
}

func NewTestCmd(accessToken string) *testCmd {
	return &testCmd{
		accessToken: accessToken,
	}
}

func (t *testCmd) Execute() *cobra.Command {
	return &cobra.Command{
		Use:   "testCmd",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: t.TestCmdRun,
	}
}

func processGoFile(path string) (map[string][]string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the file
	data := make(map[string][]string)
	currentFileName := filepath.Base(path)
	var currentCodeSection strings.Builder
	reading := false

	// Scanner to read line by line
	scanner := bufio.NewScanner(file)

	// Iterate through each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains the start comment
		if strings.Contains(line, "// start unit test") {
			reading = true
			continue
		}

		// Check if the line contains the end comment
		if strings.Contains(line, "// end unit test") {
			reading = false
			// Store the accumulated code section if it's not empty
			if currentCodeSection.Len() > 0 {
				data[currentFileName] = append(data[currentFileName], currentCodeSection.String())
				currentCodeSection.Reset() // Reset the builder for the next section
			}
			continue
		}

		// If reading flag is true, append the line to the current code section
		if reading {
			currentCodeSection.WriteString(line + "\n")
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return data, nil
}

func processCodeString(content string) ([]string, error) {
	var extractedTexts []string
	startIndex := 0

	for {
		// Find the start and end delimiters for each code block
		startDelimiter := "```"
		endDelimiter := "```"

		startIndex = strings.Index(content[startIndex:], startDelimiter)
		if startIndex == -1 {
			break // No more delimiters found
		}
		startIndex += len(startDelimiter) // Move to position after the starting delimiter

		endIndex := strings.Index(content[startIndex:], endDelimiter)
		if endIndex == -1 {
			return nil, fmt.Errorf("incomplete %s block found", endDelimiter)
		}

		// Extract the code block
		codeBlock := content[startIndex : startIndex+endIndex]
		extractedTexts = append(extractedTexts, codeBlock)

		// Move to position after the ending delimiter
		startIndex += endIndex + len(endDelimiter)
	}

	return extractedTexts, nil
}

func (t *testCmd) TestCmdRun(cmd *cobra.Command, args []string) {
	// Get the current working director
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// Walk through the directory and its subdirectories
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return nil
		}
		// Check if the file is a regular file
		if info.Mode().IsRegular() {
			// Check if the file is a Go file
			if filepath.Ext(path) == ".go" {
				// Process the Go file
				data, err := processGoFile(path)
				if err != nil {
					fmt.Println("Error processing Go file:", err)
					return nil
				}

				if data == nil {
					return nil
				}

				// Print the data
				for k, v := range data {
					fmt.Printf("File: %s\n", k)

					// Check if the data has values
					if len(v) > 0 {
						// Print the extracted code section
						fmt.Println("Extracted Code:", strings.Join(v, "\n"))
						res, err := Server(t.accessToken, strings.Join(v, "\n"))
						if err != nil {
							fmt.Println("Error:", err)
							return nil
						}
						onlyCode, err := processCodeString(res.Data.Code)
						if err != nil {
							fmt.Println("Error:", err)
							return nil
						}
						fmt.Println("Extracted Code:", strings.Join(onlyCode, "\n"))

						// Create a new file with _test.go suffix
						testName := k[:len(k)-3] + "_test.go"
						testFile, err := os.Create(testName)
						if err != nil {
							fmt.Println("Error creating test file:", err)
							return nil
						}
						defer testFile.Close()

						// Write the extracted code sections to the test file
						for _, line := range onlyCode {
							fmt.Println(line)
							_, err := fmt.Fprintln(testFile, line)
							if err != nil {
								fmt.Println("Error writing to test file:", err)
								return nil
							}
						}
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
		return
	}
}
