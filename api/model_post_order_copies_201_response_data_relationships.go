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

// POSTOrderCopies201ResponseDataRelationships struct for POSTOrderCopies201ResponseDataRelationships
type POSTOrderCopies201ResponseDataRelationships struct {
	SourceOrder GETAdyenPayments200ResponseDataInnerRelationshipsOrder `json:"source_order"`
}

// NewPOSTOrderCopies201ResponseDataRelationships instantiates a new POSTOrderCopies201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopies201ResponseDataRelationships(sourceOrder GETAdyenPayments200ResponseDataInnerRelationshipsOrder) *POSTOrderCopies201ResponseDataRelationships {
	this := POSTOrderCopies201ResponseDataRelationships{}
	this.SourceOrder = sourceOrder
	return &this
}

// NewPOSTOrderCopies201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderCopies201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopies201ResponseDataRelationshipsWithDefaults() *POSTOrderCopies201ResponseDataRelationships {
	this := POSTOrderCopies201ResponseDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value
func (o *POSTOrderCopies201ResponseDataRelationships) GetSourceOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}

	return o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value
// and a boolean to check if the value has been set.
func (o *POSTOrderCopies201ResponseDataRelationships) GetSourceOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOrder, true
}

// SetSourceOrder sets field value
func (o *POSTOrderCopies201ResponseDataRelationships) SetSourceOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.SourceOrder = v
}

func (o POSTOrderCopies201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source_order"] = o.SourceOrder
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTOrderCopies201ResponseDataRelationships struct {
	value *POSTOrderCopies201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTOrderCopies201ResponseDataRelationships) Get() *POSTOrderCopies201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationships) Set(val *POSTOrderCopies201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopies201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopies201ResponseDataRelationships(val *POSTOrderCopies201ResponseDataRelationships) *NullablePOSTOrderCopies201ResponseDataRelationships {
	return &NullablePOSTOrderCopies201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrderCopies201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopies201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}