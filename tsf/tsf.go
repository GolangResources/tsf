package tsf

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"net/url"
	"io"
	"mime/multipart"
	"bytes"
	"strconv"
)

type TGInfo struct {
	Token	string
}

func SendMSG(tg TGInfo, destList []string, msg string) ([]string, error) {
	var respList []string
	if (tg.Token == "") {
		return respList, fmt.Errorf("Check the token and try again.")
	}
	for _, v := range destList {
		resp, err := http.Get("https://api.telegram.org/bot"+tg.Token+"/sendMessage?chat_id="+v+"&parse_mode=Markdown&text="+url.QueryEscape(msg))
		if err != nil {
			return respList, err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return respList, err
		}
		respList = append(respList, string(body))
	}
	return respList, nil
}

func SendIMG(tg TGInfo, destList []string, msg string, img []byte) ([]string, error) {
	var respList []string
	if (tg.Token == "") {
		return respList, fmt.Errorf("Check the token and try again.")
	}
	for _, v := range destList {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("photo", "img.jpg")
		if err != nil {
			return respList, err
		}
		_, err = io.Copy(part, bytes.NewReader(img))
		if err != nil {
			return respList, err
		}
		_ = writer.WriteField("chat_id", v)
		_ = writer.WriteField("caption", msg)
		_ = writer.WriteField("parse_mode", "Markdown")
		err = writer.Close()
		req, err := http.NewRequest("POST", "https://api.telegram.org/bot"+tg.Token+"/sendPhoto", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		if err != nil {
			return respList, err
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return respList, err
		}
		bodyRequest := &bytes.Buffer{}
		_, err =  bodyRequest.ReadFrom(resp.Body)
		if err != nil {
			return respList, err
		}
		resp.Body.Close()
		respList = append(respList, strconv.Itoa(resp.StatusCode))
	}
	return respList, nil
}
