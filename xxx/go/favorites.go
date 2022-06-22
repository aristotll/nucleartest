package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/youseebiggirl/requests"
)

var (
	srcCookie = "SINAGLOBAL=7887897319966.217.1644161664678; UM_distinctid=17f6d2b362b973-0e4f63b661c88d-113f645d-13c680-17f6d2b362c1369; UOR=,,www.google.com; SSOLoginState=1649952044; XSRF-TOKEN=ijwzsYYvFe3AzIyD2OMZF_fi; _s_tentry=www.weibo.com; Apache=4731403497357.527.1650285486440; ULV=1650285486444:5:2:1:4731403497357.527.1650285486440:1649390707049; login_sid_t=f879f080b06346b2edb9ecf7001f223c; cross_origin_proto=SSL; appkey=; wvr=6; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWoaCjKZXjOLGBY.uSPyQey5JpX5KMhUgL.Fo-0eKMRehB0SoM2dJLoI7LadcY_PEX7eh5t; WBPSESS=u1_v-mQPMrywxGWHtCCqoFtKsqyXZylYYzkJtX2hu5hhQs6FPaOTiKRU33LBnlxLL1xPPeOCoMGzr7iEPRbA9ENkNjvSqmcb7kkEr6i7JOgTzBZRCl3Y9oM4beVRl5F1Qx9gXkjHDOs_KGLLW0FUMw==; SCF=AmOq8snt8F8rgFwawX341JCira8jmMTa_PH2ye9tx7wEKbIUPJ2ey7tZZeHpNDZsrc6nrNDqSv4U9Ouvnkl2lRQ.; SUB=_2A25PZTJoDeRhGeNN6lUZ8CrPzTuIHXVsEySgrDV8PUNbmtB-LVfzkW9NScUSzXIirGS2jKOAwnxONrTzuXbHJygA; ALF=1682077117"
	baseUrl   = "https://weibo.com/ajax/favorites/all_fav?"
	page      = 1
	ids       = make([]string, 0, 100)
	dstCookie = "XSRF-TOKEN=vXiIhmjvwyPCygYO6k9Lusdf; _ga=GA1.2.1288003011.1650532620; _gid=GA1.2.1310798656.1650532621; _ga_NNQRFSG6W5=GS1.1.1650532620.1.0.1650532621.0; SUB=_2A25PZVERDeThGeFJ7lcT-SfOyDWIHXVsE8XZrDV8PUNbmtB-LXWlkW9Nf4vOmBW--riaV2dOaY7JHb5taYeJ0Bd8; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWW0mHq0QwJ6K2aEHkMvEJL5JpX5KzhUgL.FoMNSK-E1K.Ee0.2dJLoIp7LxKML1KBLBKnLxKqL1hnLBoMNS0-feo.4eoe4; ALF=1682068672; SSOLoginState=1650532672; WBPSESS=JxldVRsUdlV8glbmXOjG-WsyitWdMeZtYBHLnP9cVFzqw8EYtrwHr_W4PV-EhPACJkISO6x06IW-fcUDgRyqnJ4u9iNexWJlU6u3Epq8GOOk0QrfMxvgl6pQzb5wmbILQUIRHmzdNiuMnpz-m6a89A==; _s_tentry=weibo.com; Apache=7474424849615.253.1650532674027; SINAGLOBAL=7474424849615.253.1650532674027; ULV=1650532674048:1:1:1:7474424849615.253.1650532674027:"
)

func getFavorites() {
	ch := make(chan struct{}, 10) // 最多 10 条线程同时爬
	done := false

	for !done {
		ch <- struct{}{}
		url := baseUrl + fmt.Sprintf("page=%v", page)
		fmt.Println(url)
		go func() {
			r := requests.GET(url, requests.WithCookie(srcCookie))
			if r.StatusCode() != http.StatusOK {
				err := fmt.Errorf("http GET status error: [%v]%v", r.StatusCode(), r.StatusText())
				log.Fatalln(err)
			}
			m := r.Map()
			data := m["data"].([]any)
			if len(data) == 0 {
				log.Println("no data, maybe is done")
				done = true
			}
			for _, d := range data {
				dd := d.(map[string]any)
				str := dd["idstr"]
				ids = append(ids, str.(string))
			}
			<-ch
		}()
		page++
	}
	saveFavoritesToFile(ids)
}

func saveFavoritesToFile(favorites []string) {
	f, err := os.OpenFile("favorites.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for _, v := range favorites {
		_, err := f.WriteString(v + "\n")
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func getFavoritesFromFile() {
	f, err := os.Open("./favorites.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		s := scan.Text()
		ids = append(ids, s)
	}
}

func transferFavoritesToAnotherAccount() {
	for _, id := range ids {
		j := fmt.Sprintf(`{"id":"%s"}`, id)
		addFavorite([]byte(j))
		time.Sleep(time.Millisecond * 500)
	}
}

func addFavorite(json_ []byte) {
	r := requests.POST("https://weibo.com/ajax/statuses/createFavorites?",
		requests.WithCookie(dstCookie),
		requests.WithHeaders(http.Header{"x-xsrf-token": []string{"vXiIhmjvwyPCygYO6k9Lusdf"}}),
		requests.WithJson(json_),
	)
	fmt.Printf("%v %v\n", r.StatusCode(), r.StatusText())
	if r.StatusCode() != http.StatusOK {
		fmt.Println(r.Text())
	}
}

func main() {
	//getFavorites()
	getFavoritesFromFile()
	transferFavoritesToAnotherAccount()
}
