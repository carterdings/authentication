package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test_Server ...
func Test_Server(t *testing.T) {
	go initAndServe("127.0.0.1:8080", "gin.log")
	time.Sleep(time.Second)

	tests := []struct {
		name string
		furl string
		body string
	}{
		{"test_CreateUser", "http://127.0.0.1:8080/user/create", `{"user_name":"cat","password":"test"}`},
		{"test_CreateRole", "http://127.0.0.1:8080/role/create", `{"role_name":"root"}`},
		{"test_AddRoleToUser", "http://127.0.0.1:8080/user/add_role", `{"user_name":"cat","password":"test","role_name":"root"}`},
		{"test_Authenticate", "http://127.0.0.1:8080/auth/authenticate", `{"user_name":"cat","password":"test"}`},
		{"test_Invalidate", "http://127.0.0.1:8080/auth/invalidate", `{"user_name":"cat","password":"test","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}`},
		{"test_CheckRole", "http://127.0.0.1:8080/user/check_role", `{"user_name":"cat","password":"test","role_name":"root","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}`},
		{"test_AllRoles", "http://127.0.0.1:8080/user/all_roles", `{"user_name":"cat","password":"test","role_name":"root","token":"expire=7200&rand=ZrCV4SLd0euk4lcHHH2cHA&token=FKunVi5yiLpGOnt5CplnT7rWtzdp-eJ4w_l9T9Yx_eUHkqBOP-ZxDHKi6nqn33JjCeSetuGlEsQ8thBU9Y5ZXG__lvBcwFhRWbWLHR_fiXQgyobrtM4bxvzXTZpGNX5Jf9ssL2YoHqeihGuHWq4DyJnqZkiVz51P5Kqh3-2WVqA&ts=1661841402&user=cat"}`},

		{"test_DeleteUser", "http://127.0.0.1:8080/user/delete", `{"user_name":"cat"}`},
		{"test_DeleteRole", "http://127.0.0.1:8080/role/delete", `{"role_name":"root"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := http.Post(tt.furl, "application/json", strings.NewReader(tt.body))
			assert.Nil(t, err)
			defer rsp.Body.Close()
			data, err := ioutil.ReadAll(rsp.Body)
			assert.Nil(t, err)
			t.Logf("furl: %s, rsp data: %s", tt.furl, string(data))
			var result map[string]interface{}
			err = json.Unmarshal(data, &result)
			assert.Nil(t, err)
			t.Logf("response: %+v", result)
		})
	}
}
