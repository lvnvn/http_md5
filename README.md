A tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

- URLs that should be requested are passed as unnamed arguments to the script or as input line when the script is running, separated by spaces. The number of URLs is unlimited.
- URLs without a scheme are considered valid and are requested through HTTP.
- Argument -parallel limits the number of parallel requests; default is 10.

#### Usage examples

```plaintext
./http_md5 http://google.com

./http_md5 -parallel 3 google.com facebook.com yahoo.com yandex.com twitter.com

./http_md5 http://google.com facebook.com
yahoo.com yandex.com http://twitter.com
```
