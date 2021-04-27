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

/**
真曹丹
 */
func BenchmarkNewStorage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewStorage()
		if err != nil {
			fmt.Println(err.Error())
		}
		//	fmt.Println(storage)
	}
}

func TestStorage(t *testing.T) {
	storage, err := NewStorage()
	if err != nil {
		fmt.Println(err.Error())
	}
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
