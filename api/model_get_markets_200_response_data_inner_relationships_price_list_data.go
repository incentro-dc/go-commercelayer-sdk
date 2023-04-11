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

// GETMarkets200ResponseDataInnerRelationshipsPriceListData struct for GETMarkets200ResponseDataInnerRelationshipsPriceListData
type GETMarkets200ResponseDataInnerRelationshipsPriceListData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsPriceListData instantiates a new GETMarkets200ResponseDataInnerRelationshipsPriceListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsPriceListData() *GETMarkets200ResponseDataInnerRelationshipsPriceListData {
	this := GETMarkets200ResponseDataInnerRelationshipsPriceListData{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsPriceListDataWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsPriceListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsPriceListDataWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsPriceListData {
	this := GETMarkets200ResponseDataInnerRelationshipsPriceListData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETMarkets200ResponseDataInnerRelationshipsPriceListData) SetId(v interface{}) {
	o.Id = v
}

func (o GETMarkets200ResponseDataInnerRelationshipsPriceListData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData struct {
	value *GETMarkets200ResponseDataInnerRelationshipsPriceListData
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) Get() *GETMarkets200ResponseDataInnerRelationshipsPriceListData {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) Set(val *GETMarkets200ResponseDataInnerRelationshipsPriceListData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsPriceListData(val *GETMarkets200ResponseDataInnerRelationshipsPriceListData) *NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsPriceListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
