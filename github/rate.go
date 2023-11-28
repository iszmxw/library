package github

import (
	"context"
	"fmt"
	"library/logger"
)

func GetRate() {
	client := InitClient()
	limits, r, err := client.RateLimits(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	fmt.Println(limits.String())
	fmt.Println(r)
}
