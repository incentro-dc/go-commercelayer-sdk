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

// GETAdyenPayments200ResponseDataInnerRelationshipsOrder struct for GETAdyenPayments200ResponseDataInnerRelationshipsOrder
type GETAdyenPayments200ResponseDataInnerRelationshipsOrder struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETAdyenPayments200ResponseDataInnerRelationshipsOrderData `json:"data,omitempty"`
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsOrder instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAdyenPayments200ResponseDataInnerRelationshipsOrder() *GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsOrder{}
	return &this
}

// NewGETAdyenPayments200ResponseDataInnerRelationshipsOrderWithDefaults instantiates a new GETAdyenPayments200ResponseDataInnerRelationshipsOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAdyenPayments200ResponseDataInnerRelationshipsOrderWithDefaults() *GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	this := GETAdyenPayments200ResponseDataInnerRelationshipsOrder{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) GetData() GETAdyenPayments200ResponseDataInnerRelationshipsOrderData {
	if o == nil || o.Data == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrderData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) GetDataOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrderData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrderData and assigns it to the Data field.
func (o *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) SetData(v GETAdyenPayments200ResponseDataInnerRelationshipsOrderData) {
	o.Data = &v
}

func (o GETAdyenPayments200ResponseDataInnerRelationshipsOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder struct {
	value *GETAdyenPayments200ResponseDataInnerRelationshipsOrder
	isSet bool
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) Get() *GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	return v.value
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) Set(val *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder(val *GETAdyenPayments200ResponseDataInnerRelationshipsOrder) *NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	return &NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder{value: val, isSet: true}
}

func (v NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAdyenPayments200ResponseDataInnerRelationshipsOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
