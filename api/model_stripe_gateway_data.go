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

// StripeGatewayData struct for StripeGatewayData
type StripeGatewayData struct {
	// The resource's type
	Type          string                          `json:"type"`
	Attributes    StripeGatewayDataAttributes     `json:"attributes"`
	Relationships *StripeGatewayDataRelationships `json:"relationships,omitempty"`
}

// NewStripeGatewayData instantiates a new StripeGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayData(type_ string, attributes StripeGatewayDataAttributes) *StripeGatewayData {
	this := StripeGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStripeGatewayDataWithDefaults instantiates a new StripeGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayDataWithDefaults() *StripeGatewayData {
	this := StripeGatewayData{}
	var type_ string = "stripe_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StripeGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripeGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StripeGatewayData) GetAttributes() StripeGatewayDataAttributes {
	if o == nil {
		var ret StripeGatewayDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayData) GetAttributesOk() (*StripeGatewayDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripeGatewayData) SetAttributes(v StripeGatewayDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripeGatewayData) GetRelationships() StripeGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret StripeGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayData) GetRelationshipsOk() (*StripeGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripeGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given StripeGatewayDataRelationships and assigns it to the Relationships field.
func (o *StripeGatewayData) SetRelationships(v StripeGatewayDataRelationships) {
	o.Relationships = &v
}

func (o StripeGatewayData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayData struct {
	value *StripeGatewayData
	isSet bool
}

func (v NullableStripeGatewayData) Get() *StripeGatewayData {
	return v.value
}

func (v *NullableStripeGatewayData) Set(val *StripeGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayData(val *StripeGatewayData) *NullableStripeGatewayData {
	return &NullableStripeGatewayData{value: val, isSet: true}
}

func (v NullableStripeGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}