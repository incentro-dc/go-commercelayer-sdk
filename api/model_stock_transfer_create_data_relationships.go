/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StockTransferCreateDataRelationships struct for StockTransferCreateDataRelationships
type StockTransferCreateDataRelationships struct {
	Sku                      InStockSubscriptionCreateDataRelationshipsSku        `json:"sku"`
	OriginStockLocation      DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"origin_stock_location"`
	DestinationStockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"destination_stock_location"`
	Shipment                 *ParcelCreateDataRelationshipsShipment               `json:"shipment,omitempty"`
	LineItem                 *LineItemOptionCreateDataRelationshipsLineItem       `json:"line_item,omitempty"`
}

// NewStockTransferCreateDataRelationships instantiates a new StockTransferCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferCreateDataRelationships(sku InStockSubscriptionCreateDataRelationshipsSku, originStockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation, destinationStockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation) *StockTransferCreateDataRelationships {
	this := StockTransferCreateDataRelationships{}
	this.Sku = sku
	this.OriginStockLocation = originStockLocation
	this.DestinationStockLocation = destinationStockLocation
	return &this
}

// NewStockTransferCreateDataRelationshipsWithDefaults instantiates a new StockTransferCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferCreateDataRelationshipsWithDefaults() *StockTransferCreateDataRelationships {
	this := StockTransferCreateDataRelationships{}
	return &this
}

// GetSku returns the Sku field value
func (o *StockTransferCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *StockTransferCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = v
}

// GetOriginStockLocation returns the OriginStockLocation field value
func (o *StockTransferCreateDataRelationships) GetOriginStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}

	return o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreateDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginStockLocation, true
}

// SetOriginStockLocation sets field value
func (o *StockTransferCreateDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.OriginStockLocation = v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value
func (o *StockTransferCreateDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}

	return o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value
// and a boolean to check if the value has been set.
func (o *StockTransferCreateDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationStockLocation, true
}

// SetDestinationStockLocation sets field value
func (o *StockTransferCreateDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.DestinationStockLocation = v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *StockTransferCreateDataRelationships) GetShipment() ParcelCreateDataRelationshipsShipment {
	if o == nil || o.Shipment == nil {
		var ret ParcelCreateDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferCreateDataRelationships) GetShipmentOk() (*ParcelCreateDataRelationshipsShipment, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *StockTransferCreateDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given ParcelCreateDataRelationshipsShipment and assigns it to the Shipment field.
func (o *StockTransferCreateDataRelationships) SetShipment(v ParcelCreateDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *StockTransferCreateDataRelationships) GetLineItem() LineItemOptionCreateDataRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret LineItemOptionCreateDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferCreateDataRelationships) GetLineItemOk() (*LineItemOptionCreateDataRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *StockTransferCreateDataRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionCreateDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *StockTransferCreateDataRelationships) SetLineItem(v LineItemOptionCreateDataRelationshipsLineItem) {
	o.LineItem = &v
}

func (o StockTransferCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sku"] = o.Sku
	}
	if true {
		toSerialize["origin_stock_location"] = o.OriginStockLocation
	}
	if true {
		toSerialize["destination_stock_location"] = o.DestinationStockLocation
	}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	return json.Marshal(toSerialize)
}

type NullableStockTransferCreateDataRelationships struct {
	value *StockTransferCreateDataRelationships
	isSet bool
}

func (v NullableStockTransferCreateDataRelationships) Get() *StockTransferCreateDataRelationships {
	return v.value
}

func (v *NullableStockTransferCreateDataRelationships) Set(val *StockTransferCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferCreateDataRelationships(val *StockTransferCreateDataRelationships) *NullableStockTransferCreateDataRelationships {
	return &NullableStockTransferCreateDataRelationships{value: val, isSet: true}
}

func (v NullableStockTransferCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
