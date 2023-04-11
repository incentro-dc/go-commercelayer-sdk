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

// GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules struct for GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules
type GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks               `json:"links,omitempty"`
	Data  *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData `json:"data,omitempty"`
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules() *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules{}
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesWithDefaults() *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) GetData() GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData {
	if o == nil || o.Data == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) GetDataOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData and assigns it to the Data field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) SetData(v GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRulesData) {
	o.Data = &v
}

func (o GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules struct {
	value *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) Get() *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) Set(val *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules(val *GETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) *NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules {
	return &NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsPromotionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
