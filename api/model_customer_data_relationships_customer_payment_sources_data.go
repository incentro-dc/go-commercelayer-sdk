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

// CustomerDataRelationshipsCustomerPaymentSourcesData struct for CustomerDataRelationshipsCustomerPaymentSourcesData
type CustomerDataRelationshipsCustomerPaymentSourcesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource's id
	Id interface{} `json:"id,omitempty"`
}

// NewCustomerDataRelationshipsCustomerPaymentSourcesData instantiates a new CustomerDataRelationshipsCustomerPaymentSourcesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerDataRelationshipsCustomerPaymentSourcesData() *CustomerDataRelationshipsCustomerPaymentSourcesData {
	this := CustomerDataRelationshipsCustomerPaymentSourcesData{}
	return &this
}

// NewCustomerDataRelationshipsCustomerPaymentSourcesDataWithDefaults instantiates a new CustomerDataRelationshipsCustomerPaymentSourcesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerDataRelationshipsCustomerPaymentSourcesDataWithDefaults() *CustomerDataRelationshipsCustomerPaymentSourcesData {
	this := CustomerDataRelationshipsCustomerPaymentSourcesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *CustomerDataRelationshipsCustomerPaymentSourcesData) SetId(v interface{}) {
	o.Id = v
}

func (o CustomerDataRelationshipsCustomerPaymentSourcesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerDataRelationshipsCustomerPaymentSourcesData struct {
	value *CustomerDataRelationshipsCustomerPaymentSourcesData
	isSet bool
}

func (v NullableCustomerDataRelationshipsCustomerPaymentSourcesData) Get() *CustomerDataRelationshipsCustomerPaymentSourcesData {
	return v.value
}

func (v *NullableCustomerDataRelationshipsCustomerPaymentSourcesData) Set(val *CustomerDataRelationshipsCustomerPaymentSourcesData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerDataRelationshipsCustomerPaymentSourcesData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerDataRelationshipsCustomerPaymentSourcesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerDataRelationshipsCustomerPaymentSourcesData(val *CustomerDataRelationshipsCustomerPaymentSourcesData) *NullableCustomerDataRelationshipsCustomerPaymentSourcesData {
	return &NullableCustomerDataRelationshipsCustomerPaymentSourcesData{value: val, isSet: true}
}

func (v NullableCustomerDataRelationshipsCustomerPaymentSourcesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerDataRelationshipsCustomerPaymentSourcesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
