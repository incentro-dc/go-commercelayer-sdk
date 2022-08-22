/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources struct for GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources
type GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources(type_ string, id string) *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSourcesWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSourcesWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources {
	this := GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources{}
	var type_ string = "customer_payment_sources"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) SetId(v string) {
	o.Id = v
}

func (o GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources struct {
	value *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) Get() *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) Set(val *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources(val *GETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsCustomerPaymentSources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
