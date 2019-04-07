package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main() {

	url := "https://www.medindia.net/doctors/drug_information/ramelteon.htm"
	resp, err := http.Get(url )
	if err != nil {
		panic(err)
	}
	var medicineArr []string
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	fmt.Println(medicineArr)
}