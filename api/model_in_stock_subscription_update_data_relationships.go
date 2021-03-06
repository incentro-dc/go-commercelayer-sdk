/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// InStockSubscriptionUpdateDataRelationships struct for InStockSubscriptionUpdateDataRelationships
type InStockSubscriptionUpdateDataRelationships struct {
	Market   *AvalaraAccountDataRelationshipsMarkets   `json:"market,omitempty"`
	Customer *CouponRecipientDataRelationshipsCustomer `json:"customer,omitempty"`
	Sku      *BundleDataRelationshipsSkus              `json:"sku,omitempty"`
}

// NewInStockSubscriptionUpdateDataRelationships instantiates a new InStockSubscriptionUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionUpdateDataRelationships() *InStockSubscriptionUpdateDataRelationships {
	this := InStockSubscriptionUpdateDataRelationships{}
	return &this
}

// NewInStockSubscriptionUpdateDataRelationshipsWithDefaults instantiates a new InStockSubscriptionUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionUpdateDataRelationshipsWithDefaults() *InStockSubscriptionUpdateDataRelationships {
	this := InStockSubscriptionUpdateDataRelationships{}
	return &this
}

// GetMarket returns the Market field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets {
	if o == nil || o.Market == nil {
		var ret AvalaraAccountDataRelationshipsMarkets
		return ret
	}
	return *o.Market
}

// GetMarketOk returns a tuple with the Market field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool) {
	if o == nil || o.Market == nil {
		return nil, false
	}
	return o.Market, true
}

// HasMarket returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasMarket() bool {
	if o != nil && o.Market != nil {
		return true
	}

	return false
}

// SetMarket gets a reference to the given AvalaraAccountDataRelationshipsMarkets and assigns it to the Market field.
func (o *InStockSubscriptionUpdateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets) {
	o.Market = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer {
	if o == nil || o.Customer == nil {
		var ret CouponRecipientDataRelationshipsCustomer
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CouponRecipientDataRelationshipsCustomer and assigns it to the Customer field.
func (o *InStockSubscriptionUpdateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer) {
	o.Customer = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *InStockSubscriptionUpdateDataRelationships) GetSku() BundleDataRelationshipsSkus {
	if o == nil || o.Sku == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionUpdateDataRelationships) GetSkuOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *InStockSubscriptionUpdateDataRelationships) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Sku field.
func (o *InStockSubscriptionUpdateDataRelationships) SetSku(v BundleDataRelationshipsSkus) {
	o.Sku = &v
}

func (o InStockSubscriptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Market != nil {
		toSerialize["market"] = o.Market
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullableInStockSubscriptionUpdateDataRelationships struct {
	value *InStockSubscriptionUpdateDataRelationships
	isSet bool
}

func (v NullableInStockSubscriptionUpdateDataRelationships) Get() *InStockSubscriptionUpdateDataRelationships {
	return v.value
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) Set(val *InStockSubscriptionUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionUpdateDataRelationships(val *InStockSubscriptionUpdateDataRelationships) *NullableInStockSubscriptionUpdateDataRelationships {
	return &NullableInStockSubscriptionUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableInStockSubscriptionUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
