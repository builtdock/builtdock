package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/builtdock/builtdock/tests/dockercli"
	"github.com/builtdock/builtdock/tests/etcdutils"
	"github.com/builtdock/builtdock/tests/utils"
)

func TestBuilder(t *testing.T) {
	var err error
	setkeys := []string{
		"/builtdock/registry/protocol",
		"/builtdock/registry/host",
		"/builtdock/registry/port",
		"/builtdock/cache/host",
		"/builtdock/cache/port",
		"/builtdock/controller/protocol",
		"/builtdock/controller/host",
		"/builtdock/controller/port",
		"/builtdock/controller/builderKey",
	}
	setdir := []string{
		"/builtdock/controller",
		"/builtdock/cache",
		"/builtdock/database",
		"/builtdock/registry",
		"/builtdock/domains",
		"/builtdock/services",
	}
	tag, etcdPort := utils.BuildTag(), utils.RandomPort()
	etcdName := "deis-etcd-" + tag
	cli, stdout, stdoutPipe := dockercli.NewClient()
	dockercli.RunTestEtcd(t, etcdName, etcdPort)
	defer cli.CmdRm("-f", etcdName)
	handler := etcdutils.InitEtcd(setdir, setkeys, etcdPort)
	etcdutils.PublishEtcd(t, handler)
	dockercli.RunDeisDataTest(t, "--name", "deis-builder-data",
		"-v", "/var/lib/docker", "builtdock/base", "true")
	ipaddr, port := utils.HostAddress(), utils.RandomPort()
	fmt.Printf("--- Run builtdock/builder:%s at %s:%s\n", tag, ipaddr, port)
	name := "deis-builder-" + tag
	defer cli.CmdRm("-f", name)
	go func() {
		_ = cli.CmdRm("-f", name)
		err = dockercli.RunContainer(cli,
			"--name", name,
			"--rm",
			"-p", port+":22",
			"-e", "PUBLISH=22",
			"-e", "STORAGE_DRIVER=aufs",
			"-e", "HOST="+ipaddr,
			"-e", "ETCD_PORT="+etcdPort,
			"-e", "PORT="+port,
			"--volumes-from", "deis-builder-data",
			"--privileged", "builtdock/builder:"+tag)
	}()
	dockercli.PrintToStdout(t, stdout, stdoutPipe, "deis-builder running")
	if err != nil {
		t.Fatal(err)
	}
	// TODO: builder needs a few seconds to wake up here--fixme!
	time.Sleep(5000 * time.Millisecond)
	dockercli.DeisServiceTest(t, name, port, "tcp")
}
