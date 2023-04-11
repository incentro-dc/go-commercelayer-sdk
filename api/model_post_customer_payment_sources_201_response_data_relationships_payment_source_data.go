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

// checks if the POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData{}

// POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData struct for POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData
type POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData instantiates a new POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData {
	this := POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData{}
	return &this
}

// NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceDataWithDefaults instantiates a new POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceDataWithDefaults() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData {
	this := POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData struct {
	value *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData
	isSet bool
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) Get() *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData {
	return v.value
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) Set(val *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData(val *POSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData {
	return &NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData{value: val, isSet: true}
}

func (v NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomerPaymentSources201ResponseDataRelationshipsPaymentSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}