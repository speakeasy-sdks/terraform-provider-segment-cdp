// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReplaceMessagingSubscriptionsInSpacesAlphaInput - Input for the endpoint.
type ReplaceMessagingSubscriptionsInSpacesAlphaInput struct {
	// An array of subscriptions.
	Subscriptions []MessagesSubscriptionRequest `json:"subscriptions"`
}

func (o *ReplaceMessagingSubscriptionsInSpacesAlphaInput) GetSubscriptions() []MessagesSubscriptionRequest {
	if o == nil {
		return []MessagesSubscriptionRequest{}
	}
	return o.Subscriptions
}
