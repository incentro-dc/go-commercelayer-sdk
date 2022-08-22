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

// OrderUpdateData struct for OrderUpdateData
type OrderUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                      `json:"id"`
	Attributes    PATCHOrdersOrderId200ResponseDataAttributes `json:"attributes"`
	Relationships *POSTOrders201ResponseDataRelationships     `json:"relationships,omitempty"`
}

// NewOrderUpdateData instantiates a new OrderUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderUpdateData(type_ string, id string, attributes PATCHOrdersOrderId200ResponseDataAttributes) *OrderUpdateData {
	this := OrderUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewOrderUpdateDataWithDefaults instantiates a new OrderUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderUpdateDataWithDefaults() *OrderUpdateData {
	this := OrderUpdateData{}
	var type_ string = "orders"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *OrderUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrderUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OrderUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *OrderUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *OrderUpdateData) GetAttributes() PATCHOrdersOrderId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHOrdersOrderId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrderUpdateData) GetAttributesOk() (*PATCHOrdersOrderId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *OrderUpdateData) SetAttributes(v PATCHOrdersOrderId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *OrderUpdateData) GetRelationships() POSTOrders201ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret POSTOrders201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderUpdateData) GetRelationshipsOk() (*POSTOrders201ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *OrderUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTOrders201ResponseDataRelationships and assigns it to the Relationships field.
func (o *OrderUpdateData) SetRelationships(v POSTOrders201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o OrderUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
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

type NullableOrderUpdateData struct {
	value *OrderUpdateData
	isSet bool
}

func (v NullableOrderUpdateData) Get() *OrderUpdateData {
	return v.value
}

func (v *NullableOrderUpdateData) Set(val *OrderUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderUpdateData(val *OrderUpdateData) *NullableOrderUpdateData {
	return &NullableOrderUpdateData{value: val, isSet: true}
}

func (v NullableOrderUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
