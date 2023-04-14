package config

import (
	kubernetes "github.com/odysseia-greek/thales"
)

type Config struct {
	Namespace string
	Job       string
	Kube      kubernetes.KubeClient
}
