package github

import (
	"context"
	"errors"
	"fmt"
	gt "github.com/google/go-github/v56/github"
	"io/ioutil"
	"library/logger"
)

func Create(message, filePath string) {
	// 请替换以下值为你的GitHub用户名、仓库名和个人访问令牌
	username := "fang54zm"
	repo := "videos"
	mail := "fang@54zm.com"
	client := InitClient()
	// 读取要上传的文件内容
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Error(err)
		return
	}
	commitAuthor := &gt.CommitAuthor{Name: gt.String(username), Email: gt.String(mail)}
	//// 检查文件是否存在
	if _, _, _, err := client.Repositories.GetContents(context.Background(), username, repo, filePath, nil); err == nil {
		logger.Error(errors.New("文件已经存在"))
		return
	}
	// 创建文件
	createOptions := &gt.RepositoryContentFileOptions{
		Message:   gt.String(message),
		Content:   fileContent,
		Committer: commitAuthor,
	}
	file, r, err := client.Repositories.CreateFile(context.Background(), username, repo, filePath, createOptions)
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(file, r)
}
