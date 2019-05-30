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
	writer.WriteField("ID", id)
	part, err := writer.CreateFormFile("zippedSource", filepath.Base(path))
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
