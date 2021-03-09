package judger

import "fmt"

func Judger(problemFile string, codeFile string, codeLanguage string) (string, error) {
	err := InitWorkRoot()
	if err != nil {
		return "", err
	}
	result, err := RunJudge(problemFile, codeFile, codeLanguage)
	if err != nil {
		return "", err
	}
	status, err := AnalysisResult(result)
	if err != nil {
		return "", err
	}
	fmt.Println(status)
	return status, nil
}
