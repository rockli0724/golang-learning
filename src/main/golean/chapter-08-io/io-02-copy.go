package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	source := "/Users/lilin/Documents/doc.txt"
	target := "/Users/lilin/Documents/doc_target.txt"
	copy := "/Users/lilin/Documents/doc_copy.txt"

	exist, err2 := pathExist(source)
	fmt.Println(exist, err2)
	file, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Printf("readFile error=%v", err)
		return
	}
	err = ioutil.WriteFile(target, file, 0666)
	if err != nil {
		fmt.Printf("writer file error=%v", err)
		return
	}
	_, errs := CopyFile(copy, source)
	if errs == nil {
		fmt.Println("copy success")
	} else {
		fmt.Printf("copy error =%v", errs)
	}

	CharCountFunc(source)
	hostname, _ := os.Hostname()
	fmt.Println("host_name=", hostname)

}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, nil
}

//拷贝文件
func CopyFile(dstFileName, srcFilename string) (written int64, err error) {
	sourceFile, err := os.Open(srcFilename)
	if err != nil {
		fmt.Printf("open file error=%v \n", err)
	}
	defer sourceFile.Close()
	//通过srcFile 获取Reader
	reader := bufio.NewReader(sourceFile)
	//打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func CharCountFunc(path string) {
	open, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file error=%v", err)
		return
	}
	defer open.Close()

	var char CharCount
	reader := bufio.NewReader(open)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str := []rune(readString)
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				char.ChCount++
			case v == ' ' && v == '\t':
				char.SpaceCount++
			case v >= '0' && v <= '9':
				char.NumCount++
			default:
				char.OtherCount++
			}
		}
	}
	fmt.Printf("字符个数=%v 数字个数=%v 空格个数=%v 其他字符个数=%v",
		char.ChCount, char.NumCount, char.SpaceCount, char.OtherCount)

}

type CharCount struct {
	ChCount    int //英文
	NumCount   int //数字
	SpaceCount int //空格
	OtherCount int //其他
}
