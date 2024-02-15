package configurations

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	SunsamaSessionID string `envconfig:"sunsama_session_id" required:"true"`
}

// MustInitializedSpecification returns the configuration Specification
// initialized
func MustInitializedSpecification() Specification {
	s, err := InitializeSpecification()
	if err != nil {
		panic(err)
	}
	return s
}

func InitializeSpecification() (Specification, error) {
	var s Specification
	err := envconfig.Process("", &s)
	if err != nil {
		return Specification{}, fmt.Errorf(
			"failed to initialized specifications, %w",
			err,
		)
	}

	return s, nil
}
