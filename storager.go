/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import "github.com/gogf/gf/net/ghttp"

//定义接口Storager类型,方法，只有声明，没有实现，由别的类型（自定义类型）实现
type Storager interface {
	Init() (ali Storager, err error)                                                            //实例化
	GetObjectToFile()                                                                           //下载
	PutObjectFromFile(filePath string, file *ghttp.UploadFile) (filePathName string, err error) //上传

}
