package request

import (
	// "net/http"
)

type HttpToMD5Response struct {
	Url string
	Hash string
}

func NormalizeUrl(url string) string {
	// TODO
	return url
}

func HttpToMD5(url string) (HttpToMD5Response, error) {
	url = NormalizeUrl(url)
	// TODO
	return HttpToMD5Response {
		Url: url,
		Hash: "",
	}, nil
}
