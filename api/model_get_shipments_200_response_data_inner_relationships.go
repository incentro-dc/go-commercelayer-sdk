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

// GETShipments200ResponseDataInnerRelationships struct for GETShipments200ResponseDataInnerRelationships
type GETShipments200ResponseDataInnerRelationships struct {
	Order            *GETAdyenPayments200ResponseDataInnerRelationshipsOrder              `json:"order,omitempty"`
	ShippingCategory *GETShipments200ResponseDataInnerRelationshipsShippingCategory       `json:"shipping_category,omitempty"`
	StockLocation    *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation  `json:"stock_location,omitempty"`
	OriginAddress    *GETBingGeocoders200ResponseDataInnerRelationshipsAddresses          `json:"origin_address,omitempty"`
	ShippingAddress  *GETBingGeocoders200ResponseDataInnerRelationshipsAddresses          `json:"shipping_address,omitempty"`
	ShippingMethod   *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod `json:"shipping_method,omitempty"`
	DeliveryLeadTime *GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime       `json:"delivery_lead_time,omitempty"`
	// Deprecated
	ShipmentLineItems        *GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems      `json:"shipment_line_items,omitempty"`
	StockLineItems           *GETLineItems200ResponseDataInnerRelationshipsStockLineItems         `json:"stock_line_items,omitempty"`
	StockTransfers           *GETLineItems200ResponseDataInnerRelationshipsStockTransfers         `json:"stock_transfers,omitempty"`
	AvailableShippingMethods *GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod `json:"available_shipping_methods,omitempty"`
	CarrierAccounts          *GETShipments200ResponseDataInnerRelationshipsCarrierAccounts        `json:"carrier_accounts,omitempty"`
	Parcels                  *GETPackages200ResponseDataInnerRelationshipsParcels                 `json:"parcels,omitempty"`
	Attachments              *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments      `json:"attachments,omitempty"`
	Events                   *GETCustomerAddresses200ResponseDataInnerRelationshipsEvents         `json:"events,omitempty"`
}

// NewGETShipments200ResponseDataInnerRelationships instantiates a new GETShipments200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShipments200ResponseDataInnerRelationships() *GETShipments200ResponseDataInnerRelationships {
	this := GETShipments200ResponseDataInnerRelationships{}
	return &this
}

// NewGETShipments200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETShipments200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShipments200ResponseDataInnerRelationshipsWithDefaults() *GETShipments200ResponseDataInnerRelationships {
	this := GETShipments200ResponseDataInnerRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrder and assigns it to the Order field.
func (o *GETShipments200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.Order = &v
}

// GetShippingCategory returns the ShippingCategory field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategory() GETShipments200ResponseDataInnerRelationshipsShippingCategory {
	if o == nil || o.ShippingCategory == nil {
		var ret GETShipments200ResponseDataInnerRelationshipsShippingCategory
		return ret
	}
	return *o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingCategoryOk() (*GETShipments200ResponseDataInnerRelationshipsShippingCategory, bool) {
	if o == nil || o.ShippingCategory == nil {
		return nil, false
	}
	return o.ShippingCategory, true
}

// HasShippingCategory returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasShippingCategory() bool {
	if o != nil && o.ShippingCategory != nil {
		return true
	}

	return false
}

// SetShippingCategory gets a reference to the given GETShipments200ResponseDataInnerRelationshipsShippingCategory and assigns it to the ShippingCategory field.
func (o *GETShipments200ResponseDataInnerRelationships) SetShippingCategory(v GETShipments200ResponseDataInnerRelationshipsShippingCategory) {
	o.ShippingCategory = &v
}

// GetStockLocation returns the StockLocation field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocation() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation {
	if o == nil || o.StockLocation == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation
		return ret
	}
	return *o.StockLocation
}

// GetStockLocationOk returns a tuple with the StockLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockLocationOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation, bool) {
	if o == nil || o.StockLocation == nil {
		return nil, false
	}
	return o.StockLocation, true
}

