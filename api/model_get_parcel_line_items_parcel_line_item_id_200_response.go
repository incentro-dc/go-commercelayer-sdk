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

// GETParcelLineItemsParcelLineItemId200Response struct for GETParcelLineItemsParcelLineItemId200Response
type GETParcelLineItemsParcelLineItemId200Response struct {
	Data *GETParcelLineItems200ResponseDataInner `json:"data,omitempty"`
}

// NewGETParcelLineItemsParcelLineItemId200Response instantiates a new GETParcelLineItemsParcelLineItemId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETParcelLineItemsParcelLineItemId200Response() *GETParcelLineItemsParcelLineItemId200Response {
	this := GETParcelLineItemsParcelLineItemId200Response{}
	return &this
}

// NewGETParcelLineItemsParcelLineItemId200ResponseWithDefaults instantiates a new GETParcelLineItemsParcelLineItemId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETParcelLineItemsParcelLineItemId200ResponseWithDefaults() *GETParcelLineItemsParcelLineItemId200Response {
	this := GETParcelLineItemsParcelLineItemId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETParcelLineItemsParcelLineItemId200Response) GetData() GETParcelLineItems200ResponseDataInner {
	if o == nil || o.Data == nil {
		var ret GETParcelLineItems200ResponseDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETParcelLineItemsParcelLineItemId200Response) GetDataOk() (*GETParcelLineItems200ResponseDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETParcelLineItemsParcelLineItemId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETParcelLineItems200ResponseDataInner and assigns it to the Data field.
func (o *GETParcelLineItemsParcelLineItemId200Response) SetData(v GETParcelLineItems200ResponseDataInner) {
	o.Data = &v
}

func (o GETParcelLineItemsParcelLineItemId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETParcelLineItemsParcelLineItemId200Response struct {
	value *GETParcelLineItemsParcelLineItemId200Response
	isSet bool
}

func (v NullableGETParcelLineItemsParcelLineItemId200Response) Get() *GETParcelLineItemsParcelLineItemId200Response {
	return v.value
}

func (v *NullableGETParcelLineItemsParcelLineItemId200Response) Set(val *GETParcelLineItemsParcelLineItemId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGETParcelLineItemsParcelLineItemId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGETParcelLineItemsParcelLineItemId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETParcelLineItemsParcelLineItemId200Response(val *GETParcelLineItemsParcelLineItemId200Response) *NullableGETParcelLineItemsParcelLineItemId200Response {
	return &NullableGETParcelLineItemsParcelLineItemId200Response{value: val, isSet: true}
}

func (v NullableGETParcelLineItemsParcelLineItemId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETParcelLineItemsParcelLineItemId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
