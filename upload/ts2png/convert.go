package ts2png

import (
	"errors"
	"io/ioutil"
	"library/logger"
	"os"
	"path/filepath"
)

func Convert(filePath string) {
	//filePath := "./your_directory" // 替换为你的目录路径
	// Step 1: 将所有.ts文件重命名为.png
	err := renameTSFiles(filePath)
	if err != nil {
		logger.Error(errors.New("Error renaming TS files:" + err.Error()))
		return
	}
	logger.Info("TS重命名为PNG成功")

	// Step 2: 将PNG文件添加PNG文件头
	err = addPNGHeader(filePath)
	if err != nil {
		logger.Error(errors.New("Error adding PNG header:" + err.Error()))
		return
	}
	logger.Info("PNG元数据转PNG-TS成功！")
}

func renameTSFiles(directory string) error {
	fileList, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range fileList {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".ts" {
			newName := filepath.Join(directory, file.Name()[0:len(file.Name())-3]+".png")
			err := os.Rename(filepath.Join(directory, file.Name()), newName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addPNGHeader(directory string) error {
	rewritePath := filepath.Join(directory, "")
	err := os.MkdirAll(rewritePath, os.ModePerm)
	if err != nil {
		return err
	}

	fileList, err := ioutil.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, file := range fileList {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".png" {
			sourcePath := filepath.Join(directory, file.Name())
			destPath := filepath.Join(rewritePath, file.Name())

			// 将PNG文件头拷贝到新文件
			err := copyPNGHeader("/Users/johnyep/service/go/library/upload/tmp/1.png", sourcePath, destPath)
			if err != nil {
				return err
			}
		} else {
			sourcePath := filepath.Join(directory, file.Name())
			destPath := filepath.Join(rewritePath, file.Name())

			// 直接拷贝其他文件
			err := copyFile(sourcePath, destPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// copyPNGHeader 从PNG头文件和源PNG文件中读取数据，将它们合并，然后写入目标文件。
// 参数：
//
//	headerPath: 包含PNG头数据的文件路径
//	sourcePath: 包含源PNG数据的文件路径
//	destPath:   合并数据后写入的目标文件路径
//
// 返回：
//
//	如果操作成功，返回nil；否则返回相应的错误。
func copyPNGHeader(headerPath, sourcePath, destPath string) error {
	// 读取PNG头文件数据
	headerData, err := ioutil.ReadFile(headerPath)
	if err != nil {
		return err
	}

	// 读取源PNG文件数据
	sourceData, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	// 将PNG头文件数据和源PNG文件数据合并
	destData := append(headerData, sourceData...)

	// 将合并后的数据写入目标文件
	err = ioutil.WriteFile(destPath, destData, os.ModePerm)
	if err != nil {
		return err
	}

	// 返回nil表示操作成功
	return nil
}

func copyFile(sourcePath, destPath string) error {
	sourceData, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destPath, sourceData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
