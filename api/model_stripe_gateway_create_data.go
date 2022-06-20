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

// StripeGatewayCreateData struct for StripeGatewayCreateData
type StripeGatewayCreateData struct {
	// The resource's type
	Type          string                            `json:"type"`
	Attributes    StripeGatewayCreateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}            `json:"relationships,omitempty"`
}

// NewStripeGatewayCreateData instantiates a new StripeGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayCreateData(type_ string, attributes StripeGatewayCreateDataAttributes) *StripeGatewayCreateData {
	this := StripeGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewStripeGatewayCreateDataWithDefaults instantiates a new StripeGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayCreateDataWithDefaults() *StripeGatewayCreateData {
	this := StripeGatewayCreateData{}
	var type_ string = "stripe_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StripeGatewayCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StripeGatewayCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *StripeGatewayCreateData) GetAttributes() StripeGatewayCreateDataAttributes {
	if o == nil {
		var ret StripeGatewayCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StripeGatewayCreateData) GetAttributesOk() (*StripeGatewayCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StripeGatewayCreateData) SetAttributes(v StripeGatewayCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StripeGatewayCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StripeGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *StripeGatewayCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o StripeGatewayCreateData) MarshalJSON() ([]byte, error) {
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

type NullableStripeGatewayCreateData struct {
	value *StripeGatewayCreateData
	isSet bool
}

func (v NullableStripeGatewayCreateData) Get() *StripeGatewayCreateData {
	return v.value
}

func (v *NullableStripeGatewayCreateData) Set(val *StripeGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayCreateData(val *StripeGatewayCreateData) *NullableStripeGatewayCreateData {
	return &NullableStripeGatewayCreateData{value: val, isSet: true}
}

func (v NullableStripeGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}