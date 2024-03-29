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

// checks if the GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory{}

// GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory struct for GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory
type GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData `json:"data,omitempty"`
}

// NewGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory instantiates a new GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory() *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory {
	this := GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory{}
	return &this
}

// NewGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryWithDefaults instantiates a new GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryWithDefaults() *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory {
	this := GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) GetData() GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData {
	if o == nil || IsNil(o.Data) {
		var ret GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) GetDataOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData and assigns it to the Data field.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) SetData(v GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategoryData) {
	o.Data = &v
}

func (o GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory struct {
	value *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory
	isSet bool
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) Get() *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory {
	return v.value
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) Set(val *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory(val *GETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) *NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory {
	return &NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory{value: val, isSet: true}
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
