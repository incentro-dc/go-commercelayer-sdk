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

// StockTransferUpdateData struct for StockTransferUpdateData
type StockTransferUpdateData struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id            string                                                          `json:"id"`
	Attributes    PATCHStockTransfersStockTransferId200ResponseDataAttributes     `json:"attributes"`
	Relationships *PATCHStockTransfersStockTransferId200ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewStockTransferUpdateData instantiates a new StockTransferUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockTransferUpdateData(type_ string, id string, attributes PATCHStockTransfersStockTransferId200ResponseDataAttributes) *StockTransferUpdateData {
	this := StockTransferUpdateData{}
	this.Type = type_
	this.Id = id
	this.Attributes = attributes
	return &this
}

// NewStockTransferUpdateDataWithDefaults instantiates a new StockTransferUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockTransferUpdateDataWithDefaults() *StockTransferUpdateData {
	this := StockTransferUpdateData{}
	var type_ string = "stock_transfers"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *StockTransferUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StockTransferUpdateData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *StockTransferUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StockTransferUpdateData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value
func (o *StockTransferUpdateData) GetAttributes() PATCHStockTransfersStockTransferId200ResponseDataAttributes {
	if o == nil {
		var ret PATCHStockTransfersStockTransferId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateData) GetAttributesOk() (*PATCHStockTransfersStockTransferId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *StockTransferUpdateData) SetAttributes(v PATCHStockTransfersStockTransferId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *StockTransferUpdateData) GetRelationships() PATCHStockTransfersStockTransferId200ResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret PATCHStockTransfersStockTransferId200ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockTransferUpdateData) GetRelationshipsOk() (*PATCHStockTransfersStockTransferId200ResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *StockTransferUpdateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PATCHStockTransfersStockTransferId200ResponseDataRelationships and assigns it to the Relationships field.
func (o *StockTransferUpdateData) SetRelationships(v PATCHStockTransfersStockTransferId200ResponseDataRelationships) {
	o.Relationships = &v
}

func (o StockTransferUpdateData) MarshalJSON() ([]byte, error) {
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

type NullableStockTransferUpdateData struct {
	value *StockTransferUpdateData
	isSet bool
}

func (v NullableStockTransferUpdateData) Get() *StockTransferUpdateData {
	return v.value
}

func (v *NullableStockTransferUpdateData) Set(val *StockTransferUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableStockTransferUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableStockTransferUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockTransferUpdateData(val *StockTransferUpdateData) *NullableStockTransferUpdateData {
	return &NullableStockTransferUpdateData{value: val, isSet: true}
}

func (v NullableStockTransferUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockTransferUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
