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

func NormalizeUrl(url string) (string, error) {
	urlParts, _ := neturl.Parse(url)
	if urlParts.Scheme == "" {
		url = "http://" + url
	}
	_, err := neturl.Parse(url)
	return url, err
}

func HttpToMD5(url string) (HttpToMD5Response, error) {
	url, err := NormalizeUrl(url)
	if err != nil {
		return HttpToMD5Response{Url: url}, err
	}

	client := &http.Client{
	    Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return HttpToMD5Response{Url: url}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return HttpToMD5Response{Url: url}, err
	}
	resp.Body.Close()

	hash := md5.Sum([]byte(body))

	return HttpToMD5Response{
		Url:  url,
		Hash: hex.EncodeToString(hash[:]),
	}, nil
}
