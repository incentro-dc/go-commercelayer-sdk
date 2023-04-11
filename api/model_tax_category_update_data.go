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

// TaxCategoryUpdateData struct for TaxCategoryUpdateData
type TaxCategoryUpdateData struct {
	// The resource's type
	Type interface{} `json:"type"`
	// The resource's id
	Id            interface{}                                              `json:"id"`
	Attributes    PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes `json:"attributes"`
	Relationships *TaxCategoryUpdateDataRelationships                      `json:"relationships,omitempty"`
}

// NewTaxCategoryUpdateData instantiates a new TaxCategoryUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryUpdateData(type_ interface{}, id interface{}, attributes PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes) *TaxCategoryUpdateData {
	this := TaxCategoryUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewTaxCategoryUpdateDataWithDefaults instantiates a new TaxCategoryUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryUpdateDataWithDefaults() *TaxCategoryUpdateData {
	this := TaxCategoryUpdateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxCategoryUpdateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCategoryUpdateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaxCategoryUpdateData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *TaxCategoryUpdateData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaxCategoryUpdateData) GetIdOk() (*interface{}, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaxCategoryUpdateData) SetId(v interface{}) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *TaxCategoryUpdateData) GetAttributes() PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetAttributesOk() (*PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *TaxCategoryUpdateData) SetAttributes(v PATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TaxCategoryUpdateData) GetRelationships() TaxCategoryUpdateDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret TaxCategoryUpdateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdateData) GetRelationshipsOk() (*TaxCategoryUpdateDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TaxCategoryUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given TaxCategoryUpdateDataRelationships and assigns it to the Relationships field.
func (o *TaxCategoryUpdateData) SetRelationships(v TaxCategoryUpdateDataRelationships) {
	o.Relationships = &v
}

func (o TaxCategoryUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableTaxCategoryUpdateData struct {
	value *TaxCategoryUpdateData
	isSet bool
}

func (v NullableTaxCategoryUpdateData) Get() *TaxCategoryUpdateData {
	return v.value
}

func (v *NullableTaxCategoryUpdateData) Set(val *TaxCategoryUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryUpdateData(val *TaxCategoryUpdateData) *NullableTaxCategoryUpdateData {
	return &NullableTaxCategoryUpdateData{value: val, isSet: true}
}

func (v NullableTaxCategoryUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
