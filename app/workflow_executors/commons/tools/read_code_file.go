// package
package tools

import (
	"ai-developer/app/utils"
	"fmt"
)

// ReadCodeParams encapsulates the parameters for the ReadCode function.
type ReadCodeParams struct {
	FilePath  string
	StartLine int
	EndLine   int
	Lint      bool
	MaxLines  int
}

// ReadCode reads a file and returns the code between specified line numbers.
func ReadCode(params ReadCodeParams) (string, error) {
	if params.EndLine < params.StartLine {
		return "", fmt.Errorf("end_line should be greater than or equal to start_line")
	}

	if params.EndLine-params.StartLine > params.MaxLines {
		params.EndLine = params.StartLine + params.MaxLines
	}

	lines, err := utils.ReadFile(params.FilePath)
	if err != nil {
		return "", err
	}

	if params.StartLine > len(lines) || params.EndLine > len(lines) {
		return "", fmt.Errorf("the file only contains %d lines", len(lines))
	}

	startLine := params.StartLine - 1
	endLine := params.EndLine - 1
	result := fmt.Sprintf("[File: %s (%d lines total)]\n", params.FilePath, len(lines))

	for i := startLine; i <= endLine; i++ {
		result += fmt.Sprintf("%d: %s\n", i+1, lines[i])
	}

	//if params.Lint {
	//	TODO: Implement linting logic here
	//}

	return result, nil
}

//func main() {
//	params := ReadCodeParams{
//		FilePath:  "/Users/abhijeetsinha/abhijeet/Code/SuperCoder/app/utils/api_key.go",
//		StartLine: 10,
//		EndLine:   32,
//		Lint:      false,
//		MaxLines:  200,
//	}
//
//	result, err := ReadCode(params)
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println(result)
//	}
//}
