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

// PATCHStockTransfersStockTransferId200Response struct for PATCHStockTransfersStockTransferId200Response
type PATCHStockTransfersStockTransferId200Response struct {
	Data *PATCHStockTransfersStockTransferId200ResponseData `json:"data,omitempty"`
}

// NewPATCHStockTransfersStockTransferId200Response instantiates a new PATCHStockTransfersStockTransferId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHStockTransfersStockTransferId200Response() *PATCHStockTransfersStockTransferId200Response {
	this := PATCHStockTransfersStockTransferId200Response{}
	return &this
}

// NewPATCHStockTransfersStockTransferId200ResponseWithDefaults instantiates a new PATCHStockTransfersStockTransferId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHStockTransfersStockTransferId200ResponseWithDefaults() *PATCHStockTransfersStockTransferId200Response {
	this := PATCHStockTransfersStockTransferId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHStockTransfersStockTransferId200Response) GetData() PATCHStockTransfersStockTransferId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHStockTransfersStockTransferId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHStockTransfersStockTransferId200Response) GetDataOk() (*PATCHStockTransfersStockTransferId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHStockTransfersStockTransferId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHStockTransfersStockTransferId200ResponseData and assigns it to the Data field.
func (o *PATCHStockTransfersStockTransferId200Response) SetData(v PATCHStockTransfersStockTransferId200ResponseData) {
	o.Data = &v
}

func (o PATCHStockTransfersStockTransferId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHStockTransfersStockTransferId200Response struct {
	value *PATCHStockTransfersStockTransferId200Response
	isSet bool
}

func (v NullablePATCHStockTransfersStockTransferId200Response) Get() *PATCHStockTransfersStockTransferId200Response {
	return v.value
}

func (v *NullablePATCHStockTransfersStockTransferId200Response) Set(val *PATCHStockTransfersStockTransferId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHStockTransfersStockTransferId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHStockTransfersStockTransferId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHStockTransfersStockTransferId200Response(val *PATCHStockTransfersStockTransferId200Response) *NullablePATCHStockTransfersStockTransferId200Response {
	return &NullablePATCHStockTransfersStockTransferId200Response{value: val, isSet: true}
}

func (v NullablePATCHStockTransfersStockTransferId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHStockTransfersStockTransferId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
