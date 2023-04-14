package config

import (
	"github.com/odysseia-greek/plato/models"
	"github.com/odysseia-greek/plato/service"
	kubernetes "github.com/odysseia-greek/thales"
)

type Config struct {
	Namespace            string
	HttpClients          service.OdysseiaClient
	SolonCreationRequest models.SolonCreationRequest
	Kube                 kubernetes.KubeClient
}
