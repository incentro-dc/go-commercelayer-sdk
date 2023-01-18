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

// OrderData struct for OrderData
type OrderData struct {
	// The resource's type
	Type          string                                  `json:"type"`
	Attributes    GETOrders200ResponseDataInnerAttributes `json:"attributes"`
	Relationships *OrderDataRelationships                 `json:"relationships,omitempty"`
}

// NewOrderData instantiates a new OrderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderData(type_ string, attributes GETOrders200ResponseDataInnerAttributes) *OrderData {
	this := OrderData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderDataWithDefaults instantiates a new OrderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataWithDefaults() *OrderData {
	this := OrderData{}
	return &this
}

// GetType returns the Type field value
func (o *OrderData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderData) GetAttributes() GETOrders200ResponseDataInnerAttributes {
	if o == nil {
		var ret GETOrders200ResponseDataInnerAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderData) GetAttributesOk() (*GETOrders200ResponseDataInnerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderData) SetAttributes(v GETOrders200ResponseDataInnerAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderData) GetRelationships() OrderDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret OrderDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderData) GetRelationshipsOk() (*OrderDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderDataRelationships and assigns it to the Relationships field.
func (o *OrderData) SetRelationships(v OrderDataRelationships) {
	o.Relationships = &v
}

func (o OrderData) MarshalJSON() ([]byte, error) {
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

type NullableOrderData struct {
	value *OrderData
	isSet bool
}

func (v NullableOrderData) Get() *OrderData {
	return v.value
}

func (v *NullableOrderData) Set(val *OrderData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderData(val *OrderData) *NullableOrderData {
	return &NullableOrderData{value: val, isSet: true}
}

func (v NullableOrderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
