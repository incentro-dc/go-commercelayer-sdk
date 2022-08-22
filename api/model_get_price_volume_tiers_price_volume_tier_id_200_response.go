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

// GETPriceVolumeTiersPriceVolumeTierId200Response struct for GETPriceVolumeTiersPriceVolumeTierId200Response
type GETPriceVolumeTiersPriceVolumeTierId200Response struct {
	Data *GETPriceVolumeTiers200ResponseDataInner `json:"data,omitempty"`
}

// NewGETPriceVolumeTiersPriceVolumeTierId200Response instantiates a new GETPriceVolumeTiersPriceVolumeTierId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceVolumeTiersPriceVolumeTierId200Response() *GETPriceVolumeTiersPriceVolumeTierId200Response {
	this := GETPriceVolumeTiersPriceVolumeTierId200Response{}
	return &this
}

// NewGETPriceVolumeTiersPriceVolumeTierId200ResponseWithDefaults instantiates a new GETPriceVolumeTiersPriceVolumeTierId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceVolumeTiersPriceVolumeTierId200ResponseWithDefaults() *GETPriceVolumeTiersPriceVolumeTierId200Response {
	this := GETPriceVolumeTiersPriceVolumeTierId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPriceVolumeTiersPriceVolumeTierId200Response) GetData() GETPriceVolumeTiers200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETPriceVolumeTiers200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceVolumeTiersPriceVolumeTierId200Response) GetDataOk() (*GETPriceVolumeTiers200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPriceVolumeTiersPriceVolumeTierId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPriceVolumeTiers200ResponseDataInner and assigns it to the Data field.
func (o *GETPriceVolumeTiersPriceVolumeTierId200Response) SetData(v GETPriceVolumeTiers200ResponseDataInner) {
	o.Data = &v
}

func (o GETPriceVolumeTiersPriceVolumeTierId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceVolumeTiersPriceVolumeTierId200Response struct {
	value *GETPriceVolumeTiersPriceVolumeTierId200Response
	isSet bool
}

func (v NullableGETPriceVolumeTiersPriceVolumeTierId200Response) Get() *GETPriceVolumeTiersPriceVolumeTierId200Response {
	return v.value
}

func (v *NullableGETPriceVolumeTiersPriceVolumeTierId200Response) Set(val *GETPriceVolumeTiersPriceVolumeTierId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceVolumeTiersPriceVolumeTierId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceVolumeTiersPriceVolumeTierId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceVolumeTiersPriceVolumeTierId200Response(val *GETPriceVolumeTiersPriceVolumeTierId200Response) *NullableGETPriceVolumeTiersPriceVolumeTierId200Response {
	return &NullableGETPriceVolumeTiersPriceVolumeTierId200Response{value: val, isSet: true}
}

func (v NullableGETPriceVolumeTiersPriceVolumeTierId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceVolumeTiersPriceVolumeTierId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
