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

// ManualGatewayData struct for ManualGatewayData
type ManualGatewayData struct {
	// The resource's type
	Type          string                          `json:"type"`
	Attributes    ManualGatewayDataAttributes     `json:"attributes"`
	Relationships *ManualGatewayDataRelationships `json:"relationships,omitempty"`
}

// NewManualGatewayData instantiates a new ManualGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayData(type_ string, attributes ManualGatewayDataAttributes) *ManualGatewayData {
	this := ManualGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewManualGatewayDataWithDefaults instantiates a new ManualGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayDataWithDefaults() *ManualGatewayData {
	this := ManualGatewayData{}
	var type_ string = "manual_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ManualGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ManualGatewayData) GetAttributes() ManualGatewayDataAttributes {
	if o == nil {
		var ret ManualGatewayDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualGatewayData) GetAttributesOk() (*ManualGatewayDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualGatewayData) SetAttributes(v ManualGatewayDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualGatewayData) GetRelationships() ManualGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ManualGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGatewayData) GetRelationshipsOk() (*ManualGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualGatewayDataRelationships and assigns it to the Relationships field.
func (o *ManualGatewayData) SetRelationships(v ManualGatewayDataRelationships) {
	o.Relationships = &v
}

func (o ManualGatewayData) MarshalJSON() ([]byte, error) {
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

type NullableManualGatewayData struct {
	value *ManualGatewayData
	isSet bool
}

func (v NullableManualGatewayData) Get() *ManualGatewayData {
	return v.value
}

func (v *NullableManualGatewayData) Set(val *ManualGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayData(val *ManualGatewayData) *NullableManualGatewayData {
	return &NullableManualGatewayData{value: val, isSet: true}
}

func (v NullableManualGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
