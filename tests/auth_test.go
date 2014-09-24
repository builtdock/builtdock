// +build integration

package tests

import (
	"testing"

	"github.com/builtdock/builtdock/tests/utils"
)

var (
	authLoginCmd    = "auth:login http://deis.{{.Domain}} --username={{.UserName}} --password={{.Password}}"
	authLogoutCmd   = "auth:logout"
	authRegisterCmd = "auth:register http://deis.{{.Domain}} --username={{.UserName}} --password={{.Password}} --email={{.Email}}"
)

func TestAuth(t *testing.T) {
	params := authSetup(t)
	authRegisterTest(t, params)
	authLogoutTest(t, params)
	authLoginTest(t, params)
	authCancel(t, params)
}

func authSetup(t *testing.T) *utils.DeisTestConfig {
	user := utils.GetGlobalConfig()
	user.UserName, user.Password = utils.NewID(), utils.NewID()
	return user
}

func authCancel(t *testing.T, params *utils.DeisTestConfig) {
	utils.AuthCancel(t, params)
}

func authLoginTest(t *testing.T, params *utils.DeisTestConfig) {
	cmd := authLoginCmd
	utils.Execute(t, cmd, params, false, "")
	params = authSetup(t)
	utils.Execute(t, cmd, params, true, "200 OK")
}

func authLogoutTest(t *testing.T, params *utils.DeisTestConfig) {
	utils.Execute(t, authLogoutCmd, params, false, "")
}

func authRegisterTest(t *testing.T, params *utils.DeisTestConfig) {
	cmd := authRegisterCmd
	utils.Execute(t, cmd, params, false, "")
	utils.Execute(t, cmd, params, true, "Registration failed")
}
