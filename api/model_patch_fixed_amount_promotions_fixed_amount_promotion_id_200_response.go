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

// checks if the PATCHFixedAmountPromotionsFixedAmountPromotionId200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHFixedAmountPromotionsFixedAmountPromotionId200Response{}

// PATCHFixedAmountPromotionsFixedAmountPromotionId200Response struct for PATCHFixedAmountPromotionsFixedAmountPromotionId200Response
type PATCHFixedAmountPromotionsFixedAmountPromotionId200Response struct {
	Data *PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData `json:"data,omitempty"`
}

// NewPATCHFixedAmountPromotionsFixedAmountPromotionId200Response instantiates a new PATCHFixedAmountPromotionsFixedAmountPromotionId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHFixedAmountPromotionsFixedAmountPromotionId200Response() *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response {
	this := PATCHFixedAmountPromotionsFixedAmountPromotionId200Response{}
	return &this
}

// NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseWithDefaults instantiates a new PATCHFixedAmountPromotionsFixedAmountPromotionId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseWithDefaults() *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response {
	this := PATCHFixedAmountPromotionsFixedAmountPromotionId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) GetData() PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) GetDataOk() (*PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData and assigns it to the Data field.
func (o *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) SetData(v PATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseData) {
	o.Data = &v
}

func (o PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response struct {
	value *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response
	isSet bool
}

func (v NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) Get() *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response {
	return v.value
}

func (v *NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) Set(val *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response(val *PATCHFixedAmountPromotionsFixedAmountPromotionId200Response) *NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response {
	return &NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response{value: val, isSet: true}
}

func (v NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHFixedAmountPromotionsFixedAmountPromotionId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
