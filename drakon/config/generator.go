package config

import (
	"github.com/kpango/glg"
	"github.com/odysseia-greek/aristoteles"
	"github.com/odysseia-greek/aristoteles/models"
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

	testOverWrite := config.BoolFromEnv(config.EnvTestOverWrite)
	tls := config.BoolFromEnv(config.EnvTlSKey)

	var cfg models.Config

	if healthCheck {
		vaultConfig, err := config.ConfigFromVault()
		if err != nil {
			glg.Error(err)
			return nil, err
		}

		service := aristoteles.ElasticService(tls)

		cfg = models.Config{
			Service:     service,
			Username:    vaultConfig.ElasticUsername,
			Password:    vaultConfig.ElasticPassword,
			ElasticCERT: vaultConfig.ElasticCERT,
		}
	} else {
		cfg = aristoteles.ElasticConfig(env, testOverWrite, tls)
	}

	elastic, err := aristoteles.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	if healthCheck {
		err := aristoteles.HealthCheck(elastic)
		if err != nil {
			return nil, err
		}
	}

	podName := config.ParsedPodNameFromEnv()
	ns := config.StringFromEnv(config.EnvNamespace, config.DefaultNamespace)
	roles := config.SliceFromEnv(config.EnvRoles)
	indexes := config.SliceFromEnv(config.EnvIndexes)

	return &Config{
		Namespace: ns,
		PodName:   podName,
		Kube:      kube,
		Elastic:   elastic,
		Roles:     roles,
		Indexes:   indexes,
	}, nil
}
