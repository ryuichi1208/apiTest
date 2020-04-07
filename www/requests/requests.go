package main

import (
	"fmt"
	"log"
	"net/http"
)

type programInfo struct {
	programName string
	version     string
	usage       string
}

type httpRequestInfo struct {
	program     *programInfo
	httpVersion string
	httpScheme  string
}

func (pi *programInfo) initProgramInfo() {
	pi.programName = "requests"
	pi.version = "1.0.0"
	pi.usage = fmt.Sprintf("Version : %s\nUsage : %s $1 $2", pi.version, pi.programName)
}

func (hri *httpRequestInfo) initHttpRequestInfo(pi *programInfo) {
	hri.program = pi
	hri.httpScheme = "https"
	hri.httpVersion = "1.1"
}

func generateUrl(domain string, hri *httpRequestInfo) (url string) {
	url = fmt.Sprintf("%s://%s", hri.httpScheme, domain)
	return
}

func main() {
	domainList := []string{
		"google.com",
		"yahoo.co.jp",
	}

	// new(programInfo).initProgramInfo()
	hri := new(httpRequestInfo)
	hri.initHttpRequestInfo(new(programInfo).initProgramInfo())

	var url string
	for _, domain := range domainList {
		url = generateUrl(domain, hri)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Status : %d\n", resp.StatusCode)
	}
}
