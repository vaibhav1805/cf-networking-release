package server_metrics

//go:generate counterfeiter -generate

import (
	"code.cloudfoundry.org/cf-networking-helpers/metrics"
	"code.cloudfoundry.org/policy-server/store"
)

//counterfeiter:generate -o fakes/list_store.go --fake-name ListStore . listStore
type listStore interface {
	All() ([]store.Policy, error)
}

func NewTotalPoliciesSource(lister listStore) metrics.MetricSource {
	return metrics.MetricSource{
		Name: "totalPolicies",
		Unit: "",
		Getter: func() (float64, error) {
			allPolicies, err := lister.All()
			return float64(len(allPolicies)), err
		},
	}
}
