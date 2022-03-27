from __future__ import unicode_literals # 解决 json.dumps中文乱码
import requests
import json


headers = {}
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
