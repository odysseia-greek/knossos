package config

import (
	"github.com/odysseia-greek/aristoteles"
	kube "github.com/odysseia-greek/thales"
)

type Config struct {
	Namespace string
	PodName   string
	Kube      kube.KubeClient
	Elastic   aristoteles.Client
	Roles     []string
	Indexes   []string
}
