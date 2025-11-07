package main

import (
	"fmt"
	"os"
)

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)
}



func CheckFileExist(path string) {
	_, err := os.Stat(path)
	if err == nil {
		fmt.Print("ファイルが見つかりました")
		return
	}
	if os.IsNotExist(err) {
		fmt.Print("ファイルが見つかりませんでした")
		return
	}
	// fmt.Print("ファイルの確認中にエラーが発生しました")
}
