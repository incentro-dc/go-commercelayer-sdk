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

// checks if the POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers{}

// POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers struct for POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers
type POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData `json:"data,omitempty"`
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers() *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers{}
	return &this
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersWithDefaults() *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) GetData() POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData {
	if o == nil || IsNil(o.Data) {
		var ret POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) GetDataOk() (*POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData and assigns it to the Data field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) SetData(v POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiersData) {
	o.Data = &v
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers struct {
	value *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers
	isSet bool
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) Get() *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers {
	return v.value
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) Set(val *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers(val *POSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers {
	return &NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingWeightTiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
