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

// StockTransferUpdateDataRelationships struct for StockTransferUpdateDataRelationships
type StockTransferUpdateDataRelationships struct {
	Sku                      *InStockSubscriptionCreateDataRelationshipsSku        `json:"sku,omitempty"`
	OriginStockLocation      *DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"origin_stock_location,omitempty"`
	DestinationStockLocation *DeliveryLeadTimeCreateDataRelationshipsStockLocation `json:"destination_stock_location,omitempty"`
}

// NewStockTransferUpdateDataRelationships instantiates a new StockTransferUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferUpdateDataRelationships() *StockTransferUpdateDataRelationships {
	this := StockTransferUpdateDataRelationships{}
	return &this
}

// NewStockTransferUpdateDataRelationshipsWithDefaults instantiates a new StockTransferUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferUpdateDataRelationshipsWithDefaults() *StockTransferUpdateDataRelationships {
	this := StockTransferUpdateDataRelationships{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *StockTransferUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret InStockSubscriptionCreateDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *StockTransferUpdateDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given InStockSubscriptionCreateDataRelationshipsSku and assigns it to the Sku field.
func (o *StockTransferUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku) {
	o.Sku = &v
}

// GetOriginStockLocation returns the OriginStockLocation field value if set, zero value otherwise.
func (o *StockTransferUpdateDataRelationships) GetOriginStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil || o.OriginStockLocation == nil {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}
	return *o.OriginStockLocation
}

// GetOriginStockLocationOk returns a tuple with the OriginStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateDataRelationships) GetOriginStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil || o.OriginStockLocation == nil {
		return nil, false
	}
	return o.OriginStockLocation, true
}

// HasOriginStockLocation returns a boolean if a field has been set.
func (o *StockTransferUpdateDataRelationships) HasOriginStockLocation() bool {
	if o != nil && o.OriginStockLocation != nil {
		return true
	}

	return false
}

// SetOriginStockLocation gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsStockLocation and assigns it to the OriginStockLocation field.
func (o *StockTransferUpdateDataRelationships) SetOriginStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.OriginStockLocation = &v
}

// GetDestinationStockLocation returns the DestinationStockLocation field value if set, zero value otherwise.
func (o *StockTransferUpdateDataRelationships) GetDestinationStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation {
	if o == nil || o.DestinationStockLocation == nil {
		var ret DeliveryLeadTimeCreateDataRelationshipsStockLocation
		return ret
	}
	return *o.DestinationStockLocation
}

// GetDestinationStockLocationOk returns a tuple with the DestinationStockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateDataRelationships) GetDestinationStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool) {
	if o == nil || o.DestinationStockLocation == nil {
		return nil, false
	}
	return o.DestinationStockLocation, true
}

// HasDestinationStockLocation returns a boolean if a field has been set.
func (o *StockTransferUpdateDataRelationships) HasDestinationStockLocation() bool {
	if o != nil && o.DestinationStockLocation != nil {
		return true
	}

	return false
}

// SetDestinationStockLocation gets a reference to the given DeliveryLeadTimeCreateDataRelationshipsStockLocation and assigns it to the DestinationStockLocation field.
func (o *StockTransferUpdateDataRelationships) SetDestinationStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation) {
	o.DestinationStockLocation = &v
}

func (o StockTransferUpdateDataRelationships) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableStockTransferUpdateDataRelationships struct {
	value *StockTransferUpdateDataRelationships
	isSet bool
}

func (v NullableStockTransferUpdateDataRelationships) Get() *StockTransferUpdateDataRelationships {
	return v.value
}

func (v *NullableStockTransferUpdateDataRelationships) Set(val *StockTransferUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferUpdateDataRelationships(val *StockTransferUpdateDataRelationships) *NullableStockTransferUpdateDataRelationships {
	return &NullableStockTransferUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableStockTransferUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
