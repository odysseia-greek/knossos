package config

import (
	"github.com/odysseia-greek/plato/config"
	kubernetes "github.com/odysseia-greek/thales"
)

func CreateNewConfig(env string) (*Config, error) {
	healthCheck := true
	if env == "LOCAL" || env == "TEST" {
		healthCheck = false
	}

	kube, err := kubernetes.CreateKubeClient(healthCheck)
	if err != nil {
		return nil, err
	}

	ns := config.StringFromEnv(config.EnvNamespace, config.DefaultNamespace)

	service, err := config.CreateOdysseiaClient()
	if err != nil {
		return nil, err
	}

	solonRequest := config.InitCreation()

	return &Config{
		Namespace:            ns,
		HttpClients:          service,
		SolonCreationRequest: solonRequest,
		Kube:                 kube,
	}, nil
}
