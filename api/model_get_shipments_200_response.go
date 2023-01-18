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

// GETShipments200Response struct for GETShipments200Response
type GETShipments200Response struct {
	Data []GETShipments200ResponseDataInner `json:"data,omitempty"`
}

// NewGETShipments200Response instantiates a new GETShipments200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200Response() *GETShipments200Response {
	this := GETShipments200Response{}
	return &this
}

// NewGETShipments200ResponseWithDefaults instantiates a new GETShipments200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseWithDefaults() *GETShipments200Response {
	this := GETShipments200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShipments200Response) GetData() []GETShipments200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETShipments200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200Response) GetDataOk() ([]GETShipments200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShipments200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETShipments200ResponseDataInner and assigns it to the Data field.
func (o *GETShipments200Response) SetData(v []GETShipments200ResponseDataInner) {
	o.Data = v
}

func (o GETShipments200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETShipments200Response struct {
	value *GETShipments200Response
	isSet bool
}

func (v NullableGETShipments200Response) Get() *GETShipments200Response {
	return v.value
}

func (v *NullableGETShipments200Response) Set(val *GETShipments200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200Response(val *GETShipments200Response) *NullableGETShipments200Response {
	return &NullableGETShipments200Response{value: val, isSet: true}
}

func (v NullableGETShipments200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
