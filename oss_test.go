/**
 * package: gfoss
 * author: zhaiyujin
 * description:
 */
package storage

import (
	"fmt"
	"testing"
)

//操蛋v1.1.1  ---v1.1.1中修改bug
func BenchmarkNewStorage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		storage := NewStorage()
		fmt.Println(storage)
		//	fmt.Println(storage)
	}
}

func TestStorage(t *testing.T) {
	storage := NewStorage()

	fmt.Println(storage)
	/*filename, err := storage.PutObjectFromFile("jpeg", "2.jpeg")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(filename)*/
	//storage.Init()
}

/*func TestOss(t *testing.T) {
	_, err := NewAliOssServer().PutObjectFromFile("jpeg", "2.jpeg")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestLocal(t *testing.T) {
	file, err := NewLocalFile().PutObjectFromFile("jpeg", "2.jpeg")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(file)
}*/
