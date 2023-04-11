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

// ParcelLineItemUpdateData struct for ParcelLineItemUpdateData
type ParcelLineItemUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                                                         `json:"id"`
	Attributes    PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes `json:"attributes"`
	Relationships interface{}                                                                         `json:"relationships,omitempty"`
}

// NewParcelLineItemUpdateData instantiates a new ParcelLineItemUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelLineItemUpdateData(type_ interface{}, id interface{}, attributes PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes) *ParcelLineItemUpdateData {
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
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ParcelLineItemUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParcelLineItemUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ParcelLineItemUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ParcelLineItemUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParcelLineItemUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParcelLineItemUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *ParcelLineItemUpdateData) GetAttributes() PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ParcelLineItemUpdateData) GetAttributesOk() (*PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *ParcelLineItemUpdateData) SetAttributes(v PATCHBillingInfoValidationRulesBillingInfoValidationRuleId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParcelLineItemUpdateData) GetRelationships() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParcelLineItemUpdateData) GetRelationshipsOk() (*interface{}, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ParcelLineItemUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given interface{} and assigns it to the Relationships field.
func (o *ParcelLineItemUpdateData) SetRelationships(v interface{}) {
	o.Relationships = v
}

func (o ParcelLineItemUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
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
