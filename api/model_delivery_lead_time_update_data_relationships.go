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

// DeliveryLeadTimeUpdateDataRelationships struct for DeliveryLeadTimeUpdateDataRelationships
type DeliveryLeadTimeUpdateDataRelationships struct {
	StockLocation  *DeliveryLeadTimeDataRelationshipsStockLocation  `json:"stock_location,omitempty"`
	ShippingMethod *DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
}

// NewDeliveryLeadTimeUpdateDataRelationships instantiates a new DeliveryLeadTimeUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryLeadTimeUpdateDataRelationships() *DeliveryLeadTimeUpdateDataRelationships {
	this := DeliveryLeadTimeUpdateDataRelationships{}
	return &this
}

// NewDeliveryLeadTimeUpdateDataRelationshipsWithDefaults instantiates a new DeliveryLeadTimeUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryLeadTimeUpdateDataRelationshipsWithDefaults() *DeliveryLeadTimeUpdateDataRelationships {
	this := DeliveryLeadTimeUpdateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *DeliveryLeadTimeUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *DeliveryLeadTimeUpdateDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *DeliveryLeadTimeUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *DeliveryLeadTimeUpdateDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil || o.ShippingMethod == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryLeadTimeUpdateDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *DeliveryLeadTimeUpdateDataRelationships) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *DeliveryLeadTimeUpdateDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

func (o DeliveryLeadTimeUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	return json.Marshal(toSerialize)
}

type NullableDeliveryLeadTimeUpdateDataRelationships struct {
	value *DeliveryLeadTimeUpdateDataRelationships
	isSet bool
}

func (v NullableDeliveryLeadTimeUpdateDataRelationships) Get() *DeliveryLeadTimeUpdateDataRelationships {
	return v.value
}

func (v *NullableDeliveryLeadTimeUpdateDataRelationships) Set(val *DeliveryLeadTimeUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryLeadTimeUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryLeadTimeUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryLeadTimeUpdateDataRelationships(val *DeliveryLeadTimeUpdateDataRelationships) *NullableDeliveryLeadTimeUpdateDataRelationships {
	return &NullableDeliveryLeadTimeUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableDeliveryLeadTimeUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryLeadTimeUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
