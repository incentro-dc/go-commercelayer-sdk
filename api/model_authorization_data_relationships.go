/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AuthorizationDataRelationships struct for AuthorizationDataRelationships
type AuthorizationDataRelationships struct {
	Order    *AdyenPaymentDataRelationshipsOrder     `json:"order,omitempty"`
	Captures *AuthorizationDataRelationshipsCaptures `json:"captures,omitempty"`
	Voids    *AuthorizationDataRelationshipsVoids    `json:"voids,omitempty"`
}

// NewAuthorizationDataRelationships instantiates a new AuthorizationDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationDataRelationships() *AuthorizationDataRelationships {
	this := AuthorizationDataRelationships{}
	return &this
}

// NewAuthorizationDataRelationshipsWithDefaults instantiates a new AuthorizationDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationDataRelationshipsWithDefaults() *AuthorizationDataRelationships {
	this := AuthorizationDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AuthorizationDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AuthorizationDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *AuthorizationDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetCaptures returns the Captures field value if set, zero value otherwise.
func (o *AuthorizationDataRelationships) GetCaptures() AuthorizationDataRelationshipsCaptures {
	if o == nil || o.Captures == nil {
		var ret AuthorizationDataRelationshipsCaptures
		return ret
	}
	return *o.Captures
}

// GetCapturesOk returns a tuple with the Captures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationships) GetCapturesOk() (*AuthorizationDataRelationshipsCaptures, bool) {
	if o == nil || o.Captures == nil {
		return nil, false
	}
	return o.Captures, true
}

// HasCaptures returns a boolean if a field has been set.
func (o *AuthorizationDataRelationships) HasCaptures() bool {
	if o != nil && o.Captures != nil {
		return true
	}

	return false
}

// SetCaptures gets a reference to the given AuthorizationDataRelationshipsCaptures and assigns it to the Captures field.
func (o *AuthorizationDataRelationships) SetCaptures(v AuthorizationDataRelationshipsCaptures) {
	o.Captures = &v
}

// GetVoids returns the Voids field value if set, zero value otherwise.
func (o *AuthorizationDataRelationships) GetVoids() AuthorizationDataRelationshipsVoids {
	if o == nil || o.Voids == nil {
		var ret AuthorizationDataRelationshipsVoids
		return ret
	}
	return *o.Voids
}

// GetVoidsOk returns a tuple with the Voids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationDataRelationships) GetVoidsOk() (*AuthorizationDataRelationshipsVoids, bool) {
	if o == nil || o.Voids == nil {
		return nil, false
	}
	return o.Voids, true
}

// HasVoids returns a boolean if a field has been set.
func (o *AuthorizationDataRelationships) HasVoids() bool {
	if o != nil && o.Voids != nil {
		return true
	}

	return false
}

// SetVoids gets a reference to the given AuthorizationDataRelationshipsVoids and assigns it to the Voids field.
func (o *AuthorizationDataRelationships) SetVoids(v AuthorizationDataRelationshipsVoids) {
	o.Voids = &v
}

func (o AuthorizationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Captures != nil {
		toSerialize["captures"] = o.Captures
	}
	if o.Voids != nil {
		toSerialize["voids"] = o.Voids
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationDataRelationships struct {
	value *AuthorizationDataRelationships
	isSet bool
}

func (v NullableAuthorizationDataRelationships) Get() *AuthorizationDataRelationships {
	return v.value
}

func (v *NullableAuthorizationDataRelationships) Set(val *AuthorizationDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationDataRelationships(val *AuthorizationDataRelationships) *NullableAuthorizationDataRelationships {
	return &NullableAuthorizationDataRelationships{value: val, isSet: true}
}

func (v NullableAuthorizationDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
