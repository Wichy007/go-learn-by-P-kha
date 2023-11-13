package main

import (
	"fmt"
	"reflect"
)

type SaveLogInterface interface {
	SaveLog()
}

func SaveLog(st SaveLogInterface) {
	st.SaveLog()
}

type TransferBBL struct {
	name string
}

func (tBBl *TransferBBL) SaveLog() {
	fmt.Println("save to database", tBBl.name)
	fmt.Println(tBBl)
	tBBl.name = "test change name with pointer"

}

type TransferKTB struct {
	name string
}

func (tKTB TransferKTB) SaveLog() {
	fmt.Println("save to database", tKTB.name)
	tKTB.name = "test change name without pointer"

}

func main() {
	transA := TransferBBL{name: "BBL"}
	transB := TransferKTB{name: "KTB"}
	SaveLog(&transA)
	SaveLog(transB)

	fmt.Println(transA)
	fmt.Println(transB)

	var a int

	fmt.Println(reflect.TypeOf(&a))
}
