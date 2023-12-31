// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TransformationV1 - Represents a Transformation.
type TransformationV1 struct {
	// The optional Destination metadata associated with the Transformation.
	DestinationMetadataID *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation is enabled.
	Enabled bool `json:"enabled"`
	// Optional array for defining new properties in FQL. Limited to 1 property right now.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
	// The id of the Transformation.
	ID string `json:"id"`
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.
	//
	// For standard event matchers, use the following:
	//   Track -\> "event='\<eventName\>'"
	//   Identify -\> "type='identify'"
	//   Group -\> "type='group'"
	If string `json:"if"`
	// The name of the Transformation.
	Name string `json:"name"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// The Source associated with the Transformation.
	SourceID string `json:"sourceId"`
}

func (o *TransformationV1) GetDestinationMetadataID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationMetadataID
}

func (o *TransformationV1) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *TransformationV1) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil {
		return nil
	}
	return o.FqlDefinedProperties
}

func (o *TransformationV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TransformationV1) GetIf() string {
	if o == nil {
		return ""
	}
	return o.If
}

func (o *TransformationV1) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TransformationV1) GetNewEventName() *string {
	if o == nil {
		return nil
	}
	return o.NewEventName
}

func (o *TransformationV1) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil {
		return nil
	}
	return o.PropertyRenames
}

func (o *TransformationV1) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil {
		return nil
	}
	return o.PropertyValueTransformations
}

func (o *TransformationV1) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}
