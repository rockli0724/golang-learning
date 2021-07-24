package main

import (
	"fmt"
	service2 "golang-learning/chapter-07-object/service"
)

type customerView struct {
	key             string //接收用户输入
	loop            bool   //表示是否循环的显示主菜单
	customerService *service2.CustomerService
}

func (this customerView) List() {
	customers := this.customerService.List()
	fmt.Println("客户列表-----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("客户列表end-----------")

}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("---------客户信息管理系统---------")
		fmt.Println("1.添加客户---------")
		fmt.Println("2.修改客户---------")
		fmt.Println("3.删除客户---------")
		fmt.Println("4.客户列表---------")
		fmt.Println("5.退   出---------")
		fmt.Println("请选择（1-5）：")
		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			cv.List()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入有误。。。。")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("已退出系统")
}

func main() {
	view := customerView{
		key:  "",
		loop: false,
	}
	view.customerService = service2.NewCustomerService()
	view.mainMenu()

}
