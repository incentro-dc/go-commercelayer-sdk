/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETBundles200ResponseDataInnerAttributes struct for GETBundles200ResponseDataInnerAttributes
type GETBundles200ResponseDataInnerAttributes struct {
	// The bundle code, that uniquely identifies the bundle within the market.
	Code *string `json:"code,omitempty"`
	// The internal name of the bundle.
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// An internal description of the bundle.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the bundle.
	ImageUrl *string `json:"image_url,omitempty"`
	// Indicates if the bundle doesn't generate shipments (all sku_list's SKUs must be do_not_ship).
	DoNotShip *bool `json:"do_not_ship,omitempty"`
	// Indicates if the bundle doesn't track the stock inventory (all sku_list's SKUs must be do_not_track).
	DoNotTrack *bool `json:"do_not_track,omitempty"`
	// The bundle price amount for the associated market, in cents.
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The bundle price amount for the associated market, float.
	PriceAmountFloat *float32 `json:"price_amount_float,omitempty"`
	// The bundle price amount for the associated market, formatted.
	FormattedPriceAmount *string `json:"formatted_price_amount,omitempty"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents *int32 `json:"compare_at_amount_cents,omitempty"`
	// The compared price amount, float.
	CompareAtAmountFloat *float32 `json:"compare_at_amount_float,omitempty"`
	// The compared price amount, formatted.
	FormattedCompareAtAmount *string `json:"formatted_compare_at_amount,omitempty"`
	// The total number of SKUs in the bundle.
	SkusCount *int32 `json:"skus_count,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETBundles200ResponseDataInnerAttributes instantiates a new GETBundles200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETBundles200ResponseDataInnerAttributes() *GETBundles200ResponseDataInnerAttributes {
	this := GETBundles200ResponseDataInnerAttributes{}
	return &this
}

// NewGETBundles200ResponseDataInnerAttributesWithDefaults instantiates a new GETBundles200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETBundles200ResponseDataInnerAttributesWithDefaults() *GETBundles200ResponseDataInnerAttributes {
	this := GETBundles200ResponseDataInnerAttributes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GETBundles200ResponseDataInnerAttributes) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETBundles200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GETBundles200ResponseDataInnerAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GETBundles200ResponseDataInnerAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *GETBundles200ResponseDataInnerAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetDoNotShip returns the DoNotShip field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShip() bool {
	if o == nil || o.DoNotShip == nil {
		var ret bool
		return ret
	}
	return *o.DoNotShip
}

// GetDoNotShipOk returns a tuple with the DoNotShip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotShipOk() (*bool, bool) {
	if o == nil || o.DoNotShip == nil {
		return nil, false
	}
	return o.DoNotShip, true
}

// HasDoNotShip returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotShip() bool {
	if o != nil && o.DoNotShip != nil {
		return true
	}

	return false
}

// SetDoNotShip gets a reference to the given bool and assigns it to the DoNotShip field.
func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotShip(v bool) {
	o.DoNotShip = &v
}

// GetDoNotTrack returns the DoNotTrack field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrack() bool {
	if o == nil || o.DoNotTrack == nil {
		var ret bool
		return ret
	}
	return *o.DoNotTrack
}

// GetDoNotTrackOk returns a tuple with the DoNotTrack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetDoNotTrackOk() (*bool, bool) {
	if o == nil || o.DoNotTrack == nil {
		return nil, false
	}
	return o.DoNotTrack, true
}

// HasDoNotTrack returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasDoNotTrack() bool {
	if o != nil && o.DoNotTrack != nil {
		return true
	}

	return false
}

