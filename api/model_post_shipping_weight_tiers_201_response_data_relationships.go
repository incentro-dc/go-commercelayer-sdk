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

// POSTShippingWeightTiers201ResponseDataRelationships struct for POSTShippingWeightTiers201ResponseDataRelationships
type POSTShippingWeightTiers201ResponseDataRelationships struct {
	ShippingMethod GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod `json:"shipping_method"`
}

// NewPOSTShippingWeightTiers201ResponseDataRelationships instantiates a new POSTShippingWeightTiers201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingWeightTiers201ResponseDataRelationships(shippingMethod GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod) *POSTShippingWeightTiers201ResponseDataRelationships {
	this := POSTShippingWeightTiers201ResponseDataRelationships{}
	this.ShippingMethod = shippingMethod
	return &this
}

// NewPOSTShippingWeightTiers201ResponseDataRelationshipsWithDefaults instantiates a new POSTShippingWeightTiers201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingWeightTiers201ResponseDataRelationshipsWithDefaults() *POSTShippingWeightTiers201ResponseDataRelationships {
	this := POSTShippingWeightTiers201ResponseDataRelationships{}
	return &this
}

// GetShippingMethod returns the ShippingMethod field value
func (o *POSTShippingWeightTiers201ResponseDataRelationships) GetShippingMethod() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod {
	if o == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod
		return ret
	}

	return o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value
// and a boolean to check if the value has been set.
func (o *POSTShippingWeightTiers201ResponseDataRelationships) GetShippingMethodOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingMethod, true
}

// SetShippingMethod sets field value
func (o *POSTShippingWeightTiers201ResponseDataRelationships) SetShippingMethod(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod) {
	o.ShippingMethod = v
}

func (o POSTShippingWeightTiers201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTShippingWeightTiers201ResponseDataRelationships struct {
	value *POSTShippingWeightTiers201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTShippingWeightTiers201ResponseDataRelationships) Get() *POSTShippingWeightTiers201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTShippingWeightTiers201ResponseDataRelationships) Set(val *POSTShippingWeightTiers201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingWeightTiers201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingWeightTiers201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingWeightTiers201ResponseDataRelationships(val *POSTShippingWeightTiers201ResponseDataRelationships) *NullablePOSTShippingWeightTiers201ResponseDataRelationships {
	return &NullablePOSTShippingWeightTiers201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTShippingWeightTiers201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingWeightTiers201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
