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

// PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response struct for PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response
type PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response struct {
	Data *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData `json:"data,omitempty"`
}

// NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response() *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	this := PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response{}
	return &this
}

// NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseWithDefaults instantiates a new PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseWithDefaults() *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	this := PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) GetData() PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) GetDataOk() (*PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData and assigns it to the Data field.
func (o *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) SetData(v PATCHExternalTaxCalculatorsExternalTaxCalculatorId200ResponseData) {
	o.Data = &v
}

func (o PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response struct {
	value *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response
	isSet bool
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) Get() *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	return v.value
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) Set(val *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response(val *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response {
	return &NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response{value: val, isSet: true}
}

func (v NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
