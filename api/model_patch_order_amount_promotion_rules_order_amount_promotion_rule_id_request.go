/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest{}

// PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct for PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest
type PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct {
	Data PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData `json:"data"`
}

// NewPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest instantiates a new PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest(data PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData) *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest {
	this := PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestWithDefaults instantiates a new PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestWithDefaults() *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest {
	this := PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) GetData() PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData {
	if o == nil {
		var ret PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) GetDataOk() (*PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) SetData(v PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequestData) {
	o.Data = v
}

func (o PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest struct {
	value *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest
	isSet bool
}

func (v NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) Get() *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest {
	return v.value
}

func (v *NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) Set(val *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest(val *PATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) *NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest {
	return &NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest{value: val, isSet: true}
}

func (v NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHOrderAmountPromotionRulesOrderAmountPromotionRuleIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}