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

// checks if the PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships{}

// PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships struct for PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships
type PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships struct {
	Market         *POSTBillingInfoValidationRulesRequestDataRelationshipsMarket `json:"market,omitempty"`
	PaymentGateway *POSTPaymentMethodsRequestDataRelationshipsPaymentGateway     `json:"payment_gateway,omitempty"`
}

// NewPATCHPaymentMethodsPaymentMethodIdRequestDataRelationships instantiates a new PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHPaymentMethodsPaymentMethodIdRequestDataRelationships() *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships {
	this := PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships{}
	return &this
}

// NewPATCHPaymentMethodsPaymentMethodIdRequestDataRelationshipsWithDefaults instantiates a new PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHPaymentMethodsPaymentMethodIdRequestDataRelationshipsWithDefaults() *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships {
	this := PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) GetMarket() POSTBillingInfoValidationRulesRequestDataRelationshipsMarket {
	if o == nil || IsNil(o.Market) {
		var ret POSTBillingInfoValidationRulesRequestDataRelationshipsMarket
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) GetMarketOk() (*POSTBillingInfoValidationRulesRequestDataRelationshipsMarket, bool) {
	if o == nil || IsNil(o.Market) {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) HasMarket() bool {
	if o != nil && !IsNil(o.Market) {
		return true
	}

	return false
}

// SetMarket gets a reference to the given POSTBillingInfoValidationRulesRequestDataRelationshipsMarket and assigns it to the Market field.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) SetMarket(v POSTBillingInfoValidationRulesRequestDataRelationshipsMarket) {
	o.Market = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) GetPaymentGateway() POSTPaymentMethodsRequestDataRelationshipsPaymentGateway {
	if o == nil || IsNil(o.PaymentGateway) {
		var ret POSTPaymentMethodsRequestDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) GetPaymentGatewayOk() (*POSTPaymentMethodsRequestDataRelationshipsPaymentGateway, bool) {
	if o == nil || IsNil(o.PaymentGateway) {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) HasPaymentGateway() bool {
	if o != nil && !IsNil(o.PaymentGateway) {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given POSTPaymentMethodsRequestDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) SetPaymentGateway(v POSTPaymentMethodsRequestDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

func (o PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Market) {
		toSerialize["market"] = o.Market
	}
	if !IsNil(o.PaymentGateway) {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	return toSerialize, nil
}

type NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships struct {
	value *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships
	isSet bool
}

func (v NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) Get() *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships {
	return v.value
}

func (v *NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) Set(val *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships(val *PATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) *NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships {
	return &NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHPaymentMethodsPaymentMethodIdRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}