package main

import "fmt"

//1)使用 map[string]map[string]sting 的 map 类型
//2)key: 表示用户名，是唯一的，不可以重复 3)如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息,
//(包括昵称 nickname 和 密码 pwd)。
//4)编写一个函数 modifyUser(users map[string]map[string]sting, name string) 完成上述功能
func main() {
	users := make(map[string]map[string]string, 10)
	users["lilin"] = make(map[string]string, 2)
	users["lilin"]["name"] = "lilin"
	users["lilin"]["pwd"] = "888"
	users["lilin"]["nickName"] = "昵称~~~"

	modifyUser(users, "lilin")
	modifyUser(users, "lilin2")
	modifyUser(users, "lilinawe")
	fmt.Println("users=", users)
	fmt.Println("users.size=", len(users))
}

func modifyUser(users map[string]map[string]string, name string) {

	//判断user是否宝行name
	_, ok := users[name]
	//有这个用户
	if ok {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "99999"
		users[name]["nickName"] = "这是小子儿"
	}
}
