package qiniu_upload

import (
	"fmt"
	"strings"
	"testing"
)

func TestScanDir(t *testing.T) {
	d := "./"
	files := ScanDir(d)

	filePool := map[string]string{}
	for _, fn := range files {
		key := strings.Replace(fn, d, "images", 1)
		filePool[key] = fn

		fmt.Printf("key: %s , path: %s \n", key, fn)

	}

}

func TestUpload(t *testing.T) {
	key := "images/zero-width-space-2.png"
	path := "/Users/gpf/Documents/www/hugo-blog/content/images/zero-width-space-2.png"
	key, hash, err := UploadFile(key, path, GetBucket())
	if err != nil {
		t.Error(err)
	}
	t.Logf("upload success, key: %s, hash: %s \n", key, hash)
}

func TestEnv(t *testing.T) {
	t.Log("start")
	ak, sk := GetAkSk()
	t.Log(ak)
	t.Log(sk)
}
