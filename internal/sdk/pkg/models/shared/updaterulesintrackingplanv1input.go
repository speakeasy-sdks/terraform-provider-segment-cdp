// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateRulesInTrackingPlanV1Input - Updates Tracking Plan rules. Non-existent rules are added.
type UpdateRulesInTrackingPlanV1Input struct {
	// Rules to update or insert.
	Rules []UpsertRuleV1 `json:"rules"`
}