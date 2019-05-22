package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var expireTime = 0
var token = ""

type Mp3Link struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Cover    string `json:"cover"`
	Normal   string `json:"128"`
	High     string `json:"320"`
	Lossless string `json:"lossless"`
}

func generateNCTToken() (accessToken string) {
	url := "https://graph.nhaccuatui.com/v1/commons/token"
	data := []byte(`deviceinfo={"DeviceID":"dd03852ada21ec149103d02f76eb0a04","DeviceName":"HellCatVN:S7:NMF26X","OsName":"ANDROID","OsVersion":"8.0","AppName":"NCTTablet","UserName":"hellcatvn","QualityPlay":"128","QualityDownload":"128","QualityCloud":"128","Network":"WIFI","Provider":"NCTCorp"}&md5=ebd547335f855f3e4f7136f92ccc6955&timestamp=1499177482892`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Host", "graph.nhaccuatui.com")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "Keep-Alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	accessToken = raw["data"].(map[string]interface{})["accessToken"].(string)
	expireTime = int(raw["data"].(map[string]interface{})["timeExpire"].(float64))
	return
}

func getNCT(id string) (mp3Link Mp3Link) {
	if currentTime := int(time.Now().Unix()); currentTime > expireTime {
		token = generateNCTToken()
	}
	url := "https://graph.nhaccuatui.com/v1/songs/" + id + "?access_token=" + token
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	var raw map[string]interface{}
	json.Unmarshal(body, &raw)
	data := raw["data"].(map[string]interface{})
	mp3Link.Title = data["2"].(string)
	mp3Link.Artist = data["3"].(string)
	mp3Link.Cover = data["8"].(string)
	mp3Link.Normal = strings.Replace(data["11"].(string), "&download=true", "", -1)

	if high := data["12"].(string); strings.Index(high, "mp3") > 0 {
		mp3Link.High = strings.Replace(high, "&download=true", "", -1)
	}

	if lossless := data["19"].(string); strings.Index(lossless, "flac") > 0 {
		mp3Link.Lossless = strings.Replace(lossless, "&download=true", "", -1)
	}
	return
}

func getLinkNCT(url string) (mp3Link Mp3Link) {
	r, _ := regexp.Compile(`.*nhaccuatui.com\/bai-hat.*`)

	if r.MatchString(url) {
		arrPath := strings.Split(url, "/")
		lastPath := arrPath[len(arrPath)-1]

		if idArr := strings.Split(lastPath, "."); len(idArr) == 3 {
			nctId := idArr[1]
			mp3Link = getNCT(nctId)
		}
	}
	return
}

//func main() {
//	fmt.Println(getLinkNCT("https://www.nhaccuatui.com/bai-hat/until-you-shayne-ward.s4WCRxBVMdCS.html"))
//}
