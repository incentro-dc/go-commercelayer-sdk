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

// GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers struct for GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers
type GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers(type_ string, id string) *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiersWithDefaults instantiates a new GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiersWithDefaults() *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers {
	this := GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers{}
	var type_ string = "shipping_weight_tiers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) SetId(v string) {
	o.Id = v
}

func (o GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers struct {
	value *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers
	isSet bool
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) Get() *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers {
	return v.value
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) Set(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers(val *GETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers {
	return &NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers{value: val, isSet: true}
}

func (v NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingMethods200ResponseDataInnerRelationshipsShippingWeightTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}