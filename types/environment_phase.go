package t

type EnvironmentPhase string

const (
	ENVIRONMENT_PHASE_DEVELOPMENT EnvironmentPhase = "dev"
	ENVIRONMENT_PHASE_TESTING     EnvironmentPhase = "testing"
	ENVIRONMENT_PHASE_STAGING     EnvironmentPhase = "staging"
	ENVIRONMENT_PHASE_PRODUCTION  EnvironmentPhase = "prod"
)

func EnvironmentPhases() []EnvironmentPhase {
	return []EnvironmentPhase{
		ENVIRONMENT_PHASE_DEVELOPMENT,
		ENVIRONMENT_PHASE_TESTING,
		ENVIRONMENT_PHASE_STAGING,
		ENVIRONMENT_PHASE_PRODUCTION,
	}
}

func (e EnvironmentPhase) IsValid() bool {
	for _, env := range EnvironmentPhases() {
		if env == e {
			return true
		}
	}

	return false
}
