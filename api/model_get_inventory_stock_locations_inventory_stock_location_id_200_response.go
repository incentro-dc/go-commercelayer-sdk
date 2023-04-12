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

// checks if the GETInventoryStockLocationsInventoryStockLocationId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETInventoryStockLocationsInventoryStockLocationId200Response{}

// GETInventoryStockLocationsInventoryStockLocationId200Response struct for GETInventoryStockLocationsInventoryStockLocationId200Response
type GETInventoryStockLocationsInventoryStockLocationId200Response struct {
	Data *GETInventoryStockLocationsInventoryStockLocationId200ResponseData `json:"data,omitempty"`
}

// NewGETInventoryStockLocationsInventoryStockLocationId200Response instantiates a new GETInventoryStockLocationsInventoryStockLocationId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryStockLocationsInventoryStockLocationId200Response() *GETInventoryStockLocationsInventoryStockLocationId200Response {
	this := GETInventoryStockLocationsInventoryStockLocationId200Response{}
	return &this
}

// NewGETInventoryStockLocationsInventoryStockLocationId200ResponseWithDefaults instantiates a new GETInventoryStockLocationsInventoryStockLocationId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryStockLocationsInventoryStockLocationId200ResponseWithDefaults() *GETInventoryStockLocationsInventoryStockLocationId200Response {
	this := GETInventoryStockLocationsInventoryStockLocationId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETInventoryStockLocationsInventoryStockLocationId200Response) GetData() GETInventoryStockLocationsInventoryStockLocationId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret GETInventoryStockLocationsInventoryStockLocationId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200Response) GetDataOk() (*GETInventoryStockLocationsInventoryStockLocationId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETInventoryStockLocationsInventoryStockLocationId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GETInventoryStockLocationsInventoryStockLocationId200ResponseData and assigns it to the Data field.
func (o *GETInventoryStockLocationsInventoryStockLocationId200Response) SetData(v GETInventoryStockLocationsInventoryStockLocationId200ResponseData) {
	o.Data = &v
}

func (o GETInventoryStockLocationsInventoryStockLocationId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETInventoryStockLocationsInventoryStockLocationId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGETInventoryStockLocationsInventoryStockLocationId200Response struct {
	value *GETInventoryStockLocationsInventoryStockLocationId200Response
	isSet bool
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200Response) Get() *GETInventoryStockLocationsInventoryStockLocationId200Response {
	return v.value
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200Response) Set(val *GETInventoryStockLocationsInventoryStockLocationId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryStockLocationsInventoryStockLocationId200Response(val *GETInventoryStockLocationsInventoryStockLocationId200Response) *NullableGETInventoryStockLocationsInventoryStockLocationId200Response {
	return &NullableGETInventoryStockLocationsInventoryStockLocationId200Response{value: val, isSet: true}
}

func (v NullableGETInventoryStockLocationsInventoryStockLocationId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryStockLocationsInventoryStockLocationId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
