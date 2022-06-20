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

// OrderAmountPromotionRuleUpdateData struct for OrderAmountPromotionRuleUpdateData
type OrderAmountPromotionRuleUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                       `json:"id"`
	Attributes    OrderAmountPromotionRuleCreateDataAttributes `json:"attributes"`
	Relationships *OrderAmountPromotionRuleDataRelationships   `json:"relationships,omitempty"`
}

// NewOrderAmountPromotionRuleUpdateData instantiates a new OrderAmountPromotionRuleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleUpdateData(type_ string, id string, attributes OrderAmountPromotionRuleCreateDataAttributes) *OrderAmountPromotionRuleUpdateData {
	this := OrderAmountPromotionRuleUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewOrderAmountPromotionRuleUpdateDataWithDefaults instantiates a new OrderAmountPromotionRuleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleUpdateDataWithDefaults() *OrderAmountPromotionRuleUpdateData {
	this := OrderAmountPromotionRuleUpdateData{}
	var type_ string = "order_amount_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderAmountPromotionRuleUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderAmountPromotionRuleUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *OrderAmountPromotionRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderAmountPromotionRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *OrderAmountPromotionRuleUpdateData) GetAttributes() OrderAmountPromotionRuleCreateDataAttributes {
	if o == nil {
		var ret OrderAmountPromotionRuleCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleUpdateData) GetAttributesOk() (*OrderAmountPromotionRuleCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderAmountPromotionRuleUpdateData) SetAttributes(v OrderAmountPromotionRuleCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderAmountPromotionRuleUpdateData) GetRelationships() OrderAmountPromotionRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderAmountPromotionRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleUpdateData) GetRelationshipsOk() (*OrderAmountPromotionRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderAmountPromotionRuleUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderAmountPromotionRuleDataRelationships and assigns it to the Relationships field.
func (o *OrderAmountPromotionRuleUpdateData) SetRelationships(v OrderAmountPromotionRuleDataRelationships) {
	o.Relationships = &v
}

func (o OrderAmountPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleUpdateData struct {
	value *OrderAmountPromotionRuleUpdateData
	isSet bool
}

func (v NullableOrderAmountPromotionRuleUpdateData) Get() *OrderAmountPromotionRuleUpdateData {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleUpdateData) Set(val *OrderAmountPromotionRuleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleUpdateData(val *OrderAmountPromotionRuleUpdateData) *NullableOrderAmountPromotionRuleUpdateData {
	return &NullableOrderAmountPromotionRuleUpdateData{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
