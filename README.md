使用前置配置

首先在本地的环境变量中添加

export QINIU_AK=你的accessKey
export QINIU_SK=你的secretKey

然后执行
```
./upload --ctx=/Users/gpf/Documents/www/hugo-blog/content/images --key=images --bucket=hexo-blog'
```

* ctx 要上传的目录路径
* key 上传到七牛时会把目录路径替换成的key, 最终上传时的key  `key/{ctx目录下的文件}`
* bucket 指定要上传到的bucket名