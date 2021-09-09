package main

import (
	// "fmt"
	"crypto/tls"
	"encoding/json"
	// "encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

/**
*to load the env variable
 */

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurs while loading .env")
	}
	return os.Getenv((key))
}

type Template struct {
	Template_id         string
	Template_name       string
	Template_create     string
	Template_update     string
	Template_html       string
	Template_text       string
	Is_removable        string
	Template_screenshot string
}

type Response struct {
	Result []Template
}

func main() {

	baseUrl := envVariable("BASE_URL")
	token := envVariable(("AUTHENTICATION_TOKEN"))

	templateUrl := baseUrl + "template/"
	body, err := doGetRequest(templateUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
	}
	var response Response
	json.Unmarshal(body, &response)
	result := response.Result

	var required_template Template

	for _, value := range result {

		if value.Template_id == "17" {
			// Found!
			required_template = value

		}
	}
	template_html:=required_template.Template_html
	log.Println(template_html)
	
	contactUrl := baseUrl + "contact/?limit=11"
	body, err = doGetRequest(contactUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
	}

}

func doGetRequest(url string, token string) ([]byte, error) {
	//request to given url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	//add headers to request

	req.Header.Add("Authorization", token)
	req.Header.Add("Content_type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return body, err
}

