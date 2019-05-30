package methods

import (
	"bytes"
	"fmt"
	"io"
	"log"
	cr "methods"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(path, proxyserver, id, oauthtoken string) (*http.Response, error) {
	urlForUpload := "https://latitudefinancials.checkmarx.net/cxrestapi/projects/" + id + "/sourceCode/attachments"
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	readbody := &bytes.Buffer{}

	writer := multipart.NewWriter(readbody)
	// part, err := writer.CreateFormFile("ID", id)
	writer.WriteField("ID", id)

	part, err := writer.CreateFormFile("zippedSource", filepath.Base(path))

	// part, err := writer.CreateFormFile("zippedSource", filepath.Base(path))
	// writer.WriteField("zippedSource", filepath.Base(path))

	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	transport := cr.CreateTransport(proxyserver)
	client := &http.Client{
		Transport: transport,
	}

	req, err := http.NewRequest("POST", urlForUpload, readbody)
	req.Header.Add("Authorization", "Bearer "+oauthtoken)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	} else {
		readbody := &bytes.Buffer{}
		_, err := readbody.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
	}
	if resp.StatusCode != 204 {
		fmt.Println(resp.Header)
		fmt.Println("Error in file upload")
	} else {

		return resp, err
	}
	return resp, err
}

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
