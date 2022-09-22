/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCustomerAddresses200ResponseDataInnerRelationshipsAddress struct for GETCustomerAddresses200ResponseDataInnerRelationshipsAddress
type GETCustomerAddresses200ResponseDataInnerRelationshipsAddress struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData `json:"data,omitempty"`
}

// NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddress instantiates a new GETCustomerAddresses200ResponseDataInnerRelationshipsAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddress() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	this := GETCustomerAddresses200ResponseDataInnerRelationshipsAddress{}
	return &this
}

// NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressWithDefaults instantiates a new GETCustomerAddresses200ResponseDataInnerRelationshipsAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerAddresses200ResponseDataInnerRelationshipsAddressWithDefaults() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	this := GETCustomerAddresses200ResponseDataInnerRelationshipsAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) GetData() GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData {
	if o == nil || o.Data == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) GetDataOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData and assigns it to the Data field.
func (o *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) SetData(v GETCustomerAddresses200ResponseDataInnerRelationshipsAddressData) {
	o.Data = &v
}

func (o GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress struct {
	value *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress
	isSet bool
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) Get() *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	return v.value
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) Set(val *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress(val *GETCustomerAddresses200ResponseDataInnerRelationshipsAddress) *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress {
	return &NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress{value: val, isSet: true}
}

func (v NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerAddresses200ResponseDataInnerRelationshipsAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}