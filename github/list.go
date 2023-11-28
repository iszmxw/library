package github

import (
	"context"
	"fmt"
	"log"
	"strings"
)

var i = 0

func List() {
	// 请替换以下值为你的GitHub用户名、仓库名和个人访问令牌
	username := "fang54zm"
	repo := "videos"
	getContentsRecursively(username, repo, "/")
	fmt.Println(i)
	fmt.Println(i)
	fmt.Println(i)
	fmt.Println(i)
	fmt.Println(i)
}

// 递归获取目录内容的函数
func getContentsRecursively(username, repo, path string) {
	client := InitClient()
	fileContent, directoryContent, _, err := client.Repositories.GetContents(context.Background(), username, repo, path, nil)
	if err != nil {
		log.Fatal(err)
	}
	if fileContent != nil {
		// 处理文件内容
		fmt.Println(*fileContent.Name, *fileContent.Path, *fileContent.SHA, *fileContent.DownloadURL)
	}
	for _, content := range directoryContent {
		// 处理目录内容
		if *content.Type == "dir" {
			// 递归调用自身，获取子目录的内容
			getContentsRecursively(username, repo, *content.Path)
		}
		// 处理目录内容
		if *content.Type == "file" {
			// 处理文件内容
			i++
			if strings.Contains(*content.Name, ".ts") {
				fmt.Println(*content.Name, *content.Path, *content.SHA, *content.DownloadURL)
				filePath := "videos/tmp/github.m3u8"
				oldString := *content.Name
				newString := *content.DownloadURL
				ReplaceFileContent(filePath, oldString, newString)
			}
		}
	}
}
