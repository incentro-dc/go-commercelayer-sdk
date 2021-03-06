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

// PromotionData struct for PromotionData
type PromotionData struct {
	// The resource's type
	Type          string                              `json:"type"`
	Attributes    FreeShippingPromotionDataAttributes `json:"attributes"`
	Relationships *ExternalPromotionDataRelationships `json:"relationships,omitempty"`
}

// NewPromotionData instantiates a new PromotionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionData(type_ string, attributes FreeShippingPromotionDataAttributes) *PromotionData {
	this := PromotionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPromotionDataWithDefaults instantiates a new PromotionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionDataWithDefaults() *PromotionData {
	this := PromotionData{}
	var type_ string = "promotions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PromotionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromotionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromotionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PromotionData) GetAttributes() FreeShippingPromotionDataAttributes {
	if o == nil {
		var ret FreeShippingPromotionDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PromotionData) GetAttributesOk() (*FreeShippingPromotionDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PromotionData) SetAttributes(v FreeShippingPromotionDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PromotionData) GetRelationships() ExternalPromotionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ExternalPromotionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotionData) GetRelationshipsOk() (*ExternalPromotionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PromotionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ExternalPromotionDataRelationships and assigns it to the Relationships field.
func (o *PromotionData) SetRelationships(v ExternalPromotionDataRelationships) {
	o.Relationships = &v
}

func (o PromotionData) MarshalJSON() ([]byte, error) {
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

type NullablePromotionData struct {
	value *PromotionData
	isSet bool
}

func (v NullablePromotionData) Get() *PromotionData {
	return v.value
}

func (v *NullablePromotionData) Set(val *PromotionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionData(val *PromotionData) *NullablePromotionData {
	return &NullablePromotionData{value: val, isSet: true}
}

func (v NullablePromotionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
