package config

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

// KubeConfig gets the current kubernetes config
func KubeConfig() (string, error) {
	// Setup kubeconfig
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = *flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = *flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	return kubeconfig, nil
}
