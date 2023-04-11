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

// GETStockTransfers200ResponseDataInnerRelationships struct for GETStockTransfers200ResponseDataInnerRelationships
type GETStockTransfers200ResponseDataInnerRelationships struct {
	Sku                      *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku                `json:"sku,omitempty"`
	OriginStockLocation      *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation      `json:"origin_stock_location,omitempty"`
	DestinationStockLocation *GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation `json:"destination_stock_location,omitempty"`
	Shipment                 *GETParcels200ResponseDataInnerRelationshipsShipment                        `json:"shipment,omitempty"`
	LineItem                 *GETLineItemOptions200ResponseDataInnerRelationshipsLineItem                `json:"line_item,omitempty"`
	Events                   *GETAuthorizations200ResponseDataInnerRelationshipsEvents                   `json:"events,omitempty"`
}

// NewGETStockTransfers200ResponseDataInnerRelationships instantiates a new GETStockTransfers200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200ResponseDataInnerRelationships() *GETStockTransfers200ResponseDataInnerRelationships {
	this := GETStockTransfers200ResponseDataInnerRelationships{}
	return &this
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseDataInnerRelationshipsWithDefaults() *GETStockTransfers200ResponseDataInnerRelationships {
	this := GETStockTransfers200ResponseDataInnerRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetSku() GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret GETInStockSubscriptions200ResponseDataInnerRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetSkuOk() (*GETInStockSubscriptions200ResponseDataInnerRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given GETInStockSubscriptions200ResponseDataInnerRelationshipsSku and assigns it to the Sku field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetSku(v GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) {
	o.Sku = &v
}

// GetOriginStockLocation returns the OriginStockLocation field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetOriginStockLocation() GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation {
	if o == nil || o.OriginStockLocation == nil {
		var ret GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation
		return ret
	}
	return *o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetOriginStockLocationOk() (*GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation, bool) {
	if o == nil || o.OriginStockLocation == nil {
		return nil, false
	}
	return o.OriginStockLocation, true
}

// HasOriginStockLocation returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasOriginStockLocation() bool {
	if o != nil && o.OriginStockLocation != nil {
		return true
	}

	return false
}

// SetOriginStockLocation gets a reference to the given GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation and assigns it to the OriginStockLocation field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetOriginStockLocation(v GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) {
	o.OriginStockLocation = &v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetDestinationStockLocation() GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation {
	if o == nil || o.DestinationStockLocation == nil {
		var ret GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation
		return ret
	}
	return *o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetDestinationStockLocationOk() (*GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation, bool) {
	if o == nil || o.DestinationStockLocation == nil {
		return nil, false
	}
	return o.DestinationStockLocation, true
}

// HasDestinationStockLocation returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasDestinationStockLocation() bool {
	if o != nil && o.DestinationStockLocation != nil {
		return true
	}

	return false
}

// SetDestinationStockLocation gets a reference to the given GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation and assigns it to the DestinationStockLocation field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetDestinationStockLocation(v GETStockTransfers200ResponseDataInnerRelationshipsDestinationStockLocation) {
	o.DestinationStockLocation = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetShipment() GETParcels200ResponseDataInnerRelationshipsShipment {
	if o == nil || o.Shipment == nil {
		var ret GETParcels200ResponseDataInnerRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetShipmentOk() (*GETParcels200ResponseDataInnerRelationshipsShipment, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given GETParcels200ResponseDataInnerRelationshipsShipment and assigns it to the Shipment field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetShipment(v GETParcels200ResponseDataInnerRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetLineItem() GETLineItemOptions200ResponseDataInnerRelationshipsLineItem {
	if o == nil || o.LineItem == nil {
		var ret GETLineItemOptions200ResponseDataInnerRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetLineItemOk() (*GETLineItemOptions200ResponseDataInnerRelationshipsLineItem, bool) {
	if o == nil || o.LineItem == nil {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasLineItem() bool {
	if o != nil && o.LineItem != nil {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given GETLineItemOptions200ResponseDataInnerRelationshipsLineItem and assigns it to the LineItem field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetLineItem(v GETLineItemOptions200ResponseDataInnerRelationshipsLineItem) {
	o.LineItem = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetEvents() GETAuthorizations200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETAuthorizations200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) GetEventsOk() (*GETAuthorizations200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETAuthorizations200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETStockTransfers200ResponseDataInnerRelationships) SetEvents(v GETAuthorizations200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETStockTransfers200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.OriginStockLocation != nil {
		toSerialize["origin_stock_location"] = o.OriginStockLocation
	}
	if o.DestinationStockLocation != nil {
		toSerialize["destination_stock_location"] = o.DestinationStockLocation
	}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.LineItem != nil {
		toSerialize["line_item"] = o.LineItem
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockTransfers200ResponseDataInnerRelationships struct {
	value *GETStockTransfers200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationships) Get() *GETStockTransfers200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationships) Set(val *GETStockTransfers200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200ResponseDataInnerRelationships(val *GETStockTransfers200ResponseDataInnerRelationships) *NullableGETStockTransfers200ResponseDataInnerRelationships {
	return &NullableGETStockTransfers200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
