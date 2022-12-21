package ado

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func ExecutePreviewRequest(orgInfo OrganizationInfo, method string, api string, params string, body io.Reader) *http.Response {
	b64Personalaccesstoken := base64.StdEncoding.EncodeToString([]byte("ado_user:" + orgInfo.PersonalAccessToken))

	requestUri := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/%s?api-version=7.0-preview.2", orgInfo.Organization, orgInfo.Project, api)

	if params != "" {
		requestUri = requestUri + "&" + params
	}

	fmt.Printf("%s [%s] %s", time.Now().Local(), method, requestUri)

	req, err := http.NewRequest(method, requestUri, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+b64Personalaccesstoken)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" - %d\n\r", resp.StatusCode)

	if resp.StatusCode >= 300 {
		log.Fatalf("Request Failed\n\r")
	}

	return resp
}

func ExecuteRequest(orgInfo OrganizationInfo, method string, api string, params string, body io.Reader) *http.Response {
	b64Personalaccesstoken := base64.StdEncoding.EncodeToString([]byte("ado_user:" + orgInfo.PersonalAccessToken))

	requestUri := fmt.Sprintf("https://dev.azure.com/%s/%s/_apis/%s?api-version=7.0", orgInfo.Organization, orgInfo.Project, api)

	if params != "" {
		requestUri = requestUri + "&" + params
	}

	fmt.Printf("%s [%s] %s", time.Now().Local(), method, requestUri)

	req, err := http.NewRequest(method, requestUri, body)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Basic "+b64Personalaccesstoken)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" - %d\n\r", resp.StatusCode)

	if resp.StatusCode >= 300 {
		log.Fatalf("Request Failed\n\r")
	}

	return resp
}
