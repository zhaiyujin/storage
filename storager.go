/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import "github.com/gogf/gf/net/ghttp"

type Storager interface {
	Init() (ali Storager, err error)
	GetObjectToFile()
	PutObjectFromFile(filePath string, file *ghttp.UploadFile) (filePathName string, err error) //上传
	GetUrl(filename string) string
}
