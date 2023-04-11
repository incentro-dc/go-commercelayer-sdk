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

// checks if the PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{}

// PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct for PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest
type PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	Data PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData `json:"data"`
}

// NewPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest instantiates a new PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest(data PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData) *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	this := PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{}
	this.Data = data
	return &this
}

// NewPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestWithDefaults instantiates a new PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestWithDefaults() *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	this := PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) GetData() PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData {
	if o == nil {
		var ret PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) GetDataOk() (*PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) SetData(v PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequestData) {
	o.Data = v
}

func (o PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest struct {
	value *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest
	isSet bool
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Get() *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return v.value
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Set(val *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest(val *PATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest {
	return &NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest{value: val, isSet: true}
}

func (v NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCustomerPaymentSourcesCustomerPaymentSourceIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}