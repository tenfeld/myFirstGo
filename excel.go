// linux向けビルド
// GOOS=linux GOARCH=amd64 go build -o excel_linux excel.go
// mac向けビルド
// GOOS=darwin GOARCH=amd64 go build -o excel_mac excel.go
// windows向けビルド
// GOOS=windows GOARCH=amd64 go build -o excel_win excel.go
package main

import (
	"flag"
	"fmt"

	"github.com/tealeg/xlsx"
)

func main() {
	flag.Parse()
	args := flag.Args()

	//fmt.Println(args)
	if len(args) == 0 {
		fmt.Println("error: no filename")
		return
	}
	filename := args[0]

	dumpSheet(filename, "effects")
}

func dumpSheet(filename, sheetname string) {

	// エクセルを開く
	excel, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// 任意のシートを取得
	sheet, ok := excel.Sheet[sheetname]
	if ok == false {
		fmt.Println("error: sheet name ", sheetname, "is not found")
		return
	}

	// 中身を全て取得
	for _, row := range sheet.Rows {
		for _, cell := range row.Cells {
			fmt.Printf("%v\t", cell.Value)
		}
		fmt.Print("\n")
	}
}
