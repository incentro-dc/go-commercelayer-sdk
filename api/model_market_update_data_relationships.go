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

// MarketUpdateDataRelationships struct for MarketUpdateDataRelationships
type MarketUpdateDataRelationships struct {
	Merchant          *MarketCreateDataRelationshipsMerchant                        `json:"merchant,omitempty"`
	PriceList         *MarketCreateDataRelationshipsPriceList                       `json:"price_list,omitempty"`
	InventoryModel    *InventoryReturnLocationCreateDataRelationshipsInventoryModel `json:"inventory_model,omitempty"`
	SubscriptionModel *MarketCreateDataRelationshipsSubscriptionModel               `json:"subscription_model,omitempty"`
	TaxCalculator     *MarketCreateDataRelationshipsTaxCalculator                   `json:"tax_calculator,omitempty"`
	CustomerGroup     *CustomerCreateDataRelationshipsCustomerGroup                 `json:"customer_group,omitempty"`
}

// NewMarketUpdateDataRelationships instantiates a new MarketUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketUpdateDataRelationships() *MarketUpdateDataRelationships {
	this := MarketUpdateDataRelationships{}
	return &this
}

// NewMarketUpdateDataRelationshipsWithDefaults instantiates a new MarketUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketUpdateDataRelationshipsWithDefaults() *MarketUpdateDataRelationships {
	this := MarketUpdateDataRelationships{}
	return &this
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetMerchant() MarketCreateDataRelationshipsMerchant {
	if o == nil || o.Merchant == nil {
		var ret MarketCreateDataRelationshipsMerchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetMerchantOk() (*MarketCreateDataRelationshipsMerchant, bool) {
	if o == nil || o.Merchant == nil {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasMerchant() bool {
	if o != nil && o.Merchant != nil {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given MarketCreateDataRelationshipsMerchant and assigns it to the Merchant field.
func (o *MarketUpdateDataRelationships) SetMerchant(v MarketCreateDataRelationshipsMerchant) {
	o.Merchant = &v
}

// GetPriceList returns the PriceList field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetPriceList() MarketCreateDataRelationshipsPriceList {
	if o == nil || o.PriceList == nil {
		var ret MarketCreateDataRelationshipsPriceList
		return ret
	}
	return *o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetPriceListOk() (*MarketCreateDataRelationshipsPriceList, bool) {
	if o == nil || o.PriceList == nil {
		return nil, false
	}
	return o.PriceList, true
}

// HasPriceList returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasPriceList() bool {
	if o != nil && o.PriceList != nil {
		return true
	}

	return false
}

// SetPriceList gets a reference to the given MarketCreateDataRelationshipsPriceList and assigns it to the PriceList field.
func (o *MarketUpdateDataRelationships) SetPriceList(v MarketCreateDataRelationshipsPriceList) {
	o.PriceList = &v
}

// GetInventoryModel returns the InventoryModel field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetInventoryModel() InventoryReturnLocationCreateDataRelationshipsInventoryModel {
	if o == nil || o.InventoryModel == nil {
		var ret InventoryReturnLocationCreateDataRelationshipsInventoryModel
		return ret
	}
	return *o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetInventoryModelOk() (*InventoryReturnLocationCreateDataRelationshipsInventoryModel, bool) {
	if o == nil || o.InventoryModel == nil {
		return nil, false
	}
	return o.InventoryModel, true
}

// HasInventoryModel returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasInventoryModel() bool {
	if o != nil && o.InventoryModel != nil {
		return true
	}

	return false
}

// SetInventoryModel gets a reference to the given InventoryReturnLocationCreateDataRelationshipsInventoryModel and assigns it to the InventoryModel field.
func (o *MarketUpdateDataRelationships) SetInventoryModel(v InventoryReturnLocationCreateDataRelationshipsInventoryModel) {
	o.InventoryModel = &v
}

// GetSubscriptionModel returns the SubscriptionModel field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetSubscriptionModel() MarketCreateDataRelationshipsSubscriptionModel {
	if o == nil || o.SubscriptionModel == nil {
		var ret MarketCreateDataRelationshipsSubscriptionModel
		return ret
	}
	return *o.SubscriptionModel
}

// GetSubscriptionModelOk returns a tuple with the SubscriptionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetSubscriptionModelOk() (*MarketCreateDataRelationshipsSubscriptionModel, bool) {
	if o == nil || o.SubscriptionModel == nil {
		return nil, false
	}
	return o.SubscriptionModel, true
}

// HasSubscriptionModel returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasSubscriptionModel() bool {
	if o != nil && o.SubscriptionModel != nil {
		return true
	}

	return false
}

// SetSubscriptionModel gets a reference to the given MarketCreateDataRelationshipsSubscriptionModel and assigns it to the SubscriptionModel field.
func (o *MarketUpdateDataRelationships) SetSubscriptionModel(v MarketCreateDataRelationshipsSubscriptionModel) {
	o.SubscriptionModel = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetTaxCalculator() MarketCreateDataRelationshipsTaxCalculator {
	if o == nil || o.TaxCalculator == nil {
		var ret MarketCreateDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetTaxCalculatorOk() (*MarketCreateDataRelationshipsTaxCalculator, bool) {
	if o == nil || o.TaxCalculator == nil {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasTaxCalculator() bool {
	if o != nil && o.TaxCalculator != nil {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given MarketCreateDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *MarketUpdateDataRelationships) SetTaxCalculator(v MarketCreateDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *MarketUpdateDataRelationships) GetCustomerGroup() CustomerCreateDataRelationshipsCustomerGroup {
	if o == nil || o.CustomerGroup == nil {
		var ret CustomerCreateDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketUpdateDataRelationships) GetCustomerGroupOk() (*CustomerCreateDataRelationshipsCustomerGroup, bool) {
	if o == nil || o.CustomerGroup == nil {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *MarketUpdateDataRelationships) HasCustomerGroup() bool {
	if o != nil && o.CustomerGroup != nil {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given CustomerCreateDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *MarketUpdateDataRelationships) SetCustomerGroup(v CustomerCreateDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

func (o MarketUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Merchant != nil {
		toSerialize["merchant"] = o.Merchant
	}
	if o.PriceList != nil {
		toSerialize["price_list"] = o.PriceList
	}
	if o.InventoryModel != nil {
		toSerialize["inventory_model"] = o.InventoryModel
	}
	if o.SubscriptionModel != nil {
		toSerialize["subscription_model"] = o.SubscriptionModel
	}
	if o.TaxCalculator != nil {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if o.CustomerGroup != nil {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	return json.Marshal(toSerialize)
}

type NullableMarketUpdateDataRelationships struct {
	value *MarketUpdateDataRelationships
	isSet bool
}

func (v NullableMarketUpdateDataRelationships) Get() *MarketUpdateDataRelationships {
	return v.value
}

func (v *NullableMarketUpdateDataRelationships) Set(val *MarketUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketUpdateDataRelationships(val *MarketUpdateDataRelationships) *NullableMarketUpdateDataRelationships {
	return &NullableMarketUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableMarketUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
