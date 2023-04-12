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

// checks if the POSTOrders201ResponseDataRelationshipsTransactionsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsTransactionsData{}

// POSTOrders201ResponseDataRelationshipsTransactionsData struct for POSTOrders201ResponseDataRelationshipsTransactionsData
type POSTOrders201ResponseDataRelationshipsTransactionsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsTransactionsData instantiates a new POSTOrders201ResponseDataRelationshipsTransactionsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsTransactionsData() *POSTOrders201ResponseDataRelationshipsTransactionsData {
	this := POSTOrders201ResponseDataRelationshipsTransactionsData{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsTransactionsDataWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsTransactionsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsTransactionsDataWithDefaults() *POSTOrders201ResponseDataRelationshipsTransactionsData {
	this := POSTOrders201ResponseDataRelationshipsTransactionsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrders201ResponseDataRelationshipsTransactionsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrders201ResponseDataRelationshipsTransactionsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsTransactionsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsTransactionsData struct {
	value *POSTOrders201ResponseDataRelationshipsTransactionsData
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) Get() *POSTOrders201ResponseDataRelationshipsTransactionsData {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) Set(val *POSTOrders201ResponseDataRelationshipsTransactionsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsTransactionsData(val *POSTOrders201ResponseDataRelationshipsTransactionsData) *NullablePOSTOrders201ResponseDataRelationshipsTransactionsData {
	return &NullablePOSTOrders201ResponseDataRelationshipsTransactionsData{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsTransactionsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
