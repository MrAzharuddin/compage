package unitTest

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (u *UnitTestCmd) fileProcess(path string) (map[string][]string, error) {
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read the file
	data := make(map[string][]string)
	currentFileName := filepath.Base(path)
	var currentCodeSection strings.Builder
	reading := false

	// scanner to read line by line
	scanner := bufio.NewScanner(file)

	// iterate through each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		// check if the line contains the code section start
		if strings.Contains(line, "// start unit test") {
			reading = true
			continue
		}

		// check if the line contains the code section end
		if strings.Contains(line, "// end unit test") {
			reading = false

			// store the accumulated code section if it's not empty
			if currentCodeSection.Len() > 0 {
				data[currentFileName] = append(data[currentFileName], currentCodeSection.String())
				currentCodeSection.Reset() // reset the builder for the next code section
			}
			continue
		}

		// If reading flag is true, append the line to the current code section
		if reading {
			currentCodeSection.WriteString(line + "\n")
		}
	}

	// Check for any errors during the scan
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UnitTestCmd) processCodeString(content string) ([]string, error) {
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
			return nil, errors.New("incomplete code block found")
		}

		// Extract the code block
		codeBlock := content[startIndex : startIndex+endIndex]
		extractedTexts = append(extractedTexts, codeBlock)

		// Move to position after the ending delimiter
		startIndex += endIndex + len(endDelimiter)
	}

	return extractedTexts, nil
}

func (u *UnitTestCmd) processCodeBlock(fileName string, codeSections string) error {
	// Process the code sections
	code, err := u.processCodeString(codeSections)
	if err != nil {
		return err
	}

	// Create a new file with the same name as the original file and _test.go file
	testFileName := fileName[0:len(fileName)-3] + "_test.go"
	testFile, err := os.Create(testFileName)
	if err != nil {
		return err
	}
	defer testFile.Close()

	// Write the code to the test file
	for _, codeSection := range code {
		_, err := fmt.Fprintln(testFile, codeSection)
		if err != nil {
			return err
		}
	}

	return nil
}
