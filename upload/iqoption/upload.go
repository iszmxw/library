package iqoption

import (
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
)

type Res struct {
	IsSuccessful bool          `json:"isSuccessful"`
	Message      []interface{} `json:"message"`
	Result       struct {
		Id             string `json:"id"`
		Status         string `json:"status"`
		ValidateStatus string `json:"validateStatus"`
		Url            string `json:"url"`
	} `json:"result"`
}

func UploadImg(cookie, filePath string) (urls string) {
	urls = ""
	url := "https://api.iqoption.com/v1/avatars"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(filePath)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("image", filepath.Base(filePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		logger.Error(errFile1)
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
	req.Header.Add("authority", "api.iqoption.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", cookie)
	req.Header.Add("origin", "https://iqoption.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://iqoption.com/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
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
	if !Resp.IsSuccessful {
		logger.Error(errors.New(fmt.Sprintf("%v", Resp.Message)))
		return
	}
	logger.Info(Resp)
	urls = Resp.Result.Url
	return
}
