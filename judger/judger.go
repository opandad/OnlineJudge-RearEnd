package judger

func Judger() error {
	err := InitWorkRoot()
	if err != nil {
		return err
	}
	result, err := RunJudge("./data/problems/APlusB/problem.json", "./data/codes/APlusB/ac.c", "")
	if err != nil {
		return err
	}
	err = AnalysisResult(result)
	if err != nil {
		return err
	}
	return nil
}
