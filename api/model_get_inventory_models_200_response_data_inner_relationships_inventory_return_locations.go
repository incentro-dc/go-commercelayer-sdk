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

// GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations struct for GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations
type GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                      `json:"links,omitempty"`
	Data  *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData `json:"data,omitempty"`
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations{}
	return &this
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsWithDefaults instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsWithDefaults() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) GetData() GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData {
	if o == nil || o.Data == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) GetDataOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData and assigns it to the Data field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) SetData(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocationsData) {
	o.Data = &v
}

func (o GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations struct {
	value *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations
	isSet bool
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) Get() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	return v.value
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) Set(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations {
	return &NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations{value: val, isSet: true}
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryReturnLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
