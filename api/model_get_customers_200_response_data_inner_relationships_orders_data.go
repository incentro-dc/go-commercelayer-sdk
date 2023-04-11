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

// GETCustomers200ResponseDataInnerRelationshipsOrdersData struct for GETCustomers200ResponseDataInnerRelationshipsOrdersData
type GETCustomers200ResponseDataInnerRelationshipsOrdersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrdersData instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrdersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsOrdersData() *GETCustomers200ResponseDataInnerRelationshipsOrdersData {
	this := GETCustomers200ResponseDataInnerRelationshipsOrdersData{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrdersDataWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrdersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsOrdersDataWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsOrdersData {
	this := GETCustomers200ResponseDataInnerRelationshipsOrdersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrdersData) SetId(v interface{}) {
	o.Id = v
}

func (o GETCustomers200ResponseDataInnerRelationshipsOrdersData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData struct {
	value *GETCustomers200ResponseDataInnerRelationshipsOrdersData
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) Get() *GETCustomers200ResponseDataInnerRelationshipsOrdersData {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) Set(val *GETCustomers200ResponseDataInnerRelationshipsOrdersData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsOrdersData(val *GETCustomers200ResponseDataInnerRelationshipsOrdersData) *NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrdersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
