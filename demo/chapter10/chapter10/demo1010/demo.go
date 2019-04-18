package main

import "fmt"

type User struct {
	username string
	age      Age
}

type Client struct {
	name string
	age  Age
}

type Age int

func (u *User) sayName() {
	fmt.Println(u.username)
}

func (u User) sayHello() {
	fmt.Println(u.username + ",hello ")
}

func (u User) changeName(newName string) {
	u.username = newName
}

func (u *User) changeName2(newName string) {
	u.username = newName
}

func (c *Client) sayName() {
	fmt.Println(c.name)
}

func (age *Age) sayAge() {
	fmt.Println(*age)
}

func main() {
	user := User{"zhangsan", 12}
	client := Client{"lisi", 14}

	(&user).sayName()
	user.sayName() //接受者可以是某个类型的指针，也可以是某个类型实例本身

	client.sayName()
	user.age.sayAge()

	user.sayHello()
	(&user).sayHello()

	//对于调用函数的是接受者是类型还是类型的指针对于类型的方法其实是一致的，但方法本身声明的接受者类型是指针还是类型本身却对
	//操作接受者有影响，如果是接受者指针，则会改变调用者，而如果是类型，那么实际上改变的拷贝的值，下面四个例子说明这个情况
	user.changeName("wori")
	fmt.Println("user changeName(u User)  result:" + user.username)

	user.changeName2("rini")
	fmt.Println("user changeName2(u* User) result:" + user.username)

	(&user).changeName("nima")
	fmt.Println("&user changeName(u User) result:" + user.username)

	(&user).changeName2("leger")
	fmt.Println("&user changeName2(u* User) result:" + user.username)
}
