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

// checks if the OrderDataRelationshipsPaymentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDataRelationshipsPaymentSource{}

// OrderDataRelationshipsPaymentSource struct for OrderDataRelationshipsPaymentSource
type OrderDataRelationshipsPaymentSource struct {
	Data *POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData `json:"data,omitempty"`
}

// NewOrderDataRelationshipsPaymentSource instantiates a new OrderDataRelationshipsPaymentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDataRelationshipsPaymentSource() *OrderDataRelationshipsPaymentSource {
	this := OrderDataRelationshipsPaymentSource{}
	return &this
}

// NewOrderDataRelationshipsPaymentSourceWithDefaults instantiates a new OrderDataRelationshipsPaymentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDataRelationshipsPaymentSourceWithDefaults() *OrderDataRelationshipsPaymentSource {
	this := OrderDataRelationshipsPaymentSource{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *OrderDataRelationshipsPaymentSource) GetData() POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDataRelationshipsPaymentSource) GetDataOk() (*POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *OrderDataRelationshipsPaymentSource) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData and assigns it to the Data field.
func (o *OrderDataRelationshipsPaymentSource) SetData(v POSTCustomerPaymentSourcesRequestDataRelationshipsPaymentSourceData) {
	o.Data = &v
}

func (o OrderDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDataRelationshipsPaymentSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableOrderDataRelationshipsPaymentSource struct {
	value *OrderDataRelationshipsPaymentSource
	isSet bool
}

func (v NullableOrderDataRelationshipsPaymentSource) Get() *OrderDataRelationshipsPaymentSource {
	return v.value
}

func (v *NullableOrderDataRelationshipsPaymentSource) Set(val *OrderDataRelationshipsPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDataRelationshipsPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDataRelationshipsPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDataRelationshipsPaymentSource(val *OrderDataRelationshipsPaymentSource) *NullableOrderDataRelationshipsPaymentSource {
	return &NullableOrderDataRelationshipsPaymentSource{value: val, isSet: true}
}

func (v NullableOrderDataRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDataRelationshipsPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}