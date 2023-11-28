package bilibili

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"library/logger"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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
