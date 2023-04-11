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

// BundleCreateDataRelationships struct for BundleCreateDataRelationships
type BundleCreateDataRelationships struct {
	Market  *BillingInfoValidationRuleCreateDataRelationshipsMarket `json:"market,omitempty"`
	SkuList BundleCreateDataRelationshipsSkuList                    `json:"sku_list"`
}

// NewBundleCreateDataRelationships instantiates a new BundleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleCreateDataRelationships(skuList BundleCreateDataRelationshipsSkuList) *BundleCreateDataRelationships {
	this := BundleCreateDataRelationships{}
	this.SkuList = skuList
	return &this
}

// NewBundleCreateDataRelationshipsWithDefaults instantiates a new BundleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleCreateDataRelationshipsWithDefaults() *BundleCreateDataRelationships {
	this := BundleCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *BundleCreateDataRelationships) GetMarket() BillingInfoValidationRuleCreateDataRelationshipsMarket {
	if o == nil || o.Market == nil {
		var ret BillingInfoValidationRuleCreateDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationships) GetMarketOk() (*BillingInfoValidationRuleCreateDataRelationshipsMarket, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *BundleCreateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given BillingInfoValidationRuleCreateDataRelationshipsMarket and assigns it to the Market field.
func (o *BundleCreateDataRelationships) SetMarket(v BillingInfoValidationRuleCreateDataRelationshipsMarket) {
	o.Market = &v
}

// GetSkuList returns the SkuList field value
func (o *BundleCreateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList {
	if o == nil {
		var ret BundleCreateDataRelationshipsSkuList
		return ret
	}

	return o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value
// and a boolean to check if the value has been set.
func (o *BundleCreateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkuList, true
}

// SetSkuList sets field value
func (o *BundleCreateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList) {
	o.SkuList = v
}

func (o BundleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if true {
		toSerialize["sku_list"] = o.SkuList
	}
	return json.Marshal(toSerialize)
}

type NullableBundleCreateDataRelationships struct {
	value *BundleCreateDataRelationships
	isSet bool
}

func (v NullableBundleCreateDataRelationships) Get() *BundleCreateDataRelationships {
	return v.value
}

func (v *NullableBundleCreateDataRelationships) Set(val *BundleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleCreateDataRelationships(val *BundleCreateDataRelationships) *NullableBundleCreateDataRelationships {
	return &NullableBundleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableBundleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
