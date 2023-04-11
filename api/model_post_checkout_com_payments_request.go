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

// checks if the POSTCheckoutComPaymentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCheckoutComPaymentsRequest{}

// POSTCheckoutComPaymentsRequest struct for POSTCheckoutComPaymentsRequest
type POSTCheckoutComPaymentsRequest struct {
	Data POSTCheckoutComPaymentsRequestData `json:"data"`
}

// NewPOSTCheckoutComPaymentsRequest instantiates a new POSTCheckoutComPaymentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCheckoutComPaymentsRequest(data POSTCheckoutComPaymentsRequestData) *POSTCheckoutComPaymentsRequest {
	this := POSTCheckoutComPaymentsRequest{}
	this.Data = data
	return &this
}

// NewPOSTCheckoutComPaymentsRequestWithDefaults instantiates a new POSTCheckoutComPaymentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCheckoutComPaymentsRequestWithDefaults() *POSTCheckoutComPaymentsRequest {
	this := POSTCheckoutComPaymentsRequest{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCheckoutComPaymentsRequest) GetData() POSTCheckoutComPaymentsRequestData {
	if o == nil {
		var ret POSTCheckoutComPaymentsRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCheckoutComPaymentsRequest) GetDataOk() (*POSTCheckoutComPaymentsRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCheckoutComPaymentsRequest) SetData(v POSTCheckoutComPaymentsRequestData) {
	o.Data = v
}

func (o POSTCheckoutComPaymentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCheckoutComPaymentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCheckoutComPaymentsRequest struct {
	value *POSTCheckoutComPaymentsRequest
	isSet bool
}

func (v NullablePOSTCheckoutComPaymentsRequest) Get() *POSTCheckoutComPaymentsRequest {
	return v.value
}

func (v *NullablePOSTCheckoutComPaymentsRequest) Set(val *POSTCheckoutComPaymentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCheckoutComPaymentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCheckoutComPaymentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCheckoutComPaymentsRequest(val *POSTCheckoutComPaymentsRequest) *NullablePOSTCheckoutComPaymentsRequest {
	return &NullablePOSTCheckoutComPaymentsRequest{value: val, isSet: true}
}

func (v NullablePOSTCheckoutComPaymentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCheckoutComPaymentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}