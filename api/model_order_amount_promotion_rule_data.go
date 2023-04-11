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

// OrderAmountPromotionRuleData struct for OrderAmountPromotionRuleData
type OrderAmountPromotionRuleData struct {
	// The resource's type
	Type          interface{}                                                `json:"type"`
	Attributes    GETOrderAmountPromotionRules200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *OrderAmountPromotionRuleDataRelationships                 `json:"relationships,omitempty"`
}

// NewOrderAmountPromotionRuleData instantiates a new OrderAmountPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleData(type_ interface{}, attributes GETOrderAmountPromotionRules200ResponseDataInnerAttributes) *OrderAmountPromotionRuleData {
	this := OrderAmountPromotionRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderAmountPromotionRuleDataWithDefaults instantiates a new OrderAmountPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleDataWithDefaults() *OrderAmountPromotionRuleData {
	this := OrderAmountPromotionRuleData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderAmountPromotionRuleData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderAmountPromotionRuleData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderAmountPromotionRuleData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderAmountPromotionRuleData) GetAttributes() GETOrderAmountPromotionRules200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETOrderAmountPromotionRules200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleData) GetAttributesOk() (*GETOrderAmountPromotionRules200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderAmountPromotionRuleData) SetAttributes(v GETOrderAmountPromotionRules200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleData) GetRelationships() OrderAmountPromotionRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderAmountPromotionRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleData) GetRelationshipsOk() (*OrderAmountPromotionRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderAmountPromotionRuleDataRelationships and assigns it to the Relationships field.
func (o *OrderAmountPromotionRuleData) SetRelationships(v OrderAmountPromotionRuleDataRelationships) {
	o.Relationships = &v
}

func (o OrderAmountPromotionRuleData) MarshalJSON() ([]byte, error) {
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

type NullableOrderAmountPromotionRuleData struct {
	value *OrderAmountPromotionRuleData
	isSet bool
}

func (v NullableOrderAmountPromotionRuleData) Get() *OrderAmountPromotionRuleData {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleData) Set(val *OrderAmountPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleData(val *OrderAmountPromotionRuleData) *NullableOrderAmountPromotionRuleData {
	return &NullableOrderAmountPromotionRuleData{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
