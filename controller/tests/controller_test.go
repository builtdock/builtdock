package tests

import (
	"fmt"
	"testing"

	"github.com/builtdock/builtdock/tests/dockercli"
	"github.com/builtdock/builtdock/tests/etcdutils"
	"github.com/builtdock/builtdock/tests/mock"
	"github.com/builtdock/builtdock/tests/utils"
)

func TestController(t *testing.T) {
	var err error
	setkeys := []string{
		"/builtdock/registry/protocol",
		"builtdock/registry/host",
		"/builtdock/registry/port",
		"/builtdock/cache/host",
		"/builtdock/cache/port",
	}
	setdir := []string{
		"/builtdock/controller",
		"/builtdock/cache",
		"/builtdock/database",
		"/builtdock/registry",
		"/builtdock/domains",
	}
	tag, etcdPort := utils.BuildTag(), utils.RandomPort()
	etcdName := "deis-etcd-" + tag
	cli, stdout, stdoutPipe := dockercli.NewClient()
	dockercli.RunTestEtcd(t, etcdName, etcdPort)
	defer cli.CmdRm("-f", etcdName)
	handler := etcdutils.InitEtcd(setdir, setkeys, etcdPort)
	etcdutils.PublishEtcd(t, handler)
	mock.RunMockDatabase(t, tag, etcdPort, utils.RandomPort())
	defer cli.CmdRm("-f", "deis-test-database-"+tag)
	host, port := utils.HostAddress(), utils.RandomPort()
	fmt.Printf("--- Run builtdock/controller:%s at %s:%s\n", tag, host, port)
	name := "deis-controller-" + tag
	defer cli.CmdRm("-f", name)
	go func() {
		cli.CmdRm("-f", name)
		err = dockercli.RunContainer(cli,
			"--name", name,
			"--rm",
			"-p", port+":8000",
			"-e", "PUBLISH="+port,
			"-e", "HOST="+host,
			"-e", "ETCD_PORT="+etcdPort,
			"builtdock/controller:"+tag)
	}()
	dockercli.PrintToStdout(t, stdout, stdoutPipe, "Booting")
	if err != nil {
		t.Fatal(err)
	}
	dockercli.DeisServiceTest(t, name, port, "http")
}
