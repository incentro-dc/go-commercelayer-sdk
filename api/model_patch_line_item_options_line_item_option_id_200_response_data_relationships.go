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

// PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships struct for PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships
type PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships struct {
	SkuOption *GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption `json:"sku_option,omitempty"`
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships() *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships{}
	return &this
}

// NewPATCHLineItemOptionsLineItemOptionId200ResponseDataRelationshipsWithDefaults instantiates a new PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHLineItemOptionsLineItemOptionId200ResponseDataRelationshipsWithDefaults() *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships {
	this := PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships{}
	return &this
}

// GetSkuOption returns the SkuOption field value if set, zero value otherwise.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) GetSkuOption() GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption {
	if o == nil || o.SkuOption == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption
		return ret
	}
	return *o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) GetSkuOptionOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption, bool) {
	if o == nil || o.SkuOption == nil {
		return nil, false
	}
	return o.SkuOption, true
}

// HasSkuOption returns a boolean if a field has been set.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) HasSkuOption() bool {
	if o != nil && o.SkuOption != nil {
		return true
	}

	return false
}

// SetSkuOption gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption and assigns it to the SkuOption field.
func (o *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) SetSkuOption(v GETLineItemOptions200ResponseDataInnerRelationshipsSkuOption) {
	o.SkuOption = &v
}

func (o PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuOption != nil {
		toSerialize["sku_option"] = o.SkuOption
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships struct {
	value *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships
	isSet bool
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) Get() *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships {
	return v.value
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) Set(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships(val *PATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships {
	return &NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHLineItemOptionsLineItemOptionId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}