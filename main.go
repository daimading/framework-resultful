package main

import (
	"fmt"
	"time"
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

	//copy 数组的类型必须相同。
}
