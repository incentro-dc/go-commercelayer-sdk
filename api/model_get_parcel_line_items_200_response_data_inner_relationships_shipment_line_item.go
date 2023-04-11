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

// GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem struct for GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem
type GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks              `json:"links,omitempty"`
	Data  *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData `json:"data,omitempty"`
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem{}
	return &this
}

// NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemWithDefaults instantiates a new GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemWithDefaults() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem {
	this := GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) GetData() GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData {
	if o == nil || o.Data == nil {
		var ret GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) GetDataOk() (*GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData and assigns it to the Data field.
func (o *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) SetData(v GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItemData) {
	o.Data = &v
}

func (o GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem struct {
	value *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem
	isSet bool
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) Get() *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem {
	return v.value
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) Set(val *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem(val *GETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem {
	return &NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem{value: val, isSet: true}
}

func (v NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItems200ResponseDataInnerRelationshipsShipmentLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
