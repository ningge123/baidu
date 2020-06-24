package baidu

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// tokenAPI 获取带 token 的 API 地址
func tokenApi(api, token string) (string, error) {
	queries := requestQueries{
		"access_token": token,
	}

	return encodeURL(api, queries)
}

// encodeURL add and encode parameters.
func encodeURL(api string, params requestQueries) (string, error) {
	url, err := url.Parse(api)

	if err != nil {
		return "", err
	}

	query := url.Query()

	for k, v := range params {
		query.Set(k, v)
	}

	url.RawQuery = query.Encode()

	return url.String(), nil
}

// base 64 压缩图片
func Base64EncodeImage(filePath string) (string, error)  {
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	sourceBuffer := make([]byte, 500000)
	n, err := file.Read(sourceBuffer)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sourceBuffer[:n]), nil
}

func postForm(url string, params url.Values, response interface{}) error {
	// Prepare a form that you will submit to that URL.
	request, err := http.PostForm(url, params)

	if err != nil {
		return err
	}
	defer request.Body.Close()

	return json.NewDecoder(request.Body).Decode(response)
}

// postJSON perform a HTTP/POST request with json body
func postJson(url string, params interface{}, response interface{}) error {
	resp, err := postJSONWithBody(url, params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}

// postJSONWithBody return with http body.
func postJSONWithBody(url string, params interface{}) (*http.Response, error) {
	b := &bytes.Buffer{}
	if params != nil {
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(params)
		if err != nil {
			return nil, err
		}
	}

	return http.Post(url, "application/json; charset=utf-8", b)
}


func getJSON(url string, response interface{}) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(response)
}