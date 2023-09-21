// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationMetadataActionFieldV1Type - The data type for this value.
type DestinationMetadataActionFieldV1Type string

const (
	DestinationMetadataActionFieldV1TypeBoolean  DestinationMetadataActionFieldV1Type = "BOOLEAN"
	DestinationMetadataActionFieldV1TypeDatetime DestinationMetadataActionFieldV1Type = "DATETIME"
	DestinationMetadataActionFieldV1TypeHidden   DestinationMetadataActionFieldV1Type = "HIDDEN"
	DestinationMetadataActionFieldV1TypeInteger  DestinationMetadataActionFieldV1Type = "INTEGER"
	DestinationMetadataActionFieldV1TypeNumber   DestinationMetadataActionFieldV1Type = "NUMBER"
	DestinationMetadataActionFieldV1TypeObject   DestinationMetadataActionFieldV1Type = "OBJECT"
	DestinationMetadataActionFieldV1TypePassword DestinationMetadataActionFieldV1Type = "PASSWORD"
	DestinationMetadataActionFieldV1TypeString   DestinationMetadataActionFieldV1Type = "STRING"
	DestinationMetadataActionFieldV1TypeText     DestinationMetadataActionFieldV1Type = "TEXT"
)

func (e DestinationMetadataActionFieldV1Type) ToPointer() *DestinationMetadataActionFieldV1Type {
	return &e
}

func (e *DestinationMetadataActionFieldV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BOOLEAN":
		fallthrough
	case "DATETIME":
		fallthrough
	case "HIDDEN":
		fallthrough
	case "INTEGER":
		fallthrough
	case "NUMBER":
		fallthrough
	case "OBJECT":
		fallthrough
	case "PASSWORD":
		fallthrough
	case "STRING":
		fallthrough
	case "TEXT":
		*e = DestinationMetadataActionFieldV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMetadataActionFieldV1Type: %v", v)
	}
}

// DestinationMetadataActionFieldV1 - Represents a field used in configuring an action.
type DestinationMetadataActionFieldV1 struct {
	// Whether this field allows null values.
	AllowNull bool `json:"allowNull"`
	// A list of machine-readable value/label pairs to populate a static dropdown.
	Choices interface{} `json:"choices,omitempty"`
	// A default value that is saved the first time an action is created.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// A human-readable description of this value. You can use Markdown.
	Description string `json:"description"`
	// Whether this field should execute a dynamic request to fetch choices to populate a dropdown. When true, `choices` is ignored.
	Dynamic bool `json:"dynamic"`
	// A unique machine-readable key for the field. Should ideally match the expected key in the action\'s API request.
	FieldKey string `json:"fieldKey"`
	// The primary key of the field.
	ID string `json:"id"`
	// A human-readable label for this value.
	Label string `json:"label"`
	// Whether a user can provide multiples of this field.
	Multiple bool `json:"multiple"`
	// An example value displayed but not saved.
	Placeholder *string `json:"placeholder,omitempty"`
	// Whether this field is required.
	Required bool `json:"required"`
	// The order this particular field is (used in the UI for displaying the fields in a specified order).
	SortOrder float64 `json:"sortOrder"`
	// The data type for this value.
	Type DestinationMetadataActionFieldV1Type `json:"type"`
}
