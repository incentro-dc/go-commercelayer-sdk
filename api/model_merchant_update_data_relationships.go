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

// MerchantUpdateDataRelationships struct for MerchantUpdateDataRelationships
type MerchantUpdateDataRelationships struct {
	Address *BingGeocoderDataRelationshipsAddresses `json:"address,omitempty"`
}

// NewMerchantUpdateDataRelationships instantiates a new MerchantUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantUpdateDataRelationships() *MerchantUpdateDataRelationships {
	this := MerchantUpdateDataRelationships{}
	return &this
}

// NewMerchantUpdateDataRelationshipsWithDefaults instantiates a new MerchantUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantUpdateDataRelationshipsWithDefaults() *MerchantUpdateDataRelationships {
	this := MerchantUpdateDataRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MerchantUpdateDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.Address == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantUpdateDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MerchantUpdateDataRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the Address field.
func (o *MerchantUpdateDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.Address = &v
}

func (o MerchantUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantUpdateDataRelationships struct {
	value *MerchantUpdateDataRelationships
	isSet bool
}

func (v NullableMerchantUpdateDataRelationships) Get() *MerchantUpdateDataRelationships {
	return v.value
}

func (v *NullableMerchantUpdateDataRelationships) Set(val *MerchantUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantUpdateDataRelationships(val *MerchantUpdateDataRelationships) *NullableMerchantUpdateDataRelationships {
	return &NullableMerchantUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableMerchantUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
