# storage

## 介绍
gf图片上传功能,支持本地上传和阿里云oss上传。

## 主要功能
- 默认是本地文件上传。
- 通过配置支持阿里云oss图片上传。

## 使用说明

配置文件说明：
```toml
[storage]
    driver =    "alioss"
[storage.local]
    root    =   "upload"
    url = "http://192.168.0.126:8199/upload"
[storage.alioss]
    accessId    =   ""
    accessKey   =   ""
    bucket  =   ""
    endpoint    =   "oss-cn-beijing.aliyuncs.com"
    url =   ""
```
 需要设置驱动,当设置local,可以不配置local


### 代码使用方式


```go

fileName, err := storage.NewStorage().PutObjectFromFile(upPath, file)
    if err != nil {
     fmt.Println(err.error())
    }
    fmt.Println(fileName)

```

- file 类型是*ghttp.UploadFile
- upPath 为文件路径
