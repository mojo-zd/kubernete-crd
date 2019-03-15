package client

import (
	"testing"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/golang/glog"
	"flag"
	"gitee.com/wisecloud/crd/pkg/client/clientset/versioned"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kr/pretty"
)

var (
	kuberconfig = flag.String("kubeconfig", "/Users/mojo/.kube/config", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func Test_Client(t *testing.T) {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	crdClient, err := versioned.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	list, err := crdClient.Wise2cV1().Tenants().List(v1.ListOptions{})
	pretty.Println("----", list)
	if err != nil {
		glog.Fatalf("Error listing all databases: %v", err)
	}

	for _, db := range list.Items {
		rl := db.Spec.Quota.Hard["limits.memory"]
		cpu := db.Spec.Quota.Hard["limits.cpu"]
		pretty.Println("memory:",rl.Value())
		pretty.Println(cpu)
	}
}
