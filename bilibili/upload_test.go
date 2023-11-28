package bilibili

import (
	"fmt"
	"testing"
)

func TestUploadImg(t *testing.T) {
	csrf := "477603d6058f90ef6b29c2286691abf6"
	cookie := "SESSDATA=8887890f%2C1716689918%2C5745c%2Ab2CjAmIEhXb31NLV9Jh2B1KtMkQ6HyanbjDZZeC5RGt79ic8LkZAfrn6s-AHUx4C0_D6MSVloxYTVnbVZyN05SejZoUU03UTRNVTk4cXB2Q1hiRU9yeHFRRjVKajEyazVpVXRFRWFQeWVJWVc1M0JSYUtSOWpiREQ2R1ZMbmpEZEROYkQ2ellIUnN3IIEC;"
	filePath := "/Users/johnyep/Documents/GitHub/videos/1/02/out00003.png"
	urls := UploadImg(csrf, cookie, filePath)
	fmt.Println(urls)
}
