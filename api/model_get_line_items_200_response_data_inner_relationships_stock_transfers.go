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

// GETLineItems200ResponseDataInnerRelationshipsStockTransfers struct for GETLineItems200ResponseDataInnerRelationshipsStockTransfers
type GETLineItems200ResponseDataInnerRelationshipsStockTransfers struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *GETLineItems200ResponseDataInnerRelationshipsStockTransfersData `json:"data,omitempty"`
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfers instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfers() *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfers{}
	return &this
}

// NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersWithDefaults instantiates a new GETLineItems200ResponseDataInnerRelationshipsStockTransfers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLineItems200ResponseDataInnerRelationshipsStockTransfersWithDefaults() *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	this := GETLineItems200ResponseDataInnerRelationshipsStockTransfers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetData() GETLineItems200ResponseDataInnerRelationshipsStockTransfersData {
	if o == nil || o.Data == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsStockTransfersData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) GetDataOk() (*GETLineItems200ResponseDataInnerRelationshipsStockTransfersData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsStockTransfersData and assigns it to the Data field.
func (o *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) SetData(v GETLineItems200ResponseDataInnerRelationshipsStockTransfersData) {
	o.Data = &v
}

func (o GETLineItems200ResponseDataInnerRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers struct {
	value *GETLineItems200ResponseDataInnerRelationshipsStockTransfers
	isSet bool
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Get() *GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	return v.value
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Set(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers(val *GETLineItems200ResponseDataInnerRelationshipsStockTransfers) *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	return &NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers{value: val, isSet: true}
}

func (v NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLineItems200ResponseDataInnerRelationshipsStockTransfers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
