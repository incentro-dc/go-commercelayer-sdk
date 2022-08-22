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

// ParcelCreateData struct for ParcelCreateData
type ParcelCreateData struct {
	// The resource's type
	Type          string                                   `json:"type"`
	Attributes    POSTParcels201ResponseDataAttributes     `json:"attributes"`
	Relationships *POSTParcels201ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewParcelCreateData instantiates a new ParcelCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelCreateData(type_ string, attributes POSTParcels201ResponseDataAttributes) *ParcelCreateData {
	this := ParcelCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewParcelCreateDataWithDefaults instantiates a new ParcelCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelCreateDataWithDefaults() *ParcelCreateData {
	this := ParcelCreateData{}
	var type_ string = "parcels"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ParcelCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelCreateData) GetAttributes() POSTParcels201ResponseDataAttributes {
	if o == nil {
		var ret POSTParcels201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelCreateData) GetAttributesOk() (*POSTParcels201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelCreateData) SetAttributes(v POSTParcels201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelCreateData) GetRelationships() POSTParcels201ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret POSTParcels201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelCreateData) GetRelationshipsOk() (*POSTParcels201ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTParcels201ResponseDataRelationships and assigns it to the Relationships field.
func (o *ParcelCreateData) SetRelationships(v POSTParcels201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o ParcelCreateData) MarshalJSON() ([]byte, error) {
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

type NullableParcelCreateData struct {
	value *ParcelCreateData
	isSet bool
}

func (v NullableParcelCreateData) Get() *ParcelCreateData {
	return v.value
}

func (v *NullableParcelCreateData) Set(val *ParcelCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelCreateData(val *ParcelCreateData) *NullableParcelCreateData {
	return &NullableParcelCreateData{value: val, isSet: true}
}

func (v NullableParcelCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
