package utils

// Extend this struct to add more languages
type Languages struct {
	Go     string `json:"go"`
	DotNet string `json:"dotnet"`
}

// Extend this variable to add more languages
var AvailableLanguages = Languages{
	Go:     "go",
	DotNet: "dotnet",
}

const BACKEND_LLM_URL = "http://localhost:8000/api"

const PROMPTS_CONFIG_FILE = "prompts.json"

const UnitTestDir = "genai_unit_test"
const DocGenerationDir = "genai_doc_generation"
