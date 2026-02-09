package httplib

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"lmsapieng/include/common/globaldef"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var httpClient *http.Client
var httpClientOnce sync.Once
var tlsHttpClient *http.Client
var tlsHttpClientOnce sync.Once

func getHttpClient(TimeOut int) *http.Client {
	httpClientOnce.Do(func() {
		//trace.Lg("Allocating httpClient for the first time")
		httpClient = &http.Client{Timeout: time.Second * time.Duration(TimeOut)}
	})
	return httpClient
}

func getTlsHttpClient(serverName, serverCrt string, TimeOut int) (*http.Client, error) {
	var err error
	tlsHttpClientOnce.Do(func() {
		// trace.Lg( "Allocating tlsHttpClient for the first time")
		serverCrtFile := fmt.Sprintf("%s/%s/%s", globaldef.GetAppBaseDir(), "config/tls/certificates/client", serverCrt)
		var cert []byte
		cert, err = os.ReadFile(serverCrtFile)
		if err != nil {
			//trace.Lg("io.ReadFile() failed for serverCrtFile:%s with err:%s", serverCrtFile, err)
			return
		}
		certPool := x509.NewCertPool()
		certPool.AppendCertsFromPEM(cert)
		config := &tls.Config{
			RootCAs:    certPool,
			ServerName: serverName,
			MinVersion: tls.VersionTLS13,
		}
		tlsHttpClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: config,
			},
			Timeout: time.Second * time.Duration(TimeOut),
		}
	})
	return tlsHttpClient, err
}

func SendGETRequest(URL string, TimeOut int, respData *[]byte, StatusCode *int, headerList ...string) int {
	//trace.Lg("SendGETRequest() called for URL(%s) TimeOut(%d)", URL, TimeOut)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		//trace.Lg("http.NewRequest() failed with err(%s)", err)
		return -1
	}
	for _, header := range headerList {
		headerStr := strings.Split(header, ":")
		req.Header.Set(headerStr[0], headerStr[1])
	}
	client := getHttpClient(TimeOut)
	resp, err := client.Do(req)
	if err != nil {
		//trace.Lg("client.Do() failed with err(%s)", err)
		return -1
	}
	defer resp.Body.Close()
	*StatusCode = resp.StatusCode
	*respData, err = io.ReadAll(resp.Body)
	if err != nil {
		//trace.Lg("io.ReadAll() failed with err(%s)", err)
		return -1
	}
	//trace.Lg("sendGETRequest() Success with respData(%s)", *respData)
	return 1
}

func SendPOSTRequest(reqData []byte, URL string, ContentType string, TimeOut int, respData *[]byte, StatusCode *int, headerList ...string) int {
	//trace.Lg("SendPOSTRequest() called for URL(%s) ContentType(%s) TimeOut(%d)", URL, ContentType, TimeOut)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqData))
	if err != nil {
		//trace.Lg("http.NewRequest() failed with err(%s)", err)
		return -1
	}
	req.Header.Set("Content-Type", ContentType)
	for _, header := range headerList {
		headerStr := strings.Split(header, ":")
		req.Header.Set(headerStr[0], headerStr[1])
	}
	client := getHttpClient(TimeOut)
	resp, err := client.Do(req)
	if err != nil {
		//trace.Lg("httpClient.Do() failed with err(%s)", err)
		if os.IsTimeout(err) {
			return -2
		}
		return -1
	}
	defer resp.Body.Close()
	*StatusCode = resp.StatusCode
	*respData, err = io.ReadAll(resp.Body)
	if err != nil {
		//trace.Lg("io.ReadAll() failed with err(%s)", err)
		return -1
	}
	//trace.Lg("sendPOSTRequest() Success with respData(%s)", *respData)
	return 1
}

func SendPOSTTLSRequest(serverName string, serverCrt string, reqData []byte, URL string, ContentType string, TimeOut int, respData *[]byte, StatusCode *int, headerList ...string) int {
	// trace.Lg("SendPOSTTLSRequest() called for serverName:%s serverCrt:%s URL:%s", serverName, serverCrt, URL)
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqData))
	if err != nil {
		//trace.Lg("http.NewRequest() failed with err(%s)", err)
		return -1
	}
	req.Header.Set("Content-Type", ContentType)
	for _, header := range headerList {
		headerStr := strings.Split(header, ":")
		req.Header.Set(headerStr[0], headerStr[1])
	}
	client, err := getTlsHttpClient(serverName, serverCrt, TimeOut)
	if err != nil {
		//trace.Lg("getTlsHttpClient() failed with err(%s)", err)
		return -1
	}
	resp, err := client.Do(req)
	if err != nil {
		//trace.Lg("httpClient.Do() failed with err(%s)", err)
		if os.IsTimeout(err) {
			return -2
		}
		return -1
	}
	defer resp.Body.Close()
	*StatusCode = resp.StatusCode
	*respData, err = io.ReadAll(resp.Body)
	if err != nil {
		//trace.Lg("io.ReadAll() failed with err(%s)", err)
		return -1
	}
	return 1
}
