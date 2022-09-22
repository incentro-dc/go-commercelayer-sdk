/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct for GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments
type GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                       `json:"links,omitempty"`
	Data  []GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner `json:"data,omitempty"`
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{}
	return &this
}

// NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsWithDefaults instantiates a new GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsWithDefaults() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	this := GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetData() []GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner {
	if o == nil || o.Data == nil {
		var ret []GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) GetDataOk() ([]GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner and assigns it to the Data field.
func (o *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) SetData(v []GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePaymentsDataInner) {
	o.Data = v
}

func (o GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments struct {
	value *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments
	isSet bool
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Get() *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	return v.value
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Set(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments(val *GETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments {
	return &NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments{value: val, isSet: true}
}

func (v NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBraintreeGateways200ResponseDataInnerRelationshipsBraintreePayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}