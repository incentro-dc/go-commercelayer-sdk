/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData struct for GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData
type GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData{}
	return &this
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceDataWithDefaults instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceDataWithDefaults() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) SetId(v string) {
	o.Id = &v
}

func (o GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData struct {
	value *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData
	isSet bool
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) Get() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData {
	return v.value
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) Set(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData {
	return &NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData{value: val, isSet: true}
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
