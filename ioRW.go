package main

import (
	"fmt"

	"io/ioutil"
	"log"
)

//var name string
//var age int
//var inputReader *bufio.Reader

func main() {
	//fmt.Println("请输入你的名字，年龄：")
	////fmt.Scanln(&name,&age)
	////fmt.Println(name,age)
	//inputReader := bufio.NewReader(os.Stdin)
	//input,err := inputReader.ReadString('\n')
	//if err!= nil{
	//	log.Println(err.Error())
	//}
	//fmt.Println(input)
	fmt.Println("打开文件。。。")
	//inputfile,err := os.Open("goTest/hello.go")
	//if err!= nil{
	//	log.Printf("文件打开错误%s",err.Error())
	//}
	//defer inputfile.Close()
	fmt.Println("读取文件。。。")
	hello, err := ioutil.ReadFile("goTest/hello.go")
	if err != nil {
		log.Printf("读取失败%s", err.Error())
	}

	for i, _ := range hello {
		fmt.Print(string(hello[i]))

	}

	fmt.Println("将该文件写入其它文件中。。。")
	if err := ioutil.WriteFile("goTest/input.dat", hello, 0777); err != nil {
		log.Println("write error: ", err.Error())
	}
	fmt.Println("写入成功！！！")

}
