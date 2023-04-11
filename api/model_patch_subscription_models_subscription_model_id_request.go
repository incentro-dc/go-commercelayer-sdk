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

// checks if the PATCHSubscriptionModelsSubscriptionModelIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSubscriptionModelsSubscriptionModelIdRequest{}

// PATCHSubscriptionModelsSubscriptionModelIdRequest struct for PATCHSubscriptionModelsSubscriptionModelIdRequest
type PATCHSubscriptionModelsSubscriptionModelIdRequest struct {
	Data PATCHSubscriptionModelsSubscriptionModelIdRequestData `json:"data"`
}

// NewPATCHSubscriptionModelsSubscriptionModelIdRequest instantiates a new PATCHSubscriptionModelsSubscriptionModelIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSubscriptionModelsSubscriptionModelIdRequest(data PATCHSubscriptionModelsSubscriptionModelIdRequestData) *PATCHSubscriptionModelsSubscriptionModelIdRequest {
	this := PATCHSubscriptionModelsSubscriptionModelIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHSubscriptionModelsSubscriptionModelIdRequestWithDefaults instantiates a new PATCHSubscriptionModelsSubscriptionModelIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSubscriptionModelsSubscriptionModelIdRequestWithDefaults() *PATCHSubscriptionModelsSubscriptionModelIdRequest {
	this := PATCHSubscriptionModelsSubscriptionModelIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHSubscriptionModelsSubscriptionModelIdRequest) GetData() PATCHSubscriptionModelsSubscriptionModelIdRequestData {
	if o == nil {
		var ret PATCHSubscriptionModelsSubscriptionModelIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHSubscriptionModelsSubscriptionModelIdRequest) GetDataOk() (*PATCHSubscriptionModelsSubscriptionModelIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHSubscriptionModelsSubscriptionModelIdRequest) SetData(v PATCHSubscriptionModelsSubscriptionModelIdRequestData) {
	o.Data = v
}

func (o PATCHSubscriptionModelsSubscriptionModelIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSubscriptionModelsSubscriptionModelIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHSubscriptionModelsSubscriptionModelIdRequest struct {
	value *PATCHSubscriptionModelsSubscriptionModelIdRequest
	isSet bool
}

func (v NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) Get() *PATCHSubscriptionModelsSubscriptionModelIdRequest {
	return v.value
}

func (v *NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) Set(val *PATCHSubscriptionModelsSubscriptionModelIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSubscriptionModelsSubscriptionModelIdRequest(val *PATCHSubscriptionModelsSubscriptionModelIdRequest) *NullablePATCHSubscriptionModelsSubscriptionModelIdRequest {
	return &NullablePATCHSubscriptionModelsSubscriptionModelIdRequest{value: val, isSet: true}
}

func (v NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSubscriptionModelsSubscriptionModelIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}