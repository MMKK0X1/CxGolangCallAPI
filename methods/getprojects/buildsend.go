package methods

import (
	"fmt"
	"io/ioutil"
	"log"
	"mtypes"

	cr "methods"
	"net/http"
	"net/url"
	"strings"
)

//Buildandsend ,,,
func Buildandsend(sc mtypes.ConnectStruct, au mtypes.Authparams, bodyortoken string) ([]byte, int, error) {

	//Create Transport based on proxy configuration
	transport := cr.CreateTransport(sc.Proxyserver)

	//Parse Url
	urlStr := sc.Destinationserver + "/" + sc.Action.Url
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}
	//Create Client
	client := &http.Client{
		Transport: transport,
	}

	//generating the HTTP POST request

	var request *http.Request

	if strings.ToUpper(sc.Action.Method) == "POST" {
		//POST must have body
		request, err = http.NewRequest(sc.Action.Method, url.String(), strings.NewReader(bodyortoken))
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Method Sent", sc.Action.Method)

	} else if strings.ToUpper(sc.Action.Method) == "GET" {

		//GET not body, but need hearders
		request, err = http.NewRequest(sc.Action.Method, url.String(), nil)
		request.Header.Set("Authorization", "Bearer "+bodyortoken)
		request.Header.Set("Accept", "application/json;v=1.0")
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Method Sent", sc.Action.Method)

	} else {
		fmt.Println("No Method defined", sc.Action.Method)

	}

	//calling the URL
	//getting the response
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	statusCode := response.StatusCode
	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Println(err)
	}
	return data, statusCode, err

}
