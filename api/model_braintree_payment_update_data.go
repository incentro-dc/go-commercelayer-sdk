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

// BraintreePaymentUpdateData struct for BraintreePaymentUpdateData
type BraintreePaymentUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                               `json:"id"`
	Attributes    BraintreePaymentUpdateDataAttributes `json:"attributes"`
	Relationships *AdyenPaymentUpdateDataRelationships `json:"relationships,omitempty"`
}

// NewBraintreePaymentUpdateData instantiates a new BraintreePaymentUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBraintreePaymentUpdateData(type_ string, id string, attributes BraintreePaymentUpdateDataAttributes) *BraintreePaymentUpdateData {
	this := BraintreePaymentUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewBraintreePaymentUpdateDataWithDefaults instantiates a new BraintreePaymentUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBraintreePaymentUpdateDataWithDefaults() *BraintreePaymentUpdateData {
	this := BraintreePaymentUpdateData{}
	var type_ string = "braintree_payments"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *BraintreePaymentUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BraintreePaymentUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BraintreePaymentUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BraintreePaymentUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *BraintreePaymentUpdateData) GetAttributes() BraintreePaymentUpdateDataAttributes {
	if o == nil {
		var ret BraintreePaymentUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BraintreePaymentUpdateData) GetAttributesOk() (*BraintreePaymentUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BraintreePaymentUpdateData) SetAttributes(v BraintreePaymentUpdateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BraintreePaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret AdyenPaymentUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BraintreePaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BraintreePaymentUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenPaymentUpdateDataRelationships and assigns it to the Relationships field.
func (o *BraintreePaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships) {
	o.Relationships = &v
}

func (o BraintreePaymentUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableBraintreePaymentUpdateData struct {
	value *BraintreePaymentUpdateData
	isSet bool
}

func (v NullableBraintreePaymentUpdateData) Get() *BraintreePaymentUpdateData {
	return v.value
}

func (v *NullableBraintreePaymentUpdateData) Set(val *BraintreePaymentUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableBraintreePaymentUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableBraintreePaymentUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBraintreePaymentUpdateData(val *BraintreePaymentUpdateData) *NullableBraintreePaymentUpdateData {
	return &NullableBraintreePaymentUpdateData{value: val, isSet: true}
}

func (v NullableBraintreePaymentUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBraintreePaymentUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
