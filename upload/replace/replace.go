package replace

import (
	"bufio"
	"fmt"
	"io"
	"library/logger"
	"os"
	"strings"
)

func ReplaceFileContent(filePath, oldString, newString string) {
	// 以读写模式打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		logger.Error(err)
		return
	}
	defer file.Close()

	// 创建一个新的写入缓冲
	var result strings.Builder
	scanner := bufio.NewScanner(file)

	// 逐行读取文件内容
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") || strings.Contains(line, "http") {
			// 不用替换的直接写回去
			result.WriteString(line + "\n")
			continue
		}
		// 替换字符串
		newLine := strings.Replace(line, oldString, newString, -1)

		// 将替换后的行写入缓冲
		result.WriteString(newLine + "\n")
	}

	// 检查扫描是否出错
	if err := scanner.Err(); err != nil {
		logger.Error(err)
		return
	}

	// 将缓冲中的内容写回文件
	err = file.Truncate(0) // 清空文件内容
	if err != nil {
		logger.Error(err)
		return
	}
	_, err = file.Seek(0, io.SeekStart) // 将文件指针移到文件开头
	if err != nil {
		logger.Error(err)
		return
	}
	_, err = file.WriteString(result.String())
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(oldString, "文件内容替换成功!", newString)
}
