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

// ManualTaxCalculatorUpdateData struct for ManualTaxCalculatorUpdateData
type ManualTaxCalculatorUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                             `json:"id"`
	Attributes    PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes `json:"attributes"`
	Relationships *ManualTaxCalculatorCreateDataRelationships                             `json:"relationships,omitempty"`
}

// NewManualTaxCalculatorUpdateData instantiates a new ManualTaxCalculatorUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorUpdateData(type_ interface{}, id interface{}, attributes PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) *ManualTaxCalculatorUpdateData {
	this := ManualTaxCalculatorUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewManualTaxCalculatorUpdateDataWithDefaults instantiates a new ManualTaxCalculatorUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorUpdateDataWithDefaults() *ManualTaxCalculatorUpdateData {
	this := ManualTaxCalculatorUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualTaxCalculatorUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualTaxCalculatorUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ManualTaxCalculatorUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ManualTaxCalculatorUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManualTaxCalculatorUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ManualTaxCalculatorUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ManualTaxCalculatorUpdateData) GetAttributes() PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorUpdateData) GetAttributesOk() (*PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ManualTaxCalculatorUpdateData) SetAttributes(v PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ManualTaxCalculatorUpdateData) GetRelationships() ManualTaxCalculatorCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret ManualTaxCalculatorCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorUpdateData) GetRelationshipsOk() (*ManualTaxCalculatorCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ManualTaxCalculatorUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ManualTaxCalculatorCreateDataRelationships and assigns it to the Relationships field.
func (o *ManualTaxCalculatorUpdateData) SetRelationships(v ManualTaxCalculatorCreateDataRelationships) {
	o.Relationships = &v
}

func (o ManualTaxCalculatorUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorUpdateData struct {
	value *ManualTaxCalculatorUpdateData
	isSet bool
}

func (v NullableManualTaxCalculatorUpdateData) Get() *ManualTaxCalculatorUpdateData {
	return v.value
}

func (v *NullableManualTaxCalculatorUpdateData) Set(val *ManualTaxCalculatorUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorUpdateData(val *ManualTaxCalculatorUpdateData) *NullableManualTaxCalculatorUpdateData {
	return &NullableManualTaxCalculatorUpdateData{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
