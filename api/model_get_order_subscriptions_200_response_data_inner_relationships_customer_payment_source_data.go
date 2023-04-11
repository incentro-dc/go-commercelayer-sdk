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

// GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData struct for GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData
type GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData() *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData {
	this := GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData{}
	return &this
}

// NewGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceDataWithDefaults instantiates a new GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceDataWithDefaults() *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData {
	this := GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) SetId(v interface{}) {
	o.Id = v
}

func (o GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData struct {
	value *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData
	isSet bool
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) Get() *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData {
	return v.value
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) Set(val *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData(val *GETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData {
	return &NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData{value: val, isSet: true}
}

func (v NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderSubscriptions200ResponseDataInnerRelationshipsCustomerPaymentSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
