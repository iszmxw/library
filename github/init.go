package github

import (
	"context"
	gt "github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
)

var Client *gt.Client

func InitClient() *gt.Client {
	// 请替换以下值为你的GitHub用户名、仓库名和个人访问令牌
	// token := "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"
	//token := "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"
	token := "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"
	// 使用个人访问令牌进行身份验证
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	if Client != nil {
		return Client
	}
	Client = gt.NewClient(tc) // 调用递归函数
	return Client
}
