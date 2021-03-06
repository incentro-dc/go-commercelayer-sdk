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

// ShipmentDataRelationships struct for ShipmentDataRelationships
type ShipmentDataRelationships struct {
	Order            *AdyenPaymentDataRelationshipsOrder              `json:"order,omitempty"`
	ShippingCategory *ShipmentDataRelationshipsShippingCategory       `json:"shipping_category,omitempty"`
	StockLocation    *DeliveryLeadTimeDataRelationshipsStockLocation  `json:"stock_location,omitempty"`
	OriginAddress    *BingGeocoderDataRelationshipsAddresses          `json:"origin_address,omitempty"`
	ShippingAddress  *BingGeocoderDataRelationshipsAddresses          `json:"shipping_address,omitempty"`
	ShippingMethod   *DeliveryLeadTimeDataRelationshipsShippingMethod `json:"shipping_method,omitempty"`
	DeliveryLeadTime *ShipmentDataRelationshipsDeliveryLeadTime       `json:"delivery_lead_time,omitempty"`
	// Deprecated
	ShipmentLineItems        *LineItemDataRelationshipsShipmentLineItems      `json:"shipment_line_items,omitempty"`
	StockLineItems           *LineItemDataRelationshipsStockLineItems         `json:"stock_line_items,omitempty"`
	StockTransfers           *LineItemDataRelationshipsStockTransfers         `json:"stock_transfers,omitempty"`
	AvailableShippingMethods *DeliveryLeadTimeDataRelationshipsShippingMethod `json:"available_shipping_methods,omitempty"`
	CarrierAccounts          *ShipmentDataRelationshipsCarrierAccounts        `json:"carrier_accounts,omitempty"`
	Parcels                  *PackageDataRelationshipsParcels                 `json:"parcels,omitempty"`
	Attachments              *AvalaraAccountDataRelationshipsAttachments      `json:"attachments,omitempty"`
	Events                   *CustomerAddressDataRelationshipsEvents          `json:"events,omitempty"`
}

// NewShipmentDataRelationships instantiates a new ShipmentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDataRelationships() *ShipmentDataRelationships {
	this := ShipmentDataRelationships{}
	return &this
}

