package bilibili

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"library/logger"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Size int    `json:"size"`
		Url  string `json:"url"`
	} `json:"data"`
}

func UploadImg(csrf, cookie, filePath string) (urls string) {
	urls = ""
	url := "https://api.bilibili.com/x/article/creative/article/upcover"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("csrf", csrf)
	file, errFile2 := os.Open(filePath)
	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("binary", filepath.Base(filePath))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		logger.Error(errFile2)
		return
	}
	err := writer.Close()
	if err != nil {
		logger.Error(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("authority", "api.bilibili.com")
	req.Header.Add("cookie", cookie)
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	var Resp Res
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		logger.Error(err)
		return
	}
	if Resp.Code != 0 {
		logger.Error(errors.New(Resp.Message))
		return
	}
	urls = Resp.Data.Url
	return
}

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
