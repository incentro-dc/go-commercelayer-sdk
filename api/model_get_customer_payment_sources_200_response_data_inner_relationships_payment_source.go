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

// GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource struct for GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
type GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                  `json:"links,omitempty"`
	Data  *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData `json:"data,omitempty"`
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{}
	return &this
}

// NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceWithDefaults instantiates a new GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceWithDefaults() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	this := GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetData() GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData {
	if o == nil || o.Data == nil {
		var ret GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetDataOk() (*GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData and assigns it to the Data field.
func (o *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) SetData(v GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSourceData) {
	o.Data = &v
}

func (o GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource struct {
	value *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
	isSet bool
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Get() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return v.value
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Set(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return &NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{value: val, isSet: true}
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
