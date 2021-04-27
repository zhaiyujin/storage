/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"io"
	"strconv"
	"strings"
)

type Local struct {
	file   *ghttp.UploadFile
	config *localConfig
}

type localConfig struct {
	Root string
	Url  string
}

func NewLocalFile() (s Storager, err error) {
	local := Local{}
	return local.Init()
}

func (a *Local) Init() (s Storager, err error) {
	config := &localConfig{}
	err = g.Cfg().GetStruct("storage.local", config)
	if err != nil {
		return
	}
	return &Local{
		config: config,
	}, nil
}

func (a *Local) GetObjectToFile() {
	fmt.Println("ceshi local")
}

func (a *Local) PutObjectFromFile(filePath string, uploadFile *ghttp.UploadFile) (filePathName string, err error) {
	g.Dump(filePath)
	if !gfile.Exists(filePath) {

		if err = gfile.Mkdir(filePath); err != nil {
			return "", err
		}
	} else if !gfile.IsDir(filePath) {

		return "", errors.New(`parameter "dirPath" should be a directory path`)
	}

	//打开文件
	file, err := uploadFile.Open()
	if err != nil {
		return "", err
	}
	//结束后关闭文件
	defer file.Close()
	//获取完整文件名称设置随机名称
	name := gfile.Basename(uploadFile.Filename)
	name = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
	name = name + gfile.Ext(uploadFile.Filename)

	filePathName = gfile.Join(filePath, name)
	newFile, err := gfile.Create(gfile.Join(serverRoot, filePathName))
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	//glog.Printf(`save upload file: %s`, filePathName)
	if _, err := io.Copy(newFile, file); err != nil {
		return "", err
	}

	return filePathName, nil

}
