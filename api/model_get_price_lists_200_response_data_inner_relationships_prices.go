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

// GETPriceLists200ResponseDataInnerRelationshipsPrices struct for GETPriceLists200ResponseDataInnerRelationshipsPrices
type GETPriceLists200ResponseDataInnerRelationshipsPrices struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks `json:"links,omitempty"`
	Data  *GETPriceLists200ResponseDataInnerRelationshipsPricesData   `json:"data,omitempty"`
}

// NewGETPriceLists200ResponseDataInnerRelationshipsPrices instantiates a new GETPriceLists200ResponseDataInnerRelationshipsPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceLists200ResponseDataInnerRelationshipsPrices() *GETPriceLists200ResponseDataInnerRelationshipsPrices {
	this := GETPriceLists200ResponseDataInnerRelationshipsPrices{}
	return &this
}

// NewGETPriceLists200ResponseDataInnerRelationshipsPricesWithDefaults instantiates a new GETPriceLists200ResponseDataInnerRelationshipsPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceLists200ResponseDataInnerRelationshipsPricesWithDefaults() *GETPriceLists200ResponseDataInnerRelationshipsPrices {
	this := GETPriceLists200ResponseDataInnerRelationshipsPrices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) GetData() GETPriceLists200ResponseDataInnerRelationshipsPricesData {
	if o == nil || o.Data == nil {
		var ret GETPriceLists200ResponseDataInnerRelationshipsPricesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) GetDataOk() (*GETPriceLists200ResponseDataInnerRelationshipsPricesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETPriceLists200ResponseDataInnerRelationshipsPricesData and assigns it to the Data field.
func (o *GETPriceLists200ResponseDataInnerRelationshipsPrices) SetData(v GETPriceLists200ResponseDataInnerRelationshipsPricesData) {
	o.Data = &v
}

func (o GETPriceLists200ResponseDataInnerRelationshipsPrices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETPriceLists200ResponseDataInnerRelationshipsPrices struct {
	value *GETPriceLists200ResponseDataInnerRelationshipsPrices
	isSet bool
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) Get() *GETPriceLists200ResponseDataInnerRelationshipsPrices {
	return v.value
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) Set(val *GETPriceLists200ResponseDataInnerRelationshipsPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceLists200ResponseDataInnerRelationshipsPrices(val *GETPriceLists200ResponseDataInnerRelationshipsPrices) *NullableGETPriceLists200ResponseDataInnerRelationshipsPrices {
	return &NullableGETPriceLists200ResponseDataInnerRelationshipsPrices{value: val, isSet: true}
}

func (v NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceLists200ResponseDataInnerRelationshipsPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
