package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func getIPv4() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api4.tschwaak.me", nil)
	if err != nil {
		return "", errors.New("encountered an Error while creating request")
	}
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", errors.New("encountered an Error while sending request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("encountered an Error while reading response")
	}
	s := string(body)
	resp.Body.Close()
	resp.Close = true
	return s, nil
}

func getIPv6() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api6.tschwaak.me", nil)
	if err != nil {
		return "", errors.New("Encounterd an Error while creating request")
	}
	req.Close = true
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", errors.New("Encounterd an Error while sending request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("encountered an Error while reading response")
	}
	s := string(body)
	resp.Body.Close()
	resp.Close = true
	return s, nil
}
