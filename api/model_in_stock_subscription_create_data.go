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

// InStockSubscriptionCreateData struct for InStockSubscriptionCreateData
type InStockSubscriptionCreateData struct {
	// The resource's type
	Type          string                                      `json:"type"`
	Attributes    InStockSubscriptionCreateDataAttributes     `json:"attributes"`
	Relationships *InStockSubscriptionCreateDataRelationships `json:"relationships,omitempty"`
}

// NewInStockSubscriptionCreateData instantiates a new InStockSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreateData(type_ string, attributes InStockSubscriptionCreateDataAttributes) *InStockSubscriptionCreateData {
	this := InStockSubscriptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInStockSubscriptionCreateDataWithDefaults instantiates a new InStockSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateDataWithDefaults() *InStockSubscriptionCreateData {
	this := InStockSubscriptionCreateData{}
	var type_ string = "in_stock_subscriptions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *InStockSubscriptionCreateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InStockSubscriptionCreateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InStockSubscriptionCreateData) GetAttributes() InStockSubscriptionCreateDataAttributes {
	if o == nil {
		var ret InStockSubscriptionCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateData) GetAttributesOk() (*InStockSubscriptionCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InStockSubscriptionCreateData) SetAttributes(v InStockSubscriptionCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateData) GetRelationships() InStockSubscriptionCreateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret InStockSubscriptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateData) GetRelationshipsOk() (*InStockSubscriptionCreateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionCreateDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionCreateData) SetRelationships(v InStockSubscriptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionCreateData) MarshalJSON() ([]byte, error) {
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

type NullableInStockSubscriptionCreateData struct {
	value *InStockSubscriptionCreateData
	isSet bool
}

func (v NullableInStockSubscriptionCreateData) Get() *InStockSubscriptionCreateData {
	return v.value
}

func (v *NullableInStockSubscriptionCreateData) Set(val *InStockSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreateData(val *InStockSubscriptionCreateData) *NullableInStockSubscriptionCreateData {
	return &NullableInStockSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}