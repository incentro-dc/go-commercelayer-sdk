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

// checks if the POSTMarketsRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTMarketsRequestDataRelationships{}

// POSTMarketsRequestDataRelationships struct for POSTMarketsRequestDataRelationships
type POSTMarketsRequestDataRelationships struct {
	Merchant          POSTMarketsRequestDataRelationshipsMerchant                        `json:"merchant"`
	PriceList         POSTMarketsRequestDataRelationshipsPriceList                       `json:"price_list"`
	InventoryModel    POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel `json:"inventory_model"`
	SubscriptionModel *POSTMarketsRequestDataRelationshipsSubscriptionModel              `json:"subscription_model,omitempty"`
	TaxCalculator     *POSTMarketsRequestDataRelationshipsTaxCalculator                  `json:"tax_calculator,omitempty"`
	CustomerGroup     *POSTCustomersRequestDataRelationshipsCustomerGroup                `json:"customer_group,omitempty"`
}

// NewPOSTMarketsRequestDataRelationships instantiates a new POSTMarketsRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTMarketsRequestDataRelationships(merchant POSTMarketsRequestDataRelationshipsMerchant, priceList POSTMarketsRequestDataRelationshipsPriceList, inventoryModel POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel) *POSTMarketsRequestDataRelationships {
	this := POSTMarketsRequestDataRelationships{}
	this.Merchant = merchant
	this.PriceList = priceList
	this.InventoryModel = inventoryModel
	return &this
}

// NewPOSTMarketsRequestDataRelationshipsWithDefaults instantiates a new POSTMarketsRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTMarketsRequestDataRelationshipsWithDefaults() *POSTMarketsRequestDataRelationships {
	this := POSTMarketsRequestDataRelationships{}
	return &this
}

// GetMerchant returns the Merchant field value
func (o *POSTMarketsRequestDataRelationships) GetMerchant() POSTMarketsRequestDataRelationshipsMerchant {
	if o == nil {
		var ret POSTMarketsRequestDataRelationshipsMerchant
		return ret
	}

	return o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetMerchantOk() (*POSTMarketsRequestDataRelationshipsMerchant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchant, true
}

// SetMerchant sets field value
func (o *POSTMarketsRequestDataRelationships) SetMerchant(v POSTMarketsRequestDataRelationshipsMerchant) {
	o.Merchant = v
}

// GetPriceList returns the PriceList field value
func (o *POSTMarketsRequestDataRelationships) GetPriceList() POSTMarketsRequestDataRelationshipsPriceList {
	if o == nil {
		var ret POSTMarketsRequestDataRelationshipsPriceList
		return ret
	}

	return o.PriceList
}

// GetPriceListOk returns a tuple with the PriceList field value
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetPriceListOk() (*POSTMarketsRequestDataRelationshipsPriceList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PriceList, true
}

// SetPriceList sets field value
func (o *POSTMarketsRequestDataRelationships) SetPriceList(v POSTMarketsRequestDataRelationshipsPriceList) {
	o.PriceList = v
}

// GetInventoryModel returns the InventoryModel field value
func (o *POSTMarketsRequestDataRelationships) GetInventoryModel() POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel {
	if o == nil {
		var ret POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel
		return ret
	}

	return o.InventoryModel
}

// GetInventoryModelOk returns a tuple with the InventoryModel field value
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetInventoryModelOk() (*POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventoryModel, true
}

// SetInventoryModel sets field value
func (o *POSTMarketsRequestDataRelationships) SetInventoryModel(v POSTInventoryReturnLocationsRequestDataRelationshipsInventoryModel) {
	o.InventoryModel = v
}

// GetSubscriptionModel returns the SubscriptionModel field value if set, zero value otherwise.
func (o *POSTMarketsRequestDataRelationships) GetSubscriptionModel() POSTMarketsRequestDataRelationshipsSubscriptionModel {
	if o == nil || IsNil(o.SubscriptionModel) {
		var ret POSTMarketsRequestDataRelationshipsSubscriptionModel
		return ret
	}
	return *o.SubscriptionModel
}

// GetSubscriptionModelOk returns a tuple with the SubscriptionModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetSubscriptionModelOk() (*POSTMarketsRequestDataRelationshipsSubscriptionModel, bool) {
	if o == nil || IsNil(o.SubscriptionModel) {
		return nil, false
	}
	return o.SubscriptionModel, true
}

