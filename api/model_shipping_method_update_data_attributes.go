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

// ShippingMethodUpdateDataAttributes struct for ShippingMethodUpdateDataAttributes
type ShippingMethodUpdateDataAttributes struct {
	// The shipping method's name
	Name *string `json:"name,omitempty"`
	// The shipping method's scheme, one of 'flat' or 'weight_tiered'.
	Scheme *string `json:"scheme,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The price of this shipping method, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// Apply free shipping if the order amount is over this value, in cents.
	FreeOverAmountCents *int32 `json:"free_over_amount_cents,omitempty"`
	// The minimum weight for which this shipping method is available.
	MinWeight *float32 `json:"min_weight,omitempty"`
	// The maximum weight for which this shipping method is available.
	MaxWeight *float32 `json:"max_weight,omitempty"`
	// Can be one of 'gr', 'lb', or 'oz'
	UnitOfWeight *string `json:"unit_of_weight,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewShippingMethodUpdateDataAttributes instantiates a new ShippingMethodUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodUpdateDataAttributes() *ShippingMethodUpdateDataAttributes {
	this := ShippingMethodUpdateDataAttributes{}
	return &this
}

// NewShippingMethodUpdateDataAttributesWithDefaults instantiates a new ShippingMethodUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodUpdateDataAttributesWithDefaults() *ShippingMethodUpdateDataAttributes {
	this := ShippingMethodUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShippingMethodUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *ShippingMethodUpdateDataAttributes) SetScheme(v string) {
	o.Scheme = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ShippingMethodUpdateDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *ShippingMethodUpdateDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetFreeOverAmountCents returns the FreeOverAmountCents field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetFreeOverAmountCents() int32 {
	if o == nil || o.FreeOverAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.FreeOverAmountCents
}

// GetFreeOverAmountCentsOk returns a tuple with the FreeOverAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetFreeOverAmountCentsOk() (*int32, bool) {
	if o == nil || o.FreeOverAmountCents == nil {
		return nil, false
	}
	return o.FreeOverAmountCents, true
}

// HasFreeOverAmountCents returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasFreeOverAmountCents() bool {
	if o != nil && o.FreeOverAmountCents != nil {
		return true
	}

	return false
}

// SetFreeOverAmountCents gets a reference to the given int32 and assigns it to the FreeOverAmountCents field.
func (o *ShippingMethodUpdateDataAttributes) SetFreeOverAmountCents(v int32) {
	o.FreeOverAmountCents = &v
}

// GetMinWeight returns the MinWeight field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetMinWeight() float32 {
	if o == nil || o.MinWeight == nil {
		var ret float32
		return ret
	}
	return *o.MinWeight
}

// GetMinWeightOk returns a tuple with the MinWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetMinWeightOk() (*float32, bool) {
	if o == nil || o.MinWeight == nil {
		return nil, false
	}
	return o.MinWeight, true
}

// HasMinWeight returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasMinWeight() bool {
	if o != nil && o.MinWeight != nil {
		return true
	}

	return false
}

// SetMinWeight gets a reference to the given float32 and assigns it to the MinWeight field.
func (o *ShippingMethodUpdateDataAttributes) SetMinWeight(v float32) {
	o.MinWeight = &v
}

// GetMaxWeight returns the MaxWeight field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetMaxWeight() float32 {
	if o == nil || o.MaxWeight == nil {
		var ret float32
		return ret
	}
	return *o.MaxWeight
}

// GetMaxWeightOk returns a tuple with the MaxWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetMaxWeightOk() (*float32, bool) {
	if o == nil || o.MaxWeight == nil {
		return nil, false
	}
	return o.MaxWeight, true
}

// HasMaxWeight returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasMaxWeight() bool {
	if o != nil && o.MaxWeight != nil {
		return true
	}

	return false
}

// SetMaxWeight gets a reference to the given float32 and assigns it to the MaxWeight field.
func (o *ShippingMethodUpdateDataAttributes) SetMaxWeight(v float32) {
	o.MaxWeight = &v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetUnitOfWeight() string {
	if o == nil || o.UnitOfWeight == nil {
		var ret string
		return ret
	}
	return *o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetUnitOfWeightOk() (*string, bool) {
	if o == nil || o.UnitOfWeight == nil {
		return nil, false
	}
	return o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasUnitOfWeight() bool {
	if o != nil && o.UnitOfWeight != nil {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given string and assigns it to the UnitOfWeight field.
func (o *ShippingMethodUpdateDataAttributes) SetUnitOfWeight(v string) {
	o.UnitOfWeight = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ShippingMethodUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ShippingMethodUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShippingMethodUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShippingMethodUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ShippingMethodUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ShippingMethodUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.FreeOverAmountCents != nil {
		toSerialize["free_over_amount_cents"] = o.FreeOverAmountCents
	}
	if o.MinWeight != nil {
		toSerialize["min_weight"] = o.MinWeight
	}
	if o.MaxWeight != nil {
		toSerialize["max_weight"] = o.MaxWeight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
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

type NullableShippingMethodUpdateDataAttributes struct {
	value *ShippingMethodUpdateDataAttributes
	isSet bool
}

func (v NullableShippingMethodUpdateDataAttributes) Get() *ShippingMethodUpdateDataAttributes {
	return v.value
}

func (v *NullableShippingMethodUpdateDataAttributes) Set(val *ShippingMethodUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodUpdateDataAttributes(val *ShippingMethodUpdateDataAttributes) *NullableShippingMethodUpdateDataAttributes {
	return &NullableShippingMethodUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableShippingMethodUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
