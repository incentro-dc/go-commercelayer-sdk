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

// checks if the PATCHSkusSkuIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSkusSkuIdRequest{}

// PATCHSkusSkuIdRequest struct for PATCHSkusSkuIdRequest
type PATCHSkusSkuIdRequest struct {
	Data PATCHSkusSkuIdRequestData `json:"data"`
}

// NewPATCHSkusSkuIdRequest instantiates a new PATCHSkusSkuIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkusSkuIdRequest(data PATCHSkusSkuIdRequestData) *PATCHSkusSkuIdRequest {
	this := PATCHSkusSkuIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHSkusSkuIdRequestWithDefaults instantiates a new PATCHSkusSkuIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkusSkuIdRequestWithDefaults() *PATCHSkusSkuIdRequest {
	this := PATCHSkusSkuIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHSkusSkuIdRequest) GetData() PATCHSkusSkuIdRequestData {
	if o == nil {
		var ret PATCHSkusSkuIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHSkusSkuIdRequest) GetDataOk() (*PATCHSkusSkuIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHSkusSkuIdRequest) SetData(v PATCHSkusSkuIdRequestData) {
	o.Data = v
}

func (o PATCHSkusSkuIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSkusSkuIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHSkusSkuIdRequest struct {
	value *PATCHSkusSkuIdRequest
	isSet bool
}

func (v NullablePATCHSkusSkuIdRequest) Get() *PATCHSkusSkuIdRequest {
	return v.value
}

func (v *NullablePATCHSkusSkuIdRequest) Set(val *PATCHSkusSkuIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkusSkuIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkusSkuIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkusSkuIdRequest(val *PATCHSkusSkuIdRequest) *NullablePATCHSkusSkuIdRequest {
	return &NullablePATCHSkusSkuIdRequest{value: val, isSet: true}
}

func (v NullablePATCHSkusSkuIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkusSkuIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}