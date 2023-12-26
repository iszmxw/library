package itab

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"library/logger"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Img  string `json:"img"`
}

func generateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result += string(chars[randomIndex.Int64()])
	}
	return result
}

func UploadImg(filePath string) (urls string) {
	urls = ""
	url := "http://localhost/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	defer file.Close()
	part2, errFile2 := writer.CreateFormFile("file", filepath.Base(filePath))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		logger.Error(errFile2)
		return
	}
	filename := fmt.Sprintf("user-website-icon/%s/%s.png", time.Now().Format("20060102"), generateRandomString(16))
	fmt.Println(filename)
	_ = writer.WriteField("filename", filename)
	err = writer.Close()
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
	if Resp.Code != 1 {
		logger.Error(errors.New(Resp.Msg))
		return
	}
	urls = Resp.Img
	return
}
