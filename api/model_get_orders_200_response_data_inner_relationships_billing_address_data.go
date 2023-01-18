/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETOrders200ResponseDataInnerRelationshipsBillingAddressData struct for GETOrders200ResponseDataInnerRelationshipsBillingAddressData
type GETOrders200ResponseDataInnerRelationshipsBillingAddressData struct {
	// The resource's type
	Type *string `json:"type,omitempty"`
	// The resource ID
	Id *string `json:"id,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsBillingAddressData instantiates a new GETOrders200ResponseDataInnerRelationshipsBillingAddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsBillingAddressData() *GETOrders200ResponseDataInnerRelationshipsBillingAddressData {
	this := GETOrders200ResponseDataInnerRelationshipsBillingAddressData{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsBillingAddressDataWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsBillingAddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsBillingAddressDataWithDefaults() *GETOrders200ResponseDataInnerRelationshipsBillingAddressData {
	this := GETOrders200ResponseDataInnerRelationshipsBillingAddressData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) SetId(v string) {
	o.Id = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsBillingAddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData struct {
	value *GETOrders200ResponseDataInnerRelationshipsBillingAddressData
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) Get() *GETOrders200ResponseDataInnerRelationshipsBillingAddressData {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) Set(val *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData(val *GETOrders200ResponseDataInnerRelationshipsBillingAddressData) *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData {
	return &NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
