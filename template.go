package main

import (
	"encoding/json"
	"log"
)
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
type TemplateResponse struct {
	Result []Template
}
func getTemplate(baseUrl string,token string) (Template,error) {

	templateUrl := baseUrl + "template/"
	body, err := doGetRequest(templateUrl, token)
	//request to given url

	if err != nil {
		log.Fatalf(err.Error())
		return Template{},err
	}
	var templateResponse TemplateResponse
	json.Unmarshal(body, &templateResponse)
	result := templateResponse.Result

	var required_template Template

	for _, value := range result {

		if value.Template_id == "17" {
			// Found!
			required_template = value
			break
		}
	}
	return required_template,nil
}