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

// GETCouponRecipients200ResponseDataInnerRelationshipsCustomer struct for GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
type GETCouponRecipients200ResponseDataInnerRelationshipsCustomer struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETCouponRecipients200ResponseDataInnerRelationshipsCustomer instantiates a new GETCouponRecipients200ResponseDataInnerRelationshipsCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponRecipients200ResponseDataInnerRelationshipsCustomer(type_ string, id string) *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	this := GETCouponRecipients200ResponseDataInnerRelationshipsCustomer{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETCouponRecipients200ResponseDataInnerRelationshipsCustomerWithDefaults instantiates a new GETCouponRecipients200ResponseDataInnerRelationshipsCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponRecipients200ResponseDataInnerRelationshipsCustomerWithDefaults() *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	this := GETCouponRecipients200ResponseDataInnerRelationshipsCustomer{}
	var type_ string = "customers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) SetId(v string) {
	o.Id = v
}

func (o GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer struct {
	value *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
	isSet bool
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) Get() *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	return v.value
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) Set(val *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer(val *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) *NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	return &NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer{value: val, isSet: true}
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationshipsCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
