package testProject

import (
	"time"
	"fmt"
)

func main() {
	currencyTime := time.Now()
	date := time.Date(currencyTime.Year(), currencyTime.Month(), currencyTime.Day(), 00, 00, 00, 00, time.Local)
	fmt.Println(date)

	addDate := currencyTime.AddDate(00, 00, 1)
	fmt.Println(addDate)

	fmt.Println("sss", "sss")

	var array = []int{0: 1, 1: 2}
	var arrays = []int{}
	for i, v := range array {
		fmt.Println(i)                           //下标
		fmt.Println(v)                           //下标所对应的值
		fmt.Println("array length:", len(array)) //数组的长度
	}
	//将上一个数组的值copy到这个数组不会改变原数组的值数据
	arrays = array
	for i, v := range arrays {
		fmt.Println("arrays", i)                 //下标
		fmt.Println("arrays", v)                 //下标所对应的值
		fmt.Println("array length:", len(array)) //数组的长度
	}

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}

	user := [...]User{0: {Name: "张三", Age: 18, Sex: "男"}}
	per := [len(user)]Person{}
	for i, v := range user {
		person := Person{
			Name: v.Name,
			Age:  v.Age,
			Sex:  v.Sex,
		}
		fmt.Println(i)
		per[i] = person
	}
	for _, v := range per {
		fmt.Println(v.Name)
		fmt.Println(v.Age)
		fmt.Println(v.Sex)
	}
	type UserType struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	type Address struct {
		AddName   string `json:"add_name"`
		AddNumber int    `json:"add_number"`
	}
	type Class struct {
		ClassName string   `json:"class_name"`
		Teacher   string   `json:"teacher"`
		UserType  UserType `json:"user_type"`
		Address   Address  `json:"address"`
	}

	user_type := UserType{
		ID:   1,
		Name: "大家好",
		Age:  11,
		Sex:  "男",
	}
	address := Address{
		AddName:   "莲花小区",
		AddNumber: 666,
	}
	class := Class{
		ClassName: "幼儿园",
		Teacher:   "丁涛",
		UserType:  user_type,
		Address:   address,
	}
	fmt.Println("student:", user_type)
	fmt.Println("class student:", class)

	go func() {
		{
		}
	}()

	type School struct {
		SchoolName    string `json:"stu_name"`
		SchoolNumber  int    `json:"stu_number"`
		SchoolAddress string `json:"school_address"`
		SchoolTile    string `json:"school_tile"`
	}
	var school School
	school.SchoolName = "兰大"
	school.SchoolNumber = 12345
	school.SchoolAddress = "兰州"
	school.SchoolTile = "创新，奋进"
}
