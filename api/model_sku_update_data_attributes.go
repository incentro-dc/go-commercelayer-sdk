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

// SkuUpdateDataAttributes struct for SkuUpdateDataAttributes
type SkuUpdateDataAttributes struct {
	// The SKU code, that uniquely identifies the SKU within the organization.
	Code *string `json:"code,omitempty"`
	// The internal name of the SKU.
	Name *string `json:"name,omitempty"`
	// An internal description of the SKU.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the SKU.
	ImageUrl *string `json:"image_url,omitempty"`
	// The number of pieces that compose the SKU. This is useful to describe sets and bundles.
	PiecesPerPack *int32 `json:"pieces_per_pack,omitempty"`
	// The weight of the SKU. If present, it will be used to calculate the shipping rates.
	Weight *float32 `json:"weight,omitempty"`
	// Can be one of 'gr', or 'oz'
	UnitOfWeight *string `json:"unit_of_weight,omitempty"`
	// The Harmonized System Code used by customs to identify the products shipped across international borders.
	HsTariffNumber *string `json:"hs_tariff_number,omitempty"`
	// Indicates if the SKU doesn't generate shipments.
	DoNotShip *bool `json:"do_not_ship,omitempty"`
	// Indicates if the SKU doesn't track the stock inventory.
	DoNotTrack *bool `json:"do_not_track,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewSkuUpdateDataAttributes instantiates a new SkuUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuUpdateDataAttributes() *SkuUpdateDataAttributes {
	this := SkuUpdateDataAttributes{}
	return &this
}

// NewSkuUpdateDataAttributesWithDefaults instantiates a new SkuUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuUpdateDataAttributesWithDefaults() *SkuUpdateDataAttributes {
	this := SkuUpdateDataAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SkuUpdateDataAttributes) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SkuUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SkuUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *SkuUpdateDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetPiecesPerPack returns the PiecesPerPack field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetPiecesPerPack() int32 {
	if o == nil || o.PiecesPerPack == nil {
		var ret int32
		return ret
	}
	return *o.PiecesPerPack
}

// GetPiecesPerPackOk returns a tuple with the PiecesPerPack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetPiecesPerPackOk() (*int32, bool) {
	if o == nil || o.PiecesPerPack == nil {
		return nil, false
	}
	return o.PiecesPerPack, true
}

// HasPiecesPerPack returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasPiecesPerPack() bool {
	if o != nil && o.PiecesPerPack != nil {
		return true
	}

	return false
}

// SetPiecesPerPack gets a reference to the given int32 and assigns it to the PiecesPerPack field.
func (o *SkuUpdateDataAttributes) SetPiecesPerPack(v int32) {
	o.PiecesPerPack = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *SkuUpdateDataAttributes) SetWeight(v float32) {
	o.Weight = &v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetUnitOfWeight() string {
	if o == nil || o.UnitOfWeight == nil {
		var ret string
		return ret
	}
	return *o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetUnitOfWeightOk() (*string, bool) {
	if o == nil || o.UnitOfWeight == nil {
		return nil, false
	}
	return o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasUnitOfWeight() bool {
	if o != nil && o.UnitOfWeight != nil {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given string and assigns it to the UnitOfWeight field.
func (o *SkuUpdateDataAttributes) SetUnitOfWeight(v string) {
	o.UnitOfWeight = &v
}

// GetHsTariffNumber returns the HsTariffNumber field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetHsTariffNumber() string {
	if o == nil || o.HsTariffNumber == nil {
		var ret string
		return ret
	}
	return *o.HsTariffNumber
}

// GetHsTariffNumberOk returns a tuple with the HsTariffNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetHsTariffNumberOk() (*string, bool) {
	if o == nil || o.HsTariffNumber == nil {
		return nil, false
	}
	return o.HsTariffNumber, true
}

// HasHsTariffNumber returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasHsTariffNumber() bool {
	if o != nil && o.HsTariffNumber != nil {
		return true
	}

	return false
}

// SetHsTariffNumber gets a reference to the given string and assigns it to the HsTariffNumber field.
func (o *SkuUpdateDataAttributes) SetHsTariffNumber(v string) {
	o.HsTariffNumber = &v
}

// GetDoNotShip returns the DoNotShip field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetDoNotShip() bool {
	if o == nil || o.DoNotShip == nil {
		var ret bool
		return ret
	}
	return *o.DoNotShip
}

// GetDoNotShipOk returns a tuple with the DoNotShip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetDoNotShipOk() (*bool, bool) {
	if o == nil || o.DoNotShip == nil {
		return nil, false
	}
	return o.DoNotShip, true
}

// HasDoNotShip returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasDoNotShip() bool {
	if o != nil && o.DoNotShip != nil {
		return true
	}

	return false
}

// SetDoNotShip gets a reference to the given bool and assigns it to the DoNotShip field.
func (o *SkuUpdateDataAttributes) SetDoNotShip(v bool) {
	o.DoNotShip = &v
}

// GetDoNotTrack returns the DoNotTrack field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetDoNotTrack() bool {
	if o == nil || o.DoNotTrack == nil {
		var ret bool
		return ret
	}
	return *o.DoNotTrack
}

// GetDoNotTrackOk returns a tuple with the DoNotTrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetDoNotTrackOk() (*bool, bool) {
	if o == nil || o.DoNotTrack == nil {
		return nil, false
	}
	return o.DoNotTrack, true
}

// HasDoNotTrack returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasDoNotTrack() bool {
	if o != nil && o.DoNotTrack != nil {
		return true
	}

	return false
}

// SetDoNotTrack gets a reference to the given bool and assigns it to the DoNotTrack field.
func (o *SkuUpdateDataAttributes) SetDoNotTrack(v bool) {
	o.DoNotTrack = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *SkuUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *SkuUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SkuUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SkuUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SkuUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SkuUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.PiecesPerPack != nil {
		toSerialize["pieces_per_pack"] = o.PiecesPerPack
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.HsTariffNumber != nil {
		toSerialize["hs_tariff_number"] = o.HsTariffNumber
	}
	if o.DoNotShip != nil {
		toSerialize["do_not_ship"] = o.DoNotShip
	}
	if o.DoNotTrack != nil {
		toSerialize["do_not_track"] = o.DoNotTrack
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

type NullableSkuUpdateDataAttributes struct {
	value *SkuUpdateDataAttributes
	isSet bool
}

func (v NullableSkuUpdateDataAttributes) Get() *SkuUpdateDataAttributes {
	return v.value
}

func (v *NullableSkuUpdateDataAttributes) Set(val *SkuUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuUpdateDataAttributes(val *SkuUpdateDataAttributes) *NullableSkuUpdateDataAttributes {
	return &NullableSkuUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableSkuUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
