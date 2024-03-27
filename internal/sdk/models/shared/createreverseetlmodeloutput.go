// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateReverseEtlModelOutputReverseEtlModel - The created Model.
type CreateReverseEtlModelOutputReverseEtlModel struct {
	// A longer, more descriptive explanation of the Model.
	Description string `json:"description"`
	// Indicates whether the Model should have syncs enabled. When disabled, no
	// syncs will be triggered, regardless of the enabled status of the attached
	// destinations/subscriptions.
	Enabled bool `json:"enabled"`
	// The id of the Model.
	ID string `json:"id"`
	// A short, human-readable description of the Model.
	Name string `json:"name"`
	// The SQL query that will be executed to extract data from the connected
	// Source.
	Query string `json:"query"`
	// Indicates the column named in `query` that should be used to uniquely
	// identify the extracted records.
	QueryIdentifierColumn string `json:"queryIdentifierColumn"`
	// Depending on the chosen strategy, configures the schedule for this model.
	ScheduleConfig map[string]interface{} `json:"scheduleConfig"`
	// Determines the strategy used for triggering syncs, which will be used in
	// conjunction with scheduleConfig.
	//
	// Possible values: "manual", "periodic", "specific_days".
	ScheduleStrategy string `json:"scheduleStrategy"`
	// The id for the attached Source.
	SourceID string `json:"sourceId"`
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetQueryIdentifierColumn() string {
	if o == nil {
		return ""
	}
	return o.QueryIdentifierColumn
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetScheduleConfig() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.ScheduleConfig
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetScheduleStrategy() string {
	if o == nil {
		return ""
	}
	return o.ScheduleStrategy
}

func (o *CreateReverseEtlModelOutputReverseEtlModel) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateReverseEtlModelOutput - Defines the results of creating a Model.
type CreateReverseEtlModelOutput struct {
	// The created Model.
	ReverseEtlModel CreateReverseEtlModelOutputReverseEtlModel `json:"reverseEtlModel"`
}

func (o *CreateReverseEtlModelOutput) GetReverseEtlModel() CreateReverseEtlModelOutputReverseEtlModel {
	if o == nil {
		return CreateReverseEtlModelOutputReverseEtlModel{}
	}
	return o.ReverseEtlModel
}