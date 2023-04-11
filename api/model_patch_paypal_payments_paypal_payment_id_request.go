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

// checks if the PATCHPaypalPaymentsPaypalPaymentIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPaypalPaymentsPaypalPaymentIdRequest{}

// PATCHPaypalPaymentsPaypalPaymentIdRequest struct for PATCHPaypalPaymentsPaypalPaymentIdRequest
type PATCHPaypalPaymentsPaypalPaymentIdRequest struct {
	Data PATCHPaypalPaymentsPaypalPaymentIdRequestData `json:"data"`
}

// NewPATCHPaypalPaymentsPaypalPaymentIdRequest instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaypalPaymentsPaypalPaymentIdRequest(data PATCHPaypalPaymentsPaypalPaymentIdRequestData) *PATCHPaypalPaymentsPaypalPaymentIdRequest {
	this := PATCHPaypalPaymentsPaypalPaymentIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHPaypalPaymentsPaypalPaymentIdRequestWithDefaults instantiates a new PATCHPaypalPaymentsPaypalPaymentIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaypalPaymentsPaypalPaymentIdRequestWithDefaults() *PATCHPaypalPaymentsPaypalPaymentIdRequest {
	this := PATCHPaypalPaymentsPaypalPaymentIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequest) GetData() PATCHPaypalPaymentsPaypalPaymentIdRequestData {
	if o == nil {
		var ret PATCHPaypalPaymentsPaypalPaymentIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequest) GetDataOk() (*PATCHPaypalPaymentsPaypalPaymentIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHPaypalPaymentsPaypalPaymentIdRequest) SetData(v PATCHPaypalPaymentsPaypalPaymentIdRequestData) {
	o.Data = v
}

func (o PATCHPaypalPaymentsPaypalPaymentIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPaypalPaymentsPaypalPaymentIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHPaypalPaymentsPaypalPaymentIdRequest struct {
	value *PATCHPaypalPaymentsPaypalPaymentIdRequest
	isSet bool
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) Get() *PATCHPaypalPaymentsPaypalPaymentIdRequest {
	return v.value
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) Set(val *PATCHPaypalPaymentsPaypalPaymentIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaypalPaymentsPaypalPaymentIdRequest(val *PATCHPaypalPaymentsPaypalPaymentIdRequest) *NullablePATCHPaypalPaymentsPaypalPaymentIdRequest {
	return &NullablePATCHPaypalPaymentsPaypalPaymentIdRequest{value: val, isSet: true}
}

func (v NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaypalPaymentsPaypalPaymentIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}