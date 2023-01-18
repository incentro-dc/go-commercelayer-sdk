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

// GETLineItems200ResponseDataInnerRelationshipsLineItemOptions struct for GETLineItems200ResponseDataInnerRelationshipsLineItemOptions
type GETLineItems200ResponseDataInnerRelationshipsLineItemOptions struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData `json:"data,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsLineItemOptions instantiates a new GETLineItems200ResponseDataInnerRelationshipsLineItemOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsLineItemOptions() *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions {
	this := GETLineItems200ResponseDataInnerRelationshipsLineItemOptions{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsLineItemOptionsWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsLineItemOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsLineItemOptionsWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions {
	this := GETLineItems200ResponseDataInnerRelationshipsLineItemOptions{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) GetData() GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData {
	if o == nil || o.Data == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) GetDataOk() (*GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData and assigns it to the Data field.
func (o *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) SetData(v GETLineItems200ResponseDataInnerRelationshipsLineItemOptionsData) {
	o.Data = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions struct {
	value *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) Get() *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) Set(val *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions(val *GETLineItems200ResponseDataInnerRelationshipsLineItemOptions) *NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsLineItemOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
