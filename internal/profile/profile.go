package profile

import (
	"os"

	"gopkg.in/yaml.v3"
)

type RuntimeProfile struct {
	APIVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`

	Application struct {
		Name        string `yaml:"name" json:"name"`
		Environment string `yaml:"environment" json:"environment"`
		Owner       string `yaml:"owner" json:"owner"`
	} `yaml:"application" json:"application"`

	Availability struct {
		Tier             string `yaml:"tier" json:"tier"`
		MultiRegion      bool   `yaml:"multi_region" json:"multi_region"`
		FailoverRequired bool   `yaml:"failover_required" json:"failover_required"`
		HealthCheckPath  string `yaml:"health_check_path" json:"health_check_path"`
		RTOMinutes       int    `yaml:"rto_minutes" json:"rto_minutes"`
		RPOMinutes       int    `yaml:"rpo_minutes" json:"rpo_minutes"`
	} `yaml:"availability" json:"availability"`

	Observability struct {
		LogsRequired    bool `yaml:"logs_required" json:"logs_required"`
		MetricsRequired bool `yaml:"metrics_required" json:"metrics_required"`
		TracingRequired bool `yaml:"tracing_required" json:"tracing_required"`
	} `yaml:"observability" json:"observability"`
}

func Load(path string) (*RuntimeProfile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var p RuntimeProfile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}
