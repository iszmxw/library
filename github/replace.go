package github

import (
	"fmt"
	"io/ioutil"
	"library/logger"
	"strings"
)

func ReplaceFileContent(filePath, oldString, newString string) {
	// 读取文件内容
	//filePath := "example.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error(err)
		return
	}

	// 将文件内容转换为字符串
	fileContent := string(content)

	// 替换字符串
	newContent := strings.Replace(fileContent, oldString, newString, -1)

	// 将替换后的内容写回文件
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println("文件内容替换成功!")
}
