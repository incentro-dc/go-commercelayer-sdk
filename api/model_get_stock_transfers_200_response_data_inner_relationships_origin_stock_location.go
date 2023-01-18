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

// GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation struct for GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation
type GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData `json:"data,omitempty"`
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation {
	this := GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation{}
	return &this
}

// NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationWithDefaults instantiates a new GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationWithDefaults() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation {
	this := GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) GetData() GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData {
	if o == nil || o.Data == nil {
		var ret GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) GetDataOk() (*GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData and assigns it to the Data field.
func (o *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) SetData(v GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocationData) {
	o.Data = &v
}

func (o GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation struct {
	value *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation
	isSet bool
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) Get() *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation {
	return v.value
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) Set(val *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation(val *GETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation {
	return &NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation{value: val, isSet: true}
}

func (v NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockTransfers200ResponseDataInnerRelationshipsOriginStockLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
