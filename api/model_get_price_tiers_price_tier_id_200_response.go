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

// GETPriceTiersPriceTierId200Response struct for GETPriceTiersPriceTierId200Response
type GETPriceTiersPriceTierId200Response struct {
	Data *GETPriceTiers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPriceTiersPriceTierId200Response instantiates a new GETPriceTiersPriceTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceTiersPriceTierId200Response() *GETPriceTiersPriceTierId200Response {
	this := GETPriceTiersPriceTierId200Response{}
	return &this
}

// NewGETPriceTiersPriceTierId200ResponseWithDefaults instantiates a new GETPriceTiersPriceTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceTiersPriceTierId200ResponseWithDefaults() *GETPriceTiersPriceTierId200Response {
	this := GETPriceTiersPriceTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPriceTiersPriceTierId200Response) GetData() GETPriceTiers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETPriceTiers200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceTiersPriceTierId200Response) GetDataOk() (*GETPriceTiers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPriceTiersPriceTierId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPriceTiers200ResponseDataInner and assigns it to the Data field.
func (o *GETPriceTiersPriceTierId200Response) SetData(v GETPriceTiers200ResponseDataInner) {
	o.Data = &v
}

func (o GETPriceTiersPriceTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceTiersPriceTierId200Response struct {
	value *GETPriceTiersPriceTierId200Response
	isSet bool
}

func (v NullableGETPriceTiersPriceTierId200Response) Get() *GETPriceTiersPriceTierId200Response {
	return v.value
}

func (v *NullableGETPriceTiersPriceTierId200Response) Set(val *GETPriceTiersPriceTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceTiersPriceTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceTiersPriceTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceTiersPriceTierId200Response(val *GETPriceTiersPriceTierId200Response) *NullableGETPriceTiersPriceTierId200Response {
	return &NullableGETPriceTiersPriceTierId200Response{value: val, isSet: true}
}

func (v NullableGETPriceTiersPriceTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceTiersPriceTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}