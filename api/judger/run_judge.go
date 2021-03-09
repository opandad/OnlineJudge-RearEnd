package judger

import (
	commonStructs "github.com/LanceLRQ/deer-common/structs"
	"github.com/LanceLRQ/deer-common/utils"
	"github.com/LanceLRQ/deer-executor/v2/executor"
	uuid "github.com/satori/go.uuid"
)

func RunJudge(conf, codeFile, codeLang string) (*commonStructs.JudgeResult, error) {
	session, err := executor.NewSession(conf)
	if err != nil {
		return nil, err
	}
	session.CodeFile = codeFile
	session.CodeLangName = codeLang
	session.SessionRoot = "/tmp"
	session.SessionId = uuid.NewV1().String()
	sessionDir, err := utils.GetSessionDir(session.SessionRoot, session.SessionId)
	if err != nil {
		return nil, err
	}
	session.SessionDir = sessionDir
	defer session.Clean()
	// start judge
	judgeResult := session.RunJudge()
	return &judgeResult, err
}
