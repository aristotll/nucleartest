from __future__ import unicode_literals # 解决 json.dumps中文乱码
import requests
import json


headers = {'cookie': 'SINAGLOBAL=7887897319966.217.1644161664678; UM_distinctid=17f6d2b362b973-0e4f63b661c88d-113f645d-13c680-17f6d2b362c1369; SSOLoginState=1647576641; XSRF-TOKEN=FmpPDPJousz-sYfmtuIypglT; _s_tentry=www.google.com; UOR=,,www.google.com; Apache=1307078468349.6272.1648138820433; ULV=1648138820449:3:2:1:1307078468349.6272.1648138820433:1646804022969; wvr=6; wb_view_log_5317804367=1440*9002; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WWoaCjKZXjOLGBY.uSPyQey5JpX5KMhUgL.Fo-0eKMRehB0SoM2dJLoI7LadcY_PEX7eh5t; ALF=1679917802; SCF=AmOq8snt8F8rgFwawX341JCira8jmMTa_PH2ye9tx7wEmR4K9lDFXgfuM-gfGXkMFGsVof2MT0vc3rZGji6GhoI.; SUB=_2A25PRD87DeRhGeNN6lUZ8CrPzTuIHXVsMBfzrDV8PUNbmtAKLRWlkW9NScUSzXELytByTiF9C2vyJJv0rmqzj7RW; WBStorage=f4f1148c|undefined; webim_unReadCount=%7B%22time%22%3A1648382231526%2C%22dm_pub_total%22%3A0%2C%22chat_group_client%22%3A0%2C%22chat_group_notice%22%3A0%2C%22allcountNum%22%3A75%2C%22msgbox%22%3A0%7D; WBPSESS=u1_v-mQPMrywxGWHtCCqoFtKsqyXZylYYzkJtX2hu5hhQs6FPaOTiKRU33LBnlxLXEbvq_NSJezbLaQJwc5AlDbjOwD-jKvnjexh5mTH3AoXEtUx1W8vF5Aip1YBX4G4X9y-n5osali9o5qz-s6UbA=='}
# 所有数据保存在此，之后直接将其转换为 str 一次写入磁盘
allFollows = []
totalNumber = 0
gotNumber = 0


class Follow:
    id: int
    name: str
    description: str
    followersCount: int
    location: str

    def string(self) -> str:
        return "id=%d, name=%s, description=%s, followerCount=%d, location=%s" % (self.id, self.name, self.description, self.followersCount, self.location)


def filterAndSave(users: list):
    global gotNumber
    for u in users:
        # 返回的 json 字段非常多，这里只拿我们需要的几个字段
        fo = Follow()
        fo.id = u['id']
        fo.name = u['name']
        fo.location = u['location']
        fo.followersCount = u['followers_count']
        fo.description = u['description']
        #print(fo.string())
        # 重新序列化
        newj = json.dumps(fo.__dict__, ensure_ascii=False)
        print(newj)
        # 添加到 list 中
        allFollows.append(newj)
        gotNumber += 1


def saveToFile(data: str):
    f = open('follow.json', mode='w+', encoding="utf-8")
    f.write(data)
    f.close()


def gotit():
    page = 1
    nextCursor = 0

    while True:
        if page == 1:
            r = requests.get(
                'https://weibo.com/ajax/profile/followContent?sortType=all', headers=headers)
            j = json.loads(r.text)
            nextCursor = j['data']['follows']['next_cursor']
            totalNumber = j['data']['total_number']
            users = j['data']['follows']['users']
            print("[page=1]users len: %d" % (len(users)))
            filterAndSave(users=users)
            page += 1
        else:
            url = 'https://weibo.com/ajax/profile/followContent?page=%d&next_cursor=%s' % (
                page, nextCursor)
            r = requests.get(url, headers=headers)
            j = json.loads(r.text)
            nextCursor = j['data']['follows']['next_cursor']
            users = j['data']['follows']['users']
            print("[page=!1]users len: %d" % (len(users)))
            filterAndSave(users=users)
            page += 1
            if j['data']['follows']['next_cursor'] == 0:
                break
    # 一次性写入磁盘减少 IO 次数
    str = '\n'.join(allFollows)
    saveToFile(str)
    print('total: %d, finish: %d\n' % (totalNumber, gotNumber))


if __name__ == '__main__':
    gotit()
