package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func main() {

	url := "https://www.medindia.net/doctors/drug_information/home.asp?alpha="
	resp, err := http.Get(url + "A")
	if err != nil {
		panic(err)
	}
	var medicineArr []string
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	bodyStringified := string(body)
	// fmt.Println(bodyStringified)
	lbody := len(bodyStringified)
	sub := "<h4><a class='bold' href"
	lsub := len(sub)
	// var loc int32
	for i := 0; i< lbody - lsub; i++ {
		if sub == bodyStringified[i: i+lsub] {
			till := i+lsub
			fmt.Println("fist -> " + string(bodyStringified[till-100: till+1]))
			// for j := 1;  ; j++ {
			// 	if '"' == (bodyStringified[till + j]) {
			// 		temp := bodyStringified[till: j]
			// 		medicineArr = append(medicineArr, temp)
			// 		break
			// 	}
			// }
		}
	}
	fmt.Println(medicineArr)
}