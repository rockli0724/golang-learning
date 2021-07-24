package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir := "/Users/lilin/Documents/doc.txt"
	path := "/Users/lilin/Documents/doc_2.txt"
	readFile(dir)
	readContext(dir)
	readAndWriteFile(path)
	readAndWriteReplaceFile(path)

}

func readFile(dir string) {

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("openFIle error", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(readString)
	}
	fmt.Println("ending..111111.........")
}

func readAndWriteFile(savePath string) {

	file, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("openFIle error", err)
	}
	defer file.Close()

	str := "hello lilin\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//写入缓存
		writer.WriteString(str)
	}
	//写完之后，需要flush() 刷数据至文件
	writer.Flush()
	fmt.Println("ending..111111.........")
}
func readAndWriteReplaceFile(savePath string) {

	file, err := os.OpenFile(savePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("openFIle error", err)
	}
	defer file.Close()

	str := "替换呢绒内容\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		//写入缓存
		writer.WriteString(str)
	}
	//写完之后，需要flush() 刷数据至文件
	writer.Flush()
	fmt.Println("replace..111111.........")
}

//不适用于读取大文件  ioutil.ReadFile()
func readContext(dir string) {
	file, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Println("readContext error.", err)
	}
	//把读取的内容显示到终端
	fmt.Printf("%v", string(file))
	fmt.Println("ending 2.........")
}
