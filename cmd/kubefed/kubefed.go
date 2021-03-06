package main

import (
	"os"

	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	_ "k8s.io/kubernetes/pkg/version/prometheus"        // for version metric registration

	"github.com/openshift/origin/pkg/federation/kubefed"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	cmd := kubefed.NewKubeFedCommand(os.Stdin, os.Stdout, os.Stderr)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
