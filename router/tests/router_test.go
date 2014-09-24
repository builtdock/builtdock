package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/builtdock/builtdock/tests/dockercli"
	"github.com/builtdock/builtdock/tests/etcdutils"
	"github.com/builtdock/builtdock/tests/utils"
)

func TestRouter(t *testing.T) {
	var err error
	setkeys := []string{
		"builtdock/controller/host",
		"/builtdock/controller/port",
		"/builtdock/builder/host",
		"/builtdock/builder/port",
	}
	setdir := []string{
		"/builtdock/controller",
		"/builtdock/router",
		"/builtdock/database",
		"/builtdock/services",
		"/builtdock/builder",
		"/builtdock/domains",
	}
	tag, etcdPort := utils.BuildTag(), utils.RandomPort()
	etcdName := "deis-etcd-" + tag
	cli, stdout, stdoutPipe := dockercli.NewClient()
	dockercli.RunTestEtcd(t, etcdName, etcdPort)
	defer cli.CmdRm("-f", etcdName)
	handler := etcdutils.InitEtcd(setdir, setkeys, etcdPort)
	etcdutils.PublishEtcd(t, handler)
	host, port := utils.HostAddress(), utils.RandomPort()
	fmt.Printf("--- Run builtdock/router:%s at %s:%s\n", tag, host, port)
	name := "deis-router-" + tag
	go func() {
		_ = cli.CmdRm("-f", name)
		err = dockercli.RunContainer(cli,
			"--name", name,
			"--rm",
			"-p", port+":80",
			"-p", utils.RandomPort()+":2222",
			"-e", "PUBLISH="+port,
			"-e", "HOST="+host,
			"-e", "ETCD_PORT="+etcdPort,
			"builtdock/router:"+tag)
	}()
	dockercli.PrintToStdout(t, stdout, stdoutPipe, "deis-router running")
	if err != nil {
		t.Fatal(err)
	}
	// FIXME: nginx needs a couple seconds to wake up here
	time.Sleep(2000 * time.Millisecond)
	dockercli.DeisServiceTest(t, name, port, "http")
	_ = cli.CmdRm("-f", name)
}
