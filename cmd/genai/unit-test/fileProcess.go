package unitTest

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/intelops/compage/cmd/internal/utils"
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
		u.logger.Error("Error processing code sections:", err)
		return err
	}

	// Save the test file
	err = u.processAndSaveTestFile(fileName, code)
	if err != nil {
		u.logger.Error("Error saving test file:", err)
		return err
	}

	return nil
}

// Based on language, create a test file in `genai_unit_test` folder, create this folder if it doesn't exist.
func (u *UnitTestCmd) processAndSaveTestFile(fileName string, code []string) error {
	if config == nil {
		return fmt.Errorf("config is nil")
	}
	if config.Language == "" {
		return fmt.Errorf("language is empty")
	}
	if len(fileName) == 0 {
		return fmt.Errorf("fileName is empty")
	}
	if code == nil {
		return fmt.Errorf("code is nil")
	}

	var testfileName string
	if config.Language == utils.AvailableLanguages.Go {
		testfileName = fileName[0:len(fileName)-3] + "_test.go"
	} else if config.Language == utils.AvailableLanguages.DotNet {
		testfileName = fileName[0:len(fileName)-3] + "Test.cs"
	} else {
		return fmt.Errorf("language not supported")
	}

	// Create the `genai_unit_test` folder if it doesn't exist
	if _, err := os.Stat(utils.UnitTestDir); os.IsNotExist(err) {
		err := os.Mkdir(utils.UnitTestDir, 0755)
		if err != nil {
			u.logger.Error("Error creating test folder:", err)
			return err
		}
	}

	// Create the test file
	testFile, err := os.Create(utils.UnitTestDir + "/" + testfileName)
	if err != nil {
		u.logger.Error("Error creating test file:", err)
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
