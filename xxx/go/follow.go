package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/youseebiggirl/requests"
)

var fs = make([]string, 0, 100)
var dstCookie = "XSRF-TOKEN=vXiIhmjvwyPCygYO6k9Lusdf; _ga=GA1.2.1288003011.1650532620; _gid=GA1.2.1310798656.1650532621; _ga_NNQRFSG6W5=GS1.1.1650532620.1.0.1650532621.0; SUB=_2A25PZVERDeThGeFJ7lcT-SfOyDWIHXVsE8XZrDV8PUNbmtB-LXWlkW9Nf4vOmBW--riaV2dOaY7JHb5taYeJ0Bd8; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWW0mHq0QwJ6K2aEHkMvEJL5JpX5KzhUgL.FoMNSK-E1K.Ee0.2dJLoIp7LxKML1KBLBKnLxKqL1hnLBoMNS0-feo.4eoe4; ALF=1682068672; SSOLoginState=1650532672; WBPSESS=JxldVRsUdlV8glbmXOjG-WsyitWdMeZtYBHLnP9cVFzqw8EYtrwHr_W4PV-EhPACJkISO6x06IW-fcUDgRyqnJ4u9iNexWJlU6u3Epq8GOOk0QrfMxvgl6pQzb5wmbILQUIRHmzdNiuMnpz-m6a89A==; _s_tentry=weibo.com; Apache=7474424849615.253.1650532674027; SINAGLOBAL=7474424849615.253.1650532674027; ULV=1650532674048:1:1:1:7474424849615.253.1650532674027:"

func transferFollowsToAnotherAccount() {
	for _, id := range fs {
		j := fmt.Sprintf(`{"id":"%s"}`, id)
		_ = j
		time.Sleep(time.Millisecond * 500)
	}
}

func getFollowsFromFile() {
	b, err := os.ReadFile("../follow.json")
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]any)
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatal(err)
	}
	data := m["data"].([]any)
	for _, v := range data {
		vv := v.(map[string]any)
		id := vv["id"]
		idstr := strconv.FormatFloat(id.(float64), 'f', 0, 64)
		j := fmt.Sprintf(`{"friend_uid": "%v"}`, idstr)
		//log.Println(j)
		follow([]byte(j))
	}
}

func follow(json_ []byte) {
	r := requests.POST("https://weibo.com/ajax/friendships/create",
		requests.WithCookie(dstCookie),
		requests.WithJson(json_),
		requests.WithHeaders(http.Header{"x-xsrf-token": []string{"vXiIhmjvwyPCygYO6k9Lusdf"}}),
	)
	fmt.Printf("%v %v\n", r.StatusCode(), r.StatusText())
	if r.StatusCode() != http.StatusOK {
		fmt.Println(r.Text())
	}
}

func main() {
	getFollowsFromFile()
}
