package back5

import "fmt"

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
	tBBl.name = "test"
	fmt.Println("save to database", tBBl.name)

}

type TransferKTB struct {
	name string
}

func (tKTB TransferKTB) SaveLog() {
	fmt.Println("save to database", tKTB.name)

}

func back5() {
	transA := TransferBBL{name: "BBL"}
	transB := TransferKTB{name: "KTB"}
	SaveLog(&transA)
	SaveLog(transB)

	fmt.Println(transA)
}
