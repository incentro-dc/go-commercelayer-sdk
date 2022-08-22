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

// PATCHGiftCardsGiftCardId200Response struct for PATCHGiftCardsGiftCardId200Response
type PATCHGiftCardsGiftCardId200Response struct {
	Data *PATCHGiftCardsGiftCardId200ResponseData `json:"data,omitempty"`
}

// NewPATCHGiftCardsGiftCardId200Response instantiates a new PATCHGiftCardsGiftCardId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHGiftCardsGiftCardId200Response() *PATCHGiftCardsGiftCardId200Response {
	this := PATCHGiftCardsGiftCardId200Response{}
	return &this
}

// NewPATCHGiftCardsGiftCardId200ResponseWithDefaults instantiates a new PATCHGiftCardsGiftCardId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHGiftCardsGiftCardId200ResponseWithDefaults() *PATCHGiftCardsGiftCardId200Response {
	this := PATCHGiftCardsGiftCardId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHGiftCardsGiftCardId200Response) GetData() PATCHGiftCardsGiftCardId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHGiftCardsGiftCardId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHGiftCardsGiftCardId200Response) GetDataOk() (*PATCHGiftCardsGiftCardId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHGiftCardsGiftCardId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHGiftCardsGiftCardId200ResponseData and assigns it to the Data field.
func (o *PATCHGiftCardsGiftCardId200Response) SetData(v PATCHGiftCardsGiftCardId200ResponseData) {
	o.Data = &v
}

func (o PATCHGiftCardsGiftCardId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHGiftCardsGiftCardId200Response struct {
	value *PATCHGiftCardsGiftCardId200Response
	isSet bool
}

func (v NullablePATCHGiftCardsGiftCardId200Response) Get() *PATCHGiftCardsGiftCardId200Response {
	return v.value
}

func (v *NullablePATCHGiftCardsGiftCardId200Response) Set(val *PATCHGiftCardsGiftCardId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHGiftCardsGiftCardId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHGiftCardsGiftCardId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHGiftCardsGiftCardId200Response(val *PATCHGiftCardsGiftCardId200Response) *NullablePATCHGiftCardsGiftCardId200Response {
	return &NullablePATCHGiftCardsGiftCardId200Response{value: val, isSet: true}
}

func (v NullablePATCHGiftCardsGiftCardId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHGiftCardsGiftCardId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}