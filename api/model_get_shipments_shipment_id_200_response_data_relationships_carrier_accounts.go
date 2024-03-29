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

// checks if the GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts{}

// GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts struct for GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts
type GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                `json:"links,omitempty"`
	Data  *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData `json:"data,omitempty"`
}

// NewGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts instantiates a new GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts() *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts {
	this := GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts{}
	return &this
}

// NewGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsWithDefaults instantiates a new GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsWithDefaults() *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts {
	this := GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) GetData() GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData {
	if o == nil || IsNil(o.Data) {
		var ret GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) GetDataOk() (*GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData and assigns it to the Data field.
func (o *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) SetData(v GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccountsData) {
	o.Data = &v
}

func (o GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts struct {
	value *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts
	isSet bool
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) Get() *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts {
	return v.value
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) Set(val *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts(val *GETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) *NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts {
	return &NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts{value: val, isSet: true}
}

func (v NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipmentsShipmentId200ResponseDataRelationshipsCarrierAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
