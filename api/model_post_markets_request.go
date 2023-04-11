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

// checks if the POSTMarketsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarketsRequest{}

// POSTMarketsRequest struct for POSTMarketsRequest
type POSTMarketsRequest struct {
	Data POSTMarketsRequestData `json:"data"`
}

// NewPOSTMarketsRequest instantiates a new POSTMarketsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarketsRequest(data POSTMarketsRequestData) *POSTMarketsRequest {
	this := POSTMarketsRequest{}
	this.Data = data
	return &this
}

// NewPOSTMarketsRequestWithDefaults instantiates a new POSTMarketsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarketsRequestWithDefaults() *POSTMarketsRequest {
	this := POSTMarketsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTMarketsRequest) GetData() POSTMarketsRequestData {
	if o == nil {
		var ret POSTMarketsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequest) GetDataOk() (*POSTMarketsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTMarketsRequest) SetData(v POSTMarketsRequestData) {
	o.Data = v
}

func (o POSTMarketsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarketsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTMarketsRequest struct {
	value *POSTMarketsRequest
	isSet bool
}

func (v NullablePOSTMarketsRequest) Get() *POSTMarketsRequest {
	return v.value
}

func (v *NullablePOSTMarketsRequest) Set(val *POSTMarketsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarketsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarketsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarketsRequest(val *POSTMarketsRequest) *NullablePOSTMarketsRequest {
	return &NullablePOSTMarketsRequest{value: val, isSet: true}
}

func (v NullablePOSTMarketsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarketsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}