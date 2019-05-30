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

	fmt.Println("Start :")
	// st := new(Heelo)
	// fmt.Println(st)

	// proxyserver := "http://localhost:8090"
	// destinationserver := "https://latitudefinancials.checkmarx.net"
	// getTokenURI := "cxrestapi/auth/identity/connect/token"
	// submitTokenUri := "/cxrestapi/projects"

	// username := "testuser"
	// password := "Password1!"
	// token2 := "hello toke"
	// proxyserver = ""

	// fmt.Println("hwerwe is is :",de.A)
	// Getcxlinks("asdf")
	destinationserver := flag.String("cxserver", "127.0.0.1", "Checkmarx server destination URL (\"no URI\")")
	fullProxyURL := flag.String("proxy", "", "ProxyServer without Authentication destination URL (\"no URI\")")
	user := flag.String("user", "defaultUser", "Username  for Authentication against Checkmarx server  (\"no URI\")")
	pass := flag.String("pass", "defaultPassword", "Password for Authentication against Checkmarx server  (\"no URI\")")
	oauthtoken := flag.String("oauthtoken", "", "oauth2 token to be used in requests")
	// action := flag.String("action", "defaultAction", "Type of action perfomed on Checkmarx server (start scan/ get results)  (\"no URI\")")
	//action Start Scan= scan
	action := flag.String("action", "getprojects", "Type of action to perform:\ngetprojects (list of projects)\\SastScan (scan project)\n get (get scan results)\nGettoken (!!! Sensitive Present Oauth2 Token !!!) ")

	filelocation := flag.String("filelocation", "defaultfilelocation", "Full Path to Local system ZIP file location  (\"no URI\")")
	projectID := flag.String("projectID", "defaultprojectID", "projectID on Checkmarx server under which scan will be executed  (\"no URI\")")
	cxpreset := flag.String("cxpreset", "1", "projectID on Checkmarx server under which scan will be executed  (\"no URI\")")

	_ = cxpreset
	flag.Parse()
	//	var svar string
	//	flag.StringVar(&svar, "svar", "bar", "a string var")

	// *destinationserver = "https://destination.checkmarx.net"
	fmt.Println("*destinationserver ", *destinationserver)
	fmt.Println("fullProxyURL:", *fullProxyURL)
	fmt.Println("user:", *user)
	fmt.Println("pass:", *pass)
	fmt.Println("oauthtoken:", *oauthtoken)
	fmt.Println("action:", *action)
	fmt.Println("filelocation:", *filelocation)
	fmt.Println("projectID:", *projectID)
	fmt.Println("cxpreset:", *cxpreset)

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
	// *oauthtoken = j.GetTokenfromBody(data)

	// Use the token for exection of commands agains Checkmarx server

	switch strings.ToLower(*action) {
	case strings.ToLower("getToken"):
		fmt.Println(*oauthtoken)

	case strings.ToLower("SastScan"):
		fmt.Println("Chosen action: ", *action)
		fmt.Println("Start command execution of :", *action)
		fmt.Println(responseConnectStruct.Action.GetAction(*action))
		//define scan settings
		rp := st.ScanSettings(responseConnectStruct, responseAuthparams, *cxpreset, *oauthtoken)
		if rp.StatusCode != 200 {
			fmt.Println("Could not set scan settings")
		}

		r2, err := up.UploadFile(*filelocation, responseConnectStruct.Proxyserver, *projectID, *oauthtoken)
		if r2.StatusCode != 204 {

			fmt.Println(err)
		}
		fmt.Printf("Settings were set: Return status good is 204=: %s, triggering scan of Project: %s", r2.Status, *projectID)
		fmt.Print((st.SastScan(responseConnectStruct, *oauthtoken)))

		fmt.Println("Scan started")

		data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
		if err != nil {
			log.Println(err)
		}
		// Check for successful Authentication
		if statusCode != 200 {
			panic("Can't authenticate to Checkmarx servers, check url and credentials")
		} else {

			// Outbody
			// fmt.Println("Execution output: \n", string(data))

			// 	Future Placeholder for project id parcer
			// 	var cxj *mtypes.CxJresponse
			// tmp := cxj.Parsecxresponse(data)

		}

	case strings.ToLower("getprojects"):
		fmt.Println("Chosen action: ", *action)
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
			tmp := cxj.ParsecxResponse(data)
			for i := range tmp {
				fmt.Printf("Application: %v, ProjectID: %v \n", cxj.ParsecxResponse(data)[i].Name, cxj.ParsecxResponse(data)[i].ID)
			}
			*action = "GetPreset"
			responseConnectStruct, responseAuthparams = mtypes.SetConnectDetails(*destinationserver, *fullProxyURL, *user, *pass, *oauthtoken, temp.GetAction(*action), *filelocation, *projectID)
			data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
			if err != nil {
				log.Println(err)
			}
			fmt.Print("Awailable presets: ")
			tmp2 := cxj.ParsecxResponse(data)
			for i := range tmp2 {
				fmt.Printf(" %s:%d\n", cxj.ParsecxResponse(data)[i].Name, cxj.ParsecxResponse(data)[i].ID)
			}

		}

	case strings.ToLower("UploadFile"):
		// *projectID = "24"
		// *filelocation = "d:/tmp/main.zip"

		// urlForUpload := "https://latitudefinancials.checkmarx.net/cxrestapi/projects/" + *projectID + "/sourceCode/attachments"

		// request, err := up.UploadFile(urlForUpload, *filelocation, *fullProxyURL, *projectID)
		// request.Header.Add("Authorization", "Bearer "+*oauthtoken)

		// if err != nil {
		// 	log.Fatal(err)
		// }
		// //Create Transport based on proxy configuration
		// transport := cr.CreateTransport(*fullProxyURL)
		// client := &http.Client{
		// 	Transport: transport,
		// }
		// resp, err := client.Do(request)
		// if err != nil {
		// 	log.Fatal(err)
		// } else {
		// 	readbody := &bytes.Buffer{}
		// 	_, err := readbody.ReadFrom(resp.Body)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	resp.Body.Close()
		// }
		// if resp.StatusCode != 204 {
		// 	fmt.Println(resp.Header)
		// 	fmt.Println("File wasn't uploaded")
		// } else {
		// 	fmt.Println("Triggering scan")
		// 	*action = "SastScan"
		// 	responseConnectStruct, responseAuthparams = mtypes.SetConnectDetails(*destinationserver, *fullProxyURL, *user, *pass, *oauthtoken, temp.GetAction(*action), *filelocation, *projectID)
		// 	data, statusCode, err = methods.Buildandsend(responseConnectStruct, responseAuthparams, mtypes.StringBuilder(responseConnectStruct, responseAuthparams))
		// }

	default:
		return

	}

	// fmt.Println("tail:", flag.Args())
}
