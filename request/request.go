package request

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	neturl "net/url"
	"time"
)

type HttpToMD5Response struct {
	Url  string
	Hash string
}

// ClientInterface is implemented by Client and MockClient
type ClientInterface interface {
	MakeRequest(string) ([]byte, error)
}

type Client struct{}

func NormalizeUrl(url string) (string, error) {
	urlParts, _ := neturl.Parse(url)
	if urlParts.Scheme == "" {
		url = "http://" + url
	}
	_, err := neturl.Parse(url)
	return url, err
}

func (c Client) MakeRequest(url string) ([]byte, error) {
	body := make([]byte, 0)
	client := &http.Client{
	    Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return body, err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	resp.Body.Close()
	return body, nil
}

func HttpToMD5(url string, client ClientInterface) (HttpToMD5Response, error) {
	url, err := NormalizeUrl(url)
	if err != nil {
		return HttpToMD5Response{Url: url}, err
	}

	body, err := client.MakeRequest(url)
	if err != nil {
		return HttpToMD5Response{Url: url}, err
	}

	hash := md5.Sum(body)

	return HttpToMD5Response{
		Url:  url,
		Hash: hex.EncodeToString(hash[:]),
	}, nil
}
