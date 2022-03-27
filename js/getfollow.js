const axios = require('axios')
const fs = require('fs')

axios.get('https://weibo.com/ajax/profile/followContent?sortType=all', {
  headers: {
    'cookie': 'SINAGLOBAL=7887897319966.217.1644161664678; UM_distinctid=17f6d2b362b973-0e4f63b661c88d-113f645d-13c680-17f6d2b362c1369; SSOLoginState=1647576641; XSRF-TOKEN=FmpPDPJousz-sYfmtuIypglT; _s_tentry=www.google.com; UOR=,,www.google.com; Apache=1307078468349.6272.1648138820433; ULV=1648138820449:3:2:1:1307078468349.6272.1648138820433:1646804022969; wvr=6; wb_view_log_5317804367=1440*9002; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWoaCjKZXjOLGBY.uSPyQey5JpX5KMhUgL.Fo-0eKMRehB0SoM2dJLoI7LadcY_PEX7eh5t; ALF=1679917802; SCF=AmOq8snt8F8rgFwawX341JCira8jmMTa_PH2ye9tx7wEmR4K9lDFXgfuM-gfGXkMFGsVof2MT0vc3rZGji6GhoI.; SUB=_2A25PRD87DeRhGeNN6lUZ8CrPzTuIHXVsMBfzrDV8PUNbmtAKLRWlkW9NScUSzXELytByTiF9C2vyJJv0rmqzj7RW; WBStorage=f4f1148c|undefined; webim_unReadCount=%7B%22time%22%3A1648382231526%2C%22dm_pub_total%22%3A0%2C%22chat_group_client%22%3A0%2C%22chat_group_notice%22%3A0%2C%22allcountNum%22%3A75%2C%22msgbox%22%3A0%7D; WBPSESS=u1_v-mQPMrywxGWHtCCqoFtKsqyXZylYYzkJtX2hu5hhQs6FPaOTiKRU33LBnlxLXEbvq_NSJezbLaQJwc5AlDbjOwD-jKvnjexh5mTH3AoXEtUx1W8vF5Aip1YBX4G4X9y-n5osali9o5qz-s6UbA=='
  }
}).then(res => {
  console.log(`statusCode: ${res.status}`)
  //console.log(res.data.data.follows)
  var jsonData = res.data.data.follows
  //console.log(jsonData)
  var jsonObj = JSON.parse(jsonData)
  var jsonContent = JSON.stringify(jsonObj)

  fs.writeFileSync('./output.json', jsonContent, function (err) {
    if (err) {
      console.log("write data to file error: " + err)
    }
    console.log("write data done")
  })
})






