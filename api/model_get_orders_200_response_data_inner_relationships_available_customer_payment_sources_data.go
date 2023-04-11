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

// GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData struct for GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData
type GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesDataWithDefaults() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData {
	this := GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) SetId(v interface{}) {
	o.Id = v
}

func (o GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData struct {
	value *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) Get() *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) Set(val *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData(val *GETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData {
	return &NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsAvailableCustomerPaymentSourcesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
