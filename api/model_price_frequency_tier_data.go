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

// PriceFrequencyTierData struct for PriceFrequencyTierData
type PriceFrequencyTierData struct {
	// The resource's type
	Type          interface{}                                          `json:"type"`
	Attributes    GETPriceFrequencyTiers200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *PriceFrequencyTierDataRelationships                 `json:"relationships,omitempty"`
}

// NewPriceFrequencyTierData instantiates a new PriceFrequencyTierData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFrequencyTierData(type_ interface{}, attributes GETPriceFrequencyTiers200ResponseDataInnerAttributes) *PriceFrequencyTierData {
	this := PriceFrequencyTierData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceFrequencyTierDataWithDefaults instantiates a new PriceFrequencyTierData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFrequencyTierDataWithDefaults() *PriceFrequencyTierData {
	this := PriceFrequencyTierData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceFrequencyTierData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceFrequencyTierData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceFrequencyTierData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceFrequencyTierData) GetAttributes() GETPriceFrequencyTiers200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETPriceFrequencyTiers200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierData) GetAttributesOk() (*GETPriceFrequencyTiers200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceFrequencyTierData) SetAttributes(v GETPriceFrequencyTiers200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceFrequencyTierData) GetRelationships() PriceFrequencyTierDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PriceFrequencyTierDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierData) GetRelationshipsOk() (*PriceFrequencyTierDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceFrequencyTierData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceFrequencyTierDataRelationships and assigns it to the Relationships field.
func (o *PriceFrequencyTierData) SetRelationships(v PriceFrequencyTierDataRelationships) {
	o.Relationships = &v
}

func (o PriceFrequencyTierData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullablePriceFrequencyTierData struct {
	value *PriceFrequencyTierData
	isSet bool
}

func (v NullablePriceFrequencyTierData) Get() *PriceFrequencyTierData {
	return v.value
}

func (v *NullablePriceFrequencyTierData) Set(val *PriceFrequencyTierData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFrequencyTierData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFrequencyTierData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFrequencyTierData(val *PriceFrequencyTierData) *NullablePriceFrequencyTierData {
	return &NullablePriceFrequencyTierData{value: val, isSet: true}
}

func (v NullablePriceFrequencyTierData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFrequencyTierData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
