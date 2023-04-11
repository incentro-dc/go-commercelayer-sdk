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

// checks if the PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships{}

// PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships struct for PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships
type PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships struct {
	Order *POSTAdyenPaymentsRequestDataRelationshipsOrder `json:"order,omitempty"`
}

// NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships instantiates a new PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships() *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships {
	this := PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships{}
	return &this
}

// NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationshipsWithDefaults instantiates a new PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationshipsWithDefaults() *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships {
	this := PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) GetOrder() POSTAdyenPaymentsRequestDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPaymentsRequestDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) GetOrderOk() (*POSTAdyenPaymentsRequestDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPaymentsRequestDataRelationshipsOrder and assigns it to the Order field.
func (o *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) SetOrder(v POSTAdyenPaymentsRequestDataRelationshipsOrder) {
	o.Order = &v
}

func (o PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships struct {
	value *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) Get() *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) Set(val *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships(val *PATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) *NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships {
	return &NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAdyenPaymentsAdyenPaymentIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}