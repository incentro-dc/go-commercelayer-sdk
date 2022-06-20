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

// BraintreeGatewayData struct for BraintreeGatewayData
type BraintreeGatewayData struct {
	// The resource's type
	Type          string                             `json:"type"`
	Attributes    BraintreeGatewayDataAttributes     `json:"attributes"`
	Relationships *BraintreeGatewayDataRelationships `json:"relationships,omitempty"`
}

// NewBraintreeGatewayData instantiates a new BraintreeGatewayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreeGatewayData(type_ string, attributes BraintreeGatewayDataAttributes) *BraintreeGatewayData {
	this := BraintreeGatewayData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewBraintreeGatewayDataWithDefaults instantiates a new BraintreeGatewayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreeGatewayDataWithDefaults() *BraintreeGatewayData {
	this := BraintreeGatewayData{}
	var type_ string = "braintree_gateways"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BraintreeGatewayData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreeGatewayData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreeGatewayData) GetAttributes() BraintreeGatewayDataAttributes {
	if o == nil {
		var ret BraintreeGatewayDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayData) GetAttributesOk() (*BraintreeGatewayDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreeGatewayData) SetAttributes(v BraintreeGatewayDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreeGatewayData) GetRelationships() BraintreeGatewayDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret BraintreeGatewayDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreeGatewayData) GetRelationshipsOk() (*BraintreeGatewayDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreeGatewayData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BraintreeGatewayDataRelationships and assigns it to the Relationships field.
func (o *BraintreeGatewayData) SetRelationships(v BraintreeGatewayDataRelationships) {
	o.Relationships = &v
}

func (o BraintreeGatewayData) MarshalJSON() ([]byte, error) {
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

type NullableBraintreeGatewayData struct {
	value *BraintreeGatewayData
	isSet bool
}

func (v NullableBraintreeGatewayData) Get() *BraintreeGatewayData {
	return v.value
}

func (v *NullableBraintreeGatewayData) Set(val *BraintreeGatewayData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreeGatewayData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreeGatewayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreeGatewayData(val *BraintreeGatewayData) *NullableBraintreeGatewayData {
	return &NullableBraintreeGatewayData{value: val, isSet: true}
}

func (v NullableBraintreeGatewayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreeGatewayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}