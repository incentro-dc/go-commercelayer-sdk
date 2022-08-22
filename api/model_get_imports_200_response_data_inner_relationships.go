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

// GETImports200ResponseDataInnerRelationships struct for GETImports200ResponseDataInnerRelationships
type GETImports200ResponseDataInnerRelationships struct {
	Events *GETCustomerAddresses200ResponseDataInnerRelationshipsEvents `json:"events,omitempty"`
}

// NewGETImports200ResponseDataInnerRelationships instantiates a new GETImports200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETImports200ResponseDataInnerRelationships() *GETImports200ResponseDataInnerRelationships {
	this := GETImports200ResponseDataInnerRelationships{}
	return &this
}

// NewGETImports200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETImports200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETImports200ResponseDataInnerRelationshipsWithDefaults() *GETImports200ResponseDataInnerRelationships {
	this := GETImports200ResponseDataInnerRelationships{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETImports200ResponseDataInnerRelationships) GetEvents() GETCustomerAddresses200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETImports200ResponseDataInnerRelationships) GetEventsOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETImports200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETImports200ResponseDataInnerRelationships) SetEvents(v GETCustomerAddresses200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETImports200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableGETImports200ResponseDataInnerRelationships struct {
	value *GETImports200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETImports200ResponseDataInnerRelationships) Get() *GETImports200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETImports200ResponseDataInnerRelationships) Set(val *GETImports200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETImports200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETImports200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETImports200ResponseDataInnerRelationships(val *GETImports200ResponseDataInnerRelationships) *NullableGETImports200ResponseDataInnerRelationships {
	return &NullableGETImports200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETImports200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETImports200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
