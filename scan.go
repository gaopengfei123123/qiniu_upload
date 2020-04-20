package qiniu_upload

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ScanDir 获取目录下全部文件路径
func ScanDir(filepath string) []string {
	result := []string{}
	filepath = strings.TrimRight(filepath, "/")

	files, _ := ioutil.ReadDir(filepath)
	for _, f := range files {
		fName := f.Name()
		if string(fName[0]) == "." {
			continue
		}
		pwd := fmt.Sprintf("%s/%s", filepath, fName)
		if f.IsDir() {
			c := ScanDir(pwd)
			result = append(result, c...)
		} else {
			pwd := fmt.Sprintf("%s/%s", filepath, f.Name())
			result = append(result, pwd)
		}
	}

	return result
}

// GetKey 获取上传key
func GetKey(basePath, key, filePath string) string {
	return strings.Replace(filePath, basePath, key, 1)
}
