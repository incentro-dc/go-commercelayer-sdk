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

// GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments struct for GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments
type GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks           `json:"links,omitempty"`
	Data  *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData `json:"data,omitempty"`
}

// NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments instantiates a new GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments {
	this := GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments{}
	return &this
}

// NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsWithDefaults instantiates a new GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsWithDefaults() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments {
	this := GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) GetData() GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData {
	if o == nil || o.Data == nil {
		var ret GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) GetDataOk() (*GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData and assigns it to the Data field.
func (o *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) SetData(v GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPaymentsData) {
	o.Data = &v
}

func (o GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments struct {
	value *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments
	isSet bool
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) Get() *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments {
	return v.value
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) Set(val *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments(val *GETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments {
	return &NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments{value: val, isSet: true}
}

func (v NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETKlarnaGateways200ResponseDataInnerRelationshipsKlarnaPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
