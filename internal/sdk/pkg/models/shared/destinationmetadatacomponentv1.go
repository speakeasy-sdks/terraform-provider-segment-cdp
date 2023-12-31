// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Owner - The owner of this component. Either 'SEGMENT' or 'PARTNER'.
type Owner string

const (
	OwnerPartner Owner = "PARTNER"
	OwnerSegment Owner = "SEGMENT"
)

func (e Owner) ToPointer() *Owner {
	return &e
}

func (e *Owner) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARTNER":
		fallthrough
	case "SEGMENT":
		*e = Owner(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Owner: %v", v)
	}
}

// DestinationMetadataComponentV1Type - The component type.
type DestinationMetadataComponentV1Type string

const (
	DestinationMetadataComponentV1TypeAndroid DestinationMetadataComponentV1Type = "ANDROID"
	DestinationMetadataComponentV1TypeBrowser DestinationMetadataComponentV1Type = "BROWSER"
	DestinationMetadataComponentV1TypeIos     DestinationMetadataComponentV1Type = "IOS"
	DestinationMetadataComponentV1TypeServer  DestinationMetadataComponentV1Type = "SERVER"
)

func (e DestinationMetadataComponentV1Type) ToPointer() *DestinationMetadataComponentV1Type {
	return &e
}

func (e *DestinationMetadataComponentV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ANDROID":
		fallthrough
	case "BROWSER":
		fallthrough
	case "IOS":
		fallthrough
	case "SERVER":
		*e = DestinationMetadataComponentV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataComponentV1Type: %v", v)
	}
}

// DestinationMetadataComponentV1 - Represents a component this Destination provides.
type DestinationMetadataComponentV1 struct {
	// Link to the repository hosting the code for this component.
	Code string `json:"code"`
	// The owner of this component. Either 'SEGMENT' or 'PARTNER'.
	Owner *Owner `json:"owner,omitempty"`
	// The component type.
	Type DestinationMetadataComponentV1Type `json:"type"`
}

func (o *DestinationMetadataComponentV1) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *DestinationMetadataComponentV1) GetOwner() *Owner {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *DestinationMetadataComponentV1) GetType() DestinationMetadataComponentV1Type {
	if o == nil {
		return DestinationMetadataComponentV1Type("")
	}
	return o.Type
}
