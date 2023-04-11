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

// OrderSubscriptionData struct for OrderSubscriptionData
type OrderSubscriptionData struct {
	// The resource's type
	Type          interface{}                                         `json:"type"`
	Attributes    GETOrderSubscriptions200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *OrderSubscriptionDataRelationships                 `json:"relationships,omitempty"`
}

// NewOrderSubscriptionData instantiates a new OrderSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionData(type_ interface{}, attributes GETOrderSubscriptions200ResponseDataInnerAttributes) *OrderSubscriptionData {
	this := OrderSubscriptionData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderSubscriptionDataWithDefaults instantiates a new OrderSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionDataWithDefaults() *OrderSubscriptionData {
	this := OrderSubscriptionData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderSubscriptionData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderSubscriptionData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderSubscriptionData) GetAttributes() GETOrderSubscriptions200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETOrderSubscriptions200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionData) GetAttributesOk() (*GETOrderSubscriptions200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderSubscriptionData) SetAttributes(v GETOrderSubscriptions200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderSubscriptionData) GetRelationships() OrderSubscriptionDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderSubscriptionDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionData) GetRelationshipsOk() (*OrderSubscriptionDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderSubscriptionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderSubscriptionDataRelationships and assigns it to the Relationships field.
func (o *OrderSubscriptionData) SetRelationships(v OrderSubscriptionDataRelationships) {
	o.Relationships = &v
}

func (o OrderSubscriptionData) MarshalJSON() ([]byte, error) {
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

type NullableOrderSubscriptionData struct {
	value *OrderSubscriptionData
	isSet bool
}

func (v NullableOrderSubscriptionData) Get() *OrderSubscriptionData {
	return v.value
}

func (v *NullableOrderSubscriptionData) Set(val *OrderSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionData(val *OrderSubscriptionData) *NullableOrderSubscriptionData {
	return &NullableOrderSubscriptionData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
