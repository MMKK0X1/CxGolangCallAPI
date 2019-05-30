// [_Command-line flags_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.

package main

import (
	"cxurls"
	"flag"
	"fmt"
	"log"
	methods "methods/getprojects"
	st "methods/set"
	up "methods/upload"
	"mtypes"
	"strings"
)

func main() {

	fmt.Println("\n")

	destinationserver := flag.String("cxserver", "127.0.0.1", "Checkmarx server destination URL (\"no URI\")")
	fullProxyURL := flag.String("proxy", "", "ProxyServer without Authentication destination URL (\"no URI\")")
	user := flag.String("user", "defaultUser", "Username  for Authentication against Checkmarx server  (\"no URI\")")
	pass := flag.String("pass", "defaultPassword", "Password for Authentication against Checkmarx server  (\"no URI\")")
	oauthtoken := flag.String("oauthtoken", "", "oauth2 token to be used in requests")
	action := flag.String("action", "getprojects", "Type of action to perform:\ngetprojects (list of projects)\\SastScan (scan project)\n get (get scan results)\nGettoken (!!! Sensitive Present Oauth2 Token !!!) ")
	filelocation := flag.String("filelocation", "defaultfilelocation", "Full Path to Local system ZIP file location  (\"no URI\")")
	projectID := flag.String("projectID", "defaultprojectID", "projectID on Checkmarx server under which scan will be executed  (\"no URI\")")
	cxpreset := flag.String("cxpreset", "1", "projectID on Checkmarx server under which scan will be executed  (\"no URI\")")

	flag.Parse()

	//First step to get Oauth2 Token for following commands
	temp := cxurls.Cxmetadata{Action: "getToken"}
	actionObject := temp.GetAction("getToken")

	responseConnectStruct, responseAuthparams := mtypes.SetConnectDetails(*destinationserver, *fullProxyURL, *user, *pass, *oauthtoken, actionObject, *filelocation, *projectID)
	data, statusCode, err := methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
	if err != nil {
		log.Println(err)
	}
	// Check for successful Authentication
	if statusCode != 200 {
		panic("Can't authenticate to Checkmarx servers, check url and credentials")
	}
	temp = cxurls.Cxmetadata{Action: *action}
	actionObject = temp.GetAction(*action)

	//run bodyparcer to extract Authenticaiton token
	var j *mtypes.Jresponse
	*oauthtoken = j.GetTokenfromBody(data).Access_token

	// Use the token for exection of commands agains Checkmarx server

	switch strings.ToLower(*action) {
	case strings.ToLower("getToken"):
		fmt.Println(*oauthtoken)

	case strings.ToLower("SastScan"):
		fmt.Println("Starting action: ", *action)
		fmt.Println("Start command execution of :", *action)
		responseConnectStruct.Action.GetAction(*action)
		//define scan settings
		rp := st.ScanSettings(responseConnectStruct, responseAuthparams, *cxpreset, *oauthtoken)
		if rp.StatusCode != 200 {
			fmt.Println("Could not set scan settings")
		}

		r2, err := up.UploadFile(*filelocation, responseConnectStruct.Proxyserver, *projectID, *oauthtoken)
		if r2.StatusCode != 204 {

			fmt.Println(err)
		} else {
			fmt.Printf("Settings set successfully for ProjectID: %s\n", (*projectID))
		}

		if resbody, ercode := st.SastScan(responseConnectStruct, *oauthtoken); ercode != 201 {
			fmt.Printf("Couldn't start the scan, error =%s", err)
		} else {
			var j *mtypes.CxJresponseScan

			fmt.Printf("Scan started for projectID: %s\n", *projectID)
			fmt.Printf("Check for results: %s/cxrestapi%s\n", *destinationserver, j.ParcecxResponseNonByte(resbody).Link.Uri)
		}

		data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
		if err != nil {
			log.Println(err)
		}
		// Check for successful Authentication
		if statusCode != 200 {
			panic("Can't authenticate to Checkmarx servers, check url and credentials")
		} else {

		}

	case strings.ToLower("getprojects"):
		fmt.Println("Starting action: ", *action)
		fmt.Println("Start command execution of :", *action)
		responseConnectStruct, responseAuthparams = mtypes.SetConnectDetails(*destinationserver, *fullProxyURL, *user, *pass, *oauthtoken, temp.GetAction(*action), *filelocation, *projectID)
		data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
		if err != nil {
			log.Println(err)
		}
		// Check for successful Authentication
		if statusCode != 200 {
			panic("Can't authenticate to Checkmarx servers, check url and credentials")
		} else {
			//Project id parcer
			fmt.Print("List of availabe applications:\n")
			var cxj *mtypes.CxJresponseScan
			tmp := cxj.ParcecxResponse(data)
			for i := range tmp {
				fmt.Printf("Application: %v, ProjectID: %v \n", cxj.ParcecxResponse(data)[i].Name, cxj.ParcecxResponse(data)[i].ID)
			}
			*action = "GetPreset"
			responseConnectStruct, responseAuthparams = mtypes.SetConnectDetails(*destinationserver, *fullProxyURL, *user, *pass, *oauthtoken, temp.GetAction(*action), *filelocation, *projectID)
			data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
			if err != nil {
				log.Println(err)
			}
			fmt.Print("Awailable presets: ")
			tmp2 := cxj.ParcecxResponse(data)
			for i := range tmp2 {
				fmt.Printf(" %s:%d\n", cxj.ParcecxResponse(data)[i].Name, cxj.ParcecxResponse(data)[i].ID)
			}

		}

	case strings.ToLower("UploadFile"):

	default:
		return

	}
	fmt.Println("\n")

}
