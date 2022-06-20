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

// CustomerGroupUpdateData struct for CustomerGroupUpdateData
type CustomerGroupUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                            `json:"id"`
	Attributes    CustomerGroupUpdateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}            `json:"relationships,omitempty"`
}

// NewCustomerGroupUpdateData instantiates a new CustomerGroupUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupUpdateData(type_ string, id string, attributes CustomerGroupUpdateDataAttributes) *CustomerGroupUpdateData {
	this := CustomerGroupUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewCustomerGroupUpdateDataWithDefaults instantiates a new CustomerGroupUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupUpdateDataWithDefaults() *CustomerGroupUpdateData {
	this := CustomerGroupUpdateData{}
	var type_ string = "customer_groups"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CustomerGroupUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerGroupUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerGroupUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerGroupUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerGroupUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerGroupUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *CustomerGroupUpdateData) GetAttributes() CustomerGroupUpdateDataAttributes {
	if o == nil {
		var ret CustomerGroupUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *CustomerGroupUpdateData) GetAttributesOk() (*CustomerGroupUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *CustomerGroupUpdateData) SetAttributes(v CustomerGroupUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerGroupUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerGroupUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *CustomerGroupUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o CustomerGroupUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableCustomerGroupUpdateData struct {
	value *CustomerGroupUpdateData
	isSet bool
}

func (v NullableCustomerGroupUpdateData) Get() *CustomerGroupUpdateData {
	return v.value
}

func (v *NullableCustomerGroupUpdateData) Set(val *CustomerGroupUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupUpdateData(val *CustomerGroupUpdateData) *NullableCustomerGroupUpdateData {
	return &NullableCustomerGroupUpdateData{value: val, isSet: true}
}

func (v NullableCustomerGroupUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}