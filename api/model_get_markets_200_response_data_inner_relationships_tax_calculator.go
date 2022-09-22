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

// GETMarkets200ResponseDataInnerRelationshipsTaxCalculator struct for GETMarkets200ResponseDataInnerRelationshipsTaxCalculator
type GETMarkets200ResponseDataInnerRelationshipsTaxCalculator struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData `json:"data,omitempty"`
}

// NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculator instantiates a new GETMarkets200ResponseDataInnerRelationshipsTaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculator() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator {
	this := GETMarkets200ResponseDataInnerRelationshipsTaxCalculator{}
	return &this
}

// NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorWithDefaults instantiates a new GETMarkets200ResponseDataInnerRelationshipsTaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETMarkets200ResponseDataInnerRelationshipsTaxCalculatorWithDefaults() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator {
	this := GETMarkets200ResponseDataInnerRelationshipsTaxCalculator{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) GetData() GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData {
	if o == nil || o.Data == nil {
		var ret GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) GetDataOk() (*GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData and assigns it to the Data field.
func (o *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) SetData(v GETMarkets200ResponseDataInnerRelationshipsTaxCalculatorData) {
	o.Data = &v
}

func (o GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator struct {
	value *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator
	isSet bool
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) Get() *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator {
	return v.value
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) Set(val *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator(val *GETMarkets200ResponseDataInnerRelationshipsTaxCalculator) *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator {
	return &NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator{value: val, isSet: true}
}

func (v NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETMarkets200ResponseDataInnerRelationshipsTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}