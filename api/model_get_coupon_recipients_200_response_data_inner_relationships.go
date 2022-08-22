/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCouponRecipients200ResponseDataInnerRelationships struct for GETCouponRecipients200ResponseDataInnerRelationships
type GETCouponRecipients200ResponseDataInnerRelationships struct {
	Customer    *GETCouponRecipients200ResponseDataInnerRelationshipsCustomer   `json:"customer,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETCouponRecipients200ResponseDataInnerRelationships instantiates a new GETCouponRecipients200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponRecipients200ResponseDataInnerRelationships() *GETCouponRecipients200ResponseDataInnerRelationships {
	this := GETCouponRecipients200ResponseDataInnerRelationships{}
	return &this
}

// NewGETCouponRecipients200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETCouponRecipients200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponRecipients200ResponseDataInnerRelationshipsWithDefaults() *GETCouponRecipients200ResponseDataInnerRelationships {
	this := GETCouponRecipients200ResponseDataInnerRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) GetCustomer() GETCouponRecipients200ResponseDataInnerRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret GETCouponRecipients200ResponseDataInnerRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) GetCustomerOk() (*GETCouponRecipients200ResponseDataInnerRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given GETCouponRecipients200ResponseDataInnerRelationshipsCustomer and assigns it to the Customer field.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) SetCustomer(v GETCouponRecipients200ResponseDataInnerRelationshipsCustomer) {
	o.Customer = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETCouponRecipients200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETCouponRecipients200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponRecipients200ResponseDataInnerRelationships struct {
	value *GETCouponRecipients200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationships) Get() *GETCouponRecipients200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationships) Set(val *GETCouponRecipients200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponRecipients200ResponseDataInnerRelationships(val *GETCouponRecipients200ResponseDataInnerRelationships) *NullableGETCouponRecipients200ResponseDataInnerRelationships {
	return &NullableGETCouponRecipients200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETCouponRecipients200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponRecipients200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
