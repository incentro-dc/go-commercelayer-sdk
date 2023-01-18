/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPromotionDataRelationshipsOrderAmountPromotionRule struct for ExternalPromotionDataRelationshipsOrderAmountPromotionRule
type ExternalPromotionDataRelationshipsOrderAmountPromotionRule struct {
	Data *ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData `json:"data,omitempty"`
}

// NewExternalPromotionDataRelationshipsOrderAmountPromotionRule instantiates a new ExternalPromotionDataRelationshipsOrderAmountPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionDataRelationshipsOrderAmountPromotionRule() *ExternalPromotionDataRelationshipsOrderAmountPromotionRule {
	this := ExternalPromotionDataRelationshipsOrderAmountPromotionRule{}
	return &this
}

// NewExternalPromotionDataRelationshipsOrderAmountPromotionRuleWithDefaults instantiates a new ExternalPromotionDataRelationshipsOrderAmountPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionDataRelationshipsOrderAmountPromotionRuleWithDefaults() *ExternalPromotionDataRelationshipsOrderAmountPromotionRule {
	this := ExternalPromotionDataRelationshipsOrderAmountPromotionRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) GetData() ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData {
	if o == nil || o.Data == nil {
		var ret ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) GetDataOk() (*ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData and assigns it to the Data field.
func (o *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) SetData(v ExternalPromotionDataRelationshipsOrderAmountPromotionRuleData) {
	o.Data = &v
}

func (o ExternalPromotionDataRelationshipsOrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule struct {
	value *ExternalPromotionDataRelationshipsOrderAmountPromotionRule
	isSet bool
}

func (v NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) Get() *ExternalPromotionDataRelationshipsOrderAmountPromotionRule {
	return v.value
}

func (v *NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) Set(val *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionDataRelationshipsOrderAmountPromotionRule(val *ExternalPromotionDataRelationshipsOrderAmountPromotionRule) *NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule {
	return &NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule{value: val, isSet: true}
}

func (v NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionDataRelationshipsOrderAmountPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
