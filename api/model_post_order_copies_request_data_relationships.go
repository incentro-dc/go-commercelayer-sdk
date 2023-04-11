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

// checks if the POSTOrderCopiesRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopiesRequestDataRelationships{}

// POSTOrderCopiesRequestDataRelationships struct for POSTOrderCopiesRequestDataRelationships
type POSTOrderCopiesRequestDataRelationships struct {
	SourceOrder POSTAdyenPaymentsRequestDataRelationshipsOrder `json:"source_order"`
}

// NewPOSTOrderCopiesRequestDataRelationships instantiates a new POSTOrderCopiesRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopiesRequestDataRelationships(sourceOrder POSTAdyenPaymentsRequestDataRelationshipsOrder) *POSTOrderCopiesRequestDataRelationships {
	this := POSTOrderCopiesRequestDataRelationships{}
	this.SourceOrder = sourceOrder
	return &this
}

// NewPOSTOrderCopiesRequestDataRelationshipsWithDefaults instantiates a new POSTOrderCopiesRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopiesRequestDataRelationshipsWithDefaults() *POSTOrderCopiesRequestDataRelationships {
	this := POSTOrderCopiesRequestDataRelationships{}
	return &this
}

// GetSourceOrder returns the SourceOrder field value
func (o *POSTOrderCopiesRequestDataRelationships) GetSourceOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder {
	if o == nil {
		var ret POSTAdyenPaymentsRequestDataRelationshipsOrder
		return ret
	}

	return o.SourceOrder
}

// GetSourceOrderOk returns a tuple with the SourceOrder field value
// and a boolean to check if the value has been set.
func (o *POSTOrderCopiesRequestDataRelationships) GetSourceOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceOrder, true
}

// SetSourceOrder sets field value
func (o *POSTOrderCopiesRequestDataRelationships) SetSourceOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder) {
	o.SourceOrder = v
}

func (o POSTOrderCopiesRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopiesRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_order"] = o.SourceOrder
	return toSerialize, nil
}

type NullablePOSTOrderCopiesRequestDataRelationships struct {
	value *POSTOrderCopiesRequestDataRelationships
	isSet bool
}

func (v NullablePOSTOrderCopiesRequestDataRelationships) Get() *POSTOrderCopiesRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTOrderCopiesRequestDataRelationships) Set(val *POSTOrderCopiesRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopiesRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopiesRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopiesRequestDataRelationships(val *POSTOrderCopiesRequestDataRelationships) *NullablePOSTOrderCopiesRequestDataRelationships {
	return &NullablePOSTOrderCopiesRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrderCopiesRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopiesRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}