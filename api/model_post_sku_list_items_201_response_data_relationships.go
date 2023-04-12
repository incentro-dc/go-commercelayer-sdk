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

// checks if the POSTSkuListItems201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSkuListItems201ResponseDataRelationships{}

// POSTSkuListItems201ResponseDataRelationships struct for POSTSkuListItems201ResponseDataRelationships
type POSTSkuListItems201ResponseDataRelationships struct {
	SkuList *POSTBundles201ResponseDataRelationshipsSkuList          `json:"sku_list,omitempty"`
	Sku     *POSTInStockSubscriptions201ResponseDataRelationshipsSku `json:"sku,omitempty"`
}

// NewPOSTSkuListItems201ResponseDataRelationships instantiates a new POSTSkuListItems201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListItems201ResponseDataRelationships() *POSTSkuListItems201ResponseDataRelationships {
	this := POSTSkuListItems201ResponseDataRelationships{}
	return &this
}

// NewPOSTSkuListItems201ResponseDataRelationshipsWithDefaults instantiates a new POSTSkuListItems201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListItems201ResponseDataRelationshipsWithDefaults() *POSTSkuListItems201ResponseDataRelationships {
	this := POSTSkuListItems201ResponseDataRelationships{}
	return &this
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataRelationships) GetSkuList() POSTBundles201ResponseDataRelationshipsSkuList {
	if o == nil || IsNil(o.SkuList) {
		var ret POSTBundles201ResponseDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataRelationships) GetSkuListOk() (*POSTBundles201ResponseDataRelationshipsSkuList, bool) {
	if o == nil || IsNil(o.SkuList) {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataRelationships) HasSkuList() bool {
	if o != nil && !IsNil(o.SkuList) {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given POSTBundles201ResponseDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *POSTSkuListItems201ResponseDataRelationships) SetSkuList(v POSTBundles201ResponseDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *POSTSkuListItems201ResponseDataRelationships) GetSku() POSTInStockSubscriptions201ResponseDataRelationshipsSku {
	if o == nil || IsNil(o.Sku) {
		var ret POSTInStockSubscriptions201ResponseDataRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListItems201ResponseDataRelationships) GetSkuOk() (*POSTInStockSubscriptions201ResponseDataRelationshipsSku, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *POSTSkuListItems201ResponseDataRelationships) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given POSTInStockSubscriptions201ResponseDataRelationshipsSku and assigns it to the Sku field.
func (o *POSTSkuListItems201ResponseDataRelationships) SetSku(v POSTInStockSubscriptions201ResponseDataRelationshipsSku) {
	o.Sku = &v
}

func (o POSTSkuListItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSkuListItems201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkuList) {
		toSerialize["sku_list"] = o.SkuList
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullablePOSTSkuListItems201ResponseDataRelationships struct {
	value *POSTSkuListItems201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTSkuListItems201ResponseDataRelationships) Get() *POSTSkuListItems201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTSkuListItems201ResponseDataRelationships) Set(val *POSTSkuListItems201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListItems201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListItems201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListItems201ResponseDataRelationships(val *POSTSkuListItems201ResponseDataRelationships) *NullablePOSTSkuListItems201ResponseDataRelationships {
	return &NullablePOSTSkuListItems201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSkuListItems201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListItems201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
