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

// checks if the CouponRecipientDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponRecipientDataRelationships{}

// CouponRecipientDataRelationships struct for CouponRecipientDataRelationships
type CouponRecipientDataRelationships struct {
	Customer    *CouponRecipientDataRelationshipsCustomer   `json:"customer,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewCouponRecipientDataRelationships instantiates a new CouponRecipientDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponRecipientDataRelationships() *CouponRecipientDataRelationships {
	this := CouponRecipientDataRelationships{}
	return &this
}

// NewCouponRecipientDataRelationshipsWithDefaults instantiates a new CouponRecipientDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponRecipientDataRelationshipsWithDefaults() *CouponRecipientDataRelationships {
	this := CouponRecipientDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CouponRecipientDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponRecipientDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CouponRecipientDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *CouponRecipientDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CouponRecipientDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponRecipientDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CouponRecipientDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *CouponRecipientDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o CouponRecipientDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponRecipientDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullableCouponRecipientDataRelationships struct {
	value *CouponRecipientDataRelationships
	isSet bool
}

func (v NullableCouponRecipientDataRelationships) Get() *CouponRecipientDataRelationships {
	return v.value
}

func (v *NullableCouponRecipientDataRelationships) Set(val *CouponRecipientDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponRecipientDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponRecipientDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponRecipientDataRelationships(val *CouponRecipientDataRelationships) *NullableCouponRecipientDataRelationships {
	return &NullableCouponRecipientDataRelationships{value: val, isSet: true}
}

func (v NullableCouponRecipientDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponRecipientDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
