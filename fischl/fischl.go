package fischl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

type Fischl struct {
	Client  *http.Client
	Header  map[string]string
	request *http.Request
}

func NewFischl() *Fischl {
	fischl := new(Fischl)
	fischl.Client = &http.Client{}
	return fischl
}

func (f *Fischl) HttpGet(url string) ([]byte, error) {
	f.request, _ = http.NewRequest("GET", url, nil)
	return f.run()
}

func (f *Fischl) HttpPost(url, data string) ([]byte, error) {
	f.Header = map[string]string{
		"Content-Type": "application/json",
	}
	body := bytes.NewBuffer([]byte(data))
	f.request, _ = http.NewRequest("POST", url, body)
	return f.run()
}

func (f *Fischl) HttpPut(url, data string) ([]byte, error) {
	body := strings.NewReader(data)
	f.request, _ = http.NewRequest("PUT", url, body)
	return f.run()
}

func (f *Fischl) run() ([]byte, error) {
	f.request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	if len(f.Header) != 0 {
		for k, v := range f.Header {
			f.request.Header.Set(k, v)
		}
	}
	info, err := f.Client.Do(f.request)
	if err != nil {
		return nil, err
	}
	defer info.Body.Close()
	return ioutil.ReadAll(info.Body)
}
