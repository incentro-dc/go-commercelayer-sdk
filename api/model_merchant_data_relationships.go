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

// MerchantDataRelationships struct for MerchantDataRelationships
type MerchantDataRelationships struct {
	Address     *BingGeocoderDataRelationshipsAddresses     `json:"address,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewMerchantDataRelationships instantiates a new MerchantDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDataRelationships() *MerchantDataRelationships {
	this := MerchantDataRelationships{}
	return &this
}

// NewMerchantDataRelationshipsWithDefaults instantiates a new MerchantDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDataRelationshipsWithDefaults() *MerchantDataRelationships {
	this := MerchantDataRelationships{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MerchantDataRelationships) GetAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.Address == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDataRelationships) GetAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MerchantDataRelationships) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the Address field.
func (o *MerchantDataRelationships) SetAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.Address = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *MerchantDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *MerchantDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *MerchantDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o MerchantDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantDataRelationships struct {
	value *MerchantDataRelationships
	isSet bool
}

func (v NullableMerchantDataRelationships) Get() *MerchantDataRelationships {
	return v.value
}

func (v *NullableMerchantDataRelationships) Set(val *MerchantDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDataRelationships(val *MerchantDataRelationships) *NullableMerchantDataRelationships {
	return &NullableMerchantDataRelationships{value: val, isSet: true}
}

func (v NullableMerchantDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