// NewShipmentDataRelationshipsWithDefaults instantiates a new ShipmentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDataRelationshipsWithDefaults() *ShipmentDataRelationships {
	this := ShipmentDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *ShipmentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret ShipmentDataRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given ShipmentDataRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *ShipmentDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetStockLocation() DeliveryLeadTimeDataRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret DeliveryLeadTimeDataRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeDataRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given DeliveryLeadTimeDataRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *ShipmentDataRelationships) SetStockLocation(v DeliveryLeadTimeDataRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetOriginAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.OriginAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetOriginAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.OriginAddress == nil {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasOriginAddress() bool {
	if o != nil && o.OriginAddress != nil {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the OriginAddress field.
func (o *ShipmentDataRelationships) SetOriginAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.OriginAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses {
	if o == nil || o.ShippingAddress == nil {
		var ret BingGeocoderDataRelationshipsAddresses
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given BingGeocoderDataRelationshipsAddresses and assigns it to the ShippingAddress field.
func (o *ShipmentDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses) {
	o.ShippingAddress = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetShippingMethod() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil || o.ShippingMethod == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetShippingMethodOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given DeliveryLeadTimeDataRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *ShipmentDataRelationships) SetShippingMethod(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

// GetDeliveryLeadTime returns the DeliveryLeadTime field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetDeliveryLeadTime() ShipmentDataRelationshipsDeliveryLeadTime {
	if o == nil || o.DeliveryLeadTime == nil {
		var ret ShipmentDataRelationshipsDeliveryLeadTime
		return ret
	}
	return *o.DeliveryLeadTime
}

// GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetDeliveryLeadTimeOk() (*ShipmentDataRelationshipsDeliveryLeadTime, bool) {
	if o == nil || o.DeliveryLeadTime == nil {
		return nil, false
	}
	return o.DeliveryLeadTime, true
}

// HasDeliveryLeadTime returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasDeliveryLeadTime() bool {
	if o != nil && o.DeliveryLeadTime != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTime gets a reference to the given ShipmentDataRelationshipsDeliveryLeadTime and assigns it to the DeliveryLeadTime field.
func (o *ShipmentDataRelationships) SetDeliveryLeadTime(v ShipmentDataRelationshipsDeliveryLeadTime) {
	o.DeliveryLeadTime = &v
}

// GetShipmentLineItems returns the ShipmentLineItems field value if set, zero value otherwise.
// Deprecated
func (o *ShipmentDataRelationships) GetShipmentLineItems() LineItemDataRelationshipsShipmentLineItems {
	if o == nil || o.ShipmentLineItems == nil {
		var ret LineItemDataRelationshipsShipmentLineItems
		return ret
	}
	return *o.ShipmentLineItems
}

// GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ShipmentDataRelationships) GetShipmentLineItemsOk() (*LineItemDataRelationshipsShipmentLineItems, bool) {
	if o == nil || o.ShipmentLineItems == nil {
		return nil, false
	}
	return o.ShipmentLineItems, true
}

// HasShipmentLineItems returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasShipmentLineItems() bool {
	if o != nil && o.ShipmentLineItems != nil {
		return true
	}

	return false
}

// SetShipmentLineItems gets a reference to the given LineItemDataRelationshipsShipmentLineItems and assigns it to the ShipmentLineItems field.
// Deprecated
func (o *ShipmentDataRelationships) SetShipmentLineItems(v LineItemDataRelationshipsShipmentLineItems) {
	o.ShipmentLineItems = &v
}

// GetStockLineItems returns the StockLineItems field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetStockLineItems() LineItemDataRelationshipsStockLineItems {
	if o == nil || o.StockLineItems == nil {
		var ret LineItemDataRelationshipsStockLineItems
		return ret
	}
	return *o.StockLineItems
}

// GetStockLineItemsOk returns a tuple with the StockLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetStockLineItemsOk() (*LineItemDataRelationshipsStockLineItems, bool) {
	if o == nil || o.StockLineItems == nil {
		return nil, false
	}
	return o.StockLineItems, true
}

// HasStockLineItems returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasStockLineItems() bool {
	if o != nil && o.StockLineItems != nil {
		return true
	}

	return false
}

