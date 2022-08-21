package account

import (
	"bytes"
	"demo/service/gateway/api/internal/types"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

var BaseUrl = "http://localhost:8888/api/account/"

func httpTest(method string, url string, request any, response any, t *testing.T) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		t.Error(err.Error())
	}
	req, err := http.NewRequest(method, BaseUrl+url, bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	if resp.StatusCode != 200 {
		t.Errorf("http status code is %d", resp.StatusCode)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err.Error())
	}
	err = resp.Body.Close()
	if err != nil {
		t.Error(err.Error())
	}
	err = json.Unmarshal(respBody, response)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLogin(t *testing.T) {
	loginRequest := &types.LoginRequest{
		Name:     "foliet",
		Password: "1234567",
	}
	loginResponse := &types.LoginResponse{}
	httpTest(http.MethodPost, "login", loginRequest, loginResponse, t)
	if loginResponse.Code != 0 {
		t.Errorf("body status code is %d\n%s", loginResponse.Code, loginResponse.Msg)
	}
}
