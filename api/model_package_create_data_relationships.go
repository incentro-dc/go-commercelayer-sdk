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

// PackageCreateDataRelationships struct for PackageCreateDataRelationships
type PackageCreateDataRelationships struct {
	StockLocation DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location"`
}

// NewPackageCreateDataRelationships instantiates a new PackageCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageCreateDataRelationships(stockLocation DeliveryLeadTimeDataRelationshipsStockLocation) *PackageCreateDataRelationships {
	this := PackageCreateDataRelationships{}
	this.StockLocation = stockLocation
	return &this
}

// NewPackageCreateDataRelationshipsWithDefaults instantiates a new PackageCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageCreateDataRelationshipsWithDefaults() *PackageCreateDataRelationships {
	this := PackageCreateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value
func (o *PackageCreateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}

	return o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value
// and a boolean to check if the value has been set.
func (o *PackageCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StockLocation, true
}

// SetStockLocation sets field value
func (o *PackageCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = v
}

func (o PackageCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["stock_location"] = o.StockLocation
	}
	return json.Marshal(toSerialize)
}

type NullablePackageCreateDataRelationships struct {
	value *PackageCreateDataRelationships
	isSet bool
}

func (v NullablePackageCreateDataRelationships) Get() *PackageCreateDataRelationships {
	return v.value
}

func (v *NullablePackageCreateDataRelationships) Set(val *PackageCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageCreateDataRelationships(val *PackageCreateDataRelationships) *NullablePackageCreateDataRelationships {
	return &NullablePackageCreateDataRelationships{value: val, isSet: true}
}

func (v NullablePackageCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
