package controllers

import (
	"io"
	"log"

	"github.com/devetek/caddy-dataprovider/utils"
)

// get link API
func getLinkProductDetailAPI(target string) string {
	apiMap := make(map[string]string)

	apiMap["a"] = "https://mocki.io/v1/145e3cc7-bc70-450f-b05f-6c46877492a9"
	apiMap["b"] = "https://mocki.io/v1/f0509897-dda7-4dad-b6d4-519658f9ed8b"
	apiMap["c"] = "https://mocki.io/v1/b3777049-c4c3-4ffc-9eed-c9d2d92da465"
	apiMap["default"] = "http://localhost:8080/api/product/default"

	selectedData := apiMap[target]

	if selectedData == "" {
		return apiMap["default"]
	}

	return selectedData
}

func GetDataPDP(productName string) string {
	apiResp, err := utils.Fetcher(getLinkProductDetailAPI(productName), "")
	if err != nil {
		log.Println(err)
		return ""
	}

	bodyBytes, err := io.ReadAll(apiResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	defer apiResp.Body.Close()

	return bodyString
}
