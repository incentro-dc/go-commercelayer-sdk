/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETStockItems200ResponseDataInnerRelationships struct for GETStockItems200ResponseDataInnerRelationships
type GETStockItems200ResponseDataInnerRelationships struct {
	StockLocation *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation `json:"stock_location,omitempty"`
	Sku           *GETInStockSubscriptions200ResponseDataInnerRelationshipsSku        `json:"sku,omitempty"`
	Attachments   *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments     `json:"attachments,omitempty"`
}

// NewGETStockItems200ResponseDataInnerRelationships instantiates a new GETStockItems200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockItems200ResponseDataInnerRelationships() *GETStockItems200ResponseDataInnerRelationships {
	this := GETStockItems200ResponseDataInnerRelationships{}
	return &this
}

// NewGETStockItems200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETStockItems200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockItems200ResponseDataInnerRelationshipsWithDefaults() *GETStockItems200ResponseDataInnerRelationships {
	this := GETStockItems200ResponseDataInnerRelationships{}
	return &this
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *GETStockItems200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *GETStockItems200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GETStockItems200ResponseDataInnerRelationships) GetSku() GETInStockSubscriptions200ResponseDataInnerRelationshipsSku {
	if o == nil || o.Sku == nil {
		var ret GETInStockSubscriptions200ResponseDataInnerRelationshipsSku
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) GetSkuOk() (*GETInStockSubscriptions200ResponseDataInnerRelationshipsSku, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given GETInStockSubscriptions200ResponseDataInnerRelationshipsSku and assigns it to the Sku field.
func (o *GETStockItems200ResponseDataInnerRelationships) SetSku(v GETInStockSubscriptions200ResponseDataInnerRelationshipsSku) {
	o.Sku = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETStockItems200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETStockItems200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETStockItems200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETStockItems200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETStockItems200ResponseDataInnerRelationships struct {
	value *GETStockItems200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETStockItems200ResponseDataInnerRelationships) Get() *GETStockItems200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETStockItems200ResponseDataInnerRelationships) Set(val *GETStockItems200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockItems200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockItems200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockItems200ResponseDataInnerRelationships(val *GETStockItems200ResponseDataInnerRelationships) *NullableGETStockItems200ResponseDataInnerRelationships {
	return &NullableGETStockItems200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETStockItems200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockItems200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
