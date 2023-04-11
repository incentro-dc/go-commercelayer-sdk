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

// PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct for PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response
type PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct {
	Data *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData `json:"data,omitempty"`
}

// NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response instantiates a new PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	this := PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response{}
	return &this
}

// NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseWithDefaults instantiates a new PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseWithDefaults() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	this := PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) GetData() PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) GetDataOk() (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData and assigns it to the Data field.
func (o *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) SetData(v PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseData) {
	o.Data = &v
}

func (o PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct {
	value *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response
	isSet bool
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Get() *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	return v.value
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Set(val *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response(val *PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	return &NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response{value: val, isSet: true}
}

func (v NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHBillingInfoValidationRulesBillingInfoValidationRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
