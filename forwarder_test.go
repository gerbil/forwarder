package forwarder

import (
	"context"
	"os"

	"fmt"
	"testing"

	"github.com/namsral/flag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func TestBasic(t *testing.T) {
	var kubecfg string
	flag.StringVar(&kubecfg, "kubeconfig", "./kubeconfig", `

	the path of kubeconfig, default is '~/.kube/config'
	you can configure kubeconfig by environment variable: KUBECONFIG=./kubeconfig, 
	or provide a option: --kubeconfig=./kubeconfig

	`)
	flag.Parse()
	fmt.Printf("kubecfg: %v\n", kubecfg)

	options := []*Option{
		{
			LocalPort:   8080,
			RemotePort:  80,
			ServiceName: "my-nginx-svc",
		},
		{
			// LocalPort: 8081,
			// RemotePort:   80,
			Source: "po/my-nginx-66b6c48dd5-ttdb2",
		},
	}

	stream := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	ret, err := WithForwarders(context.Background(), stream, options, kubecfg)
	if err != nil {
		panic(err)
	}
	defer ret.Close()
	ports, err := ret.Ready()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ports: %+v\n", ports)
	ret.Wait()

}
