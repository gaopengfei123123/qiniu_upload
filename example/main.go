package main

import (
	"flag"
	"fmt"
	"qiniu_upload"
)

var ctx = flag.String("ctx", ".", "请输入上传文件目录, 默认当前文件")
var key = flag.String("key", "images", "请设置key值, 将目录下的文件公用这一个key, 默认 images")
var bucket = flag.String("bucket", "hexo-blog", "设置bucket, 默认 default")

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(*ctx, *key)

	files := qiniu_upload.ScanDir(*ctx)
	for _, fn := range files {
		key := qiniu_upload.GetKey(*ctx, *key, fn)
		key, _, err := qiniu_upload.UploadFile(key, fn, qiniu_upload.GetBucket())
		if err != nil {
			fmt.Printf("上传失败: %s \n", fn)
			continue
		}
		fmt.Printf("上传成功: key: %s \n filePath: %s \n", key, fn)
	}
}
