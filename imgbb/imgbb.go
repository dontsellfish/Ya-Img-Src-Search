package imgbb

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func Post(apiKey string, imgPath string, opts ...interface{}) (displayUrl string, err error) {
	buffer, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return
	}

	payload := url.Values{"key": {apiKey}, "image": {base64.StdEncoding.EncodeToString(buffer)}}

	for _, opt := range opts {
		switch opt.(type) {
		case int:
			payload["expiration"] = []string{strconv.Itoa(opt.(int))}
		case string:
			payload["name"] = []string{opts[1].(string)}
		}
	}

	resp, err := http.PostForm("https://api.imgbb.com/1/upload", payload)
	if err != nil {
		return
	}

	buffer, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	println(string(buffer))

	var imgbbResp imgbbApiResponse
	err = json.Unmarshal(buffer, &imgbbResp)
	if err != nil {
		return
	}

	if imgbbResp.Success {
		return imgbbResp.Data.DisplayUrl, nil
	} else {
		err = newError(string(buffer))
		return
	}
}

type fieldImgbbApiResponeData struct {
	DisplayUrl string `json:"url,omitempty"`
}

type imgbbApiResponse struct {
	Data    fieldImgbbApiResponeData `json:"data,omitempty"`
	Status  int                      `json:"status"`
	Success bool                     `json:"success"`
}

type dynamicGolangError struct {
	body string
}

func newError(body string) dynamicGolangError {
	return dynamicGolangError{body: body}
}

func (err dynamicGolangError) Error() string {
	return err.body
}
