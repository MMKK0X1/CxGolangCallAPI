package cxurls

import (
	"strings"
)

type StartScan struct {
	CxOrigin              string
	ProjectId             string
	PresetId              string
	EngineConfigurationId string
	Id                    string
	IsIncremental         bool
	IsPublic              bool
	ForceScan             bool
	Comment               string
}

type Cxmetadata struct {
	Action    string
	Method    string
	Url       string
	ProjectID string
	StartScan StartScan
}

//Getlink ,,

var Gettoken = Cxmetadata{
	Action:    "getToken",
	Method:    "POST",
	Url:       "cxrestapi/auth/identity/connect/token",
	ProjectID: "",
}
var GetProjects = Cxmetadata{
	Action:    "getProjects",
	Method:    "GET",
	Url:       "cxrestapi/projects",
	ProjectID: "",
}
var UploadFile = Cxmetadata{
	Action:    "UploadFile",
	Method:    "POST",
	Url:       "cxrestapi/projects/{id}/sourceCode/attachments",
	ProjectID: "",
}
var ScanSettings = Cxmetadata{
	Action:    "ScanSettings",
	Method:    "POST",
	Url:       "cxrestapi/sast/scanSettings",
	ProjectID: "",
}
var SastScan = Cxmetadata{
	Action:    "SastScan",
	Method:    "POST",
	Url:       "cxrestapi/sast/scans",
	ProjectID: "",
}
var ScanStat = Cxmetadata{
	Action:    "ScanStat",
	Method:    "GET",
	Url:       "cxrestapi/sast/scans/[ID]",
	ProjectID: "",
}
var GetPreset = Cxmetadata{
	Action:    "GetPreset",
	Method:    "GET",
	Url:       "cxrestapi/sast/presets",
	ProjectID: "",
}

func (c *Cxmetadata) Raction() *string {
	return &c.Action
}

func (c *Cxmetadata) GetAction(s string) Cxmetadata {

	var Getlink = map[string]Cxmetadata{
		"GetProjects":  GetProjects,
		"Gettoken":     Gettoken,
		"UploadFile":   UploadFile,
		"ScanSettings": ScanSettings,
		"SastScan":     SastScan,
		"GetPreset":    GetPreset,
		"ScanStat":     ScanStat,
	}
	for k, v := range Getlink {

		if (strings.ToLower(s)) == (strings.ToLower(k)) {
			return v
		}
	}
	scanval := StartScan{}

	x := Cxmetadata{"Action Not found", "Action Not found", "Action not  found", "Action not found", scanval}
	return x

}
