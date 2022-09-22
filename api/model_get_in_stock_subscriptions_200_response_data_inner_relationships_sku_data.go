/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData struct for GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData
type GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData {
	this := GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData{}
	return &this
}

// NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuDataWithDefaults instantiates a new GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuDataWithDefaults() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData {
	this := GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) SetId(v string) {
	o.Id = &v
}

func (o GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData struct {
	value *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData
	isSet bool
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) Get() *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData {
	return v.value
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) Set(val *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData(val *GETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData {
	return &NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData{value: val, isSet: true}
}

func (v NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInStockSubscriptions200ResponseDataInnerRelationshipsSkuData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}