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

// PATCHSkuOptionsSkuOptionId200Response struct for PATCHSkuOptionsSkuOptionId200Response
type PATCHSkuOptionsSkuOptionId200Response struct {
	Data *PATCHSkuOptionsSkuOptionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHSkuOptionsSkuOptionId200Response instantiates a new PATCHSkuOptionsSkuOptionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuOptionsSkuOptionId200Response() *PATCHSkuOptionsSkuOptionId200Response {
	this := PATCHSkuOptionsSkuOptionId200Response{}
	return &this
}

// NewPATCHSkuOptionsSkuOptionId200ResponseWithDefaults instantiates a new PATCHSkuOptionsSkuOptionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuOptionsSkuOptionId200ResponseWithDefaults() *PATCHSkuOptionsSkuOptionId200Response {
	this := PATCHSkuOptionsSkuOptionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHSkuOptionsSkuOptionId200Response) GetData() PATCHSkuOptionsSkuOptionId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHSkuOptionsSkuOptionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuOptionsSkuOptionId200Response) GetDataOk() (*PATCHSkuOptionsSkuOptionId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHSkuOptionsSkuOptionId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHSkuOptionsSkuOptionId200ResponseData and assigns it to the Data field.
func (o *PATCHSkuOptionsSkuOptionId200Response) SetData(v PATCHSkuOptionsSkuOptionId200ResponseData) {
	o.Data = &v
}

func (o PATCHSkuOptionsSkuOptionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHSkuOptionsSkuOptionId200Response struct {
	value *PATCHSkuOptionsSkuOptionId200Response
	isSet bool
}

func (v NullablePATCHSkuOptionsSkuOptionId200Response) Get() *PATCHSkuOptionsSkuOptionId200Response {
	return v.value
}

func (v *NullablePATCHSkuOptionsSkuOptionId200Response) Set(val *PATCHSkuOptionsSkuOptionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuOptionsSkuOptionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuOptionsSkuOptionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuOptionsSkuOptionId200Response(val *PATCHSkuOptionsSkuOptionId200Response) *NullablePATCHSkuOptionsSkuOptionId200Response {
	return &NullablePATCHSkuOptionsSkuOptionId200Response{value: val, isSet: true}
}

func (v NullablePATCHSkuOptionsSkuOptionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuOptionsSkuOptionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
