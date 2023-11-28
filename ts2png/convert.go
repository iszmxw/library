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
			err := copyPNGHeader("/Users/johnyep/service/go/library/ts2png/1.png", sourcePath, destPath)
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

func copyPNGHeader(headerPath, sourcePath, destPath string) error {
	headerData, err := ioutil.ReadFile(headerPath)
	if err != nil {
		return err
	}

	sourceData, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	destData := append(headerData, sourceData...)

	err = ioutil.WriteFile(destPath, destData, os.ModePerm)
	if err != nil {
		return err
	}

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
