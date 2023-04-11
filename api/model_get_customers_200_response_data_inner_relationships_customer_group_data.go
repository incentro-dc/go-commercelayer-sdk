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

// GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData struct for GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData
type GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupDataWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerGroupDataWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) SetId(v interface{}) {
	o.Id = v
}

func (o GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData struct {
	value *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) Get() *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) Set(val *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData(val *GETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
