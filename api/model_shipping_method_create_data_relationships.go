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

// ShippingMethodCreateDataRelationships struct for ShippingMethodCreateDataRelationships
type ShippingMethodCreateDataRelationships struct {
	Market              *AvalaraAccountDataRelationshipsMarkets             `json:"market,omitempty"`
	ShippingZone        *ShippingMethodDataRelationshipsShippingZone        `json:"shipping_zone,omitempty"`
	ShippingCategory    *ShipmentDataRelationshipsShippingCategory          `json:"shipping_category,omitempty"`
	StockLocation       *DeliveryLeadTimeDataRelationshipsStockLocation     `json:"stock_location,omitempty"`
	ShippingMethodTiers *ShippingMethodDataRelationshipsShippingMethodTiers `json:"shipping_method_tiers,omitempty"`
}

// NewShippingMethodCreateDataRelationships instantiates a new ShippingMethodCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodCreateDataRelationships() *ShippingMethodCreateDataRelationships {
	this := ShippingMethodCreateDataRelationships{}
	return &this
}

// NewShippingMethodCreateDataRelationshipsWithDefaults instantiates a new ShippingMethodCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodCreateDataRelationshipsWithDefaults() *ShippingMethodCreateDataRelationships {
	this := ShippingMethodCreateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *ShippingMethodCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetShippingZone returns the ShippingZone field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingZone() ShippingMethodDataRelationshipsShippingZone {
	if o == nil || o.ShippingZone == nil {
		var ret ShippingMethodDataRelationshipsShippingZone
		return ret
	}
	return *o.ShippingZone
}

// GetShippingZoneOk returns a tuple with the ShippingZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingZoneOk() (*ShippingMethodDataRelationshipsShippingZone, bool) {
	if o == nil || o.ShippingZone == nil {
		return nil, false
	}
	return o.ShippingZone, true
}

// HasShippingZone returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingZone() bool {
	if o != nil && o.ShippingZone != nil {
		return true
	}

	return false
}

// SetShippingZone gets a reference to the given ShippingMethodDataRelationshipsShippingZone and assigns it to the ShippingZone field.
func (o *ShippingMethodCreateDataRelationships) SetShippingZone(v ShippingMethodDataRelationshipsShippingZone) {
	o.ShippingZone = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShippingMethodCreateDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ShippingMethodCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetShippingMethodTiers returns the ShippingMethodTiers field value if set, zero value otherwise.
func (o *ShippingMethodCreateDataRelationships) GetShippingMethodTiers() ShippingMethodDataRelationshipsShippingMethodTiers {
	if o == nil || o.ShippingMethodTiers == nil {
		var ret ShippingMethodDataRelationshipsShippingMethodTiers
		return ret
	}
	return *o.ShippingMethodTiers
}

// GetShippingMethodTiersOk returns a tuple with the ShippingMethodTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodCreateDataRelationships) GetShippingMethodTiersOk() (*ShippingMethodDataRelationshipsShippingMethodTiers, bool) {
	if o == nil || o.ShippingMethodTiers == nil {
		return nil, false
	}
	return o.ShippingMethodTiers, true
}

// HasShippingMethodTiers returns a boolean if a field has been set.
func (o *ShippingMethodCreateDataRelationships) HasShippingMethodTiers() bool {
	if o != nil && o.ShippingMethodTiers != nil {
		return true
	}

	return false
}

// SetShippingMethodTiers gets a reference to the given ShippingMethodDataRelationshipsShippingMethodTiers and assigns it to the ShippingMethodTiers field.
func (o *ShippingMethodCreateDataRelationships) SetShippingMethodTiers(v ShippingMethodDataRelationshipsShippingMethodTiers) {
	o.ShippingMethodTiers = &v
}

func (o ShippingMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.ShippingZone != nil {
		toSerialize["shipping_zone"] = o.ShippingZone
	}
	if o.ShippingCategory != nil {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.ShippingMethodTiers != nil {
		toSerialize["shipping_method_tiers"] = o.ShippingMethodTiers
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodCreateDataRelationships struct {
	value *ShippingMethodCreateDataRelationships
	isSet bool
}

func (v NullableShippingMethodCreateDataRelationships) Get() *ShippingMethodCreateDataRelationships {
	return v.value
}

func (v *NullableShippingMethodCreateDataRelationships) Set(val *ShippingMethodCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodCreateDataRelationships(val *ShippingMethodCreateDataRelationships) *NullableShippingMethodCreateDataRelationships {
	return &NullableShippingMethodCreateDataRelationships{value: val, isSet: true}
}

func (v NullableShippingMethodCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
