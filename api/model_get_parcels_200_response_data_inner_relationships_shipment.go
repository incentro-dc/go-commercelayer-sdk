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

// GETParcels200ResponseDataInnerRelationshipsShipment struct for GETParcels200ResponseDataInnerRelationshipsShipment
type GETParcels200ResponseDataInnerRelationshipsShipment struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETParcels200ResponseDataInnerRelationshipsShipmentData    `json:"data,omitempty"`
}

// NewGETParcels200ResponseDataInnerRelationshipsShipment instantiates a new GETParcels200ResponseDataInnerRelationshipsShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcels200ResponseDataInnerRelationshipsShipment() *GETParcels200ResponseDataInnerRelationshipsShipment {
	this := GETParcels200ResponseDataInnerRelationshipsShipment{}
	return &this
}

// NewGETParcels200ResponseDataInnerRelationshipsShipmentWithDefaults instantiates a new GETParcels200ResponseDataInnerRelationshipsShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcels200ResponseDataInnerRelationshipsShipmentWithDefaults() *GETParcels200ResponseDataInnerRelationshipsShipment {
	this := GETParcels200ResponseDataInnerRelationshipsShipment{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) GetData() GETParcels200ResponseDataInnerRelationshipsShipmentData {
	if o == nil || o.Data == nil {
		var ret GETParcels200ResponseDataInnerRelationshipsShipmentData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) GetDataOk() (*GETParcels200ResponseDataInnerRelationshipsShipmentData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcels200ResponseDataInnerRelationshipsShipmentData and assigns it to the Data field.
func (o *GETParcels200ResponseDataInnerRelationshipsShipment) SetData(v GETParcels200ResponseDataInnerRelationshipsShipmentData) {
	o.Data = &v
}

func (o GETParcels200ResponseDataInnerRelationshipsShipment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcels200ResponseDataInnerRelationshipsShipment struct {
	value *GETParcels200ResponseDataInnerRelationshipsShipment
	isSet bool
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsShipment) Get() *GETParcels200ResponseDataInnerRelationshipsShipment {
	return v.value
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsShipment) Set(val *GETParcels200ResponseDataInnerRelationshipsShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcels200ResponseDataInnerRelationshipsShipment(val *GETParcels200ResponseDataInnerRelationshipsShipment) *NullableGETParcels200ResponseDataInnerRelationshipsShipment {
	return &NullableGETParcels200ResponseDataInnerRelationshipsShipment{value: val, isSet: true}
}

func (v NullableGETParcels200ResponseDataInnerRelationshipsShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcels200ResponseDataInnerRelationshipsShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
