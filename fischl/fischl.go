package fischl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type Fischl struct {
	Client *http.Client
}

func NewFischl() *Fischl {
	fischl := new(Fischl)
	fischl.Client = &http.Client{}
	return fischl
}

func (f *Fischl) HttpGet(url string) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	info, _ := f.Client.Do(request)
	defer info.Body.Close()
	return ioutil.ReadAll(info.Body)
}

func (f *Fischl) HttpPost(url, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	request, _ := http.NewRequest("POST", url, body)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	request.Header.Set("Content-Type", "application/json")
	info, _ := f.Client.Do(request)
	defer info.Body.Close()
	return ioutil.ReadAll(info.Body)
}

func (f *Fischl) HttpPut(url, data string) ([]byte, error) {
	body := strings.NewReader(data)
	request, _ := http.NewRequest("PUT", url, body)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	info, _ := f.Client.Do(request)
	defer info.Body.Close()
	return ioutil.ReadAll(info.Body)
}
