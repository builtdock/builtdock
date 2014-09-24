package tests

import (
	"fmt"
	"testing"

	"github.com/builtdock/builtdock/tests/dockercli"
	"github.com/builtdock/builtdock/tests/etcdutils"
	"github.com/builtdock/builtdock/tests/utils"
)

func TestRegistry(t *testing.T) {
	var err error
	setkeys := []string{
		"/builtdock/cache/host",
		"/builtdock/cache/port",
	}
	setdir := []string{
		"/builtdock/cache",
	}
	tag, etcdPort := utils.BuildTag(), utils.RandomPort()
	etcdName := "deis-etcd-" + tag
	cli, stdout, stdoutPipe := dockercli.NewClient()
	dockercli.RunTestEtcd(t, etcdName, etcdPort)
	defer cli.CmdRm("-f", etcdName)
	handler := etcdutils.InitEtcd(setdir, setkeys, etcdPort)
	etcdutils.PublishEtcd(t, handler)
	dockercli.RunDeisDataTest(t, "--name", "deis-registry-data",
		"-v", "/data", "builtdock/base", "/bin/true")
	host, port := utils.HostAddress(), utils.RandomPort()
	fmt.Printf("--- Run builtdock/registry:%s at %s:%s\n", tag, host, port)
	name := "deis-registry-" + tag
	defer cli.CmdRm("-f", name)
	go func() {
		_ = cli.CmdRm("-f", name)
		err = dockercli.RunContainer(cli,
			"--name", name,
			"--rm",
			"-p", port+":5000",
			"-e", "PUBLISH="+port,
			"-e", "HOST="+host,
			"-e", "ETCD_PORT="+etcdPort,
			"--volumes-from", "deis-registry-data",
			"builtdock/registry:"+tag)
	}()
	dockercli.PrintToStdout(t, stdout, stdoutPipe, "Booting")
	if err != nil {
		t.Fatal(err)
	}
	dockercli.DeisServiceTest(t, name, port, "http")
}
