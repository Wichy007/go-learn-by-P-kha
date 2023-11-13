package main

import (
	"fmt"
	"time"
)

func Getmember() {
	fmt.Println("Please wait ...")
	time.Sleep(3 * time.Second)
}

func IsInSystem(username string) bool {
	return true
}

func GetUserDetail(username string) (int, string) {
	return 201, "manager"
}

func CheckLogin(username string, password string) {
	if IsInSystem(username) {
		fmt.Println("user is found in system")
		code, role := GetUserDetail(username)
		greetingTxt := fmt.Sprintf("hello %s code number %d", role, code)
		fmt.Println(greetingTxt)
	}
}

func LogEnd() {
	CheckServeResponse()
	time.Now()
	fmt.Println("program end")
	fmt.Println(time.Now())
}

func CheckServeResponse() {
	fmt.Println("check server time")
	time.Sleep(3 * time.Second)
	panic("oh no")
}

func Back2() {
	// defer ประกาศก่อนทำทีหลัง
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	defer LogEnd()

	Getmember()
	var username string
	// fmt.Scan(&username)
	CheckLogin(username, "password")
}
