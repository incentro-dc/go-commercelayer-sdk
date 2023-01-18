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

// PATCHPricesPriceId200Response struct for PATCHPricesPriceId200Response
type PATCHPricesPriceId200Response struct {
	Data *PATCHPricesPriceId200ResponseData `json:"data,omitempty"`
}

// NewPATCHPricesPriceId200Response instantiates a new PATCHPricesPriceId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPricesPriceId200Response() *PATCHPricesPriceId200Response {
	this := PATCHPricesPriceId200Response{}
	return &this
}

// NewPATCHPricesPriceId200ResponseWithDefaults instantiates a new PATCHPricesPriceId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPricesPriceId200ResponseWithDefaults() *PATCHPricesPriceId200Response {
	this := PATCHPricesPriceId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHPricesPriceId200Response) GetData() PATCHPricesPriceId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHPricesPriceId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPricesPriceId200Response) GetDataOk() (*PATCHPricesPriceId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHPricesPriceId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHPricesPriceId200ResponseData and assigns it to the Data field.
func (o *PATCHPricesPriceId200Response) SetData(v PATCHPricesPriceId200ResponseData) {
	o.Data = &v
}

func (o PATCHPricesPriceId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPricesPriceId200Response struct {
	value *PATCHPricesPriceId200Response
	isSet bool
}

func (v NullablePATCHPricesPriceId200Response) Get() *PATCHPricesPriceId200Response {
	return v.value
}

func (v *NullablePATCHPricesPriceId200Response) Set(val *PATCHPricesPriceId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPricesPriceId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPricesPriceId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPricesPriceId200Response(val *PATCHPricesPriceId200Response) *NullablePATCHPricesPriceId200Response {
	return &NullablePATCHPricesPriceId200Response{value: val, isSet: true}
}

func (v NullablePATCHPricesPriceId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPricesPriceId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
