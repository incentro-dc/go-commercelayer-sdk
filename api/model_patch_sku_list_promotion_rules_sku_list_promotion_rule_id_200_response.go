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

// PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response struct for PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response
type PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response struct {
	Data *PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData `json:"data,omitempty"`
}

// NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200Response instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200Response() *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response {
	this := PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response{}
	return &this
}

// NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseWithDefaults instantiates a new PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseWithDefaults() *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response {
	this := PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) GetData() PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData {
	if o == nil || o.Data == nil {
		var ret PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) GetDataOk() (*PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData and assigns it to the Data field.
func (o *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) SetData(v PATCHSkuListPromotionRulesSkuListPromotionRuleId200ResponseData) {
	o.Data = &v
}

func (o PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response struct {
	value *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response
	isSet bool
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) Get() *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response {
	return v.value
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) Set(val *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response(val *PATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response {
	return &NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response{value: val, isSet: true}
}

func (v NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuListPromotionRulesSkuListPromotionRuleId200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
