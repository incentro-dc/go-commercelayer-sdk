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

// checks if the CustomerAddressCreateDataRelationshipsAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddressCreateDataRelationshipsAddress{}

// CustomerAddressCreateDataRelationshipsAddress struct for CustomerAddressCreateDataRelationshipsAddress
type CustomerAddressCreateDataRelationshipsAddress struct {
	Data BingGeocoderDataRelationshipsAddressesData `json:"data"`
}

// NewCustomerAddressCreateDataRelationshipsAddress instantiates a new CustomerAddressCreateDataRelationshipsAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddressCreateDataRelationshipsAddress(data BingGeocoderDataRelationshipsAddressesData) *CustomerAddressCreateDataRelationshipsAddress {
	this := CustomerAddressCreateDataRelationshipsAddress{}
	this.Data = data
	return &this
}

// NewCustomerAddressCreateDataRelationshipsAddressWithDefaults instantiates a new CustomerAddressCreateDataRelationshipsAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressCreateDataRelationshipsAddressWithDefaults() *CustomerAddressCreateDataRelationshipsAddress {
	this := CustomerAddressCreateDataRelationshipsAddress{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerAddressCreateDataRelationshipsAddress) GetData() BingGeocoderDataRelationshipsAddressesData {
	if o == nil {
		var ret BingGeocoderDataRelationshipsAddressesData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerAddressCreateDataRelationshipsAddress) GetDataOk() (*BingGeocoderDataRelationshipsAddressesData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerAddressCreateDataRelationshipsAddress) SetData(v BingGeocoderDataRelationshipsAddressesData) {
	o.Data = v
}

func (o CustomerAddressCreateDataRelationshipsAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddressCreateDataRelationshipsAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCustomerAddressCreateDataRelationshipsAddress struct {
	value *CustomerAddressCreateDataRelationshipsAddress
	isSet bool
}

func (v NullableCustomerAddressCreateDataRelationshipsAddress) Get() *CustomerAddressCreateDataRelationshipsAddress {
	return v.value
}

func (v *NullableCustomerAddressCreateDataRelationshipsAddress) Set(val *CustomerAddressCreateDataRelationshipsAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddressCreateDataRelationshipsAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddressCreateDataRelationshipsAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddressCreateDataRelationshipsAddress(val *CustomerAddressCreateDataRelationshipsAddress) *NullableCustomerAddressCreateDataRelationshipsAddress {
	return &NullableCustomerAddressCreateDataRelationshipsAddress{value: val, isSet: true}
}

func (v NullableCustomerAddressCreateDataRelationshipsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddressCreateDataRelationshipsAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
