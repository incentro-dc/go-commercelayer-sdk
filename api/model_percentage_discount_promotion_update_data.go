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

// PercentageDiscountPromotionUpdateData struct for PercentageDiscountPromotionUpdateData
type PercentageDiscountPromotionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                          `json:"id"`
	Attributes    PercentageDiscountPromotionUpdateDataAttributes `json:"attributes"`
	Relationships *FixedPricePromotionUpdateDataRelationships     `json:"relationships,omitempty"`
}

// NewPercentageDiscountPromotionUpdateData instantiates a new PercentageDiscountPromotionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPercentageDiscountPromotionUpdateData(type_ string, id string, attributes PercentageDiscountPromotionUpdateDataAttributes) *PercentageDiscountPromotionUpdateData {
	this := PercentageDiscountPromotionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewPercentageDiscountPromotionUpdateDataWithDefaults instantiates a new PercentageDiscountPromotionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPercentageDiscountPromotionUpdateDataWithDefaults() *PercentageDiscountPromotionUpdateData {
	this := PercentageDiscountPromotionUpdateData{}
	var type_ string = "percentage_discount_promotions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *PercentageDiscountPromotionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PercentageDiscountPromotionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PercentageDiscountPromotionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PercentageDiscountPromotionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *PercentageDiscountPromotionUpdateData) GetAttributes() PercentageDiscountPromotionUpdateDataAttributes {
	if o == nil {
		var ret PercentageDiscountPromotionUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionUpdateData) GetAttributesOk() (*PercentageDiscountPromotionUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PercentageDiscountPromotionUpdateData) SetAttributes(v PercentageDiscountPromotionUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PercentageDiscountPromotionUpdateData) GetRelationships() FixedPricePromotionUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret FixedPricePromotionUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PercentageDiscountPromotionUpdateData) GetRelationshipsOk() (*FixedPricePromotionUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PercentageDiscountPromotionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given FixedPricePromotionUpdateDataRelationships and assigns it to the Relationships field.
func (o *PercentageDiscountPromotionUpdateData) SetRelationships(v FixedPricePromotionUpdateDataRelationships) {
	o.Relationships = &v
}

func (o PercentageDiscountPromotionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullablePercentageDiscountPromotionUpdateData struct {
	value *PercentageDiscountPromotionUpdateData
	isSet bool
}

func (v NullablePercentageDiscountPromotionUpdateData) Get() *PercentageDiscountPromotionUpdateData {
	return v.value
}

func (v *NullablePercentageDiscountPromotionUpdateData) Set(val *PercentageDiscountPromotionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePercentageDiscountPromotionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePercentageDiscountPromotionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePercentageDiscountPromotionUpdateData(val *PercentageDiscountPromotionUpdateData) *NullablePercentageDiscountPromotionUpdateData {
	return &NullablePercentageDiscountPromotionUpdateData{value: val, isSet: true}
}

func (v NullablePercentageDiscountPromotionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePercentageDiscountPromotionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}