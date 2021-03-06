/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InStockSubscriptionData struct for InStockSubscriptionData
type InStockSubscriptionData struct {
	// The resource's type
	Type          string                                `json:"type"`
	Attributes    InStockSubscriptionDataAttributes     `json:"attributes"`
	Relationships *InStockSubscriptionDataRelationships `json:"relationships,omitempty"`
}

// NewInStockSubscriptionData instantiates a new InStockSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionData(type_ string, attributes InStockSubscriptionDataAttributes) *InStockSubscriptionData {
	this := InStockSubscriptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInStockSubscriptionDataWithDefaults instantiates a new InStockSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionDataWithDefaults() *InStockSubscriptionData {
	this := InStockSubscriptionData{}
	var type_ string = "in_stock_subscriptions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InStockSubscriptionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InStockSubscriptionData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InStockSubscriptionData) GetAttributes() InStockSubscriptionDataAttributes {
	if o == nil {
		var ret InStockSubscriptionDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionData) GetAttributesOk() (*InStockSubscriptionDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InStockSubscriptionData) SetAttributes(v InStockSubscriptionDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionData) GetRelationships() InStockSubscriptionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InStockSubscriptionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionData) GetRelationshipsOk() (*InStockSubscriptionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionData) SetRelationships(v InStockSubscriptionDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionData) MarshalJSON() ([]byte, error) {
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

type NullableInStockSubscriptionData struct {
	value *InStockSubscriptionData
	isSet bool
}

func (v NullableInStockSubscriptionData) Get() *InStockSubscriptionData {
	return v.value
}

func (v *NullableInStockSubscriptionData) Set(val *InStockSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionData(val *InStockSubscriptionData) *NullableInStockSubscriptionData {
	return &NullableInStockSubscriptionData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