// HasStockLocation returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasStockLocation() bool {
	if o != nil && o.StockLocation != nil {
		return true
	}

	return false
}

// SetStockLocation gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation and assigns it to the StockLocation field.
func (o *GETShipments200ResponseDataInnerRelationships) SetStockLocation(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsStockLocation) {
	o.StockLocation = &v
}

// GetOriginAddress returns the OriginAddress field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddress() GETBingGeocoders200ResponseDataInnerRelationshipsAddresses {
	if o == nil || o.OriginAddress == nil {
		var ret GETBingGeocoders200ResponseDataInnerRelationshipsAddresses
		return ret
	}
	return *o.OriginAddress
}

// GetOriginAddressOk returns a tuple with the OriginAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetOriginAddressOk() (*GETBingGeocoders200ResponseDataInnerRelationshipsAddresses, bool) {
	if o == nil || o.OriginAddress == nil {
		return nil, false
	}
	return o.OriginAddress, true
}

// HasOriginAddress returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasOriginAddress() bool {
	if o != nil && o.OriginAddress != nil {
		return true
	}

	return false
}

// SetOriginAddress gets a reference to the given GETBingGeocoders200ResponseDataInnerRelationshipsAddresses and assigns it to the OriginAddress field.
func (o *GETShipments200ResponseDataInnerRelationships) SetOriginAddress(v GETBingGeocoders200ResponseDataInnerRelationshipsAddresses) {
	o.OriginAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddress() GETBingGeocoders200ResponseDataInnerRelationshipsAddresses {
	if o == nil || o.ShippingAddress == nil {
		var ret GETBingGeocoders200ResponseDataInnerRelationshipsAddresses
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingAddressOk() (*GETBingGeocoders200ResponseDataInnerRelationshipsAddresses, bool) {
	if o == nil || o.ShippingAddress == nil {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasShippingAddress() bool {
	if o != nil && o.ShippingAddress != nil {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given GETBingGeocoders200ResponseDataInnerRelationshipsAddresses and assigns it to the ShippingAddress field.
func (o *GETShipments200ResponseDataInnerRelationships) SetShippingAddress(v GETBingGeocoders200ResponseDataInnerRelationshipsAddresses) {
	o.ShippingAddress = &v
}

// GetShippingMethod returns the ShippingMethod field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethod() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod {
	if o == nil || o.ShippingMethod == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod
		return ret
	}
	return *o.ShippingMethod
}

// GetShippingMethodOk returns a tuple with the ShippingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetShippingMethodOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod, bool) {
	if o == nil || o.ShippingMethod == nil {
		return nil, false
	}
	return o.ShippingMethod, true
}

// HasShippingMethod returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasShippingMethod() bool {
	if o != nil && o.ShippingMethod != nil {
		return true
	}

	return false
}

// SetShippingMethod gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod and assigns it to the ShippingMethod field.
func (o *GETShipments200ResponseDataInnerRelationships) SetShippingMethod(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod) {
	o.ShippingMethod = &v
}

// GetDeliveryLeadTime returns the DeliveryLeadTime field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTime() GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime {
	if o == nil || o.DeliveryLeadTime == nil {
		var ret GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime
		return ret
	}
	return *o.DeliveryLeadTime
}

// GetDeliveryLeadTimeOk returns a tuple with the DeliveryLeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetDeliveryLeadTimeOk() (*GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime, bool) {
	if o == nil || o.DeliveryLeadTime == nil {
		return nil, false
	}
	return o.DeliveryLeadTime, true
}

// HasDeliveryLeadTime returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasDeliveryLeadTime() bool {
	if o != nil && o.DeliveryLeadTime != nil {
		return true
	}

	return false
}

