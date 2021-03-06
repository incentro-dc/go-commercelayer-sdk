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

// InventoryReturnLocationDataRelationships struct for InventoryReturnLocationDataRelationships
type InventoryReturnLocationDataRelationships struct {
	StockLocation  *DeliveryLeadTimeDataRelationshipsStockLocation         `json:"stock_location,omitempty"`
	InventoryModel *InventoryReturnLocationDataRelationshipsInventoryModel `json:"inventory_model,omitempty"`
}

// NewInventoryReturnLocationDataRelationships instantiates a new InventoryReturnLocationDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryReturnLocationDataRelationships() *InventoryReturnLocationDataRelationships {
	this := InventoryReturnLocationDataRelationships{}
	return &this
}

// NewInventoryReturnLocationDataRelationshipsWithDefaults instantiates a new InventoryReturnLocationDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryReturnLocationDataRelationshipsWithDefaults() *InventoryReturnLocationDataRelationships {
	this := InventoryReturnLocationDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *InventoryReturnLocationDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *InventoryReturnLocationDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *InventoryReturnLocationDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetInventoryModel returns the InventoryModel field value if set, zero value otherwise.
func (o *InventoryReturnLocationDataRelationships) GetInventoryModel() InventoryReturnLocationDataRelationshipsInventoryModel {
	if o == nil || o.InventoryModel == nil {
		var ret InventoryReturnLocationDataRelationshipsInventoryModel
		return ret
	}
	return *o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryReturnLocationDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationDataRelationshipsInventoryModel, bool) {
	if o == nil || o.InventoryModel == nil {
		return nil, false
	}
	return o.InventoryModel, true
}

// HasInventoryModel returns a boolean if a field has been set.
func (o *InventoryReturnLocationDataRelationships) HasInventoryModel() bool {
	if o != nil && o.InventoryModel != nil {
		return true
	}

	return false
}

// SetInventoryModel gets a reference to the given InventoryReturnLocationDataRelationshipsInventoryModel and assigns it to the InventoryModel field.
func (o *InventoryReturnLocationDataRelationships) SetInventoryModel(v InventoryReturnLocationDataRelationshipsInventoryModel) {
	o.InventoryModel = &v
}

func (o InventoryReturnLocationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.InventoryModel != nil {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryReturnLocationDataRelationships struct {
	value *InventoryReturnLocationDataRelationships
	isSet bool
}

func (v NullableInventoryReturnLocationDataRelationships) Get() *InventoryReturnLocationDataRelationships {
	return v.value
}

func (v *NullableInventoryReturnLocationDataRelationships) Set(val *InventoryReturnLocationDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryReturnLocationDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryReturnLocationDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryReturnLocationDataRelationships(val *InventoryReturnLocationDataRelationships) *NullableInventoryReturnLocationDataRelationships {
	return &NullableInventoryReturnLocationDataRelationships{value: val, isSet: true}
}

func (v NullableInventoryReturnLocationDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryReturnLocationDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
