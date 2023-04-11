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

// checks if the POSTFreeShippingPromotionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTFreeShippingPromotionsRequest{}

// POSTFreeShippingPromotionsRequest struct for POSTFreeShippingPromotionsRequest
type POSTFreeShippingPromotionsRequest struct {
	Data POSTFreeShippingPromotionsRequestData `json:"data"`
}

// NewPOSTFreeShippingPromotionsRequest instantiates a new POSTFreeShippingPromotionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTFreeShippingPromotionsRequest(data POSTFreeShippingPromotionsRequestData) *POSTFreeShippingPromotionsRequest {
	this := POSTFreeShippingPromotionsRequest{}
	this.Data = data
	return &this
}

// NewPOSTFreeShippingPromotionsRequestWithDefaults instantiates a new POSTFreeShippingPromotionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTFreeShippingPromotionsRequestWithDefaults() *POSTFreeShippingPromotionsRequest {
	this := POSTFreeShippingPromotionsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTFreeShippingPromotionsRequest) GetData() POSTFreeShippingPromotionsRequestData {
	if o == nil {
		var ret POSTFreeShippingPromotionsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTFreeShippingPromotionsRequest) GetDataOk() (*POSTFreeShippingPromotionsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTFreeShippingPromotionsRequest) SetData(v POSTFreeShippingPromotionsRequestData) {
	o.Data = v
}

func (o POSTFreeShippingPromotionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTFreeShippingPromotionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTFreeShippingPromotionsRequest struct {
	value *POSTFreeShippingPromotionsRequest
	isSet bool
}

func (v NullablePOSTFreeShippingPromotionsRequest) Get() *POSTFreeShippingPromotionsRequest {
	return v.value
}

func (v *NullablePOSTFreeShippingPromotionsRequest) Set(val *POSTFreeShippingPromotionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTFreeShippingPromotionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTFreeShippingPromotionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTFreeShippingPromotionsRequest(val *POSTFreeShippingPromotionsRequest) *NullablePOSTFreeShippingPromotionsRequest {
	return &NullablePOSTFreeShippingPromotionsRequest{value: val, isSet: true}
}

func (v NullablePOSTFreeShippingPromotionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTFreeShippingPromotionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}