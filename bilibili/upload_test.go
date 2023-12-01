package bilibili

import (
	"fmt"
	"io/ioutil"
	"library/logger"
	"strings"
	"testing"
)

func TestUploadImg(t *testing.T) {
	csrf := "477603d6058f90ef6b29c2286691abf6"
	cookie := "SESSDATA=8887890f%2C1716689918%2C5745c%2Ab2CjAmIEhXb31NLV9Jh2B1KtMkQ6HyanbjDZZeC5RGt79ic8LkZAfrn6s-AHUx4C0_D6MSVloxYTVnbVZyN05SejZoUU03UTRNVTk4cXB2Q1hiRU9yeHFRRjVKajEyazVpVXRFRWFQeWVJWVc1M0JSYUtSOWpiREQ2R1ZMbmpEZEROYkQ2ellIUnN3IIEC;"
	filePath := "/Users/johnyep/service/go/library/bilibili/tmp/WechatIMG729.jpg"
	urls := UploadImg(csrf, cookie, filePath)
	fmt.Println(urls)
}

func TestBathUploadImg(t *testing.T) {
	// 替换以下路径为你要读取的文件夹路径
	dirPath := "/Users/johnyep/service/go/library/bilibili/video"
	ScanFile(dirPath)
}

func ScanFile(dirPath string) {
	m3u8 := "/Users/johnyep/service/go/library/bilibili/video/out.m3u8"
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logger.Error(err)
		return
	}
	for _, file := range files {
		path := fmt.Sprintf("%s/%s", dirPath, file.Name())
		if file.IsDir() {
			ScanFile(path)
		} else {
			if !strings.Contains(file.Name(), ".png") {
				continue
			}
			// 读取文件内容
			content, err := ioutil.ReadFile(m3u8)
			if err != nil {
				logger.Error(err)
				return
			}
			// 不包含就得文件名称，说明已经被替换处理过了
			if !strings.Contains(string(content), file.Name()) {
				continue
			}
			// 上传图片
			csrf := "477603d6058f90ef6b29c2286691abf6"
			cookie := "SESSDATA=8887890f%2C1716689918%2C5745c%2Ab2CjAmIEhXb31NLV9Jh2B1KtMkQ6HyanbjDZZeC5RGt79ic8LkZAfrn6s-AHUx4C0_D6MSVloxYTVnbVZyN05SejZoUU03UTRNVTk4cXB2Q1hiRU9yeHFRRjVKajEyazVpVXRFRWFQeWVJWVc1M0JSYUtSOWpiREQ2R1ZMbmpEZEROYkQ2ellIUnN3IIEC;"
			urls := UploadImg(csrf, cookie, path)
			if len(urls) == 0 {
				continue
			}
			// 替换原有 m3u8 的文件地址
			ReplaceFileContent(m3u8, file.Name(), urls)
		}
	}
}

func TestReplaceFileContent(t *testing.T) {
	ReplaceFileContent("/Users/johnyep/service/go/library/bilibili/tmp/luoji01.m3u8", "out00000.png", "tt.args.newString")
}
