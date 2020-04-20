package qiniu_upload

import (
	"context"
	"os"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// Bucket 默认bucket
var bucket = "hexo-blog"

func init() {
	// mac := qbox.NewMac(accessKey, secretKey)
}

// UploadFile 上传
func UploadFile(key, filePath, bucket string) (fileKey string, fileHash string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	ak, sk := GetAkSk()

	mac := qbox.NewMac(ak, sk)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		return
	}
	fileKey = ret.Key
	fileHash = ret.Hash
	return
}

// SetBucket 设置域
func SetBucket(bk string) {
	bucket = bk
}

// GetBucket 获取域
func GetBucket() string {
	return bucket
}

// GetAkSk 获取七牛配置
func GetAkSk() (ak, sk string) {
	ak = os.Getenv("QINIU_AK")
	sk = os.Getenv("QINIU_SK")
	return
}
