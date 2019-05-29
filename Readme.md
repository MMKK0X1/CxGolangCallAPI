
# Project: Secure Static code analysis (Checkmarx) scanning api invocation
During my engagement I have been looking for solution to run secure static code analysis integrated as part of development processes. While there are a few plugins such as "Checkmarx CxSAST Plugin" this requires Jenkins, which is not the case. 

Building another plugin for addition automation tool only provides partial solution and lacks versatility. 

Next I have looked at  CxConsole (CxSAST CLI). This tool requires JVM to be installed on the machine where the scan is invoked, and since most of the project are not Java based, it was feciable solution

Besides, most of the machines are short lived, docker builders, so the solution should be portable and as small as possible, especially when docker scratch was used as base. 

:boom: 
Therefore, after taking into an account the above, I went home and end-up to build a client. 

**Cross-platform**  
The client is build on golang and can be compiled . - :+1:    				

**Fast and small**

 Run natively   - :+1:

**Fully portable**  
Can be used by developers, automation tools, orchestration tools, etc - :+1:

At this stage command line client support only a few "action" but can be extended to whole Checkmarx API suit. 

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites
Basic understanding about Secure Static code analysis process
Golang compiler and package manager

### Installing

Download and Compile the project:

download the project

For example Windows compilation command 
```
  % GOOS=windows GOARCH=amd64 go build -o scan.exe
```

The following paramets are accepted:

*cxserver:
   Checkmarx address is expected in URL for only no URI's
*proxy :
  Proxy server url ( for example "http://localhost") - Authentication is not supported at this stage, feel free to request if needed

*action: 
  -getprojects (default) - lists all projects the users has access to.
  -Gettoken   - print Oauth2 checkmarx token. Can be used for following testings, for example getting reports. 
  -SastScan - perfoms checkmarx scan 
      -projectid (required)
      -preset   (required)
      -filelocation (required) fullpath to zip file to be uploaded  
*user
  -username
*pass
  -password
*filelocation
  -full path including filename to zip file to be scanned
*projectID
  -project relavant to the scan. This information can be found in "getprojects"
*cxpreset
  -scanning preset (rules) to be used for this scan. Full list is shown as part of getprojects

Command example:

```
scan.exe -cxserver="https://checkmarxserveraddress.net" -proxy="localhost:8090" -action="SastScan" -user="user" -pass="password" filelocation="D:\tmp\main.zip" projectID="24" -cxpreset="1"
```

