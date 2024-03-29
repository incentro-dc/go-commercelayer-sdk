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

// checks if the POSTStockTransfers201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockTransfers201ResponseDataRelationships{}

// POSTStockTransfers201ResponseDataRelationships struct for POSTStockTransfers201ResponseDataRelationships
type POSTStockTransfers201ResponseDataRelationships struct {
	Sku                      *POSTInStockSubscriptions201ResponseDataRelationshipsSku                `json:"sku,omitempty"`
	OriginStockLocation      *POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation      `json:"origin_stock_location,omitempty"`
	DestinationStockLocation *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation `json:"destination_stock_location,omitempty"`
	Shipment                 *POSTParcels201ResponseDataRelationshipsShipment                        `json:"shipment,omitempty"`
	LineItem                 *POSTLineItemOptions201ResponseDataRelationshipsLineItem                `json:"line_item,omitempty"`
	Events                   *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents     `json:"events,omitempty"`
}

// NewPOSTStockTransfers201ResponseDataRelationships instantiates a new POSTStockTransfers201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockTransfers201ResponseDataRelationships() *POSTStockTransfers201ResponseDataRelationships {
	this := POSTStockTransfers201ResponseDataRelationships{}
	return &this
}

// NewPOSTStockTransfers201ResponseDataRelationshipsWithDefaults instantiates a new POSTStockTransfers201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockTransfers201ResponseDataRelationshipsWithDefaults() *POSTStockTransfers201ResponseDataRelationships {
	this := POSTStockTransfers201ResponseDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

// GetOriginStockLocation returns the OriginStockLocation field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocation() POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation {
	if o == nil || IsNil(o.OriginStockLocation) {
		var ret POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation
		return ret
	}
	return *o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetOriginStockLocationOk() (*POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation, bool) {
	if o == nil || IsNil(o.OriginStockLocation) {
		return nil, false
	}
	return o.OriginStockLocation, true
}

// HasOriginStockLocation returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasOriginStockLocation() bool {
	if o != nil && !IsNil(o.OriginStockLocation) {
		return true
	}

	return false
}

// SetOriginStockLocation gets a reference to the given POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation and assigns it to the OriginStockLocation field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetOriginStockLocation(v POSTStockTransfers201ResponseDataRelationshipsOriginStockLocation) {
	o.OriginStockLocation = &v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocation() POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation {
	if o == nil || IsNil(o.DestinationStockLocation) {
		var ret POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation
		return ret
	}
	return *o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetDestinationStockLocationOk() (*POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation, bool) {
	if o == nil || IsNil(o.DestinationStockLocation) {
		return nil, false
	}
	return o.DestinationStockLocation, true
}

// HasDestinationStockLocation returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasDestinationStockLocation() bool {
	if o != nil && !IsNil(o.DestinationStockLocation) {
		return true
	}

	return false
}

// SetDestinationStockLocation gets a reference to the given POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation and assigns it to the DestinationStockLocation field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetDestinationStockLocation(v POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocation) {
	o.DestinationStockLocation = &v
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetShipment() POSTParcels201ResponseDataRelationshipsShipment {
	if o == nil || IsNil(o.Shipment) {
		var ret POSTParcels201ResponseDataRelationshipsShipment
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetShipmentOk() (*POSTParcels201ResponseDataRelationshipsShipment, bool) {
	if o == nil || IsNil(o.Shipment) {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasShipment() bool {
	if o != nil && !IsNil(o.Shipment) {
		return true
	}

	return false
}

// SetShipment gets a reference to the given POSTParcels201ResponseDataRelationshipsShipment and assigns it to the Shipment field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetShipment(v POSTParcels201ResponseDataRelationshipsShipment) {
	o.Shipment = &v
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItem() POSTLineItemOptions201ResponseDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret POSTLineItemOptions201ResponseDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetLineItemOk() (*POSTLineItemOptions201ResponseDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given POSTLineItemOptions201ResponseDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetLineItem(v POSTLineItemOptions201ResponseDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *POSTStockTransfers201ResponseDataRelationships) GetEvents() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) GetEventsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *POSTStockTransfers201ResponseDataRelationships) SetEvents(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsEvents) {
	o.Events = &v
}

func (o POSTStockTransfers201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockTransfers201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.OriginStockLocation) {
		toSerialize["origin_stock_location"] = o.OriginStockLocation
	}
	if !IsNil(o.DestinationStockLocation) {
		toSerialize["destination_stock_location"] = o.DestinationStockLocation
	}
	if !IsNil(o.Shipment) {
		toSerialize["shipment"] = o.Shipment
	}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullablePOSTStockTransfers201ResponseDataRelationships struct {
	value *POSTStockTransfers201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTStockTransfers201ResponseDataRelationships) Get() *POSTStockTransfers201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationships) Set(val *POSTStockTransfers201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockTransfers201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockTransfers201ResponseDataRelationships(val *POSTStockTransfers201ResponseDataRelationships) *NullablePOSTStockTransfers201ResponseDataRelationships {
	return &NullablePOSTStockTransfers201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTStockTransfers201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
