/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShipmentData struct for ShipmentData
type ShipmentData struct {
	// The resource's type
	Type          string                                     `json:"type"`
	Attributes    GETShipments200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *ShipmentDataRelationships                 `json:"relationships,omitempty"`
}

// NewShipmentData instantiates a new ShipmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentData(type_ string, attributes GETShipments200ResponseDataInnerAttributes) *ShipmentData {
	this := ShipmentData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewShipmentDataWithDefaults instantiates a new ShipmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataWithDefaults() *ShipmentData {
	this := ShipmentData{}
	return &this
}

// GetType returns the Type field value
func (o *ShipmentData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipmentData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipmentData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ShipmentData) GetAttributes() GETShipments200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETShipments200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ShipmentData) GetAttributesOk() (*GETShipments200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ShipmentData) SetAttributes(v GETShipments200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ShipmentData) GetRelationships() ShipmentDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ShipmentDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentData) GetRelationshipsOk() (*ShipmentDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ShipmentData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ShipmentDataRelationships and assigns it to the Relationships field.
func (o *ShipmentData) SetRelationships(v ShipmentDataRelationships) {
	o.Relationships = &v
}

func (o ShipmentData) MarshalJSON() ([]byte, error) {
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

type NullableShipmentData struct {
	value *ShipmentData
	isSet bool
}

func (v NullableShipmentData) Get() *ShipmentData {
	return v.value
}

func (v *NullableShipmentData) Set(val *ShipmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentData(val *ShipmentData) *NullableShipmentData {
	return &NullableShipmentData{value: val, isSet: true}
}

func (v NullableShipmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
