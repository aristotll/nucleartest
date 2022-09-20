package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type FollowResponse struct {
	Data struct {
		Follows struct {
			Users []struct {
				ID          int64  `json:"id"`
				ScreenName  string `json:"screen_name"`
				Name        string `json:"name"`
				Location    string `json:"location"`
				Description string `json:"description"`
			} `json:"follows"`
			NextCursor int `json:"next_cursor"`
		} `json:"data"`
		TotalNumber int `json:"total_number"`
		Ok          int `json:"ok"`
	}
}

func gotit() *FollowResponse {
	//var f FollowResponse
	var (
		page        = 1
		nextCursor  = 0
		totalNumber = 0
		gotNumber   = 0
	)
	const cookie = ""

	for {
		if page == 1 {
			b := httpGet("https://weibo.com/ajax/profile/followContent?sortType=all", cookie)
			fmt.Println(b)
			nextCursor = b.Data.Follows.NextCursor
			totalNumber = b.Data.TotalNumber

			b1, err := json.Marshal(b)
			if err != nil {
				panic(err)
			}
			saveToFile(b1)
			gotNumber += len(b.Data.Follows.Users)
			page++
		} else {
			_url := fmt.Sprintf("https://weibo.com/ajax/profile/followContent?page=%v&next_cursor=%v", page, nextCursor)
			b := httpGet(_url, cookie)
			nextCursor = b.Data.Follows.NextCursor
			b1, err := json.Marshal(b)
			if err != nil {
				panic(err)
			}
			saveToFile(b1)
			gotNumber += len(b.Data.Follows.Users)

			// 没有数据了
			if nextCursor == 0 {
				break
			}
			page++
		}
	}
	if totalNumber == gotNumber {
		fmt.Printf("total %v, got %v, all is done\n", totalNumber, gotNumber)
	} else {
		fmt.Printf("total %v, got %v, some is wrong\n", totalNumber, gotNumber)
	}

	return nil
}

func httpGet(url, cookie string) *FollowResponse {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("cookie", cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fr := &FollowResponse{}
	if err := json.Unmarshal(b, fr); err != nil {
		panic(err)
	}

	return fr
}

func saveToFile(data []byte) {
	f, err := os.OpenFile("./follow.json", os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	gotit()
}
