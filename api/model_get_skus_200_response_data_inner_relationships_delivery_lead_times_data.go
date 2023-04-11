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

// GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData struct for GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData
type GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData{}
	return &this
}

// NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataWithDefaults instantiates a new GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesDataWithDefaults() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData {
	this := GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) SetId(v interface{}) {
	o.Id = v
}

func (o GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData struct {
	value *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData
	isSet bool
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) Get() *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData {
	return v.value
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) Set(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData(val *GETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData {
	return &NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData{value: val, isSet: true}
}

func (v NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETSkus200ResponseDataInnerRelationshipsDeliveryLeadTimesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