// SetDeliveryLeadTime gets a reference to the given GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime and assigns it to the DeliveryLeadTime field.
func (o *GETShipments200ResponseDataInnerRelationships) SetDeliveryLeadTime(v GETShipments200ResponseDataInnerRelationshipsDeliveryLeadTime) {
	o.DeliveryLeadTime = &v
}

// GetShipmentLineItems returns the ShipmentLineItems field value if set, zero value otherwise.
// Deprecated
func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItems() GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems {
	if o == nil || o.ShipmentLineItems == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems
		return ret
	}
	return *o.ShipmentLineItems
}

// GetShipmentLineItemsOk returns a tuple with the ShipmentLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GETShipments200ResponseDataInnerRelationships) GetShipmentLineItemsOk() (*GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems, bool) {
	if o == nil || o.ShipmentLineItems == nil {
		return nil, false
	}
	return o.ShipmentLineItems, true
}

// HasShipmentLineItems returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasShipmentLineItems() bool {
	if o != nil && o.ShipmentLineItems != nil {
		return true
	}

	return false
}

// SetShipmentLineItems gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems and assigns it to the ShipmentLineItems field.
// Deprecated
func (o *GETShipments200ResponseDataInnerRelationships) SetShipmentLineItems(v GETLineItems200ResponseDataInnerRelationshipsShipmentLineItems) {
	o.ShipmentLineItems = &v
}

// GetStockLineItems returns the StockLineItems field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItems() GETLineItems200ResponseDataInnerRelationshipsStockLineItems {
	if o == nil || o.StockLineItems == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsStockLineItems
		return ret
	}
	return *o.StockLineItems
}

// GetStockLineItemsOk returns a tuple with the StockLineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockLineItemsOk() (*GETLineItems200ResponseDataInnerRelationshipsStockLineItems, bool) {
	if o == nil || o.StockLineItems == nil {
		return nil, false
	}
	return o.StockLineItems, true
}

// HasStockLineItems returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasStockLineItems() bool {
	if o != nil && o.StockLineItems != nil {
		return true
	}

	return false
}

// SetStockLineItems gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsStockLineItems and assigns it to the StockLineItems field.
func (o *GETShipments200ResponseDataInnerRelationships) SetStockLineItems(v GETLineItems200ResponseDataInnerRelationshipsStockLineItems) {
	o.StockLineItems = &v
}

// GetStockTransfers returns the StockTransfers field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfers() GETLineItems200ResponseDataInnerRelationshipsStockTransfers {
	if o == nil || o.StockTransfers == nil {
		var ret GETLineItems200ResponseDataInnerRelationshipsStockTransfers
		return ret
	}
	return *o.StockTransfers
}

// GetStockTransfersOk returns a tuple with the StockTransfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetStockTransfersOk() (*GETLineItems200ResponseDataInnerRelationshipsStockTransfers, bool) {
	if o == nil || o.StockTransfers == nil {
		return nil, false
	}
	return o.StockTransfers, true
}

// HasStockTransfers returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasStockTransfers() bool {
	if o != nil && o.StockTransfers != nil {
		return true
	}

	return false
}

// SetStockTransfers gets a reference to the given GETLineItems200ResponseDataInnerRelationshipsStockTransfers and assigns it to the StockTransfers field.
func (o *GETShipments200ResponseDataInnerRelationships) SetStockTransfers(v GETLineItems200ResponseDataInnerRelationshipsStockTransfers) {
	o.StockTransfers = &v
}

// GetAvailableShippingMethods returns the AvailableShippingMethods field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethods() GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod {
	if o == nil || o.AvailableShippingMethods == nil {
		var ret GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod
		return ret
	}
	return *o.AvailableShippingMethods
}

// GetAvailableShippingMethodsOk returns a tuple with the AvailableShippingMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetAvailableShippingMethodsOk() (*GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod, bool) {
	if o == nil || o.AvailableShippingMethods == nil {
		return nil, false
	}
	return o.AvailableShippingMethods, true
}

