package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	CreateGBT2260Table()
}
func CreateGBT2260Table() {
	file, err := os.Open("./data/GBT2260-202304.csv")
	if err != nil {
		fmt.Println("Gen File Error:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var line string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break //文件读完了就结束
		} else if err != nil {
			fmt.Println("Error:", err)
		}
		code := record[0]
		name := record[1]
		line = line + "{\"" + code + "\",\"" + name + "\"},"
	}
	// 官方提供数据有缺失，手工补充直辖市，市辖区编码
	line = line + "{\"110100\",\"市辖区\"},{\"120100\",\"市辖区\"},{\"310100\",\"市辖区\"},{\"500100\",\"市辖区\"},"
	content := "package gbt2260;func GetGbt2260Table() [][]string {gbt2260Table := [][]string{" + line + "};return gbt2260Table;}"
	gbtFile, err := os.Create("gbt2260Table.go")
	if err != nil {
		fmt.Println("Gen File Error:", err)
		return
	}
	defer gbtFile.Close()
	_, err = gbtFile.Write([]byte(content))
	if err != nil {
		fmt.Println("write file error", err)
		return
	}
}