// HasSubscriptionModel returns a boolean if a field has been set.
func (o *POSTMarketsRequestDataRelationships) HasSubscriptionModel() bool {
	if o != nil && !IsNil(o.SubscriptionModel) {
		return true
	}

	return false
}

// SetSubscriptionModel gets a reference to the given POSTMarketsRequestDataRelationshipsSubscriptionModel and assigns it to the SubscriptionModel field.
func (o *POSTMarketsRequestDataRelationships) SetSubscriptionModel(v POSTMarketsRequestDataRelationshipsSubscriptionModel) {
	o.SubscriptionModel = &v
}

// GetTaxCalculator returns the TaxCalculator field value if set, zero value otherwise.
func (o *POSTMarketsRequestDataRelationships) GetTaxCalculator() POSTMarketsRequestDataRelationshipsTaxCalculator {
	if o == nil || IsNil(o.TaxCalculator) {
		var ret POSTMarketsRequestDataRelationshipsTaxCalculator
		return ret
	}
	return *o.TaxCalculator
}

// GetTaxCalculatorOk returns a tuple with the TaxCalculator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetTaxCalculatorOk() (*POSTMarketsRequestDataRelationshipsTaxCalculator, bool) {
	if o == nil || IsNil(o.TaxCalculator) {
		return nil, false
	}
	return o.TaxCalculator, true
}

// HasTaxCalculator returns a boolean if a field has been set.
func (o *POSTMarketsRequestDataRelationships) HasTaxCalculator() bool {
	if o != nil && !IsNil(o.TaxCalculator) {
		return true
	}

	return false
}

// SetTaxCalculator gets a reference to the given POSTMarketsRequestDataRelationshipsTaxCalculator and assigns it to the TaxCalculator field.
func (o *POSTMarketsRequestDataRelationships) SetTaxCalculator(v POSTMarketsRequestDataRelationshipsTaxCalculator) {
	o.TaxCalculator = &v
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *POSTMarketsRequestDataRelationships) GetCustomerGroup() POSTCustomersRequestDataRelationshipsCustomerGroup {
	if o == nil || IsNil(o.CustomerGroup) {
		var ret POSTCustomersRequestDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTMarketsRequestDataRelationships) GetCustomerGroupOk() (*POSTCustomersRequestDataRelationshipsCustomerGroup, bool) {
	if o == nil || IsNil(o.CustomerGroup) {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *POSTMarketsRequestDataRelationships) HasCustomerGroup() bool {
	if o != nil && !IsNil(o.CustomerGroup) {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given POSTCustomersRequestDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *POSTMarketsRequestDataRelationships) SetCustomerGroup(v POSTCustomersRequestDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

func (o POSTMarketsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTMarketsRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchant"] = o.Merchant
	toSerialize["price_list"] = o.PriceList
	toSerialize["inventory_model"] = o.InventoryModel
	if !IsNil(o.SubscriptionModel) {
		toSerialize["subscription_model"] = o.SubscriptionModel
	}
	if !IsNil(o.TaxCalculator) {
		toSerialize["tax_calculator"] = o.TaxCalculator
	}
	if !IsNil(o.CustomerGroup) {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	return toSerialize, nil
}

type NullablePOSTMarketsRequestDataRelationships struct {
	value *POSTMarketsRequestDataRelationships
	isSet bool
}

func (v NullablePOSTMarketsRequestDataRelationships) Get() *POSTMarketsRequestDataRelationships {
	return v.value
}

func (v *NullablePOSTMarketsRequestDataRelationships) Set(val *POSTMarketsRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTMarketsRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTMarketsRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTMarketsRequestDataRelationships(val *POSTMarketsRequestDataRelationships) *NullablePOSTMarketsRequestDataRelationships {
	return &NullablePOSTMarketsRequestDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTMarketsRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTMarketsRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}