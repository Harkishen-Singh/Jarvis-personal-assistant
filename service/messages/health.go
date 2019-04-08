package messages

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)

type medicineResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Result []string `json:"result"`
}

type medicineslist struct {
	A []string `json:"A"`
	B []string `json:"B"`
	C []string `json:"C"`
	D []string `json:"D"`
	E []string `json:"E"`
	F []string `json:"F"`
	G []string `json:"G"`
	H []string `json:"H"`
	I []string `json:"I"`
	J []string `json:"Jarr"`
	K []string `json:"K"`
	L []string `json:"L"`
	M []string `json:"M"`
	N []string `json:"N"`
	O []string `json:"O"`
	P []string `json:"P"`
	Q []string `json:"Q"`
	R []string `json:"R"`
	S []string `json:"S"`
	T []string `json:"T"`
	U []string `json:"U"`
	V []string `json:"V"`
	W []string `json:"W"`
	X []string `json:"X"`
	Y []string `json:"Y"`
	Z []string `json:"Z"`
}

var (
	medicineParser medicineslist
)

func init() {

	fmt.Println("Loading health-medicine JSON parsers....")
	medicineFile, err := os.Open("store/medicine_database.json")
	bytvalMF, _ := ioutil.ReadAll(medicineFile)
	if err != nil   {
		panic(err)
	}
	err1 := json.Unmarshal(bytvalMF, &medicineParser)
	if err1 != nil {
		panic(err1)
	}
}

// HealthController controls tasks related to health services
func HealthController(medicine string,  res http.ResponseWriter) (speech string) {

	fmt.Println("Health controllers")
	medicine = strings.TrimSpace(medicine)
	firstAlpha := medicine[0]
	if firstAlpha == 'R' {
		for i := 0; i < len(medicineParser.R) ; i++ {
			if medicine == medicineParser.R[i] {
				// handleResponse(scrapMedicineLog(&medicineParser.A[i]), res)
				fmt.Println("inside")
				x := "Ramelteon"
				speech = handleResponse(scrapMedicineLog(&x), res)
			}
		}
	}
	return
}

func scrapMedicineLog(medicine *string) string {

	directory, _ := os.Getwd()
	fmt.Println("health-medicine request")
	fmt.Println(" medicine-name -> " + *medicine + " direc -> " + directory)
	result, err := exec.Command("node", "subprocesses/health_medicine.js", *medicine).Output()
	if err != nil {
		panic(err)
	}
	stringified := string(result)
	fmt.Println("result is" , stringified)
	return processScrapLog(&stringified)
}

func processScrapLog(log *string) (data string) {
	logStr := string(*log)
	llogStr := len(logStr)
	subl := "data ->"
	lsubl := len(subl)
	for i:=0; i< llogStr - lsubl; i++ {
		if subl == logStr[i: i+lsubl] {
			data = logStr[i+lsubl: llogStr]
			break
		}
	}
	fmt.Println("scrapped is -> ", data)
	return
}

func handleResponse(data string, res http.ResponseWriter) string {

	resp := medicineResponse{
		Status: true,
		Message: data,
		Result: nil,
	}
	send, _ := json.Marshal(resp)
	res.Write(send)
	return "generic medicine " + data[0: 500]
}