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

// GETOrders200ResponseDataInnerRelationshipsBillingAddress struct for GETOrders200ResponseDataInnerRelationshipsBillingAddress
type GETOrders200ResponseDataInnerRelationshipsBillingAddress struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETOrders200ResponseDataInnerRelationshipsBillingAddressData `json:"data,omitempty"`
}

// NewGETOrders200ResponseDataInnerRelationshipsBillingAddress instantiates a new GETOrders200ResponseDataInnerRelationshipsBillingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrders200ResponseDataInnerRelationshipsBillingAddress() *GETOrders200ResponseDataInnerRelationshipsBillingAddress {
	this := GETOrders200ResponseDataInnerRelationshipsBillingAddress{}
	return &this
}

// NewGETOrders200ResponseDataInnerRelationshipsBillingAddressWithDefaults instantiates a new GETOrders200ResponseDataInnerRelationshipsBillingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrders200ResponseDataInnerRelationshipsBillingAddressWithDefaults() *GETOrders200ResponseDataInnerRelationshipsBillingAddress {
	this := GETOrders200ResponseDataInnerRelationshipsBillingAddress{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) GetData() GETOrders200ResponseDataInnerRelationshipsBillingAddressData {
	if o == nil || o.Data == nil {
		var ret GETOrders200ResponseDataInnerRelationshipsBillingAddressData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) GetDataOk() (*GETOrders200ResponseDataInnerRelationshipsBillingAddressData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETOrders200ResponseDataInnerRelationshipsBillingAddressData and assigns it to the Data field.
func (o *GETOrders200ResponseDataInnerRelationshipsBillingAddress) SetData(v GETOrders200ResponseDataInnerRelationshipsBillingAddressData) {
	o.Data = &v
}

func (o GETOrders200ResponseDataInnerRelationshipsBillingAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress struct {
	value *GETOrders200ResponseDataInnerRelationshipsBillingAddress
	isSet bool
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) Get() *GETOrders200ResponseDataInnerRelationshipsBillingAddress {
	return v.value
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) Set(val *GETOrders200ResponseDataInnerRelationshipsBillingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrders200ResponseDataInnerRelationshipsBillingAddress(val *GETOrders200ResponseDataInnerRelationshipsBillingAddress) *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress {
	return &NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress{value: val, isSet: true}
}

func (v NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrders200ResponseDataInnerRelationshipsBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}