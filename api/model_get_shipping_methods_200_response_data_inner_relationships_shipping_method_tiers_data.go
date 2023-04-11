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

// GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData struct for GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData
type GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData() *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData{}
	return &this
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersDataWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersDataWithDefaults() *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData struct {
	value *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData
	isSet bool
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) Get() *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData {
	return v.value
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) Set(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData {
	return &NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData{value: val, isSet: true}
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingMethodTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
