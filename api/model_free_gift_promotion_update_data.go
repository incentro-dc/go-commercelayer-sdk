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

// FreeGiftPromotionUpdateData struct for FreeGiftPromotionUpdateData
type FreeGiftPromotionUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                                      `json:"id"`
	Attributes    PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes         `json:"attributes"`
	Relationships *PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewFreeGiftPromotionUpdateData instantiates a new FreeGiftPromotionUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotionUpdateData(type_ string, id string, attributes PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) *FreeGiftPromotionUpdateData {
	this := FreeGiftPromotionUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewFreeGiftPromotionUpdateDataWithDefaults instantiates a new FreeGiftPromotionUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionUpdateDataWithDefaults() *FreeGiftPromotionUpdateData {
	this := FreeGiftPromotionUpdateData{}
	var type_ string = "free_gift_promotions"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *FreeGiftPromotionUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FreeGiftPromotionUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *FreeGiftPromotionUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FreeGiftPromotionUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *FreeGiftPromotionUpdateData) GetAttributes() PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionUpdateData) GetAttributesOk() (*PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *FreeGiftPromotionUpdateData) SetAttributes(v PATCHFreeGiftPromotionsFreeGiftPromotionId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *FreeGiftPromotionUpdateData) GetRelationships() PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionUpdateData) GetRelationshipsOk() (*PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *FreeGiftPromotionUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships and assigns it to the Relationships field.
func (o *FreeGiftPromotionUpdateData) SetRelationships(v PATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataRelationships) {
	o.Relationships = &v
}

func (o FreeGiftPromotionUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableFreeGiftPromotionUpdateData struct {
	value *FreeGiftPromotionUpdateData
	isSet bool
}

func (v NullableFreeGiftPromotionUpdateData) Get() *FreeGiftPromotionUpdateData {
	return v.value
}

func (v *NullableFreeGiftPromotionUpdateData) Set(val *FreeGiftPromotionUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotionUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotionUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotionUpdateData(val *FreeGiftPromotionUpdateData) *NullableFreeGiftPromotionUpdateData {
	return &NullableFreeGiftPromotionUpdateData{value: val, isSet: true}
}

func (v NullableFreeGiftPromotionUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotionUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
