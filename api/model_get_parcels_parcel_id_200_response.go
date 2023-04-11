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

// GETParcelsParcelId200Response struct for GETParcelsParcelId200Response
type GETParcelsParcelId200Response struct {
	Data *GETParcels200ResponseDataInner `json:"data,omitempty"`
}

// NewGETParcelsParcelId200Response instantiates a new GETParcelsParcelId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelsParcelId200Response() *GETParcelsParcelId200Response {
	this := GETParcelsParcelId200Response{}
	return &this
}

// NewGETParcelsParcelId200ResponseWithDefaults instantiates a new GETParcelsParcelId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelsParcelId200ResponseWithDefaults() *GETParcelsParcelId200Response {
	this := GETParcelsParcelId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcelsParcelId200Response) GetData() GETParcels200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETParcels200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelsParcelId200Response) GetDataOk() (*GETParcels200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcelsParcelId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcels200ResponseDataInner and assigns it to the Data field.
func (o *GETParcelsParcelId200Response) SetData(v GETParcels200ResponseDataInner) {
	o.Data = &v
}

func (o GETParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelsParcelId200Response struct {
	value *GETParcelsParcelId200Response
	isSet bool
}

func (v NullableGETParcelsParcelId200Response) Get() *GETParcelsParcelId200Response {
	return v.value
}

func (v *NullableGETParcelsParcelId200Response) Set(val *GETParcelsParcelId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelsParcelId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelsParcelId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelsParcelId200Response(val *GETParcelsParcelId200Response) *NullableGETParcelsParcelId200Response {
	return &NullableGETParcelsParcelId200Response{value: val, isSet: true}
}

func (v NullableGETParcelsParcelId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelsParcelId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
