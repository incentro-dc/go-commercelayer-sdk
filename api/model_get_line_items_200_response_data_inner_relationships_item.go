/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETLineItems200ResponseDataInnerRelationshipsItem struct for GETLineItems200ResponseDataInnerRelationshipsItem
type GETLineItems200ResponseDataInnerRelationshipsItem struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETLineItems200ResponseDataInnerRelationshipsItemData      `json:"data,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsItem instantiates a new GETLineItems200ResponseDataInnerRelationshipsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsItem() *GETLineItems200ResponseDataInnerRelationshipsItem {
	this := GETLineItems200ResponseDataInnerRelationshipsItem{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsItemWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsItemWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsItem {
	this := GETLineItems200ResponseDataInnerRelationshipsItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) GetData() GETLineItems200ResponseDataInnerRelationshipsItemData {
	if o == nil || o.Data == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) GetDataOk() (*GETLineItems200ResponseDataInnerRelationshipsItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsItemData and assigns it to the Data field.
func (o *GETLineItems200ResponseDataInnerRelationshipsItem) SetData(v GETLineItems200ResponseDataInnerRelationshipsItemData) {
	o.Data = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsItem struct {
	value *GETLineItems200ResponseDataInnerRelationshipsItem
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItem) Get() *GETLineItems200ResponseDataInnerRelationshipsItem {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItem) Set(val *GETLineItems200ResponseDataInnerRelationshipsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsItem(val *GETLineItems200ResponseDataInnerRelationshipsItem) *NullableGETLineItems200ResponseDataInnerRelationshipsItem {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsItem{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
