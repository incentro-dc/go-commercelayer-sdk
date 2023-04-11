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

// checks if the OrderSubscriptionItemCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSubscriptionItemCreateData{}

// OrderSubscriptionItemCreateData struct for OrderSubscriptionItemCreateData
type OrderSubscriptionItemCreateData struct {
	// The resource's type
	Type          interface{}                                     `json:"type"`
	Attributes    POSTOrderSubscriptionItemsRequestDataAttributes `json:"attributes"`
	Relationships *OrderSubscriptionItemCreateDataRelationships   `json:"relationships,omitempty"`
}

// NewOrderSubscriptionItemCreateData instantiates a new OrderSubscriptionItemCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSubscriptionItemCreateData(type_ interface{}, attributes POSTOrderSubscriptionItemsRequestDataAttributes) *OrderSubscriptionItemCreateData {
	this := OrderSubscriptionItemCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewOrderSubscriptionItemCreateDataWithDefaults instantiates a new OrderSubscriptionItemCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSubscriptionItemCreateDataWithDefaults() *OrderSubscriptionItemCreateData {
	this := OrderSubscriptionItemCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *OrderSubscriptionItemCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderSubscriptionItemCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderSubscriptionItemCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *OrderSubscriptionItemCreateData) GetAttributes() POSTOrderSubscriptionItemsRequestDataAttributes {
	if o == nil {
		var ret POSTOrderSubscriptionItemsRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemCreateData) GetAttributesOk() (*POSTOrderSubscriptionItemsRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderSubscriptionItemCreateData) SetAttributes(v POSTOrderSubscriptionItemsRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderSubscriptionItemCreateData) GetRelationships() OrderSubscriptionItemCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret OrderSubscriptionItemCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSubscriptionItemCreateData) GetRelationshipsOk() (*OrderSubscriptionItemCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderSubscriptionItemCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given OrderSubscriptionItemCreateDataRelationships and assigns it to the Relationships field.
func (o *OrderSubscriptionItemCreateData) SetRelationships(v OrderSubscriptionItemCreateDataRelationships) {
	o.Relationships = &v
}

func (o OrderSubscriptionItemCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSubscriptionItemCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableOrderSubscriptionItemCreateData struct {
	value *OrderSubscriptionItemCreateData
	isSet bool
}

func (v NullableOrderSubscriptionItemCreateData) Get() *OrderSubscriptionItemCreateData {
	return v.value
}

func (v *NullableOrderSubscriptionItemCreateData) Set(val *OrderSubscriptionItemCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSubscriptionItemCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSubscriptionItemCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSubscriptionItemCreateData(val *OrderSubscriptionItemCreateData) *NullableOrderSubscriptionItemCreateData {
	return &NullableOrderSubscriptionItemCreateData{value: val, isSet: true}
}

func (v NullableOrderSubscriptionItemCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSubscriptionItemCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}