package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	scrapper "github.com/Harkishen-Singh/Jarvis-personal-assistant/service/utils"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type medicineResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
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
	J []string `json:"J"`
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

type symptomsObj struct {
	Type string
	Link string
}

type symptomslist struct {
	symp []symptomsObj
}

var (
	medicineParser medicineslist
	symptomParser  symptomslist
)

func init() {
	prefix := ""
	if os.Getenv("ENV") == "test" {
		prefix = "../"
	}
	fmt.Println("Loading health-medicine JSON parsers....")
	medicineFile, err := os.Open(fmt.Sprintf("%sstatic/medicine_database.json", prefix))
	bytvalMF, _ := ioutil.ReadAll(medicineFile)
	if err != nil {
		logger.Error(err)
	}
	err1 := json.Unmarshal(bytvalMF, &medicineParser)
	if err1 != nil {
		logger.Error(err1)
	}

	fmt.Println("Loading health-symptoms JSON parsers....")
	medicineFileSymp, err := os.Open(fmt.Sprintf("%sstatic/health_symptoms.json", prefix))
	bytvalMF2, _ := ioutil.ReadAll(medicineFileSymp)
	if err != nil {
		logger.Error(err)
	}
	err1 = json.Unmarshal(bytvalMF2, &symptomParser.symp)
	if err1 != nil {
		logger.Error(err1)
	}
}

// HealthMedController controls tasks related to medicines
func HealthMedController(medicine string, res http.ResponseWriter) (speech string) {

	medicine = strings.ToLower(strings.TrimSpace(medicine))
	medicine = strings.Title(medicine)
	firstAlpha := medicine[0]
	if firstAlpha == 'A' {
		for i := 0; i < len(medicineParser.A); i++ {
			if medicine == medicineParser.A[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.A[i]), res)
			}
		}
	} else if firstAlpha == 'B' {
		for i := 0; i < len(medicineParser.B); i++ {
			if medicine == medicineParser.B[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.B[i]), res)
			}
		}
	} else if firstAlpha == 'C' {
		for i := 0; i < len(medicineParser.C); i++ {
			if medicine == medicineParser.C[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.C[i]), res)
			}
		}
	} else if firstAlpha == 'D' {
		for i := 0; i < len(medicineParser.D); i++ {
			if medicine == medicineParser.D[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.D[i]), res)
			}
		}
	} else if firstAlpha == 'E' {
		for i := 0; i < len(medicineParser.E); i++ {
			if medicine == medicineParser.E[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.E[i]), res)
			}
		}
	} else if firstAlpha == 'F' {
		for i := 0; i < len(medicineParser.F); i++ {
			if medicine == medicineParser.F[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.F[i]), res)
			}
		}
	} else if firstAlpha == 'G' {
		for i := 0; i < len(medicineParser.G); i++ {
			if medicine == medicineParser.G[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.G[i]), res)
			}
		}
	} else if firstAlpha == 'H' {
		for i := 0; i < len(medicineParser.H); i++ {
			if medicine == medicineParser.H[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.H[i]), res)
			}
		}
	} else if firstAlpha == 'I' {
		for i := 0; i < len(medicineParser.I); i++ {
			if medicine == medicineParser.I[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.I[i]), res)
			}
		}
	} else if firstAlpha == 'J' {
		for i := 0; i < len(medicineParser.J); i++ {
			if medicine == medicineParser.J[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.J[i]), res)
			}
		}
	} else if firstAlpha == 'K' {
		for i := 0; i < len(medicineParser.K); i++ {
			if medicine == medicineParser.K[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.K[i]), res)
			}
		}
	} else if firstAlpha == 'L' {
		for i := 0; i < len(medicineParser.L); i++ {
			if medicine == medicineParser.L[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.L[i]), res)
			}
		}
	} else if firstAlpha == 'M' {
		for i := 0; i < len(medicineParser.M); i++ {
			if medicine == medicineParser.M[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.M[i]), res)
			}
		}
	} else if firstAlpha == 'N' {
		for i := 0; i < len(medicineParser.N); i++ {
			if medicine == medicineParser.N[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.N[i]), res)
			}
		}
	} else if firstAlpha == 'O' {
		for i := 0; i < len(medicineParser.O); i++ {
			if medicine == medicineParser.O[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.O[i]), res)
			}
		}
	} else if firstAlpha == 'P' {
		for i := 0; i < len(medicineParser.P); i++ {
			if medicine == medicineParser.P[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.P[i]), res)
			}
		}
	} else if firstAlpha == 'Q' {
		for i := 0; i < len(medicineParser.Q); i++ {
			if medicine == medicineParser.Q[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.Q[i]), res)
			}
		}
	} else if firstAlpha == 'R' {
		for i := 0; i < len(medicineParser.R); i++ {
			if medicine == medicineParser.R[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.R[i]), res)
			}
		}
	} else if firstAlpha == 'S' {
		for i := 0; i < len(medicineParser.S); i++ {
			if medicine == medicineParser.S[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.S[i]), res)
			}
		}
	} else if firstAlpha == 'T' {
		for i := 0; i < len(medicineParser.T); i++ {
			if medicine == medicineParser.T[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.T[i]), res)
			}
		}
	} else if firstAlpha == 'U' {
		for i := 0; i < len(medicineParser.U); i++ {
			if medicine == medicineParser.U[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.U[i]), res)
			}
		}
	} else if firstAlpha == 'V' {
		for i := 0; i < len(medicineParser.V); i++ {
			if medicine == medicineParser.V[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.V[i]), res)
			}
		}
	} else if firstAlpha == 'W' {
		for i := 0; i < len(medicineParser.W); i++ {
			if medicine == medicineParser.W[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.W[i]), res)
			}
		}
	} else if firstAlpha == 'X' {
		for i := 0; i < len(medicineParser.X); i++ {
			if medicine == medicineParser.X[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.X[i]), res)
			}
		}
	} else if firstAlpha == 'Y' {
		for i := 0; i < len(medicineParser.Y); i++ {
			if medicine == medicineParser.Y[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.Y[i]), res)
			}
		}
	} else if firstAlpha == 'Z' {
		for i := 0; i < len(medicineParser.Z); i++ {
			if medicine == medicineParser.Z[i] {
				speech = handleResponse(1, scrapMedicineLog(&medicineParser.Z[i]), res)
			}
		}
	} else {
		speech = "invalid entry. please try again"
	}
	return
}

