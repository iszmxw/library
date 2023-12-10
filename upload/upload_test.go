package upload

import (
	"fmt"
	"library/upload/bili"
	"library/upload/ffmpeg"
	"library/upload/iqoption"
	"library/upload/replace"
	"library/upload/ts2png"
	"testing"
)

// 视频切片
func TestFfmpeg(t *testing.T) {
	// 临时工作目录
	tmpdir := "/Users/johnyep/service/go/library/upload/tmp/video"
	// 需要处理的视频
	file := "/Users/johnyep/service/go/library/upload/tmp/video/lj-02-01.mp4"
	fileSize := 4 // 设置为 0 表示使用默认的切片大小，这里为 4M
	ffmpeg.Ffmpeg(tmpdir, file, fileSize)
}

// ts 文件转 png 进行伪装
func TestConvert(t *testing.T) {
	ts2png.Convert("/Users/johnyep/service/go/library/upload/tmp/video")
}

// 测试上传图片到 bilibili
func TestUploadImg(t *testing.T) {
	csrf := "477603d6058f90ef6b29c2286691abf6"
	cookie := "SESSDATA=8887890f%2C1716689918%2C5745c%2Ab2CjAmIEhXb31NLV9Jh2B1KtMkQ6HyanbjDZZeC5RGt79ic8LkZAfrn6s-AHUx4C0_D6MSVloxYTVnbVZyN05SejZoUU03UTRNVTk4cXB2Q1hiRU9yeHFRRjVKajEyazVpVXRFRWFQeWVJWVc1M0JSYUtSOWpiREQ2R1ZMbmpEZEROYkQ2ellIUnN3IIEC;"
	filePath := "/Users/johnyep/service/go/library/upload/tmp/1.png"
	urls := bili.UploadImg(csrf, cookie, filePath)
	fmt.Println(urls)
}

// 测试上传图片到 bilibili
func TestIqOptionUploadImg(t *testing.T) {
	cookie := "platform=15; geo=cn; referrer=https://iq-option.com/; _gcl_au=1.1.1487996637.1701657696; _fbp=fb.1.1701657696390.245648631; _gid=GA1.2.1313805749.1701657675; _gac_UA-44367767-1=1.1701657678.Cj0KCQiA67CrBhC1ARIsACKAa8TsxY_lv8b0tOfRpCF3-MfBV1KFC6iJorrfc9AB1MNUjEwYF-ZVHsoaAkbzEALw_wcB; _hjFirstSeen=1; _hjIncludedInSessionSample_3225446=0; _hjSession_3225446=eyJpZCI6IjIyNWUzMjJkLTgxMzQtNGQwMi1iM2I0LWI5ODkxOTc0ZmE5NiIsImNyZWF0ZWQiOjE3MDE2NTc2OTY2OTMsImluU2FtcGxlIjpmYWxzZSwic2Vzc2lvbml6ZXJCZXRhRW5hYmxlZCI6ZmFsc2V9; _hjAbsoluteSessionInProgress=1; aff_history=[{\"aff\":\"0\",\"afftrack\":null,\"aff_model\":null,\"date\":1701657697030,\"landing\":null}]; aff=1; aff_model=; affextra=; afftrack=GAD_ALL_EN_01_Brand_Web_1708286313_kwd-1247123479__CONVTRANSFR__clickid-Cj0KCQiA67CrBhC1ARIsACKAa8TsxY_lv8b0tOfRpCF3-MfBV1KFC6iJorrfc9AB1MNUjEwYF-ZVHsoaAkbzEALw_wcB; retrack=; identity=901b317a53bee6be26e4437608f80e9fc28555193d58c6a1b74b07f8385825e3d6f91c22bf297ee67d724f9cc640dbf808a3777689e69fb16127f8c5e576b3c6536d2fa1ae20d0200831ff4e3aad6d9c9f8a141cd0d017289fb9ab0a30a253a4272da0bd9115d8e9e3d2689c8d862bebae7bbfa37af986cb19cf02bd11b27500f32ab591465630e11056c1ceab425c2fb8776bff3f8619a6c5ff06d1e9fc06bb4d3619e62c31c6b6e6be07e14bce9d11db07b45f8dfe3206e6be07e14bce9d11fecfc8cc5ec99f427e58a3757ce1eb443d755b4816b7b5da; landing=iqoption.com; lang=zh_CN; pll_language=cn; _ym_uid=1691402666529935045; _ym_d=1701657698; AffTrackGroup=BT_GL_WEB; _hjSessionUser_3225446=eyJpZCI6IjY4NDYyNjBiLWIwYTgtNTRjZS1iYWMzLTUwYTc0MGJmN2RjNyIsImNyZWF0ZWQiOjE3MDE2NTc2OTY2OTIsImV4aXN0aW5nIjp0cnVlfQ==; _ym_isad=2; _ym_visorc=b; device_id=9xHnsDadcq4HYph03JIg; afUserId=6f9ed222-176b-4d0f-9874-afaaf17c10a0-p; AF_SYNC=1701657702118; ssid=cc8fcc7dce34500ce4d31d48d759101e; _ga=GA1.1.754698029.1701657675; MgidSensorNVis=5; MgidSensorHref=https://iqoption.com/cn/profile/personal; _ga_BH1SENMS6L=GS1.1.1701657696.1.1.1701657718.38.0.0"
	filePath := "/Users/johnyep/service/go/library/upload/tmp/1.png"
	urls := iqoption.UploadImg(cookie, filePath)
	fmt.Println(urls)
}

// 测试批量上传图片
func TestBathUploadImg(t *testing.T) {
	// 替换以下路径为你要读取的文件夹路径
	dirPath := "/Users/johnyep/service/go/library/upload/tmp/video"
	m3u8Path := "/Users/johnyep/service/go/library/upload/tmp/video/itab.m3u8"
	ffmpeg.ScanFile(dirPath, m3u8Path)
}

func TestReplaceFileContent(t *testing.T) {
	replace.ReplaceFileContent("/Users/johnyep/service/go/library/upload/tmp/luoji01.m3u8", "out00000.png", "tt.args.newString")
}
