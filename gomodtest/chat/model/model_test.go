package model

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func Test(t *testing.T) {
	session := &Session{
		ID:    10086,
		Name:  "中国移动聊天室",
		msgs:  make([]*Message, 0, 100),
		users: make([]*User, 0, 10),
	}
	for i := 0; i < 5; i++ {
		session.PushMessage(CreateFakeMsg())
	}
	u1 := CreateFakeUser()
	u2 := CreateFakeUser()
	session.users = append(session.users, u1, u2)
	log.Printf("%v %v \n", u1.ID, u2.ID)

	go func() {
		err := u1.Read(session)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := u1.Write()
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := u2.Read(session)
		if err != nil {
			panic(err)
		}
	}()
	go func() {
		err := u2.Write()
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Hour)
}

func TestOfflineMsg(t *testing.T) {
	// 上线后能够收到下线期间的消息

}
