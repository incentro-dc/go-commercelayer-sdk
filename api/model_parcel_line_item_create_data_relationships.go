/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ParcelLineItemCreateDataRelationships struct for ParcelLineItemCreateDataRelationships
type ParcelLineItemCreateDataRelationships struct {
	Parcel        PackageDataRelationshipsParcels         `json:"parcel"`
	StockLineItem LineItemDataRelationshipsStockLineItems `json:"stock_line_item"`
	// Deprecated
	ShipmentLineItem *LineItemDataRelationshipsShipmentLineItems `json:"shipment_line_item,omitempty"`
}

// NewParcelLineItemCreateDataRelationships instantiates a new ParcelLineItemCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemCreateDataRelationships(parcel PackageDataRelationshipsParcels, stockLineItem LineItemDataRelationshipsStockLineItems) *ParcelLineItemCreateDataRelationships {
	this := ParcelLineItemCreateDataRelationships{}
	this.Parcel = parcel
	this.StockLineItem = stockLineItem
	return &this
}

// NewParcelLineItemCreateDataRelationshipsWithDefaults instantiates a new ParcelLineItemCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemCreateDataRelationshipsWithDefaults() *ParcelLineItemCreateDataRelationships {
	this := ParcelLineItemCreateDataRelationships{}
	return &this
}

// GetParcel returns the Parcel field value
func (o *ParcelLineItemCreateDataRelationships) GetParcel() PackageDataRelationshipsParcels {
	if o == nil {
		var ret PackageDataRelationshipsParcels
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateDataRelationships) GetParcelOk() (*PackageDataRelationshipsParcels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *ParcelLineItemCreateDataRelationships) SetParcel(v PackageDataRelationshipsParcels) {
	o.Parcel = v
}

// GetStockLineItem returns the StockLineItem field value
func (o *ParcelLineItemCreateDataRelationships) GetStockLineItem() LineItemDataRelationshipsStockLineItems {
	if o == nil {
		var ret LineItemDataRelationshipsStockLineItems
		return ret
	}

	return o.StockLineItem
}

// GetStockLineItemOk returns a tuple with the StockLineItem field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemCreateDataRelationships) GetStockLineItemOk() (*LineItemDataRelationshipsStockLineItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLineItem, true
}

// SetStockLineItem sets field value
func (o *ParcelLineItemCreateDataRelationships) SetStockLineItem(v LineItemDataRelationshipsStockLineItems) {
	o.StockLineItem = v
}

// GetShipmentLineItem returns the ShipmentLineItem field value if set, zero value otherwise.
// Deprecated
func (o *ParcelLineItemCreateDataRelationships) GetShipmentLineItem() LineItemDataRelationshipsShipmentLineItems {
	if o == nil || o.ShipmentLineItem == nil {
		var ret LineItemDataRelationshipsShipmentLineItems
		return ret
	}
	return *o.ShipmentLineItem
}

// GetShipmentLineItemOk returns a tuple with the ShipmentLineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ParcelLineItemCreateDataRelationships) GetShipmentLineItemOk() (*LineItemDataRelationshipsShipmentLineItems, bool) {
	if o == nil || o.ShipmentLineItem == nil {
		return nil, false
	}
	return o.ShipmentLineItem, true
}

// HasShipmentLineItem returns a boolean if a field has been set.
func (o *ParcelLineItemCreateDataRelationships) HasShipmentLineItem() bool {
	if o != nil && o.ShipmentLineItem != nil {
		return true
	}

	return false
}

// SetShipmentLineItem gets a reference to the given LineItemDataRelationshipsShipmentLineItems and assigns it to the ShipmentLineItem field.
// Deprecated
func (o *ParcelLineItemCreateDataRelationships) SetShipmentLineItem(v LineItemDataRelationshipsShipmentLineItems) {
	o.ShipmentLineItem = &v
}

func (o ParcelLineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["parcel"] = o.Parcel
	}
	if true {
		toSerialize["stock_line_item"] = o.StockLineItem
	}
	if o.ShipmentLineItem != nil {
		toSerialize["shipment_line_item"] = o.ShipmentLineItem
	}
	return json.Marshal(toSerialize)
}

type NullableParcelLineItemCreateDataRelationships struct {
	value *ParcelLineItemCreateDataRelationships
	isSet bool
}

func (v NullableParcelLineItemCreateDataRelationships) Get() *ParcelLineItemCreateDataRelationships {
	return v.value
}

func (v *NullableParcelLineItemCreateDataRelationships) Set(val *ParcelLineItemCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemCreateDataRelationships(val *ParcelLineItemCreateDataRelationships) *NullableParcelLineItemCreateDataRelationships {
	return &NullableParcelLineItemCreateDataRelationships{value: val, isSet: true}
}

func (v NullableParcelLineItemCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}