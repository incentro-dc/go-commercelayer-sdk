/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuListPromotionRuleCreateData struct for SkuListPromotionRuleCreateData
type SkuListPromotionRuleCreateData struct {
	// The resource's type
	Type          string                                       `json:"type"`
	Attributes    SkuListPromotionRuleCreateDataAttributes     `json:"attributes"`
	Relationships *SkuListPromotionRuleCreateDataRelationships `json:"relationships,omitempty"`
}

// NewSkuListPromotionRuleCreateData instantiates a new SkuListPromotionRuleCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleCreateData(type_ string, attributes SkuListPromotionRuleCreateDataAttributes) *SkuListPromotionRuleCreateData {
	this := SkuListPromotionRuleCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListPromotionRuleCreateDataWithDefaults instantiates a new SkuListPromotionRuleCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleCreateDataWithDefaults() *SkuListPromotionRuleCreateData {
	this := SkuListPromotionRuleCreateData{}
	var type_ string = "sku_list_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SkuListPromotionRuleCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListPromotionRuleCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListPromotionRuleCreateData) GetAttributes() SkuListPromotionRuleCreateDataAttributes {
	if o == nil {
		var ret SkuListPromotionRuleCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateData) GetAttributesOk() (*SkuListPromotionRuleCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListPromotionRuleCreateData) SetAttributes(v SkuListPromotionRuleCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListPromotionRuleCreateData) GetRelationships() SkuListPromotionRuleCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SkuListPromotionRuleCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreateData) GetRelationshipsOk() (*SkuListPromotionRuleCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListPromotionRuleCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SkuListPromotionRuleCreateDataRelationships and assigns it to the Relationships field.
func (o *SkuListPromotionRuleCreateData) SetRelationships(v SkuListPromotionRuleCreateDataRelationships) {
	o.Relationships = &v
}

func (o SkuListPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListPromotionRuleCreateData struct {
	value *SkuListPromotionRuleCreateData
	isSet bool
}

func (v NullableSkuListPromotionRuleCreateData) Get() *SkuListPromotionRuleCreateData {
	return v.value
}

func (v *NullableSkuListPromotionRuleCreateData) Set(val *SkuListPromotionRuleCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleCreateData(val *SkuListPromotionRuleCreateData) *NullableSkuListPromotionRuleCreateData {
	return &NullableSkuListPromotionRuleCreateData{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}