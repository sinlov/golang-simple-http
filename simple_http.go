package simple_http

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
	"sort"
)

func byte2String(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func Post_String(urlStr string, headParams map[string](string), bodyParams map[string](string)) (string, error) {
	res, err := Post(urlStr, headParams, bodyParams)
	if err != nil {
		return "", err
	}
	resStr := byte2String(res)
	return resStr, err
}

func Post(urlStr string, headParams map[string](string), bodyParams map[string](string)) ([]byte, error) {
	client := &http.Client{}
	bodyP := ""
	if bodyParams != nil {
		for k, v := range bodyParams {
			strings.Join(bodyP, k)
			strings.Join(bodyP, "=")
			strings.Join(bodyP, v)
		}
	}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(bodyP))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headParams != nil {
		for k, v := range headParams {
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func Get_String(urlStr string, headParams map[string](string)) (string, error) {
	res, err := Get(urlStr, headParams)
	if err != nil {
		return "", err
	}
	resStr := byte2String(res)
	return resStr, err
}

func Get(urlStr string, headParams map[string](string)) ([]byte, error) {
	u, _ := url.Parse(urlStr)
	q := u.Query()
	if headParams != nil {
		for k, v := range headParams {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	if err != nil {
		return nil, err
	}
	resByte, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return resByte, nil
}

func Get_Sort(urlStr string, headParams map[string](string)) ([]byte, error) {
	u, _ := url.Parse(urlStr)
	q := u.Query()
	if headParams != nil {
		var keys []string
		for k := range headParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			q.Set(k, headParams[k])
		}
	}
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	if err != nil {
		return nil, err
	}
	resByte, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return resByte, nil
}
