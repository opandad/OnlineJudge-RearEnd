package judger

import (
	"github.com/LanceLRQ/deer-common/constants"
	commonStructs "github.com/LanceLRQ/deer-common/structs"
)

func AnalysisResult(result *commonStructs.JudgeResult) (string, error) {
	name, ok := constants.FlagMeansMap[result.JudgeResult]
	if !ok {
		name = "Unknown"
	}
	if result.JudgeResult != 0 {
		//return error name
		return name, nil
	}

	//return ac
	return name, nil
}
