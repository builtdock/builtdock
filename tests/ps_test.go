// +build integration

package tests

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/builtdock/builtdock/tests/utils"
)

var (
	psListCmd  = "ps:list --app={{.AppName}}"
	psScaleCmd = "ps:scale web={{.ProcessNum}} --app={{.AppName}}"
)

func TestPs(t *testing.T) {
	params := psSetup(t)
	psScaleTest(t, params)
	appsOpenTest(t, params)
	psListTest(t, params, false)
	utils.AppsDestroyTest(t, params)
	utils.Execute(t, psScaleCmd, params, true, "404 NOT FOUND")
}

func psSetup(t *testing.T) *utils.DeisTestConfig {
	cfg := utils.GetGlobalConfig()
	cfg.AppName = "pssample"
	utils.Execute(t, authLoginCmd, cfg, false, "")
	utils.Execute(t, gitCloneCmd, cfg, false, "")
	if err := utils.Chdir(cfg.ExampleApp); err != nil {
		t.Fatal(err)
	}
	utils.Execute(t, appsCreateCmd, cfg, false, "")
	utils.Execute(t, gitPushCmd, cfg, false, "")
	if err := utils.Chdir(".."); err != nil {
		t.Fatal(err)
	}
	return cfg
}

func psListTest(t *testing.T, params *utils.DeisTestConfig, notflag bool) {
	output := "web.2 up (v2)"
	if strings.Contains(params.ExampleApp, "dockerfile") {
		output = strings.Replace(output, "web", "cmd", 1)
	}
	utils.CheckList(t, psListCmd, params, output, notflag)
}

func psScaleTest(t *testing.T, params *utils.DeisTestConfig) {
	cmd := psScaleCmd
	if strings.Contains(params.ExampleApp, "dockerfile") {
		cmd = strings.Replace(cmd, "web", "cmd", 1)
	}
	utils.Execute(t, cmd, params, false, "")
	// Regression test for https://github.com/builtdock/builtdock/pull/1347
	// Ensure that systemd unitfile droppings are cleaned up.
	sshCmd := exec.Command("ssh",
		"-o", "StrictHostKeyChecking=no",
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "PasswordAuthentication=no",
		"core@deis."+params.Domain, "ls")
	out, err := sshCmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), ".service") {
		t.Fatalf("systemd files left on filesystem: \n%s", out)
	}
}
