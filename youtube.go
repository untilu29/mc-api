package main

import (
	"context"
	"mc-api/youtube"
)

type YoutubeLink struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Cover    string `json:"cover"`
	Normal   string `json:"128"`
	High     string `json:"320"`
	Lossless string `json:"lossless"`
}

func getLinkYoutube(url string) (mp3Link string) {
	var client = youtube.Client{Debug: false}
	video, _ := client.GetVideoContext(context.Background(), url)

	var format = video.Formats.FindByItag(140)

	mp3Link, err := client.GetStreamURLContext(context.Background(), video, format)
	if err != nil {
		print(err)
	}
	return
}

//func main() {
//	fmt.Println(getLinkYoutube("https://www.youtube.com/watch?v=Zs8CnpoiyaQ"))
//}
