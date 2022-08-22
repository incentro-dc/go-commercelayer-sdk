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

// POSTInventoryReturnLocations201ResponseDataRelationships struct for POSTInventoryReturnLocations201ResponseDataRelationships
type POSTInventoryReturnLocations201ResponseDataRelationships struct {
	StockLocation  GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation         `json:"stock_location"`
	InventoryModel GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel `json:"inventory_model"`
}

// NewPOSTInventoryReturnLocations201ResponseDataRelationships instantiates a new POSTInventoryReturnLocations201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTInventoryReturnLocations201ResponseDataRelationships(stockLocation GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, inventoryModel GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) *POSTInventoryReturnLocations201ResponseDataRelationships {
	this := POSTInventoryReturnLocations201ResponseDataRelationships{}
	this.StockLocation = stockLocation
	this.InventoryModel = inventoryModel
	return &this
}

// NewPOSTInventoryReturnLocations201ResponseDataRelationshipsWithDefaults instantiates a new POSTInventoryReturnLocations201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTInventoryReturnLocations201ResponseDataRelationshipsWithDefaults() *POSTInventoryReturnLocations201ResponseDataRelationships {
	this := POSTInventoryReturnLocations201ResponseDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}

	return o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value
// and a boolean to check if the value has been set.
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLocation, true
}

// SetStockLocation sets field value
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = v
}

// GetInventoryModel returns the InventoryModel field value
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) GetInventoryModel() GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel {
	if o == nil {
		var ret GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel
		return ret
	}

	return o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value
// and a boolean to check if the value has been set.
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) GetInventoryModelOk() (*GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryModel, true
}

// SetInventoryModel sets field value
func (o *POSTInventoryReturnLocations201ResponseDataRelationships) SetInventoryModel(v GETInventoryReturnLocations200ResponseDataInnerRelationshipsInventoryModel) {
	o.InventoryModel = v
}

func (o POSTInventoryReturnLocations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stock_location"] = o.StockLocation
	}
	if true {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTInventoryReturnLocations201ResponseDataRelationships struct {
	value *POSTInventoryReturnLocations201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTInventoryReturnLocations201ResponseDataRelationships) Get() *POSTInventoryReturnLocations201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTInventoryReturnLocations201ResponseDataRelationships) Set(val *POSTInventoryReturnLocations201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTInventoryReturnLocations201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTInventoryReturnLocations201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTInventoryReturnLocations201ResponseDataRelationships(val *POSTInventoryReturnLocations201ResponseDataRelationships) *NullablePOSTInventoryReturnLocations201ResponseDataRelationships {
	return &NullablePOSTInventoryReturnLocations201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTInventoryReturnLocations201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTInventoryReturnLocations201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
