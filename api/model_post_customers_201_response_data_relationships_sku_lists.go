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

// checks if the POSTCustomers201ResponseDataRelationshipsSkuLists type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCustomers201ResponseDataRelationshipsSkuLists{}

// POSTCustomers201ResponseDataRelationshipsSkuLists struct for POSTCustomers201ResponseDataRelationshipsSkuLists
type POSTCustomers201ResponseDataRelationshipsSkuLists struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *POSTCustomers201ResponseDataRelationshipsSkuListsData  `json:"data,omitempty"`
}

// NewPOSTCustomers201ResponseDataRelationshipsSkuLists instantiates a new POSTCustomers201ResponseDataRelationshipsSkuLists object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCustomers201ResponseDataRelationshipsSkuLists() *POSTCustomers201ResponseDataRelationshipsSkuLists {
	this := POSTCustomers201ResponseDataRelationshipsSkuLists{}
	return &this
}

// NewPOSTCustomers201ResponseDataRelationshipsSkuListsWithDefaults instantiates a new POSTCustomers201ResponseDataRelationshipsSkuLists object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCustomers201ResponseDataRelationshipsSkuListsWithDefaults() *POSTCustomers201ResponseDataRelationshipsSkuLists {
	this := POSTCustomers201ResponseDataRelationshipsSkuLists{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) GetData() POSTCustomers201ResponseDataRelationshipsSkuListsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomers201ResponseDataRelationshipsSkuListsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) GetDataOk() (*POSTCustomers201ResponseDataRelationshipsSkuListsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomers201ResponseDataRelationshipsSkuListsData and assigns it to the Data field.
func (o *POSTCustomers201ResponseDataRelationshipsSkuLists) SetData(v POSTCustomers201ResponseDataRelationshipsSkuListsData) {
	o.Data = &v
}

func (o POSTCustomers201ResponseDataRelationshipsSkuLists) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCustomers201ResponseDataRelationshipsSkuLists) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCustomers201ResponseDataRelationshipsSkuLists struct {
	value *POSTCustomers201ResponseDataRelationshipsSkuLists
	isSet bool
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) Get() *POSTCustomers201ResponseDataRelationshipsSkuLists {
	return v.value
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) Set(val *POSTCustomers201ResponseDataRelationshipsSkuLists) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCustomers201ResponseDataRelationshipsSkuLists(val *POSTCustomers201ResponseDataRelationshipsSkuLists) *NullablePOSTCustomers201ResponseDataRelationshipsSkuLists {
	return &NullablePOSTCustomers201ResponseDataRelationshipsSkuLists{value: val, isSet: true}
}

func (v NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCustomers201ResponseDataRelationshipsSkuLists) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
