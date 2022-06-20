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

// PriceUpdateDataRelationships struct for PriceUpdateDataRelationships
type PriceUpdateDataRelationships struct {
	PriceList *MarketDataRelationshipsPriceList `json:"price_list,omitempty"`
	Sku       *BundleDataRelationshipsSkus      `json:"sku,omitempty"`
}

// NewPriceUpdateDataRelationships instantiates a new PriceUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceUpdateDataRelationships() *PriceUpdateDataRelationships {
	this := PriceUpdateDataRelationships{}
	return &this
}

// NewPriceUpdateDataRelationshipsWithDefaults instantiates a new PriceUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceUpdateDataRelationshipsWithDefaults() *PriceUpdateDataRelationships {
	this := PriceUpdateDataRelationships{}
	return &this
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *PriceUpdateDataRelationships) GetPriceList() MarketDataRelationshipsPriceList {
	if o == nil || o.PriceList == nil {
		var ret MarketDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateDataRelationships) GetPriceListOk() (*MarketDataRelationshipsPriceList, bool) {
	if o == nil || o.PriceList == nil {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *PriceUpdateDataRelationships) HasPriceList() bool {
	if o != nil && o.PriceList != nil {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given MarketDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *PriceUpdateDataRelationships) SetPriceList(v MarketDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *PriceUpdateDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceUpdateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *PriceUpdateDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *PriceUpdateDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

func (o PriceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PriceList != nil {
		toSerialize["price_list"] = o.PriceList
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullablePriceUpdateDataRelationships struct {
	value *PriceUpdateDataRelationships
	isSet bool
}

func (v NullablePriceUpdateDataRelationships) Get() *PriceUpdateDataRelationships {
	return v.value
}

func (v *NullablePriceUpdateDataRelationships) Set(val *PriceUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceUpdateDataRelationships(val *PriceUpdateDataRelationships) *NullablePriceUpdateDataRelationships {
	return &NullablePriceUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePriceUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