// SetDoNotTrack gets a reference to the given bool and assigns it to the DoNotTrack field.
func (o *GETBundles200ResponseDataInnerAttributes) SetDoNotTrack(v bool) {
	o.DoNotTrack = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloat() float32 {
	if o == nil || o.PriceAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetPriceAmountFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountFloat == nil {
		return nil, false
	}
	return o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasPriceAmountFloat() bool {
	if o != nil && o.PriceAmountFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given float32 and assigns it to the PriceAmountFloat field.
func (o *GETBundles200ResponseDataInnerAttributes) SetPriceAmountFloat(v float32) {
	o.PriceAmountFloat = &v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmount() string {
	if o == nil || o.FormattedPriceAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedPriceAmountOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmount == nil {
		return nil, false
	}
	return o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedPriceAmount() bool {
	if o != nil && o.FormattedPriceAmount != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given string and assigns it to the FormattedPriceAmount field.
func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedPriceAmount(v string) {
	o.FormattedPriceAmount = &v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCents() int32 {
	if o == nil || o.CompareAtAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountCentsOk() (*int32, bool) {
	if o == nil || o.CompareAtAmountCents == nil {
		return nil, false
	}
	return o.CompareAtAmountCents, true
}

// HasCompareAtAmountCents returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountCents() bool {
	if o != nil && o.CompareAtAmountCents != nil {
		return true
	}

	return false
}

// SetCompareAtAmountCents gets a reference to the given int32 and assigns it to the CompareAtAmountCents field.
func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountCents(v int32) {
	o.CompareAtAmountCents = &v
}

// GetCompareAtAmountFloat returns the CompareAtAmountFloat field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloat() float32 {
	if o == nil || o.CompareAtAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.CompareAtAmountFloat
}

// GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetCompareAtAmountFloatOk() (*float32, bool) {
	if o == nil || o.CompareAtAmountFloat == nil {
		return nil, false
	}
	return o.CompareAtAmountFloat, true
}

// HasCompareAtAmountFloat returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasCompareAtAmountFloat() bool {
	if o != nil && o.CompareAtAmountFloat != nil {
		return true
	}

	return false
}

// SetCompareAtAmountFloat gets a reference to the given float32 and assigns it to the CompareAtAmountFloat field.
func (o *GETBundles200ResponseDataInnerAttributes) SetCompareAtAmountFloat(v float32) {
	o.CompareAtAmountFloat = &v
}

// GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmount() string {
	if o == nil || o.FormattedCompareAtAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedCompareAtAmount
}

// GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetFormattedCompareAtAmountOk() (*string, bool) {
	if o == nil || o.FormattedCompareAtAmount == nil {
		return nil, false
	}
	return o.FormattedCompareAtAmount, true
}

// HasFormattedCompareAtAmount returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasFormattedCompareAtAmount() bool {
	if o != nil && o.FormattedCompareAtAmount != nil {
		return true
	}

	return false
}

// SetFormattedCompareAtAmount gets a reference to the given string and assigns it to the FormattedCompareAtAmount field.
func (o *GETBundles200ResponseDataInnerAttributes) SetFormattedCompareAtAmount(v string) {
	o.FormattedCompareAtAmount = &v
}

// GetSkusCount returns the SkusCount field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCount() int32 {
	if o == nil || o.SkusCount == nil {
		var ret int32
		return ret
	}
	return *o.SkusCount
}

// GetSkusCountOk returns a tuple with the SkusCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetSkusCountOk() (*int32, bool) {
	if o == nil || o.SkusCount == nil {
		return nil, false
	}
	return o.SkusCount, true
}

// HasSkusCount returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasSkusCount() bool {
	if o != nil && o.SkusCount != nil {
		return true
	}

	return false
}

// SetSkusCount gets a reference to the given int32 and assigns it to the SkusCount field.
func (o *GETBundles200ResponseDataInnerAttributes) SetSkusCount(v int32) {
	o.SkusCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETBundles200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETBundles200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETBundles200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETBundles200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETBundles200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETBundles200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETBundles200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETBundles200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETBundles200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.DoNotShip != nil {
		toSerialize["do_not_ship"] = o.DoNotShip
	}
	if o.DoNotTrack != nil {
		toSerialize["do_not_track"] = o.DoNotTrack
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.CompareAtAmountCents != nil {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
	}
	if o.CompareAtAmountFloat != nil {
		toSerialize["compare_at_amount_float"] = o.CompareAtAmountFloat
	}
	if o.FormattedCompareAtAmount != nil {
		toSerialize["formatted_compare_at_amount"] = o.FormattedCompareAtAmount
	}
	if o.SkusCount != nil {
		toSerialize["skus_count"] = o.SkusCount
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETBundles200ResponseDataInnerAttributes struct {
	value *GETBundles200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETBundles200ResponseDataInnerAttributes) Get() *GETBundles200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETBundles200ResponseDataInnerAttributes) Set(val *GETBundles200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETBundles200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETBundles200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETBundles200ResponseDataInnerAttributes(val *GETBundles200ResponseDataInnerAttributes) *NullableGETBundles200ResponseDataInnerAttributes {
	return &NullableGETBundles200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETBundles200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETBundles200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
