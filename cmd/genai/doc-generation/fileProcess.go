package docgeneration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (dg *DocGenerationCmd) fetchFoldersData(root string) ([]string, error) {
	result := make([]string, 0)

	excludedDirs = []string{
		"Dockerfile",
		"README.md",
	}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			dg.logger.Error("Error walking through directory:", err)
			return err
		}

		if info.IsDir() {
			return nil // Ignore directories
		}

		if dg.shouldExclude(path, excludedDirs) {
			result = append(result, path)
		}

		return nil
	})

	if err != nil {
		dg.logger.Error("Error fetching folders data:", err)
		return nil, err
	}

	return result, nil
}

func (dg *DocGenerationCmd) shouldExclude(filePath string, excludedDirs []string) bool {
	for _, excludedDir := range excludedDirs {
		if strings.Contains(filePath, excludedDir) {
			return false
		}
	}
	return true
}

func (dg *DocGenerationCmd) storeDocumentationInFile(root string, content string) error {
	if root == "" {
		dg.logger.Error("Error writing documentation file: root directory is empty")
		return fmt.Errorf("root directory is empty")
	}

	if content == "" {
		dg.logger.Error("Error writing documentation file: content is empty")
		return fmt.Errorf("content is empty")
	}

	// Write the content to the file
	filePath := filepath.Join(root, "genai_documentation.md")
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		dg.logger.Error("Error writing documentation file:", err)
		return err
	}
	return nil
}
