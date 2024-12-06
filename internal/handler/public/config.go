package public

import (
	"healthy/internal/database/models"
	"healthy/internal/database/repositories"
)

func (p *PublicHandler) GetGlobalConfigs() ([]*GlobalConfig, error) {
	conf, err := repositories.ListGlobalConfigs(p.db)
	if err != nil {
		return nil, err
	}
	out := adaptConfig(conf)
	return out, nil
}

type GlobalConfig struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func adaptConfig(input []*models.GlobalConfig) []*GlobalConfig {
	var output = make([]*GlobalConfig, len(input))
	for idx, conf := range input {
		output[idx] = &GlobalConfig{
			Key:   conf.Key,
			Value: conf.Value,
		}
	}
	return output
}
