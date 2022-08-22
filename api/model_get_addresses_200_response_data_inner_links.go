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

// GETAddresses200ResponseDataInnerLinks struct for GETAddresses200ResponseDataInnerLinks
type GETAddresses200ResponseDataInnerLinks struct {
	// URL
	Self *string `json:"self,omitempty"`
}

// NewGETAddresses200ResponseDataInnerLinks instantiates a new GETAddresses200ResponseDataInnerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAddresses200ResponseDataInnerLinks() *GETAddresses200ResponseDataInnerLinks {
	this := GETAddresses200ResponseDataInnerLinks{}
	return &this
}

// NewGETAddresses200ResponseDataInnerLinksWithDefaults instantiates a new GETAddresses200ResponseDataInnerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAddresses200ResponseDataInnerLinksWithDefaults() *GETAddresses200ResponseDataInnerLinks {
	this := GETAddresses200ResponseDataInnerLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *GETAddresses200ResponseDataInnerLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAddresses200ResponseDataInnerLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *GETAddresses200ResponseDataInnerLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *GETAddresses200ResponseDataInnerLinks) SetSelf(v string) {
	o.Self = &v
}

func (o GETAddresses200ResponseDataInnerLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	return json.Marshal(toSerialize)
}

type NullableGETAddresses200ResponseDataInnerLinks struct {
	value *GETAddresses200ResponseDataInnerLinks
	isSet bool
}

func (v NullableGETAddresses200ResponseDataInnerLinks) Get() *GETAddresses200ResponseDataInnerLinks {
	return v.value
}

func (v *NullableGETAddresses200ResponseDataInnerLinks) Set(val *GETAddresses200ResponseDataInnerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAddresses200ResponseDataInnerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAddresses200ResponseDataInnerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAddresses200ResponseDataInnerLinks(val *GETAddresses200ResponseDataInnerLinks) *NullableGETAddresses200ResponseDataInnerLinks {
	return &NullableGETAddresses200ResponseDataInnerLinks{value: val, isSet: true}
}

func (v NullableGETAddresses200ResponseDataInnerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAddresses200ResponseDataInnerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
