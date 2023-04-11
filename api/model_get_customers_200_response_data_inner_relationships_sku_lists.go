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

// GETCustomers200ResponseDataInnerRelationshipsSkuLists struct for GETCustomers200ResponseDataInnerRelationshipsSkuLists
type GETCustomers200ResponseDataInnerRelationshipsSkuLists struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETCustomers200ResponseDataInnerRelationshipsSkuListsData  `json:"data,omitempty"`
}

// NewGETCustomers200ResponseDataInnerRelationshipsSkuLists instantiates a new GETCustomers200ResponseDataInnerRelationshipsSkuLists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomers200ResponseDataInnerRelationshipsSkuLists() *GETCustomers200ResponseDataInnerRelationshipsSkuLists {
	this := GETCustomers200ResponseDataInnerRelationshipsSkuLists{}
	return &this
}

// NewGETCustomers200ResponseDataInnerRelationshipsSkuListsWithDefaults instantiates a new GETCustomers200ResponseDataInnerRelationshipsSkuLists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomers200ResponseDataInnerRelationshipsSkuListsWithDefaults() *GETCustomers200ResponseDataInnerRelationshipsSkuLists {
	this := GETCustomers200ResponseDataInnerRelationshipsSkuLists{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) GetData() GETCustomers200ResponseDataInnerRelationshipsSkuListsData {
	if o == nil || o.Data == nil {
		var ret GETCustomers200ResponseDataInnerRelationshipsSkuListsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) GetDataOk() (*GETCustomers200ResponseDataInnerRelationshipsSkuListsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomers200ResponseDataInnerRelationshipsSkuListsData and assigns it to the Data field.
func (o *GETCustomers200ResponseDataInnerRelationshipsSkuLists) SetData(v GETCustomers200ResponseDataInnerRelationshipsSkuListsData) {
	o.Data = &v
}

func (o GETCustomers200ResponseDataInnerRelationshipsSkuLists) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists struct {
	value *GETCustomers200ResponseDataInnerRelationshipsSkuLists
	isSet bool
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) Get() *GETCustomers200ResponseDataInnerRelationshipsSkuLists {
	return v.value
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) Set(val *GETCustomers200ResponseDataInnerRelationshipsSkuLists) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomers200ResponseDataInnerRelationshipsSkuLists(val *GETCustomers200ResponseDataInnerRelationshipsSkuLists) *NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists {
	return &NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists{value: val, isSet: true}
}

func (v NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomers200ResponseDataInnerRelationshipsSkuLists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
