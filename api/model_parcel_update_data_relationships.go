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

// ParcelUpdateDataRelationships struct for ParcelUpdateDataRelationships
type ParcelUpdateDataRelationships struct {
	Shipment *OrderDataRelationshipsShipments `json:"shipment,omitempty"`
	Package  *ParcelDataRelationshipsPackage  `json:"package,omitempty"`
}

// NewParcelUpdateDataRelationships instantiates a new ParcelUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelUpdateDataRelationships() *ParcelUpdateDataRelationships {
	this := ParcelUpdateDataRelationships{}
	return &this
}

// NewParcelUpdateDataRelationshipsWithDefaults instantiates a new ParcelUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelUpdateDataRelationshipsWithDefaults() *ParcelUpdateDataRelationships {
	this := ParcelUpdateDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value if set, zero value otherwise.
func (o *ParcelUpdateDataRelationships) GetShipment() OrderDataRelationshipsShipments {
	if o == nil || o.Shipment == nil {
		var ret OrderDataRelationshipsShipments
		return ret
	}
	return *o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelUpdateDataRelationships) GetShipmentOk() (*OrderDataRelationshipsShipments, bool) {
	if o == nil || o.Shipment == nil {
		return nil, false
	}
	return o.Shipment, true
}

// HasShipment returns a boolean if a field has been set.
func (o *ParcelUpdateDataRelationships) HasShipment() bool {
	if o != nil && o.Shipment != nil {
		return true
	}

	return false
}

// SetShipment gets a reference to the given OrderDataRelationshipsShipments and assigns it to the Shipment field.
func (o *ParcelUpdateDataRelationships) SetShipment(v OrderDataRelationshipsShipments) {
	o.Shipment = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *ParcelUpdateDataRelationships) GetPackage() ParcelDataRelationshipsPackage {
	if o == nil || o.Package == nil {
		var ret ParcelDataRelationshipsPackage
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelUpdateDataRelationships) GetPackageOk() (*ParcelDataRelationshipsPackage, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *ParcelUpdateDataRelationships) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given ParcelDataRelationshipsPackage and assigns it to the Package field.
func (o *ParcelUpdateDataRelationships) SetPackage(v ParcelDataRelationshipsPackage) {
	o.Package = &v
}

func (o ParcelUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Shipment != nil {
		toSerialize["shipment"] = o.Shipment
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	return json.Marshal(toSerialize)
}

type NullableParcelUpdateDataRelationships struct {
	value *ParcelUpdateDataRelationships
	isSet bool
}

func (v NullableParcelUpdateDataRelationships) Get() *ParcelUpdateDataRelationships {
	return v.value
}

func (v *NullableParcelUpdateDataRelationships) Set(val *ParcelUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelUpdateDataRelationships(val *ParcelUpdateDataRelationships) *NullableParcelUpdateDataRelationships {
	return &NullableParcelUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableParcelUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
