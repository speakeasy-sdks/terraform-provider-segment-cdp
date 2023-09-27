// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetFunctionV1OutputFunctionV1ResourceType - The Function type.
//
// Config API note: equal to `type`.
type GetFunctionV1OutputFunctionV1ResourceType string

const (
	GetFunctionV1OutputFunctionV1ResourceTypeDestination       GetFunctionV1OutputFunctionV1ResourceType = "DESTINATION"
	GetFunctionV1OutputFunctionV1ResourceTypeInsertDestination GetFunctionV1OutputFunctionV1ResourceType = "INSERT_DESTINATION"
	GetFunctionV1OutputFunctionV1ResourceTypeSource            GetFunctionV1OutputFunctionV1ResourceType = "SOURCE"
)

func (e GetFunctionV1OutputFunctionV1ResourceType) ToPointer() *GetFunctionV1OutputFunctionV1ResourceType {
	return &e
}

func (e *GetFunctionV1OutputFunctionV1ResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DESTINATION":
		fallthrough
	case "INSERT_DESTINATION":
		fallthrough
	case "SOURCE":
		*e = GetFunctionV1OutputFunctionV1ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetFunctionV1OutputFunctionV1ResourceType: %v", v)
	}
}

// GetFunctionV1OutputFunctionV1 - Represents a Function.
type GetFunctionV1OutputFunctionV1 struct {
	// The max count of the batch for this Function.
	BatchMaxCount *float64 `json:"batchMaxCount,omitempty"`
	// The catalog id of this Function.
	CatalogID *string `json:"catalogId,omitempty"`
	// The Function code.
	Code *string `json:"code,omitempty"`
	// The time this Function was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The id of the user who created this Function.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The time of this Function's last deployment.
	DeployedAt *string `json:"deployedAt,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// An identifier for this Function.
	ID *string `json:"id,omitempty"`
	// Whether the deployment of this Function is the latest version.
	IsLatestVersion *bool `json:"isLatestVersion,omitempty"`
	// The URL of the logo for this Function.
	LogoURL *string `json:"logoUrl,omitempty"`
	// The preview webhook URL for this Function.
	PreviewWebhookURL *string `json:"previewWebhookUrl,omitempty"`
	// The Function type.
	//
	// Config API note: equal to `type`.
	ResourceType *GetFunctionV1OutputFunctionV1ResourceType `json:"resourceType,omitempty"`
	// The list of settings for this Function.
	Settings []FunctionSettingV1 `json:"settings,omitempty"`
}

// GetFunctionV1Output - Gets a single Function.
type GetFunctionV1Output struct {
	// A Function object.
	Function *GetFunctionV1OutputFunctionV1 `json:"function"`
}