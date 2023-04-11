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

// SatispayGatewayCreateData struct for SatispayGatewayCreateData
type SatispayGatewayCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTManualGateways201ResponseDataAttributes `json:"attributes"`
	Relationships *SatispayGatewayCreateDataRelationships     `json:"relationships,omitempty"`
}

// NewSatispayGatewayCreateData instantiates a new SatispayGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayGatewayCreateData(type_ interface{}, attributes POSTManualGateways201ResponseDataAttributes) *SatispayGatewayCreateData {
	this := SatispayGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSatispayGatewayCreateDataWithDefaults instantiates a new SatispayGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayGatewayCreateDataWithDefaults() *SatispayGatewayCreateData {
	this := SatispayGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SatispayGatewayCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SatispayGatewayCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SatispayGatewayCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SatispayGatewayCreateData) GetAttributes() POSTManualGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTManualGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SatispayGatewayCreateData) GetAttributesOk() (*POSTManualGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SatispayGatewayCreateData) SetAttributes(v POSTManualGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SatispayGatewayCreateData) GetRelationships() SatispayGatewayCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret SatispayGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayGatewayCreateData) GetRelationshipsOk() (*SatispayGatewayCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SatispayGatewayCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SatispayGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *SatispayGatewayCreateData) SetRelationships(v SatispayGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o SatispayGatewayCreateData) MarshalJSON() ([]byte, error) {
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

type NullableSatispayGatewayCreateData struct {
	value *SatispayGatewayCreateData
	isSet bool
}

func (v NullableSatispayGatewayCreateData) Get() *SatispayGatewayCreateData {
	return v.value
}

func (v *NullableSatispayGatewayCreateData) Set(val *SatispayGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayGatewayCreateData(val *SatispayGatewayCreateData) *NullableSatispayGatewayCreateData {
	return &NullableSatispayGatewayCreateData{value: val, isSet: true}
}

func (v NullableSatispayGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
