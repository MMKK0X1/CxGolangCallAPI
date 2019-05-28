package methods

import (
	"fmt"
	"io/ioutil"
	"log"
	"mtypes"
	"strings"

	cr "methods"
	"net/http"
	"net/url"
)

//ScanSettings
func SastScan(sc mtypes.ConnectStruct, bodyortoken string) (rpx string) {

	//Create Transport based on proxy configuration
	transport := cr.CreateTransport(sc.Proxyserver)

	//Parse Url
	urlStr := sc.Destinationserver + "/cxrestapi/sast/scans"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}
	// Create Client
	client := &http.Client{
		Transport: transport,
	}

	//generating the HTTP POST request

	payload := "ProjectId=" + sc.ProjectID + "&isIncremental=" + "false" + "&isPublic=" + "false" + "&forceScan=true" + "&comment=comment for scan"

	// payload := strings.NewReader("ProjectId=24&PresetId=1&engineConfigurationId=1")

	req, _ := http.NewRequest("POST", url.String(), strings.NewReader(payload))

	req.Header.Add("Authorization", "Bearer "+bodyortoken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("cxOrigin", "automated scan test project")
	req.Header.Add("Accept", "application/json;v=1.0")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
	rpx = string(body)
	return rpx
}
