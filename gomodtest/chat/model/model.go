package model

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Session struct {
	sync.Mutex
	ID         uint64
	Name       string
	msgs       []*Message
	users      []*User
	LastMsgID  uint64
	haveNewMsg chan struct{}
}

func (s *Session) PushMessage(m *Message) {
	s.Lock()
	defer s.Unlock()
	s.msgs = append(s.msgs, m)
	s.LastMsgID++
	if s.haveNewMsg == nil {
		s.haveNewMsg = make(chan struct{}, 100)
	}
	s.haveNewMsg <- struct{}{}
}

type Message struct {
	ID   uint64
	From uint64
	To   uint64
	Data []byte
}

type User struct {
	ID       uint64
	Name     string
	Sessions map[uint64]uint64
	msgCh    chan *Message
}

func (u *User) Read(s *Session) error {
	if !userIsBelongSession(u, s) {
		return fmt.Errorf("this user is not belong this session")
	}
	if u.msgCh == nil {
		u.msgCh = make(chan *Message, 1)
	}
	for {
		if err := u.recvMsg(s); err != nil {
			return err
		}
	}
}

func (u *User) Write() error {
	for {
		select {
		case msg := <-u.msgCh:
			log.Printf("[user id=%v] recv msg: %+v\n", u.ID, msg)
		}
	}
}

func (u *User) recvMsg(s *Session) error {
	lastRecvId := u.Sessions[s.ID]
	slast := s.LastMsgID
	if lastRecvId > slast {
		return fmt.Errorf("lastRecvId is error, out of the msg size")
	}
	<-s.haveNewMsg
	for i := lastRecvId; i < slast; i++ {
		select {
		case u.msgCh <- s.msgs[i]:
			u.Sessions[s.ID]++
		}
	}
	return nil
}

func CreateFakeMsg() *Message {
	return &Message{
		ID:   rand.Uint64(),
		From: rand.Uint64(),
		To:   rand.Uint64(),
	}
}

func CreateFakeUser() *User {
	return &User{
		ID:       rand.Uint64(),
		Name:     "user:" + strconv.Itoa(rand.Intn(100)),
		Sessions: make(map[uint64]uint64),
		msgCh:    make(chan *Message, 1),
	}
}

func userIsBelongSession(u *User, s *Session) (exist bool) {
	for _, user := range s.users {
		if user.ID == u.ID {
			exist = true
			return
		}
	}
	exist = false
	return
}
