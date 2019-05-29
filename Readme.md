
# Usage example and Details <h1>


1. default logs writing location. 
  tmp\\CxAPIDataLog.json

cxserver- URL only
proxy:
action: 
    
user- username
pass- password
filelocation - full path including filename in zip format
    "d:\tmp\source.zip"
projectID
    scanner project ID, can be fetch from running getprojects action command
cxpreset
    scanning preset to be used for this scan. Full list is shown as part of getprojects

Example:
1. Compile the project:
  % GOOS=windows GOARCH=amd64 go build -o scan.exe

2. Trigget the scan by providing nessessary parameters 
scan.exe -cxserver="https://latitudefinancials.checkmarx.net" -proxy="localhost:8090" -action="SastScan" -user="testuser" -pass="Password1!" filelocation="D:\tmp\main.zip" projectID="24" -cxpreset="1"

