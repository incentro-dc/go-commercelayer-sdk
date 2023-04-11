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

// checks if the POSTParcelsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelsRequestDataRelationships{}

// POSTParcelsRequestDataRelationships struct for POSTParcelsRequestDataRelationships
type POSTParcelsRequestDataRelationships struct {
	Shipment POSTParcelsRequestDataRelationshipsShipment `json:"shipment"`
	Package  POSTParcelsRequestDataRelationshipsPackage  `json:"package"`
}

// NewPOSTParcelsRequestDataRelationships instantiates a new POSTParcelsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelsRequestDataRelationships(shipment POSTParcelsRequestDataRelationshipsShipment, package_ POSTParcelsRequestDataRelationshipsPackage) *POSTParcelsRequestDataRelationships {
	this := POSTParcelsRequestDataRelationships{}
	this.Shipment = shipment
	this.Package = package_
	return &this
}

// NewPOSTParcelsRequestDataRelationshipsWithDefaults instantiates a new POSTParcelsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelsRequestDataRelationshipsWithDefaults() *POSTParcelsRequestDataRelationships {
	this := POSTParcelsRequestDataRelationships{}
	return &this
}

// GetShipment returns the Shipment field value
func (o *POSTParcelsRequestDataRelationships) GetShipment() POSTParcelsRequestDataRelationshipsShipment {
	if o == nil {
		var ret POSTParcelsRequestDataRelationshipsShipment
		return ret
	}

	return o.Shipment
}

// GetShipmentOk returns a tuple with the Shipment field value
// and a boolean to check if the value has been set.
func (o *POSTParcelsRequestDataRelationships) GetShipmentOk() (*POSTParcelsRequestDataRelationshipsShipment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Shipment, true
}

// SetShipment sets field value
func (o *POSTParcelsRequestDataRelationships) SetShipment(v POSTParcelsRequestDataRelationshipsShipment) {
	o.Shipment = v
}

// GetPackage returns the Package field value
func (o *POSTParcelsRequestDataRelationships) GetPackage() POSTParcelsRequestDataRelationshipsPackage {
	if o == nil {
		var ret POSTParcelsRequestDataRelationshipsPackage
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *POSTParcelsRequestDataRelationships) GetPackageOk() (*POSTParcelsRequestDataRelationshipsPackage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *POSTParcelsRequestDataRelationships) SetPackage(v POSTParcelsRequestDataRelationshipsPackage) {
	o.Package = v
}

func (o POSTParcelsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipment"] = o.Shipment
	toSerialize["package"] = o.Package
	return toSerialize, nil
}

type NullablePOSTParcelsRequestDataRelationships struct {
	value *POSTParcelsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTParcelsRequestDataRelationships) Get() *POSTParcelsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTParcelsRequestDataRelationships) Set(val *POSTParcelsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelsRequestDataRelationships(val *POSTParcelsRequestDataRelationships) *NullablePOSTParcelsRequestDataRelationships {
	return &NullablePOSTParcelsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTParcelsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}