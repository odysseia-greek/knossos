package app

import (
	"encoding/json"
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/kpango/glg"
	"github.com/odysseia-greek/plato/aristoteles/configs"
	"github.com/odysseia-greek/plato/models"
	"time"
)

type PeriandrosHandler struct {
	Config   *configs.PeriandrosConfig
	Duration time.Duration
	Timeout  time.Duration
}

func (p *PeriandrosHandler) CreateUser() (bool, error) {
	healthy := p.CheckSolonHealth()
	if !healthy {
		return false, fmt.Errorf("solon not available cannot create user")
	}

	uuid := uuid2.New().String()

	response, err := p.Config.HttpClients.Solon().Register(p.Config.SolonCreationRequest, uuid)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	var solonResponse models.SolonResponse
	err = json.NewDecoder(response.Body).Decode(&solonResponse)
	if err != nil {
		return false, err
	}

	return solonResponse.Created, nil
}

func (p *PeriandrosHandler) CheckSolonHealth() bool {
	healthy := false

	ticker := time.NewTicker(p.Duration)
	timeout := time.After(p.Timeout)

	for {
		select {
		case t := <-ticker.C:
			glg.Infof("tick: %s", t)

			uuid := uuid2.New().String()

			response, err := p.Config.HttpClients.Solon().Health(uuid)
			if err != nil {
				glg.Errorf("Error getting response: %s", err)
				continue
			}

			var solonResponse models.Health
			err = json.NewDecoder(response.Body).Decode(&solonResponse)
			if err != nil {
				continue
			}

			healthy = solonResponse.Healthy
			if !healthy {
				continue
			}
			ticker.Stop()

		case <-timeout:
			ticker.Stop()
		}
		break
	}

	return healthy
}
