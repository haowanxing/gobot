package gobot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetUUID(t *testing.T) {
	cfg := Load()
	fmt.Println(cfg)
	wx, err := NewWecat(cfg)
	if err != nil {
		panic(err)
	}

	wx.Start()
}

func TestTuling(t *testing.T) {
	params := make(map[string]interface{})
	params["userid"] = "123123123"
	params["key"] = "808811ad0fd34abaa6fe800b44a9556a"
	params["info"] = "你好"

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "http://www.tuling123.com/openapi/api", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Referer", WxReferer)
	req.Header.Add("User-agent", WxUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}
func TestTulingNews(t *testing.T) {
	params := make(map[string]interface{})
	params["userid"] = "123123123"
	params["key"] = "626621c5dfd9a388a9048b89e5f162e0"
	params["info"] = "菜谱"

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", "http://www.tuling123.com/openapi/api", body)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Referer", WxReferer)
	req.Header.Add("User-agent", WxUserAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	data, _ = ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

	var reply Reply

	if err := json.Unmarshal(data, &reply); err != nil {
		fmt.Println(err)
		return
	}

	switch reply.Code {
	case 100000:
		fmt.Println(reply.Text)
	case 200000:
		fmt.Println(reply.Text + " " + reply.URL)
	case 302000:
		var res string
		var list = make([]News, 0)
		news, ok := reply.List.([]interface{})

		if ok {
			jsonList, err := json.Marshal(news)
			if err != nil {
				fmt.Println(err)
			}
			if err := json.Unmarshal(jsonList, &list); err != nil {
				fmt.Println(err)
			}
			for _, n := range list {
				res += fmt.Sprintf("%s\n%s\n", n.Article, n.DetailURL)
			}
		}
		fmt.Println(res)
	case 308000:
		var res string
		var list = make([]Menu, 0)
		menus, ok := reply.List.([]interface{})
		if ok {
			jsonList, err := json.Marshal(menus)
			if err != nil {
				fmt.Println(err)
			}
			if err := json.Unmarshal(jsonList, &list); err != nil {
				fmt.Println(err)
			}
			for _, m := range list {
				res += fmt.Sprintf("%s\n%s\n%s\n", m.Name, m.Info, m.DetailURL)
			}
		}
		fmt.Println(res)
	default:
		fmt.Println("不知...")
	}
}