// SetStockLineItems gets a reference to the given LineItemDataRelationshipsStockLineItems and assigns it to the StockLineItems field.
func (o *ShipmentDataRelationships) SetStockLineItems(v LineItemDataRelationshipsStockLineItems) {
	o.StockLineItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetStockTransfers() LineItemDataRelationshipsStockTransfers {
	if o == nil || o.StockTransfers == nil {
		var ret LineItemDataRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetStockTransfersOk() (*LineItemDataRelationshipsStockTransfers, bool) {
	if o == nil || o.StockTransfers == nil {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasStockTransfers() bool {
	if o != nil && o.StockTransfers != nil {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given LineItemDataRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *ShipmentDataRelationships) SetStockTransfers(v LineItemDataRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetAvailableShippingMethods returns the AvailableShippingMethods field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetAvailableShippingMethods() DeliveryLeadTimeDataRelationshipsShippingMethod {
	if o == nil || o.AvailableShippingMethods == nil {
		var ret DeliveryLeadTimeDataRelationshipsShippingMethod
		return ret
	}
	return *o.AvailableShippingMethods
}

// GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetAvailableShippingMethodsOk() (*DeliveryLeadTimeDataRelationshipsShippingMethod, bool) {
	if o == nil || o.AvailableShippingMethods == nil {
		return nil, false
	}
	return o.AvailableShippingMethods, true
}

// HasAvailableShippingMethods returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasAvailableShippingMethods() bool {
	if o != nil && o.AvailableShippingMethods != nil {
		return true
	}

	return false
}

// SetAvailableShippingMethods gets a reference to the given DeliveryLeadTimeDataRelationshipsShippingMethod and assigns it to the AvailableShippingMethods field.
func (o *ShipmentDataRelationships) SetAvailableShippingMethods(v DeliveryLeadTimeDataRelationshipsShippingMethod) {
	o.AvailableShippingMethods = &v
}

// GetCarrierAccounts returns the CarrierAccounts field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetCarrierAccounts() ShipmentDataRelationshipsCarrierAccounts {
	if o == nil || o.CarrierAccounts == nil {
		var ret ShipmentDataRelationshipsCarrierAccounts
		return ret
	}
	return *o.CarrierAccounts
}

// GetCarrierAccountsOk returns a tuple with the CarrierAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetCarrierAccountsOk() (*ShipmentDataRelationshipsCarrierAccounts, bool) {
	if o == nil || o.CarrierAccounts == nil {
		return nil, false
	}
	return o.CarrierAccounts, true
}

// HasCarrierAccounts returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasCarrierAccounts() bool {
	if o != nil && o.CarrierAccounts != nil {
		return true
	}

	return false
}

// SetCarrierAccounts gets a reference to the given ShipmentDataRelationshipsCarrierAccounts and assigns it to the CarrierAccounts field.
func (o *ShipmentDataRelationships) SetCarrierAccounts(v ShipmentDataRelationshipsCarrierAccounts) {
	o.CarrierAccounts = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetParcels() PackageDataRelationshipsParcels {
	if o == nil || o.Parcels == nil {
		var ret PackageDataRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetParcelsOk() (*PackageDataRelationshipsParcels, bool) {
	if o == nil || o.Parcels == nil {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasParcels() bool {
	if o != nil && o.Parcels != nil {
		return true
	}

	return false
}

// SetParcels gets a reference to the given PackageDataRelationshipsParcels and assigns it to the Parcels field.
func (o *ShipmentDataRelationships) SetParcels(v PackageDataRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShipmentDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ShipmentDataRelationships) GetEvents() CustomerAddressDataRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret CustomerAddressDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDataRelationships) GetEventsOk() (*CustomerAddressDataRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ShipmentDataRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given CustomerAddressDataRelationshipsEvents and assigns it to the Events field.
func (o *ShipmentDataRelationships) SetEvents(v CustomerAddressDataRelationshipsEvents) {
	o.Events = &v
}

func (o ShipmentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ShippingCategory != nil {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	if o.StockLocation != nil {
		toSerialize["stock_location"] = o.StockLocation
	}
	if o.OriginAddress != nil {
		toSerialize["origin_address"] = o.OriginAddress
	}
	if o.ShippingAddress != nil {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if o.ShippingMethod != nil {
		toSerialize["shipping_method"] = o.ShippingMethod
	}
	if o.DeliveryLeadTime != nil {
		toSerialize["delivery_lead_time"] = o.DeliveryLeadTime
	}
	if o.ShipmentLineItems != nil {
		toSerialize["shipment_line_items"] = o.ShipmentLineItems
	}
	if o.StockLineItems != nil {
		toSerialize["stock_line_items"] = o.StockLineItems
	}
	if o.StockTransfers != nil {
		toSerialize["stock_transfers"] = o.StockTransfers
	}
	if o.AvailableShippingMethods != nil {
		toSerialize["available_shipping_methods"] = o.AvailableShippingMethods
	}
	if o.CarrierAccounts != nil {
		toSerialize["carrier_accounts"] = o.CarrierAccounts
	}
	if o.Parcels != nil {
		toSerialize["parcels"] = o.Parcels
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentDataRelationships struct {
	value *ShipmentDataRelationships
	isSet bool
}

func (v NullableShipmentDataRelationships) Get() *ShipmentDataRelationships {
	return v.value
}

func (v *NullableShipmentDataRelationships) Set(val *ShipmentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDataRelationships(val *ShipmentDataRelationships) *NullableShipmentDataRelationships {
	return &NullableShipmentDataRelationships{value: val, isSet: true}
}

func (v NullableShipmentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
