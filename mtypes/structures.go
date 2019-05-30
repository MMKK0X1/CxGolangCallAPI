package mtypes

import (
	"cxurls"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type Connection struct {
	ConnectStruct ConnectStruct
	Authparams    Authparams
}

type Jresponse struct {
	Access_token string `json:"access_token"` // Uppercased first letter
	Expires_in   int    `json:"expires_in"`   // Uppercased first letter
	Token_type   string `json: "token_type"`
}

type CxJresponseScan struct {
	ID       int
	TeamId   string
	Name     string
	IsPublic bool
	//SourceSettingsLink string
	Link CxJLink
}

type CxJscansResponce struct {
	Id   string
	Link CxJLink
}
type CxJLink struct {
	Type string
	Rel  string
	Uri  string
}

type Authparams struct {
	Username      string
	Password      string
	Grant_type    string
	Scope         string // An unexported field is not encoded.
	Client_id     string
	Client_secret string
}

type ConnectStruct struct {
	Destinationserver string
	Proxyserver       string
	Oauthtoken        string
	Action            cxurls.Cxmetadata
	Filelocation      string
	ProjectID         string
	// Url               cxurls.Cxmetadata
}

func (c *ConnectStruct) Addprojectid(id string) *ConnectStruct {
	c.ProjectID = id
	fmt.Println("this is ConnectStruct struct", c)
	return c
}
func StringBuilder(cs ConnectStruct, au Authparams) string {
	if strings.ToUpper(cs.Action.Method) == "POST" {
		s := ("username=" + url.QueryEscape(au.Username) + "&" + "password=" + url.QueryEscape(au.Password) + "&" + "grant_type=" + au.Grant_type + "&" + "scope=" + au.Scope + "&" + "client_id=" + au.Client_id + "&" + "client_secret=" + au.Client_secret)
		return s
	} else if strings.ToUpper(cs.Action.Method) == "GET" {
		return cs.Oauthtoken
	} else {
		return "Unknown execution method"
	}
	return "Error"

	// return ("username=" + url.QueryEscape(authdata.Username) + "&" + "password=" + url.QueryEscape(authdata.Password) + "&" + "grant_type=" + authdata.Grant_type + "&" + "scope=" + authdata.Scope + "&" + "client_id=" + authdata.Client_id + "&" + "client_secret=" + authdata.Client_secret)
}

func (j *Jresponse) GetTokenfromBody(data []byte) *Jresponse {
	err := json.Unmarshal(data, &j)
	if err != nil {
		log.Println(err)
	}
	return j

}

func (cxj *CxJresponseScan) ParcecxResponse(data []byte) []CxJresponseScan {
	var arr []CxJresponseScan
	err := json.Unmarshal([]byte(data), &arr)
	if err != nil {
		log.Println(err)
	}
	return arr
}
func (cxj *CxJresponseScan) ParcecxResponseNonByte(data []byte) *CxJresponseScan {

	err := json.Unmarshal([]byte(data), &cxj)
	if err != nil {
		log.Println(err)
	}
	return cxj
}

func SetConnectDetails(ds string, ps string, us string, pass string, oa string, ac cxurls.Cxmetadata, filel string, prjid string) (ConnectStruct, Authparams) {

	sc := ConnectStruct{
		Destinationserver: ds,
		Proxyserver:       ps,
		Oauthtoken:        oa,
		Action:            ac,
		Filelocation:      filel,
		ProjectID:         prjid,
	}

	au := Authparams{
		Username:      us,
		Password:      pass,
		Grant_type:    "password",
		Scope:         "sast_rest_api",
		Client_id:     "resource_owner_client",
		Client_secret: "014DF517-39D1-4453-B7B3-9930C563627C",
	}

	return sc, au

}

//GetConnectDetails
func GetConnectDetails(sc ConnectStruct, au Authparams) (ConnectStruct, Authparams) {

	return sc, au
}
