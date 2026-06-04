package profile

func Validate(p *RuntimeProfile) []string {
	violations := []string{}

	if p.APIVersion == "" {
		violations = append(violations, "apiVersion is required")
	}

	if p.Kind != "RuntimeProfile" {
		violations = append(violations, "kind must equal RuntimeProfile")
	}

	if p.Application.Name == "" {
		violations = append(violations, "application.name is required")
	}

	if p.Application.Environment == "" {
		violations = append(violations, "application.environment is required")
	}

	if p.Application.Owner == "" {
		violations = append(violations, "application.owner is required")
	}

	if p.Availability.Tier == "critical" {
		if !p.Availability.FailoverRequired {
			violations = append(violations, "availability.failover_required must be true when availability.tier is critical")
		}

		if p.Availability.HealthCheckPath == "" {
			violations = append(violations, "availability.health_check_path is required when availability.tier is critical")
		}

		if !p.Observability.LogsRequired {
			violations = append(violations, "observability.logs_required must be true when availability.tier is critical")
		}

		if !p.Observability.MetricsRequired {
			violations = append(violations, "observability.metrics_required must be true when availability.tier is critical")
		}
	}

	return violations
}
