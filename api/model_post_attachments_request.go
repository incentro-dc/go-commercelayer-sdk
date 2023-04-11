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

// checks if the POSTAttachmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAttachmentsRequest{}

// POSTAttachmentsRequest struct for POSTAttachmentsRequest
type POSTAttachmentsRequest struct {
	Data POSTAttachmentsRequestData `json:"data"`
}

// NewPOSTAttachmentsRequest instantiates a new POSTAttachmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAttachmentsRequest(data POSTAttachmentsRequestData) *POSTAttachmentsRequest {
	this := POSTAttachmentsRequest{}
	this.Data = data
	return &this
}

// NewPOSTAttachmentsRequestWithDefaults instantiates a new POSTAttachmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAttachmentsRequestWithDefaults() *POSTAttachmentsRequest {
	this := POSTAttachmentsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTAttachmentsRequest) GetData() POSTAttachmentsRequestData {
	if o == nil {
		var ret POSTAttachmentsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTAttachmentsRequest) GetDataOk() (*POSTAttachmentsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTAttachmentsRequest) SetData(v POSTAttachmentsRequestData) {
	o.Data = v
}

func (o POSTAttachmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAttachmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTAttachmentsRequest struct {
	value *POSTAttachmentsRequest
	isSet bool
}

func (v NullablePOSTAttachmentsRequest) Get() *POSTAttachmentsRequest {
	return v.value
}

func (v *NullablePOSTAttachmentsRequest) Set(val *POSTAttachmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAttachmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAttachmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAttachmentsRequest(val *POSTAttachmentsRequest) *NullablePOSTAttachmentsRequest {
	return &NullablePOSTAttachmentsRequest{value: val, isSet: true}
}

func (v NullablePOSTAttachmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAttachmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}