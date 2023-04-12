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

// checks if the POSTOrders201ResponseDataRelationshipsAvailableFreeBundles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsAvailableFreeBundles{}

// POSTOrders201ResponseDataRelationshipsAvailableFreeBundles struct for POSTOrders201ResponseDataRelationshipsAvailableFreeBundles
type POSTOrders201ResponseDataRelationshipsAvailableFreeBundles struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks         `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsAvailableFreeBundles instantiates a new POSTOrders201ResponseDataRelationshipsAvailableFreeBundles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsAvailableFreeBundles() *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles {
	this := POSTOrders201ResponseDataRelationshipsAvailableFreeBundles{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsAvailableFreeBundlesWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsAvailableFreeBundles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsAvailableFreeBundlesWithDefaults() *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles {
	this := POSTOrders201ResponseDataRelationshipsAvailableFreeBundles{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) GetData() POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) GetDataOk() (*POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) SetData(v POSTOrders201ResponseDataRelationshipsAvailableFreeBundlesData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles struct {
	value *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) Get() *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) Set(val *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles(val *POSTOrders201ResponseDataRelationshipsAvailableFreeBundles) *NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles {
	return &NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsAvailableFreeBundles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
