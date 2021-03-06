// Package etcdutils helps test interactions with etcd.

package etcdutils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/coreos/go-etcd/etcd"
	"github.com/builtdock/builtdock/tests/utils"
)

// EtcdHandle is used to set keys and values in a test etcd instance.
type EtcdHandle struct {
	Dirs []string
	Keys []string
	C    *etcd.Client
}

func etcdClient(port string) *etcd.Client {
	machines := []string{"http://" + utils.HostAddress() + ":" + port}
	return etcd.NewClient(machines)
}

// InitEtcd configures a test etcd instance.
func InitEtcd(setdir, setkeys []string, port string) *EtcdHandle {
	cli := etcdClient(port)
	controllerHandle := new(EtcdHandle)
	controllerHandle.Dirs = setdir
	controllerHandle.Keys = setkeys
	controllerHandle.C = cli
	fmt.Println("Etcd client initialized")
	return controllerHandle
}

// SetEtcd sets an array of values into a test etcd instance.
func SetEtcd(t *testing.T, keys []string, values []string, c *etcd.Client) {
	for i, key := range keys {
		_, err := c.Set(key, values[i], 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// PublishEtcd sets canonical etcd values into a test etcd instance.
func PublishEtcd(t *testing.T, ecli *EtcdHandle) {
	fmt.Println("--- Publish etcd keys and values")
	for _, dir := range ecli.Dirs {
		_, err := ecli.C.SetDir(dir, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, key := range ecli.Keys {
		switch true {
		case (strings.Contains(key, "host")):
			_, err := ecli.C.Set(key, "172.17.8.100", 0)
			if err != nil {
				t.Fatal(err)
			}
		case (strings.Contains(key, "port")):
			_, err := ecli.C.Set(key, "10881", 0)
			if err != nil {
				t.Fatal(err)
			}
		default:
			_, err := ecli.C.Set(key, "deis", 0)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
