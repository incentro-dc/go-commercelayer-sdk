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

// GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct for GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule
type GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule(type_ string, id string) *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRuleWithDefaults() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{}
	var type_ string = "sku_list_promotion_rules"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) SetId(v string) {
	o.Id = v
}

func (o GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule struct {
	value *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Get() *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Set(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule(val *GETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule {
	return &NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}