func scrapMedicineLog(medicine *string) []string {

	directory, _ := os.Getwd()
	fmt.Println("health-medicine request")
	fmt.Println(" medicine-name -> " + *medicine + " direc -> " + directory)
	*medicine = strings.Replace(*medicine, ",", "_", -1)
	*medicine = strings.Replace(*medicine, "_ ", "_", -1)
	*medicine = strings.Replace(*medicine, " ", "_", -1)
	url := "https://druginfo.nlm.nih.gov/drugportal/name/" + *medicine
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	var result string
	resultFould := false
	contentArray := doc.Find("table.search-results")
	if len(contentArray.Nodes) <= 0 {
		logger.Error(errors.New("Search result table not found."))
	}
	content := contentArray.Eq(0)
	tableRowArray := content.Find("tr")
	for tableRowIndex := range tableRowArray.Nodes {
		tableRow := tableRowArray.Eq(tableRowIndex)
		tableDataLabelArray := tableRow.Find("td.label")
		if len(tableDataLabelArray.Nodes) > 0 {
			tableDataLabel := tableDataLabelArray.Eq(0)
			if strings.ToLower(tableDataLabel.Text()) == "description:" {
				tableData := tableRow.Find("td").Last()
				result = tableData.Text()
				resultFould = true
				break
			}
		}
	}

	if !resultFould {
		logger.Error(errors.New("Cannot find description"))
	}
	return []string{result, *medicine}
}

func processScrapLog(log *string) (data string) {
	logStr := string(*log)
	llogStr := len(logStr)
	subl := "data ->"
	lsubl := len(subl)
	for i := 0; i < llogStr-lsubl; i++ {
		if subl == logStr[i:i+lsubl] {
			data = logStr[i+lsubl : llogStr]
			break
		}
	}
	fmt.Println("scrapped is -> ", data)
	data = strings.Replace(data, "undefined", "", 1)
	return
}

func handleResponse(ctr int, data []string, res http.ResponseWriter) string {
	var resp medicineResponse
	if ctr == 1 {
		resp = medicineResponse{
			Status:  true,
			Message: "Information about the medicine : ",
			Result:  data[0],
		}
	} else {
		resp = medicineResponse{
			Status:  true,
			Message: "Help on the given symptoms : ",
			Result:  data[0],
		}
	}

	send, _ := json.Marshal(resp)
	res.Write(send)
	return "generic " + data[1]
}

func scrapSymptomsLog(sypm *string) []string {
	directory, _ := os.Getwd()
	fmt.Println("health-symptoms request")
	fmt.Println(" medicine-name -> " + *sypm + " direc -> " + directory)
	*sypm = strings.Replace(*sypm, ",", "_", -1)
	*sypm = strings.Replace(*sypm, "_ ", "_", -1)
	*sypm = strings.Replace(*sypm, " ", "_", -1)
	url := "https://www.medindia.net/drugs/medical-condition/" + *sypm
	res, err := scrapper.ScrapeClientRequest(url, nil)
	if err != nil {
		logger.Error(err)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		logger.Error(err)
	}
	result := ""
	articlesArray := doc.Find("article")
	for articleIndex := range articlesArray.Nodes {
		result = result + articlesArray.Eq(articleIndex).Text() + "\n"
	}
	fmt.Println("result is", result)
	return []string{result, *sypm}
}

// HealthSympController controls tasks related to health symptoms
func HealthSympController(symp string, res http.ResponseWriter) (speech string) {

	symp = strings.TrimSpace(symp)
	firstAlpha := symp[0]
	fmt.Println(string(firstAlpha) + " " + symp)
	for i := 0; i < len(symptomParser.symp); i++ {
		if strings.Contains(strings.ToLower(symptomParser.symp[i].Type), strings.ToLower(symp)) {
			fmt.Println("Matched with -> ", symptomParser.symp[i].Type)
			speech = handleResponse(0, scrapSymptomsLog(&symptomParser.symp[i].Link), res)
			break
		}
	}
	return
}
