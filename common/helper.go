package common

import "models"

//import "core/models"

//The following are helper methods that plugins will need to make a decision.
//Having these in core avoids the plugins from making network calls.
//This helps make the plugins do only business logic.
//This also helps make all the plugins similar, and concentrates the
//common logic into core.

// Helper can be instantiated with varios clients like entitlement, pplicy, db to make teh necessary calls.
// This helper has to be provided to the plugins.
type Helper struct{}

func (h *Helper) OneExists(relation string, partnerId string, id string) (bool, []models.Association, error) {
	return true, []models.Association{}, nil
}

// keys are the feature names the plugin must be looking for to make a decision.
// If by partner, don't send company and site ids.
func (h *Helper) HasEntitlement(partnerId string, companyId string, siteId string, keys []string) ([]models.EntitlementFeature, error) {
	return []models.EntitlementFeature{}, nil
}

// keys are the policy ids or names whatever we should query on, which the plugin must be lookinf for to make a
// decision. If by partner, don't send company and site ids.
func (h *Helper) HasPolicy(partnerId string, companyId string, siteId string, keys []string) ([]models.Policy, error) {
	return []models.Policy{}, nil
}
