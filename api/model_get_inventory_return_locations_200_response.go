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

// GETInventoryReturnLocations200Response struct for GETInventoryReturnLocations200Response
type GETInventoryReturnLocations200Response struct {
	Data []GETInventoryReturnLocations200ResponseDataInner `json:"data,omitempty"`
}

// NewGETInventoryReturnLocations200Response instantiates a new GETInventoryReturnLocations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryReturnLocations200Response() *GETInventoryReturnLocations200Response {
	this := GETInventoryReturnLocations200Response{}
	return &this
}

// NewGETInventoryReturnLocations200ResponseWithDefaults instantiates a new GETInventoryReturnLocations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryReturnLocations200ResponseWithDefaults() *GETInventoryReturnLocations200Response {
	this := GETInventoryReturnLocations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryReturnLocations200Response) GetData() []GETInventoryReturnLocations200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret []GETInventoryReturnLocations200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryReturnLocations200Response) GetDataOk() ([]GETInventoryReturnLocations200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryReturnLocations200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETInventoryReturnLocations200ResponseDataInner and assigns it to the Data field.
func (o *GETInventoryReturnLocations200Response) SetData(v []GETInventoryReturnLocations200ResponseDataInner) {
	o.Data = v
}

func (o GETInventoryReturnLocations200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryReturnLocations200Response struct {
	value *GETInventoryReturnLocations200Response
	isSet bool
}

func (v NullableGETInventoryReturnLocations200Response) Get() *GETInventoryReturnLocations200Response {
	return v.value
}

func (v *NullableGETInventoryReturnLocations200Response) Set(val *GETInventoryReturnLocations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryReturnLocations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryReturnLocations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryReturnLocations200Response(val *GETInventoryReturnLocations200Response) *NullableGETInventoryReturnLocations200Response {
	return &NullableGETInventoryReturnLocations200Response{value: val, isSet: true}
}

func (v NullableGETInventoryReturnLocations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryReturnLocations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
