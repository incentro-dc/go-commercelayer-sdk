/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct for GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response
type GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct {
	Data *GETBillingInfoValidationRules200ResponseDataInner `json:"data,omitempty"`
}

// NewGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response instantiates a new GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response() *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	this := GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response{}
	return &this
}

// NewGETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseWithDefaults instantiates a new GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseWithDefaults() *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	this := GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) GetData() GETBillingInfoValidationRules200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETBillingInfoValidationRules200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) GetDataOk() (*GETBillingInfoValidationRules200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETBillingInfoValidationRules200ResponseDataInner and assigns it to the Data field.
func (o *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) SetData(v GETBillingInfoValidationRules200ResponseDataInner) {
	o.Data = &v
}

func (o GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response struct {
	value *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response
	isSet bool
}

func (v NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Get() *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	return v.value
}

func (v *NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Set(val *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response(val *GETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) *NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response {
	return &NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response{value: val, isSet: true}
}

func (v NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBillingInfoValidationRulesBillingInfoValidationRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}