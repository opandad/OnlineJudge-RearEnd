package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

//excel第一行数据默认不读
func ReadTeam(filePath string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}
	rows, err := f.GetRows("Sheet1")
	// teamItem := len(rows) - 1
	// fmt.Println(teamItem)
	// for i, row := range rows {
	// 	if i == 0 {
	// 		continue
	// 	}

	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	return rows, nil
}
