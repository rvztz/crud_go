package service

import (
	"fmt"
	"net/http"
	"os"
)

var API_KEY string = os.Getenv("api_inegi")
var denueUrl string = "https://www.inegi.org.mx/app/api/denue/v1/consulta/buscar/%s/%f,%f/250/%s"

func GetDenueData(cond string, lat float32, lon float32, rc chan *http.Response) {
	fmtUrl := fmt.Sprintf(denueUrl, cond, lat, lon, API_KEY)
	res, err := http.Get(fmtUrl)
	if err != nil {
		fmt.Println("something wrong.")
	}

	rc <- res
}
