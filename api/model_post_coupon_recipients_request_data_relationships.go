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

// checks if the POSTCouponRecipientsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponRecipientsRequestDataRelationships{}

// POSTCouponRecipientsRequestDataRelationships struct for POSTCouponRecipientsRequestDataRelationships
type POSTCouponRecipientsRequestDataRelationships struct {
	Customer *POSTCouponRecipientsRequestDataRelationshipsCustomer `json:"customer,omitempty"`
}

// NewPOSTCouponRecipientsRequestDataRelationships instantiates a new POSTCouponRecipientsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponRecipientsRequestDataRelationships() *POSTCouponRecipientsRequestDataRelationships {
	this := POSTCouponRecipientsRequestDataRelationships{}
	return &this
}

// NewPOSTCouponRecipientsRequestDataRelationshipsWithDefaults instantiates a new POSTCouponRecipientsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponRecipientsRequestDataRelationshipsWithDefaults() *POSTCouponRecipientsRequestDataRelationships {
	this := POSTCouponRecipientsRequestDataRelationships{}
	return &this
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *POSTCouponRecipientsRequestDataRelationships) GetCustomer() POSTCouponRecipientsRequestDataRelationshipsCustomer {
	if o == nil || IsNil(o.Customer) {
		var ret POSTCouponRecipientsRequestDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCouponRecipientsRequestDataRelationships) GetCustomerOk() (*POSTCouponRecipientsRequestDataRelationshipsCustomer, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *POSTCouponRecipientsRequestDataRelationships) HasCustomer() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given POSTCouponRecipientsRequestDataRelationshipsCustomer and assigns it to the Customer field.
func (o *POSTCouponRecipientsRequestDataRelationships) SetCustomer(v POSTCouponRecipientsRequestDataRelationshipsCustomer) {
	o.Customer = &v
}

func (o POSTCouponRecipientsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponRecipientsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	return toSerialize, nil
}

type NullablePOSTCouponRecipientsRequestDataRelationships struct {
	value *POSTCouponRecipientsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTCouponRecipientsRequestDataRelationships) Get() *POSTCouponRecipientsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTCouponRecipientsRequestDataRelationships) Set(val *POSTCouponRecipientsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponRecipientsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponRecipientsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponRecipientsRequestDataRelationships(val *POSTCouponRecipientsRequestDataRelationships) *NullablePOSTCouponRecipientsRequestDataRelationships {
	return &NullablePOSTCouponRecipientsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTCouponRecipientsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponRecipientsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}