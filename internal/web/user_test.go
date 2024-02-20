package web

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"webook/internal/domian"
	"webook/internal/domian/req"
	"webook/internal/service"
	svcmocks "webook/internal/service/mock"
	"webook/ioc"
)

func TestUser_SingUp(t *testing.T) {
	testCases := []struct {
		name     string
		after    func(t *testing.T)
		mock     func(ctrl *gomock.Controller) service.UserService
		reqBody  req.UserReq
		wantCode int
		wantResp domian.Result[any]
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).
					Return(nil)
				return usersvc
			},
			reqBody: req.UserReq{
				Name:     "xiaobao",
				Password: "Aa123456",
			},
			wantCode: 200,
			wantResp: domian.Result[any]{
				Code:    200,
				Message: "注册成功",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := ioc.InitUserHandler()
			jsonStr, err := json.Marshal(testCase.reqBody)
			if err != nil {
				return
			}
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost,
				"/users/signup", bytes.NewBuffer(jsonStr))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder() //http返回值
			server.ServeHTTP(resp, req)
			assert.Equal(t, resp.Code, 200)
			var result domian.Result[any]
			assert.NoError(t, json.Unmarshal(resp.Body.Bytes(), &result))
		})

	}
}

// mock单元测试
func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usersvc := svcmocks.NewMockUserService(ctrl)
	usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).
		Return(errors.New("mock error"))
	err := usersvc.SignUp(context.Background(), domian.User{Name: "xiaobao", Password: "123"})
	t.Log(err)
}
