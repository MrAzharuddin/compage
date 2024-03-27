package init

import "github.com/intelops/compage/cmd/internal/utils"

var (
	docPrompt    = "Generate documentation for the following folder structure and also provide code flow diagram in the mermaid format. The entire documentation should be in markdown format.Folder Structure : %s"
	goPrompt     = "Write a unit test case for the following Golang programming language code using the in-built testing package in golang:%s. Make sure the unit test case you are generating is providing the imports on the top and also keep that whole test case in between three backticks(```) at the beginning and end of the unit test case."
	dotnetPrompt = "Write a unit test case for the following dotnet programming language code using the MSTest framework :%s. Make sure the unit test case you are generating is providing the proper imports of packages on the top and also keep that whole test case in between three backticks(```) at the beginning and end of the unit test case."
	language     = utils.AvailableLanguages.Go
)
