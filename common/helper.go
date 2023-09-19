package common

//The following are helper methods that plugins will need to make a decision.
//Having these in core avoids the plugins from making network calls.
//This helps make the plugins do only business logic.
//This also helps make all the plugins similar, and concentrates the
//common logic into core.

// Helper can be instantiated with varios clients like entitlement, pplicy, db to make teh necessary calls.
// This helper has to be provided to the plugins.
type Helper struct{}

func (h *Helper) OneExists(relation string, partnerId string, id string) (bool, []models.Association, error) {
	return true, []model.Association{}, nil
}

func (h *Helper) HasEntitlement(partnerId string, companyId string, siteId string) (bool, error) {
	return true, nil
}

func (h *Helper) HasPolicy(partnerId string, companyId string, siteId string) (bool, error) {
	return true, nil
}
