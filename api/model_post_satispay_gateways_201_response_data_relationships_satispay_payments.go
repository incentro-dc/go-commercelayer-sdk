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

// checks if the POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments{}

// POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments struct for POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments
type POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks               `json:"links,omitempty"`
	Data  *POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData `json:"data,omitempty"`
}

// NewPOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments instantiates a new POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments() *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments {
	this := POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments{}
	return &this
}

// NewPOSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsWithDefaults instantiates a new POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsWithDefaults() *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments {
	this := POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) GetData() POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) GetDataOk() (*POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData and assigns it to the Data field.
func (o *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) SetData(v POSTSatispayGateways201ResponseDataRelationshipsSatispayPaymentsData) {
	o.Data = &v
}

func (o POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments struct {
	value *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments
	isSet bool
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) Get() *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments {
	return v.value
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) Set(val *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments(val *POSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) *NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments {
	return &NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments{value: val, isSet: true}
}

func (v NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSatispayGateways201ResponseDataRelationshipsSatispayPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}