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

// SkuListPromotionRuleData struct for SkuListPromotionRuleData
type SkuListPromotionRuleData struct {
	// The resource's type
	Type          string                                                     `json:"type"`
	Attributes    GETSkuListPromotionRules200ResponseDataInnerAttributes     `json:"attributes"`
	Relationships *GETSkuListPromotionRules200ResponseDataInnerRelationships `json:"relationships,omitempty"`
}

// NewSkuListPromotionRuleData instantiates a new SkuListPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleData(type_ string, attributes GETSkuListPromotionRules200ResponseDataInnerAttributes) *SkuListPromotionRuleData {
	this := SkuListPromotionRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSkuListPromotionRuleDataWithDefaults instantiates a new SkuListPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleDataWithDefaults() *SkuListPromotionRuleData {
	this := SkuListPromotionRuleData{}
	var type_ string = "sku_list_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SkuListPromotionRuleData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SkuListPromotionRuleData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SkuListPromotionRuleData) GetAttributes() GETSkuListPromotionRules200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETSkuListPromotionRules200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleData) GetAttributesOk() (*GETSkuListPromotionRules200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SkuListPromotionRuleData) SetAttributes(v GETSkuListPromotionRules200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SkuListPromotionRuleData) GetRelationships() GETSkuListPromotionRules200ResponseDataInnerRelationships {
	if o == nil || o.Relationships == nil {
		var ret GETSkuListPromotionRules200ResponseDataInnerRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleData) GetRelationshipsOk() (*GETSkuListPromotionRules200ResponseDataInnerRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SkuListPromotionRuleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GETSkuListPromotionRules200ResponseDataInnerRelationships and assigns it to the Relationships field.
func (o *SkuListPromotionRuleData) SetRelationships(v GETSkuListPromotionRules200ResponseDataInnerRelationships) {
	o.Relationships = &v
}

func (o SkuListPromotionRuleData) MarshalJSON() ([]byte, error) {
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

type NullableSkuListPromotionRuleData struct {
	value *SkuListPromotionRuleData
	isSet bool
}

func (v NullableSkuListPromotionRuleData) Get() *SkuListPromotionRuleData {
	return v.value
}

func (v *NullableSkuListPromotionRuleData) Set(val *SkuListPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleData(val *SkuListPromotionRuleData) *NullableSkuListPromotionRuleData {
	return &NullableSkuListPromotionRuleData{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
