package tools

import (
	"github.com/EDDYCJY/fake-useragent"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestBuild struct {
	Timeout    time.Duration
	HttpProxy  string // "http://127.0.0.1:1081"
	Parameter  url.Values
	RequestUrl string
	Headers    map[string]string
	Cookies    []http.Cookie
	Ua         string
	RandomUa   bool
	SsProxy    string // "127.0.0.1:1080"
}

type Response struct {
	Headers map[string][]string
	Status  int
	Body    string
}

func GET(requestBuild RequestBuild) (Response, error) {

	if requestBuild.RandomUa {
		requestBuild.Ua = browser.Random()
	}

	if requestBuild.Timeout == 0 {
		requestBuild.Timeout = 5
	}

	timeout := requestBuild.Timeout * time.Second
	httpClient := &http.Client{
		Timeout: timeout,
	}

	if requestBuild.HttpProxy != "" {
		urlproxy, err := url.Parse(requestBuild.HttpProxy)
		if err == nil {
			tr := &http.Transport{Proxy: http.ProxyURL(urlproxy)}
			httpClient = &http.Client{
				Timeout:   timeout,
				Transport: tr,
			}
		}
	}

	if requestBuild.SsProxy != "" {
		dialSocksProxy, err := proxy.SOCKS5("tcp", requestBuild.SsProxy, nil, proxy.Direct)
		if err == nil {
			tr := &http.Transport{Dial: dialSocksProxy.Dial}
			httpClient = &http.Client{
				Timeout:   timeout,
				Transport: tr,
			}
		}
	}

	request, _ := http.NewRequest("GET", requestBuild.RequestUrl+"?"+requestBuild.Parameter.Encode(), nil)

	request.Header.Set("User-Agent", requestBuild.Ua)

	for k, v := range requestBuild.Headers {
		request.Header.Set(k, v)
	}

	for _, v := range requestBuild.Cookies {
		request.AddCookie(&v)
	}

	resp, err := httpClient.Do(request)

	if resp == nil {
		return Response{
			Status: 0,
		}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	response := Response{
		Body:    string(body),
		Status:  resp.StatusCode,
		Headers: resp.Header,
	}

	return response, err
}

func POST(requestBuild RequestBuild) (Response, error) {

	if requestBuild.RandomUa {
		requestBuild.Ua = browser.Random()
	}

	if requestBuild.Timeout == 0 {
		requestBuild.Timeout = 5
	}

	timeout := requestBuild.Timeout * time.Second
	httpClient := &http.Client{
		Timeout: timeout,
	}

	if requestBuild.HttpProxy != "" {
		urlproxy, err := url.Parse(requestBuild.HttpProxy)
		if err == nil {
			tr := &http.Transport{Proxy: http.ProxyURL(urlproxy)}
			httpClient = &http.Client{
				Timeout:   timeout,
				Transport: tr,
			}
		}
	}

	if requestBuild.SsProxy != "" {
		dialSocksProxy, err := proxy.SOCKS5("tcp", requestBuild.SsProxy, nil, proxy.Direct)
		if err == nil {
			tr := &http.Transport{Dial: dialSocksProxy.Dial}
			httpClient = &http.Client{
				Timeout:   timeout,
				Transport: tr,
			}
		}
	}

	request, _ := http.NewRequest("POST", requestBuild.RequestUrl, strings.NewReader(requestBuild.Parameter.Encode()))

	request.Header.Set("User-Agent", requestBuild.Ua)
	request.Header.Set("Content-Type","application/x-www-form-urlencoded")
	for k, v := range requestBuild.Headers {
		request.Header.Set(k, v)
	}

	for _, v := range requestBuild.Cookies {
		request.AddCookie(&v)
	}

	resp, err := httpClient.Do(request)

	if resp == nil {
		return Response{
			Status: 0,
		}, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	response := Response{
		Body:    string(body),
		Status:  resp.StatusCode,
		Headers: resp.Header,
	}

	return response, err
}