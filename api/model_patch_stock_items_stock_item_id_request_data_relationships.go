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

// checks if the PATCHStockItemsStockItemIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHStockItemsStockItemIdRequestDataRelationships{}

// PATCHStockItemsStockItemIdRequestDataRelationships struct for PATCHStockItemsStockItemIdRequestDataRelationships
type PATCHStockItemsStockItemIdRequestDataRelationships struct {
	StockLocation *POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *POSTInStockSubscriptionsRequestDataRelationshipsSku        `json:"sku,omitempty"`
}

// NewPATCHStockItemsStockItemIdRequestDataRelationships instantiates a new PATCHStockItemsStockItemIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStockItemsStockItemIdRequestDataRelationships() *PATCHStockItemsStockItemIdRequestDataRelationships {
	this := PATCHStockItemsStockItemIdRequestDataRelationships{}
	return &this
}

// NewPATCHStockItemsStockItemIdRequestDataRelationshipsWithDefaults instantiates a new PATCHStockItemsStockItemIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStockItemsStockItemIdRequestDataRelationshipsWithDefaults() *PATCHStockItemsStockItemIdRequestDataRelationships {
	this := PATCHStockItemsStockItemIdRequestDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) GetStockLocation() POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation {
	if o == nil || IsNil(o.StockLocation) {
		var ret POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) GetStockLocationOk() (*POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation, bool) {
	if o == nil || IsNil(o.StockLocation) {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) HasStockLocation() bool {
	if o != nil && !IsNil(o.StockLocation) {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) SetStockLocation(v POSTDeliveryLeadTimesRequestDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) GetSku() POSTInStockSubscriptionsRequestDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptionsRequestDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) GetSkuOk() (*POSTInStockSubscriptionsRequestDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptionsRequestDataRelationshipsSku and assigns it to the Sku field.
func (o *PATCHStockItemsStockItemIdRequestDataRelationships) SetSku(v POSTInStockSubscriptionsRequestDataRelationshipsSku) {
	o.Sku = &v
}

func (o PATCHStockItemsStockItemIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHStockItemsStockItemIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StockLocation) {
		toSerialize["stock_location"] = o.StockLocation
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullablePATCHStockItemsStockItemIdRequestDataRelationships struct {
	value *PATCHStockItemsStockItemIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHStockItemsStockItemIdRequestDataRelationships) Get() *PATCHStockItemsStockItemIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHStockItemsStockItemIdRequestDataRelationships) Set(val *PATCHStockItemsStockItemIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStockItemsStockItemIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStockItemsStockItemIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStockItemsStockItemIdRequestDataRelationships(val *PATCHStockItemsStockItemIdRequestDataRelationships) *NullablePATCHStockItemsStockItemIdRequestDataRelationships {
	return &NullablePATCHStockItemsStockItemIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHStockItemsStockItemIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStockItemsStockItemIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}