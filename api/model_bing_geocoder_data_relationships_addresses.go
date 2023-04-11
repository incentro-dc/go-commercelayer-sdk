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

// BingGeocoderDataRelationshipsAddresses struct for BingGeocoderDataRelationshipsAddresses
type BingGeocoderDataRelationshipsAddresses struct {
	Data *BingGeocoderDataRelationshipsAddressesData `json:"data,omitempty"`
}

// NewBingGeocoderDataRelationshipsAddresses instantiates a new BingGeocoderDataRelationshipsAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBingGeocoderDataRelationshipsAddresses() *BingGeocoderDataRelationshipsAddresses {
	this := BingGeocoderDataRelationshipsAddresses{}
	return &this
}

// NewBingGeocoderDataRelationshipsAddressesWithDefaults instantiates a new BingGeocoderDataRelationshipsAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBingGeocoderDataRelationshipsAddressesWithDefaults() *BingGeocoderDataRelationshipsAddresses {
	this := BingGeocoderDataRelationshipsAddresses{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BingGeocoderDataRelationshipsAddresses) GetData() BingGeocoderDataRelationshipsAddressesData {
	if o == nil || o.Data == nil {
		var ret BingGeocoderDataRelationshipsAddressesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BingGeocoderDataRelationshipsAddresses) GetDataOk() (*BingGeocoderDataRelationshipsAddressesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BingGeocoderDataRelationshipsAddresses) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given BingGeocoderDataRelationshipsAddressesData and assigns it to the Data field.
func (o *BingGeocoderDataRelationshipsAddresses) SetData(v BingGeocoderDataRelationshipsAddressesData) {
	o.Data = &v
}

func (o BingGeocoderDataRelationshipsAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableBingGeocoderDataRelationshipsAddresses struct {
	value *BingGeocoderDataRelationshipsAddresses
	isSet bool
}

func (v NullableBingGeocoderDataRelationshipsAddresses) Get() *BingGeocoderDataRelationshipsAddresses {
	return v.value
}

func (v *NullableBingGeocoderDataRelationshipsAddresses) Set(val *BingGeocoderDataRelationshipsAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableBingGeocoderDataRelationshipsAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableBingGeocoderDataRelationshipsAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBingGeocoderDataRelationshipsAddresses(val *BingGeocoderDataRelationshipsAddresses) *NullableBingGeocoderDataRelationshipsAddresses {
	return &NullableBingGeocoderDataRelationshipsAddresses{value: val, isSet: true}
}

func (v NullableBingGeocoderDataRelationshipsAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBingGeocoderDataRelationshipsAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
