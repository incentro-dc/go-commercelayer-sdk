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

// PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct for PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes
type PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct {
	// The stock location priority within the associated invetory model.
	Priority *int32 `json:"priority,omitempty"`
	// Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled.
	OnHold *bool `json:"on_hold,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes instantiates a new PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes() *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	this := PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{}
	return &this
}

// NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributesWithDefaults instantiates a new PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributesWithDefaults() *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	this := PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetPriority(v int32) {
	o.Priority = &v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetOnHold() bool {
	if o == nil || o.OnHold == nil {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetOnHoldOk() (*bool, bool) {
	if o == nil || o.OnHold == nil {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasOnHold() bool {
	if o != nil && o.OnHold != nil {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetOnHold(v bool) {
	o.OnHold = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.OnHold != nil {
		toSerialize["on_hold"] = o.OnHold
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes struct {
	value *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Get() *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Set(val *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes(val *PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) *NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes {
	return &NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
