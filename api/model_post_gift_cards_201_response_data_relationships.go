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

// POSTGiftCards201ResponseDataRelationships struct for POSTGiftCards201ResponseDataRelationships
type POSTGiftCards201ResponseDataRelationships struct {
	Market            *GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets     `json:"market,omitempty"`
	GiftCardRecipient *GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient `json:"gift_card_recipient,omitempty"`
}

// NewPOSTGiftCards201ResponseDataRelationships instantiates a new POSTGiftCards201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTGiftCards201ResponseDataRelationships() *POSTGiftCards201ResponseDataRelationships {
	this := POSTGiftCards201ResponseDataRelationships{}
	return &this
}

// NewPOSTGiftCards201ResponseDataRelationshipsWithDefaults instantiates a new POSTGiftCards201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTGiftCards201ResponseDataRelationshipsWithDefaults() *POSTGiftCards201ResponseDataRelationships {
	this := POSTGiftCards201ResponseDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetMarket() GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetMarketOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets and assigns it to the Market field.
func (o *POSTGiftCards201ResponseDataRelationships) SetMarket(v GETAvalaraAccounts200ResponseDataInnerRelationshipsMarkets) {
	o.Market = &v
}

// GetGiftCardRecipient returns the GiftCardRecipient field value if set, zero value otherwise.
func (o *POSTGiftCards201ResponseDataRelationships) GetGiftCardRecipient() GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient {
	if o == nil || o.GiftCardRecipient == nil {
		var ret GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient
		return ret
	}
	return *o.GiftCardRecipient
}

// GetGiftCardRecipientOk returns a tuple with the GiftCardRecipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTGiftCards201ResponseDataRelationships) GetGiftCardRecipientOk() (*GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient, bool) {
	if o == nil || o.GiftCardRecipient == nil {
		return nil, false
	}
	return o.GiftCardRecipient, true
}

// HasGiftCardRecipient returns a boolean if a field has been set.
func (o *POSTGiftCards201ResponseDataRelationships) HasGiftCardRecipient() bool {
	if o != nil && o.GiftCardRecipient != nil {
		return true
	}

	return false
}

// SetGiftCardRecipient gets a reference to the given GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient and assigns it to the GiftCardRecipient field.
func (o *POSTGiftCards201ResponseDataRelationships) SetGiftCardRecipient(v GETGiftCards200ResponseDataInnerRelationshipsGiftCardRecipient) {
	o.GiftCardRecipient = &v
}

func (o POSTGiftCards201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.GiftCardRecipient != nil {
		toSerialize["gift_card_recipient"] = o.GiftCardRecipient
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTGiftCards201ResponseDataRelationships struct {
	value *POSTGiftCards201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) Get() *POSTGiftCards201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) Set(val *POSTGiftCards201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTGiftCards201ResponseDataRelationships(val *POSTGiftCards201ResponseDataRelationships) *NullablePOSTGiftCards201ResponseDataRelationships {
	return &NullablePOSTGiftCards201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTGiftCards201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTGiftCards201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}