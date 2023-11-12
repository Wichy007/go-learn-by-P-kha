package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Address  string
	Postcode string
}

type UserProfile struct {
	FirstName string `json: firstname`
	LastName  string `json: lastname`
	Age       int
	Height    float32
	Address   Address
	// nested struct use anonymous struct
	Bill struct {
		BillAddress string
	}
}

func (u UserProfile) GetCustomerName() string {
	return fmt.Sprintf("Hello Mr. %s %s", u.FirstName, u.LastName)
}

func main() {
	user := map[string]string{}

	user["username"] = "wichy"
	user["password"] = "password"

	fmt.Println(user)
	fmt.Println(user["username"])

	userProfile := UserProfile{
		FirstName: "wichy",
		LastName:  "password",
		Age:       27,
		Height:    12.32,
	}
	fmt.Println(userProfile)

	fmt.Println(userProfile.GetCustomerName())

	//json Marshal and Unmarshal
	// byteTxtJsom, err := json.Marshal(userProfile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// formatedTxtJsom, err := json.MarshalIndent(userProfile, "this parmeter for prefix in ever row", "this parameter for space prefix in every child")
	formatedTxtJsom, err := json.MarshalIndent(userProfile, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(formatedTxtJsom))

	dataJson := `{
		"FirstName": "twin",
		"LastName": "Sombut",
		"Age": 32,
		"Height": 12.32,
		"Address": {
			"Address": "",
			"Postcode": ""
		},
		"Bill": {
			"BillAddress": ""
		}
	}`

	var ProofProfile UserProfile
	err = json.Unmarshal([]byte(dataJson), &ProofProfile)

	fmt.Println(ProofProfile)
	fmt.Println(ProofProfile.GetCustomerName())

}
