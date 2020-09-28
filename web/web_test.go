package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestWeb(t *testing.T) {
	port := 8000
	web := New(port)
	defer func() {
		os.Exit(0)
	}()
	go func() { web.Run() }()
	values := url.Values{}
	values.Set("url", "https://github.com/java-native-access/jna/blob/master/www/GettingStarted.md")
	res, err := http.Post(fmt.Sprintf("http://localhost:%d/add", port),
		"application/x-www-form-urlencoded",
		strings.NewReader(values.Encode()),
	)
	if nil != err {
		t.Error(err)
		return
	}
	if nil != res.Body {
		defer res.Body.Close()
	}
	if 200 <= res.StatusCode && 300 > res.StatusCode {
		raw, err := ioutil.ReadAll(res.Body)
		if nil != err {
			t.Error(err)
			return
		}
		res, err := http.Get(fmt.Sprintf("http://localhost:%d%s", port, string(raw)))
		if nil != err {
			t.Error(err)
			return
		}
		if nil != res.Body {
			defer res.Body.Close()
		}
		if 200 <= res.StatusCode && 300 > res.StatusCode {
			raw, err := ioutil.ReadAll(res.Body)
			if nil != err {
				t.Error(err)
				return
			}
			fmt.Println(string(raw))
		} else {
			t.Error("http status code error", res.StatusCode)
		}
	} else {
		t.Error("http status code error", res.StatusCode)
	}
}
