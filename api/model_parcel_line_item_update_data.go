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

// ParcelLineItemUpdateData struct for ParcelLineItemUpdateData
type ParcelLineItemUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                           `json:"id"`
	Attributes    AdyenPaymentCreateDataAttributes `json:"attributes"`
	Relationships map[string]interface{}           `json:"relationships,omitempty"`
}

// NewParcelLineItemUpdateData instantiates a new ParcelLineItemUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes) *ParcelLineItemUpdateData {
	this := ParcelLineItemUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewParcelLineItemUpdateDataWithDefaults instantiates a new ParcelLineItemUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelLineItemUpdateDataWithDefaults() *ParcelLineItemUpdateData {
	this := ParcelLineItemUpdateData{}
	var type_ string = "parcel_line_items"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ParcelLineItemUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelLineItemUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ParcelLineItemUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParcelLineItemUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelLineItemUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes {
	if o == nil {
		var ret AdyenPaymentCreateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelLineItemUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ParcelLineItemUpdateData) GetRelationships() map[string]interface{} {
	if o == nil || o.Relationships == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelLineItemUpdateData) GetRelationshipsOk() (map[string]interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelLineItemUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]interface{} and assigns it to the Relationships field.
func (o *ParcelLineItemUpdateData) SetRelationships(v map[string]interface{}) {
	o.Relationships = v
}

func (o ParcelLineItemUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableParcelLineItemUpdateData struct {
	value *ParcelLineItemUpdateData
	isSet bool
}

func (v NullableParcelLineItemUpdateData) Get() *ParcelLineItemUpdateData {
	return v.value
}

func (v *NullableParcelLineItemUpdateData) Set(val *ParcelLineItemUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelLineItemUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelLineItemUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelLineItemUpdateData(val *ParcelLineItemUpdateData) *NullableParcelLineItemUpdateData {
	return &NullableParcelLineItemUpdateData{value: val, isSet: true}
}

func (v NullableParcelLineItemUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelLineItemUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
