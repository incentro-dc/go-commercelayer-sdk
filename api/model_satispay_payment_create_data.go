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

// checks if the SatispayPaymentCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SatispayPaymentCreateData{}

// SatispayPaymentCreateData struct for SatispayPaymentCreateData
type SatispayPaymentCreateData struct {
	// The resource's type
	Type          interface{}                                   `json:"type"`
	Attributes    POSTSatispayPayments201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentCreateDataRelationships          `json:"relationships,omitempty"`
}

// NewSatispayPaymentCreateData instantiates a new SatispayPaymentCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSatispayPaymentCreateData(type_ interface{}, attributes POSTSatispayPayments201ResponseDataAttributes) *SatispayPaymentCreateData {
	this := SatispayPaymentCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewSatispayPaymentCreateDataWithDefaults instantiates a new SatispayPaymentCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSatispayPaymentCreateDataWithDefaults() *SatispayPaymentCreateData {
	this := SatispayPaymentCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SatispayPaymentCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SatispayPaymentCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SatispayPaymentCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SatispayPaymentCreateData) GetAttributes() POSTSatispayPayments201ResponseDataAttributes {
	if o == nil {
		var ret POSTSatispayPayments201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SatispayPaymentCreateData) GetAttributesOk() (*POSTSatispayPayments201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SatispayPaymentCreateData) SetAttributes(v POSTSatispayPayments201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SatispayPaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenPaymentCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SatispayPaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SatispayPaymentCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentCreateDataRelationships and assigns it to the Relationships field.
func (o *SatispayPaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships) {
	o.Relationships = &v
}

func (o SatispayPaymentCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SatispayPaymentCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableSatispayPaymentCreateData struct {
	value *SatispayPaymentCreateData
	isSet bool
}

func (v NullableSatispayPaymentCreateData) Get() *SatispayPaymentCreateData {
	return v.value
}

func (v *NullableSatispayPaymentCreateData) Set(val *SatispayPaymentCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableSatispayPaymentCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableSatispayPaymentCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSatispayPaymentCreateData(val *SatispayPaymentCreateData) *NullableSatispayPaymentCreateData {
	return &NullableSatispayPaymentCreateData{value: val, isSet: true}
}

func (v NullableSatispayPaymentCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSatispayPaymentCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
