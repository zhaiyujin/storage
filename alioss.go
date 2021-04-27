/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)
import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type AliOss struct {
	client *oss.Client
	bucket *oss.Bucket
	config *ossConfig
}

type ossConfig struct {
	AccessId         string
	AccessKey        string
	Bucket           string
	Endpoint         string
	EndpointInternal string
	CdnDomain        string
	Ssl              bool
	Url              string
}

func NewAliOss() (ali Storager, err error) {
	return (&AliOss{}).Init()
}

func (a *AliOss) Init() (ali Storager, err error) {
	config := &ossConfig{}
	var (
		bucket *oss.Bucket
		client *oss.Client
	)
	if err = g.Cfg().GetStruct("storage.alioss", config); err != nil {
		return
	}

	client, err = oss.New(config.Endpoint, config.AccessId, config.AccessKey)
	if err != nil {
		return
	}
	if config.Bucket != "" {
		bucket, err = client.Bucket(config.Bucket)
		if err != nil {
			return
		}
	}

	return &AliOss{
		client: client,
		bucket: bucket,
		config: config,
	}, nil

}

func (a *AliOss) GetObjectToFile() {
	fmt.Println("ceshi ali")
}

func (a *AliOss) PutObjectFromFile(filePath string, file *ghttp.UploadFile) (filePathName string, err error) {
	//获取完整文件名称设置随机名称
	local, err := (&Local{}).Init()
	filePathName, err = local.PutObjectFromFile(filePath, file)
	if err = a.bucket.PutObjectFromFile(filePathName, serverRoot+filePathName); err != nil {
		return
	}
	return

}
