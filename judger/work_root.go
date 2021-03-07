package judger

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/LanceLRQ/deer-common/constants"
	"github.com/LanceLRQ/deer-common/provider"
)

func InitWorkRoot() error {
	_, filename, _, _ := runtime.Caller(1)
	workPath, err := filepath.Abs(path.Dir(path.Dir(filename)))
	if err != nil {
		return err
	}
	err = os.Chdir(workPath)
	if err != nil {
		return err
	}
	err = provider.PlaceCompilerCommands("./compilers.json")
	if err != nil {
		return err
	}
	err = constants.PlaceMemorySizeForJIT("./jit_memory.json")
	if err != nil {
		return err
	}
	return nil
}
