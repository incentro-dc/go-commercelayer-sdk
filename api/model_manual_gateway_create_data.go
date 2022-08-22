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

// ManualGatewayCreateData struct for ManualGatewayCreateData
type ManualGatewayCreateData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    POSTManualGateways201ResponseDataAttributes `json:"attributes"`
	Relationships map[string]interface{}                      `json:"relationships,omitempty"`
}

// NewManualGatewayCreateData instantiates a new ManualGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayCreateData(type_ string, attributes POSTManualGateways201ResponseDataAttributes) *ManualGatewayCreateData {
	this := ManualGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualGatewayCreateDataWithDefaults instantiates a new ManualGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayCreateDataWithDefaults() *ManualGatewayCreateData {
	this := ManualGatewayCreateData{}
	var type_ string = "manual_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ManualGatewayCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualGatewayCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualGatewayCreateData) GetAttributes() POSTManualGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTManualGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayCreateData) GetAttributesOk() (*POSTManualGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualGatewayCreateData) SetAttributes(v POSTManualGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualGatewayCreateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGatewayCreateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ManualGatewayCreateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ManualGatewayCreateData) MarshalJSON() ([]byte, error) {
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

type NullableManualGatewayCreateData struct {
	value *ManualGatewayCreateData
	isSet bool
}

func (v NullableManualGatewayCreateData) Get() *ManualGatewayCreateData {
	return v.value
}

func (v *NullableManualGatewayCreateData) Set(val *ManualGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayCreateData(val *ManualGatewayCreateData) *NullableManualGatewayCreateData {
	return &NullableManualGatewayCreateData{value: val, isSet: true}
}

func (v NullableManualGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
