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

// PackageUpdateDataRelationships struct for PackageUpdateDataRelationships
type PackageUpdateDataRelationships struct {
	StockLocation *DeliveryLeadTimeDataRelationshipsStockLocation `json:"stock_location,omitempty"`
}

// NewPackageUpdateDataRelationships instantiates a new PackageUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageUpdateDataRelationships() *PackageUpdateDataRelationships {
	this := PackageUpdateDataRelationships{}
	return &this
}

// NewPackageUpdateDataRelationshipsWithDefaults instantiates a new PackageUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageUpdateDataRelationshipsWithDefaults() *PackageUpdateDataRelationships {
	this := PackageUpdateDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *PackageUpdateDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageUpdateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *PackageUpdateDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *PackageUpdateDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

func (o PackageUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	return json.Marshal(toSerialize)
}

type NullablePackageUpdateDataRelationships struct {
	value *PackageUpdateDataRelationships
	isSet bool
}

func (v NullablePackageUpdateDataRelationships) Get() *PackageUpdateDataRelationships {
	return v.value
}

func (v *NullablePackageUpdateDataRelationships) Set(val *PackageUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageUpdateDataRelationships(val *PackageUpdateDataRelationships) *NullablePackageUpdateDataRelationships {
	return &NullablePackageUpdateDataRelationships{value: val, isSet: true}
}

func (v NullablePackageUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
