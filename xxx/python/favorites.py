from __future__ import unicode_literals
from pickletools import int4  # 解决 json.dumps中文乱码
import requests
import json

headers = {'cookie': 'SINAGLOBAL=7887897319966.217.1644161664678; UM_distinctid=17f6d2b362b973-0e4f63b661c88d-113f645d-13c680-17f6d2b362c1369; UOR=,,www.google.com; SSOLoginState=1649952044; XSRF-TOKEN=ijwzsYYvFe3AzIyD2OMZF_fi; _s_tentry=www.weibo.com; Apache=4731403497357.527.1650285486440; ULV=1650285486444:5:2:1:4731403497357.527.1650285486440:1649390707049; login_sid_t=f879f080b06346b2edb9ecf7001f223c; cross_origin_proto=SSL; appkey=; wvr=6; SCF=AmOq8snt8F8rgFwawX341JCira8jmMTa_PH2ye9tx7wEGAlRzColYPTCCqChw9G3PbNjPA3YszaaAyGba2XEIrQ.; SUB=_2A25PW5YtDeRhGeNN6lUZ8CrPzTuIHXVsEIDlrDV8PUNbmtANLWejkW9NScUSzWrah39cjW0jPNxsKZBleyINz6BV; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWoaCjKZXjOLGBY.uSPyQey5JpX5KMhUgL.Fo-0eKMRehB0SoM2dJLoI7LadcY_PEX7eh5t; ALF=1681988093; webim_unReadCount=%7B%22time%22%3A1650452191389%2C%22dm_pub_total%22%3A0%2C%22chat_group_client%22%3A0%2C%22chat_group_notice%22%3A0%2C%22allcountNum%22%3A77%2C%22msgbox%22%3A0%7D; WBPSESS=u1_v-mQPMrywxGWHtCCqoFtKsqyXZylYYzkJtX2hu5hhQs6FPaOTiKRU33LBnlxLXEbvq_NSJezbLaQJwc5AlBry7OsStNRFYR6fIVhYVCW3ndeHg5tWDzRx2h1nnPZK-NqxGtYXPquByzfii6XMIg=='}
baseUrl = 'https://weibo.com/ajax/favorites/all_fav?'
page = 1
urls = []


def getAllFavorites():
    while True:
        global page
        url = baseUrl + 'page=%d' % (page)
        # print(url)
        r = requests.get(url, headers=headers)
        j: json = json.loads(r.text)
        data: slice = j['data']
        if len(data) == 0:
            print('end')
            break
        for d in data:
            mid: int = d['id']
            urls.append(mid)
        page += 1


def saveToFile(data: slice):
    f = open('favorites.txt', mode='w+', encoding="utf-8")
    for d in data:
        s = str(d)
        f.write(s)
        f.write('\n')
    f.close()


def addFavorites():
    headers = {
        'cookie': 'XSRF-TOKEN=vXiIhmjvwyPCygYO6k9Lusdf; _ga=GA1.2.1288003011.1650532620; _gid=GA1.2.1310798656.1650532621; _ga_NNQRFSG6W5=GS1.1.1650532620.1.0.1650532621.0; SUB=_2A25PZVERDeThGeFJ7lcT-SfOyDWIHXVsE8XZrDV8PUNbmtB-LXWlkW9Nf4vOmBW--riaV2dOaY7JHb5taYeJ0Bd8; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWW0mHq0QwJ6K2aEHkMvEJL5JpX5KzhUgL.FoMNSK-E1K.Ee0.2dJLoIp7LxKML1KBLBKnLxKqL1hnLBoMNS0-feo.4eoe4; ALF=1682068672; SSOLoginState=1650532672; WBPSESS=JxldVRsUdlV8glbmXOjG-WsyitWdMeZtYBHLnP9cVFzqw8EYtrwHr_W4PV-EhPACJkISO6x06IW-fcUDgRyqnJ4u9iNexWJlU6u3Epq8GOOk0QrfMxvgl6pQzb5wmbILQUIRHmzdNiuMnpz-m6a89A==; _s_tentry=weibo.com; Apache=7474424849615.253.1650532674027; SINAGLOBAL=7474424849615.253.1650532674027; ULV=1650532674048:1:1:1:7474424849615.253.1650532674027:',
        'x-xsrf-token': 'vXiIhmjvwyPCygYO6k9Lusdf'
    }
    data = {"id":"4760653908937563"}
    r = requests.post(
        'https://weibo.com/ajax/statuses/createFavorites?', 
        headers=headers, json=data)
    print(r.status_code)

if __name__ == '__main__':
    # getAllFavorites()
    # saveToFile(urls)
    # print(urls)
    addFavorites()
