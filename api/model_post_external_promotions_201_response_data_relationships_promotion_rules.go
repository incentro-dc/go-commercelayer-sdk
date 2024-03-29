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

// checks if the POSTExternalPromotions201ResponseDataRelationshipsPromotionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPromotions201ResponseDataRelationshipsPromotionRules{}

// POSTExternalPromotions201ResponseDataRelationshipsPromotionRules struct for POSTExternalPromotions201ResponseDataRelationshipsPromotionRules
type POSTExternalPromotions201ResponseDataRelationshipsPromotionRules struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks               `json:"links,omitempty"`
	Data  *POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData `json:"data,omitempty"`
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsPromotionRules instantiates a new POSTExternalPromotions201ResponseDataRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPromotions201ResponseDataRelationshipsPromotionRules() *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules {
	this := POSTExternalPromotions201ResponseDataRelationshipsPromotionRules{}
	return &this
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsPromotionRulesWithDefaults instantiates a new POSTExternalPromotions201ResponseDataRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPromotions201ResponseDataRelationshipsPromotionRulesWithDefaults() *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules {
	this := POSTExternalPromotions201ResponseDataRelationshipsPromotionRules{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) GetData() POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData {
	if o == nil || IsNil(o.Data) {
		var ret POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) GetDataOk() (*POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData and assigns it to the Data field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) SetData(v POSTExternalPromotions201ResponseDataRelationshipsPromotionRulesData) {
	o.Data = &v
}

func (o POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules struct {
	value *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules
	isSet bool
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) Get() *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules {
	return v.value
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) Set(val *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules(val *POSTExternalPromotions201ResponseDataRelationshipsPromotionRules) *NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules {
	return &NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
