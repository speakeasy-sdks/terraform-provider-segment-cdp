// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReplaceLabelsInSourceAlphaOutput - Response from replacing a list of labels in a Source.
type ReplaceLabelsInSourceAlphaOutput struct {
	// All labels replaced in the Source.
	Labels []LabelAlpha `json:"labels"`
}

func (o *ReplaceLabelsInSourceAlphaOutput) GetLabels() []LabelAlpha {
	if o == nil {
		return []LabelAlpha{}
	}
	return o.Labels
}