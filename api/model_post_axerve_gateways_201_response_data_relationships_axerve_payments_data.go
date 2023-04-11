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

// checks if the POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData{}

// POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData struct for POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData
type POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData instantiates a new POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData() *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData {
	this := POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData{}
	return &this
}

// NewPOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsDataWithDefaults instantiates a new POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsDataWithDefaults() *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData {
	this := POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData struct {
	value *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData
	isSet bool
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) Get() *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData {
	return v.value
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) Set(val *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData(val *POSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) *NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData {
	return &NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData{value: val, isSet: true}
}

func (v NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAxerveGateways201ResponseDataRelationshipsAxervePaymentsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}