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

// PATCHPackagesPackageId200ResponseDataRelationships struct for PATCHPackagesPackageId200ResponseDataRelationships
type PATCHPackagesPackageId200ResponseDataRelationships struct {
	StockLocation *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation `json:"stock_location,omitempty"`
}

// NewPATCHPackagesPackageId200ResponseDataRelationships instantiates a new PATCHPackagesPackageId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPackagesPackageId200ResponseDataRelationships() *PATCHPackagesPackageId200ResponseDataRelationships {
	this := PATCHPackagesPackageId200ResponseDataRelationships{}
	return &this
}

// NewPATCHPackagesPackageId200ResponseDataRelationshipsWithDefaults instantiates a new PATCHPackagesPackageId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPackagesPackageId200ResponseDataRelationshipsWithDefaults() *PATCHPackagesPackageId200ResponseDataRelationships {
	this := PATCHPackagesPackageId200ResponseDataRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *PATCHPackagesPackageId200ResponseDataRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPackagesPackageId200ResponseDataRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *PATCHPackagesPackageId200ResponseDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *PATCHPackagesPackageId200ResponseDataRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

func (o PATCHPackagesPackageId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHPackagesPackageId200ResponseDataRelationships struct {
	value *PATCHPackagesPackageId200ResponseDataRelationships
	isSet bool
}

func (v NullablePATCHPackagesPackageId200ResponseDataRelationships) Get() *PATCHPackagesPackageId200ResponseDataRelationships {
	return v.value
}

func (v *NullablePATCHPackagesPackageId200ResponseDataRelationships) Set(val *PATCHPackagesPackageId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPackagesPackageId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPackagesPackageId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPackagesPackageId200ResponseDataRelationships(val *PATCHPackagesPackageId200ResponseDataRelationships) *NullablePATCHPackagesPackageId200ResponseDataRelationships {
	return &NullablePATCHPackagesPackageId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHPackagesPackageId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPackagesPackageId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}