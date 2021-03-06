/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PriceListData struct for PriceListData
type PriceListData struct {
	// The resource's type
	Type          string                      `json:"type"`
	Attributes    PriceListDataAttributes     `json:"attributes"`
	Relationships *PriceListDataRelationships `json:"relationships,omitempty"`
}

// NewPriceListData instantiates a new PriceListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceListData(type_ string, attributes PriceListDataAttributes) *PriceListData {
	this := PriceListData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceListDataWithDefaults instantiates a new PriceListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceListDataWithDefaults() *PriceListData {
	this := PriceListData{}
	var type_ string = "price_lists"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PriceListData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PriceListData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceListData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceListData) GetAttributes() PriceListDataAttributes {
	if o == nil {
		var ret PriceListDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceListData) GetAttributesOk() (*PriceListDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceListData) SetAttributes(v PriceListDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceListData) GetRelationships() PriceListDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceListDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceListData) GetRelationshipsOk() (*PriceListDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceListData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceListDataRelationships and assigns it to the Relationships field.
func (o *PriceListData) SetRelationships(v PriceListDataRelationships) {
	o.Relationships = &v
}

func (o PriceListData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullablePriceListData struct {
	value *PriceListData
	isSet bool
}

func (v NullablePriceListData) Get() *PriceListData {
	return v.value
}

func (v *NullablePriceListData) Set(val *PriceListData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceListData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceListData(val *PriceListData) *NullablePriceListData {
	return &NullablePriceListData{value: val, isSet: true}
}

func (v NullablePriceListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
