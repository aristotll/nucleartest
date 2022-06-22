package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/levigross/grequests"
)

func main() {
	url := "https://weibo.com/ajax/favorites/all_fav"
	resp, err := grequests.Get(url, &grequests.RequestOptions{
		Headers: map[string]string{
			"Cookie": "SINAGLOBAL=7887897319966.217.1644161664678; UM_distinctid=17f6d2b362b973-0e4f63b661c88d-113f645d-13c680-17f6d2b362c1369; UOR=,,www.google.com; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWoaCjKZXjOLGBY.uSPyQey5JpX5KMhUgL.Fo-0eKMRehB0SoM2dJLoI7LadcY_PEX7eh5t; ULV=1649390707049:4:1:1:2651698005693.741.1649390707040:1648138820449; ALF=1681488044; SSOLoginState=1649952044; SCF=AmOq8snt8F8rgFwawX341JCira8jmMTa_PH2ye9tx7wEUyWnNYKVz3JT3MJgVwJYX3EUSJiIMhMqumijHM6tHqo.; SUB=_2A25PXDV9DeRhGeNN6lUZ8CrPzTuIHXVsKCG1rDV8PUNbmtAKLUjMkW9NScUSzRFTgjWmc-udaOPW_F752Qju3nJ9; XSRF-TOKEN=ijwzsYYvFe3AzIyD2OMZF_fi; WBPSESS=u1_v-mQPMrywxGWHtCCqoFtKsqyXZylYYzkJtX2hu5hhQs6FPaOTiKRU33LBnlxLL1xPPeOCoMGzr7iEPRbA9P2VbKEu517DcJAXfzT07TABuQaNnFCgczFS51PWpRQCUDZaaD7lc_dv9B6a_PpGWA==",
		},
	})
	if err != nil {
		panic(err)
	}
	//fmt.Printf("resp.String(): %v\n", resp.String())
	m := make(map[string]any)
	// if err := json.Unmarshal(resp.Bytes(), &m); err != nil {
	// 	panic(err)
	// }
	//resp.JSON(m)
	json.NewDecoder(bytes.NewBuffer(resp.Bytes())).Decode(&m)
	for k := range m {
		fmt.Printf("type: %T, val: %v\n", k, k)
	}
}
