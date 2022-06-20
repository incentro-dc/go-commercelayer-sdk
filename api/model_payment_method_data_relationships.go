/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaymentMethodDataRelationships struct for PaymentMethodDataRelationships
type PaymentMethodDataRelationships struct {
	Market         *AvalaraAccountDataRelationshipsMarkets      `json:"market,omitempty"`
	PaymentGateway *AdyenPaymentDataRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
	Attachments    *AvalaraAccountDataRelationshipsAttachments  `json:"attachments,omitempty"`
}

// NewPaymentMethodDataRelationships instantiates a new PaymentMethodDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodDataRelationships() *PaymentMethodDataRelationships {
	this := PaymentMethodDataRelationships{}
	return &this
}

// NewPaymentMethodDataRelationshipsWithDefaults instantiates a new PaymentMethodDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodDataRelationshipsWithDefaults() *PaymentMethodDataRelationships {
	this := PaymentMethodDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *PaymentMethodDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *PaymentMethodDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *PaymentMethodDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *PaymentMethodDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret AdyenPaymentDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *PaymentMethodDataRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given AdyenPaymentDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *PaymentMethodDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PaymentMethodDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PaymentMethodDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *PaymentMethodDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o PaymentMethodDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.PaymentGateway != nil {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodDataRelationships struct {
	value *PaymentMethodDataRelationships
	isSet bool
}

func (v NullablePaymentMethodDataRelationships) Get() *PaymentMethodDataRelationships {
	return v.value
}

func (v *NullablePaymentMethodDataRelationships) Set(val *PaymentMethodDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodDataRelationships(val *PaymentMethodDataRelationships) *NullablePaymentMethodDataRelationships {
	return &NullablePaymentMethodDataRelationships{value: val, isSet: true}
}

func (v NullablePaymentMethodDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
