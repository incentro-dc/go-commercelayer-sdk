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

// PaypalGatewayCreateData struct for PaypalGatewayCreateData
type PaypalGatewayCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTPaypalGateways201ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                 `json:"relationships,omitempty"`
}

// NewPaypalGatewayCreateData instantiates a new PaypalGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalGatewayCreateData(type_ interface{}, attributes POSTPaypalGateways201ResponseDataAttributes) *PaypalGatewayCreateData {
	this := PaypalGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPaypalGatewayCreateDataWithDefaults instantiates a new PaypalGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalGatewayCreateDataWithDefaults() *PaypalGatewayCreateData {
	this := PaypalGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PaypalGatewayCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalGatewayCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PaypalGatewayCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PaypalGatewayCreateData) GetAttributes() POSTPaypalGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTPaypalGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PaypalGatewayCreateData) GetAttributesOk() (*POSTPaypalGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PaypalGatewayCreateData) SetAttributes(v POSTPaypalGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaypalGatewayCreateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaypalGatewayCreateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PaypalGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *PaypalGatewayCreateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o PaypalGatewayCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
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

type NullablePaypalGatewayCreateData struct {
	value *PaypalGatewayCreateData
	isSet bool
}

func (v NullablePaypalGatewayCreateData) Get() *PaypalGatewayCreateData {
	return v.value
}

func (v *NullablePaypalGatewayCreateData) Set(val *PaypalGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalGatewayCreateData(val *PaypalGatewayCreateData) *NullablePaypalGatewayCreateData {
	return &NullablePaypalGatewayCreateData{value: val, isSet: true}
}

func (v NullablePaypalGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
