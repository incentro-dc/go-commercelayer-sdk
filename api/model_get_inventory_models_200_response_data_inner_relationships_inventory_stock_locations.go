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

// GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations struct for GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations
type GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                     `json:"links,omitempty"`
	Data  *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData `json:"data,omitempty"`
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations{}
	return &this
}

// NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsWithDefaults instantiates a new GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsWithDefaults() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	this := GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) GetData() GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData {
	if o == nil || o.Data == nil {
		var ret GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) GetDataOk() (*GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData and assigns it to the Data field.
func (o *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) SetData(v GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocationsData) {
	o.Data = &v
}

func (o GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations struct {
	value *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations
	isSet bool
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) Get() *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	return v.value
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) Set(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations(val *GETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations {
	return &NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations{value: val, isSet: true}
}

func (v NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryModels200ResponseDataInnerRelationshipsInventoryStockLocations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
