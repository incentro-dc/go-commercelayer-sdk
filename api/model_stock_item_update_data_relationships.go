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

// StockItemUpdateDataRelationships struct for StockItemUpdateDataRelationships
type StockItemUpdateDataRelationships struct {
	StockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *BundleDataRelationshipsSkus                    `json:"sku,omitempty"`
}

// NewStockItemUpdateDataRelationships instantiates a new StockItemUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockItemUpdateDataRelationships() *StockItemUpdateDataRelationships {
	this := StockItemUpdateDataRelationships{}
	return &this
}

// NewStockItemUpdateDataRelationshipsWithDefaults instantiates a new StockItemUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockItemUpdateDataRelationshipsWithDefaults() *StockItemUpdateDataRelationships {
	this := StockItemUpdateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *StockItemUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *StockItemUpdateDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *StockItemUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockItemUpdateDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockItemUpdateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockItemUpdateDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *StockItemUpdateDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

func (o StockItemUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullableStockItemUpdateDataRelationships struct {
	value *StockItemUpdateDataRelationships
	isSet bool
}

func (v NullableStockItemUpdateDataRelationships) Get() *StockItemUpdateDataRelationships {
	return v.value
}

func (v *NullableStockItemUpdateDataRelationships) Set(val *StockItemUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockItemUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockItemUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockItemUpdateDataRelationships(val *StockItemUpdateDataRelationships) *NullableStockItemUpdateDataRelationships {
	return &NullableStockItemUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableStockItemUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockItemUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
