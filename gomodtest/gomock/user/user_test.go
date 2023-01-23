package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/io/gomockt/mocks"
)

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDoer := mocks.NewMockDoer(ctrl)
	testUser := &User{Doer: mockDoer}

	// Times 声明函数调用预期执行的确切次数
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)
	testUser.Use()
}
