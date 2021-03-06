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

// CouponCodesPromotionRuleUpdateData struct for CouponCodesPromotionRuleUpdateData
type CouponCodesPromotionRuleUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                     `json:"id"`
	Attributes    AdyenPaymentCreateDataAttributes           `json:"attributes"`
	Relationships *CouponCodesPromotionRuleDataRelationships `json:"relationships,omitempty"`
}

// NewCouponCodesPromotionRuleUpdateData instantiates a new CouponCodesPromotionRuleUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes) *CouponCodesPromotionRuleUpdateData {
	this := CouponCodesPromotionRuleUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCouponCodesPromotionRuleUpdateDataWithDefaults instantiates a new CouponCodesPromotionRuleUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleUpdateDataWithDefaults() *CouponCodesPromotionRuleUpdateData {
	this := CouponCodesPromotionRuleUpdateData{}
	var type_ string = "coupon_codes_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponCodesPromotionRuleUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponCodesPromotionRuleUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CouponCodesPromotionRuleUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponCodesPromotionRuleUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CouponCodesPromotionRuleUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes {
	if o == nil {
		var ret AdyenPaymentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CouponCodesPromotionRuleUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CouponCodesPromotionRuleUpdateData) GetRelationships() CouponCodesPromotionRuleDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret CouponCodesPromotionRuleDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleUpdateData) GetRelationshipsOk() (*CouponCodesPromotionRuleDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CouponCodesPromotionRuleUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CouponCodesPromotionRuleDataRelationships and assigns it to the Relationships field.
func (o *CouponCodesPromotionRuleUpdateData) SetRelationships(v CouponCodesPromotionRuleDataRelationships) {
	o.Relationships = &v
}

func (o CouponCodesPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableCouponCodesPromotionRuleUpdateData struct {
	value *CouponCodesPromotionRuleUpdateData
	isSet bool
}

func (v NullableCouponCodesPromotionRuleUpdateData) Get() *CouponCodesPromotionRuleUpdateData {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleUpdateData) Set(val *CouponCodesPromotionRuleUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleUpdateData(val *CouponCodesPromotionRuleUpdateData) *NullableCouponCodesPromotionRuleUpdateData {
	return &NullableCouponCodesPromotionRuleUpdateData{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
