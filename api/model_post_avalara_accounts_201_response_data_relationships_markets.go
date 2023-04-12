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

// checks if the POSTAvalaraAccounts201ResponseDataRelationshipsMarkets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTAvalaraAccounts201ResponseDataRelationshipsMarkets{}

// POSTAvalaraAccounts201ResponseDataRelationshipsMarkets struct for POSTAvalaraAccounts201ResponseDataRelationshipsMarkets
type POSTAvalaraAccounts201ResponseDataRelationshipsMarkets struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData `json:"data,omitempty"`
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarkets instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsMarkets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarkets() *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsMarkets{}
	return &this
}

// NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsWithDefaults instantiates a new POSTAvalaraAccounts201ResponseDataRelationshipsMarkets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTAvalaraAccounts201ResponseDataRelationshipsMarketsWithDefaults() *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	this := POSTAvalaraAccounts201ResponseDataRelationshipsMarkets{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) GetData() POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) GetDataOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData and assigns it to the Data field.
func (o *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) SetData(v POSTAvalaraAccounts201ResponseDataRelationshipsMarketsData) {
	o.Data = &v
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets struct {
	value *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets
	isSet bool
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) Get() *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	return v.value
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) Set(val *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets(val *POSTAvalaraAccounts201ResponseDataRelationshipsMarkets) *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets {
	return &NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets{value: val, isSet: true}
}

func (v NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTAvalaraAccounts201ResponseDataRelationshipsMarkets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
