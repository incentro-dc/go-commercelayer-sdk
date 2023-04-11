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

// GETCustomers200ResponseDataInnerRelationshipsOrders struct for GETCustomers200ResponseDataInnerRelationshipsOrders
type GETCustomers200ResponseDataInnerRelationshipsOrders struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETCustomers200ResponseDataInnerRelationshipsOrdersData    `json:"data,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrders instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsOrders() *GETCustomers200ResponseDataInnerRelationshipsOrders {
	this := GETCustomers200ResponseDataInnerRelationshipsOrders{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsOrdersWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsOrdersWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsOrders {
	this := GETCustomers200ResponseDataInnerRelationshipsOrders{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) GetData() GETCustomers200ResponseDataInnerRelationshipsOrdersData {
	if o == nil || o.Data == nil {
		var ret GETCustomers200ResponseDataInnerRelationshipsOrdersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) GetDataOk() (*GETCustomers200ResponseDataInnerRelationshipsOrdersData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomers200ResponseDataInnerRelationshipsOrdersData and assigns it to the Data field.
func (o *GETCustomers200ResponseDataInnerRelationshipsOrders) SetData(v GETCustomers200ResponseDataInnerRelationshipsOrdersData) {
	o.Data = &v
}

func (o GETCustomers200ResponseDataInnerRelationshipsOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsOrders struct {
	value *GETCustomers200ResponseDataInnerRelationshipsOrders
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrders) Get() *GETCustomers200ResponseDataInnerRelationshipsOrders {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrders) Set(val *GETCustomers200ResponseDataInnerRelationshipsOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsOrders(val *GETCustomers200ResponseDataInnerRelationshipsOrders) *NullableGETCustomers200ResponseDataInnerRelationshipsOrders {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsOrders{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