// HasAvailableShippingMethods returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasAvailableShippingMethods() bool {
	if o != nil && o.AvailableShippingMethods != nil {
		return true
	}

	return false
}

// SetAvailableShippingMethods gets a reference to the given GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod and assigns it to the AvailableShippingMethods field.
func (o *GETShipments200ResponseDataInnerRelationships) SetAvailableShippingMethods(v GETDeliveryLeadTimes200ResponseDataInnerRelationshipsShippingMethod) {
	o.AvailableShippingMethods = &v
}

// GetCarrierAccounts returns the CarrierAccounts field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccounts() GETShipments200ResponseDataInnerRelationshipsCarrierAccounts {
	if o == nil || o.CarrierAccounts == nil {
		var ret GETShipments200ResponseDataInnerRelationshipsCarrierAccounts
		return ret
	}
	return *o.CarrierAccounts
}

// GetCarrierAccountsOk returns a tuple with the CarrierAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetCarrierAccountsOk() (*GETShipments200ResponseDataInnerRelationshipsCarrierAccounts, bool) {
	if o == nil || o.CarrierAccounts == nil {
		return nil, false
	}
	return o.CarrierAccounts, true
}

// HasCarrierAccounts returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasCarrierAccounts() bool {
	if o != nil && o.CarrierAccounts != nil {
		return true
	}

	return false
}

// SetCarrierAccounts gets a reference to the given GETShipments200ResponseDataInnerRelationshipsCarrierAccounts and assigns it to the CarrierAccounts field.
func (o *GETShipments200ResponseDataInnerRelationships) SetCarrierAccounts(v GETShipments200ResponseDataInnerRelationshipsCarrierAccounts) {
	o.CarrierAccounts = &v
}

// GetParcels returns the Parcels field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetParcels() GETPackages200ResponseDataInnerRelationshipsParcels {
	if o == nil || o.Parcels == nil {
		var ret GETPackages200ResponseDataInnerRelationshipsParcels
		return ret
	}
	return *o.Parcels
}

// GetParcelsOk returns a tuple with the Parcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetParcelsOk() (*GETPackages200ResponseDataInnerRelationshipsParcels, bool) {
	if o == nil || o.Parcels == nil {
		return nil, false
	}
	return o.Parcels, true
}

// HasParcels returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasParcels() bool {
	if o != nil && o.Parcels != nil {
		return true
	}

	return false
}

// SetParcels gets a reference to the given GETPackages200ResponseDataInnerRelationshipsParcels and assigns it to the Parcels field.
func (o *GETShipments200ResponseDataInnerRelationships) SetParcels(v GETPackages200ResponseDataInnerRelationshipsParcels) {
	o.Parcels = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETShipments200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETShipments200ResponseDataInnerRelationships) GetEvents() GETCustomerAddresses200ResponseDataInnerRelationshipsEvents {
	if o == nil || o.Events == nil {
		var ret GETCustomerAddresses200ResponseDataInnerRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShipments200ResponseDataInnerRelationships) GetEventsOk() (*GETCustomerAddresses200ResponseDataInnerRelationshipsEvents, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETShipments200ResponseDataInnerRelationships) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given GETCustomerAddresses200ResponseDataInnerRelationshipsEvents and assigns it to the Events field.
func (o *GETShipments200ResponseDataInnerRelationships) SetEvents(v GETCustomerAddresses200ResponseDataInnerRelationshipsEvents) {
	o.Events = &v
}

func (o GETShipments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
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

type NullableGETShipments200ResponseDataInnerRelationships struct {
	value *GETShipments200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETShipments200ResponseDataInnerRelationships) Get() *GETShipments200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETShipments200ResponseDataInnerRelationships) Set(val *GETShipments200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShipments200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShipments200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShipments200ResponseDataInnerRelationships(val *GETShipments200ResponseDataInnerRelationships) *NullableGETShipments200ResponseDataInnerRelationships {
	return &NullableGETShipments200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETShipments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShipments200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}