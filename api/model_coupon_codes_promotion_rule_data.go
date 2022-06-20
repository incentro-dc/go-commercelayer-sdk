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

// CouponCodesPromotionRuleData struct for CouponCodesPromotionRuleData
type CouponCodesPromotionRuleData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    BillingInfoValidationRuleDataAttributes    `json:"attributes"`
	Relationships *CouponCodesPromotionRuleDataRelationships `json:"relationships,omitempty"`
}

// NewCouponCodesPromotionRuleData instantiates a new CouponCodesPromotionRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleData(type_ string, attributes BillingInfoValidationRuleDataAttributes) *CouponCodesPromotionRuleData {
	this := CouponCodesPromotionRuleData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewCouponCodesPromotionRuleDataWithDefaults instantiates a new CouponCodesPromotionRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleDataWithDefaults() *CouponCodesPromotionRuleData {
	this := CouponCodesPromotionRuleData{}
	var type_ string = "coupon_codes_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponCodesPromotionRuleData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponCodesPromotionRuleData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *CouponCodesPromotionRuleData) GetAttributes() BillingInfoValidationRuleDataAttributes {
	if o == nil {
		var ret BillingInfoValidationRuleDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleData) GetAttributesOk() (*BillingInfoValidationRuleDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponCodesPromotionRuleData) SetAttributes(v BillingInfoValidationRuleDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponCodesPromotionRuleData) GetRelationships() CouponCodesPromotionRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponCodesPromotionRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleData) GetRelationshipsOk() (*CouponCodesPromotionRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponCodesPromotionRuleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponCodesPromotionRuleDataRelationships and assigns it to the Relationships field.
func (o *CouponCodesPromotionRuleData) SetRelationships(v CouponCodesPromotionRuleDataRelationships) {
	o.Relationships = &v
}

func (o CouponCodesPromotionRuleData) MarshalJSON() ([]byte, error) {
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

type NullableCouponCodesPromotionRuleData struct {
	value *CouponCodesPromotionRuleData
	isSet bool
}

func (v NullableCouponCodesPromotionRuleData) Get() *CouponCodesPromotionRuleData {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleData) Set(val *CouponCodesPromotionRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleData(val *CouponCodesPromotionRuleData) *NullableCouponCodesPromotionRuleData {
	return &NullableCouponCodesPromotionRuleData{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}