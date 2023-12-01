package bilibili

import (
	"fmt"
	"io/ioutil"
	"library/logger"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func bitRate(file string) float64 {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=bit_rate", "-of", "default=noprint_wrappers=1:nokey=1", file)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	bitRate, err := strconv.Atoi(string(output))
	if err != nil {
		fmt.Println("Error converting bit rate to integer:", err)
		return 0
	}
	return float64(bitRate)
}

func videoCodec(file string) string {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=codec_name", "-of", "default=noprint_wrappers=1:nokey=1", file)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	codecs := strings.Split(string(output), "\n")
	for _, codec := range codecs {
		if codec != "h264" {
			return "copy"
		}
	}
	return "VCODEC"
}

func safeName(file string) string {
	// 获取文件的基本名称
	baseName := filepath.Base(file)
	// 替换双引号为转义的双引号
	safeFileName := strings.ReplaceAll(baseName, `"`, `\"`)
	// 添加双引号
	safeFileName = `"` + safeFileName + `"`
	return file
}

func genSlice(file string, time int) string {
	var sub strings.Builder
	rate := bitRate(file)
	vcodec := videoCodec(file)
	maxBits := 8 * 1024 // 假设 uploader().MAX_BYTES 的值是 1024，转换为比特数
	segmentTime := 20
	if rate != 0 {
		segmentTime = Min(20, maxBits/int(rate*1.35))
	}
	segmentTimeArg := fmt.Sprintf(" -segment_time %d", segmentTime)
	if time > 0 {
		segmentTimeArg = fmt.Sprintf(" -segment_time %d", time)
	}
	// SEGMENT_TIME
	sub.WriteString(segmentTimeArg)
	return fmt.Sprintf("ffmpeg -y -i %s -c:v %s -c:a aac -bsf:v h264_mp4toannexb -map 0:v:0 -map 0:a? -f segment -segment_list out.m3u8 %s out%%05d.ts", safeName(file), vcodec, sub.String())
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sameParams(tmpdir, command string) bool {
	// 实现 sameParams 函数的逻辑，用于比较参数是否相同
	// 这里简化为比较目录和命令是否相同
	return tmpdir == command
}

func Up() {
	tmpdir := "/Users/johnyep/service/go/library/bilibili/video"
	file := "/Users/johnyep/service/go/library/bilibili/VID_20180615_194851.mp4"
	time := 4 // 设置为0表示使用默认的切片时间
	command := genSlice(file, time)

	// 切换到临时目录
	if sameParams(tmpdir, command) {
		err := os.Chdir(tmpdir)
		if err != nil {
			logger.Error(err)
			return
		}
	} else {
		// 创建临时目录
		if err := os.Mkdir(tmpdir, os.ModePerm); err != nil {
			fmt.Println("Error creating temporary directory:", err)
			return
		}
		// 切换到临时目录
		if err := os.Chdir(tmpdir); err != nil {
			fmt.Println("Error changing to temporary directory:", err)
			return
		}
		// 执行命令
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error running command:", err)
			return
		}

		// 记录命令到文件
		if err := ioutil.WriteFile("command.sh", []byte(command), os.ModePerm); err != nil {
			fmt.Println("Error writing command to file:", err)
			return
		}
	}
}
