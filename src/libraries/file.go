package libraries

import (
	"io"
	"log"
	"mime/multipart"
	"os"
)

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Print(filename + " not exist")
		exist = false
	}
	return exist
}

func Mkdir(path string) bool {
	if !CheckFileIsExist(path) {
		err1 := os.MkdirAll(path, os.ModePerm) //创建文件夹
		if err1 != nil {
			return false
		}
	}

	return true
}