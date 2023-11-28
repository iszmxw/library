package github

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	List()
}

func TestGithub(t *testing.T) {
	// List()
	message := "上传文件"
	// 准备要上传的文件信息，包括文件的路径、消息、内容和提交者的信息
	filePath := "videos/tmp/out00000.ts"
	Create(message, filePath)
}

func TestBathCreate(t *testing.T) {

	// 替换以下路径为你要读取的文件夹路径
	dirPath := "videos/tmp"
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Files in", dirPath, ":")
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", dirPath, file.Name())
		Create(path, path)
	}
	fmt.Println("ok")
}

func TestGetRate(t *testing.T) {
	// GetRate()
	filePath := "videos/tmp/out.m3u8"
	oldString := "out00000.ts"
	newString := "https://raw.githubusercontent.com/fang54zm/videos/main/"
	ReplaceFileContent(filePath, oldString, newString)
}
