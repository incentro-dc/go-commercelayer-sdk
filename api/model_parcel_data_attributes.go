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

// ParcelDataAttributes struct for ParcelDataAttributes
type ParcelDataAttributes struct {
	// Unique identifier for the parcel
	Number *string `json:"number,omitempty"`
	// The parcel weight, used to automatically calculate the tax rates from the available carrier accounts.
	Weight *float32 `json:"weight,omitempty"`
	// The unit of weight. Can be one of 'gr', or 'oz'.
	UnitOfWeight *string `json:"unit_of_weight,omitempty"`
	// When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \"EEL\" o \"PFC\".
	EelPfc *string `json:"eel_pfc,omitempty"`
	// The type of item you are sending. Can be one of 'merchandise', 'gift', 'documents', 'returned_goods', 'sample', or 'other'.
	ContentsType *string `json:"contents_type,omitempty"`
	// If you specify 'other' in the 'contents_type' attribute, you must supply a brief description in this attribute.
	ContentsExplanation *string `json:"contents_explanation,omitempty"`
	// Indicates if the provided information is accurate
	CustomsCertify *bool `json:"customs_certify,omitempty"`
	// This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this.
	CustomsSigner *string `json:"customs_signer,omitempty"`
	// In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either 'return', or 'abandon'. The value defaults to 'return'. If you pass 'abandon', you will not receive the parcel back if it cannot be delivered.
	NonDeliveryOption *string `json:"non_delivery_option,omitempty"`
	// Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of 'none', 'other', 'quarantine', or 'sanitary_phytosanitary_inspection'.
	RestrictionType *string `json:"restriction_type,omitempty"`
	// If you specify 'other' in the restriction type, you must supply a brief description of what is required.
	RestrictionComments *string `json:"restriction_comments,omitempty"`
	// Indicates if the parcel requires customs info to get the shipping rates.
	CustomsInfoRequired *bool `json:"customs_info_required,omitempty"`
	// The shipping label url, ready to be downloaded and printed.
	ShippingLabelUrl *string `json:"shipping_label_url,omitempty"`
	// The shipping label file type. One of 'application/pdf', 'application/zpl', 'application/epl2', or 'image/png'.
	ShippingLabelFileType *string `json:"shipping_label_file_type,omitempty"`
	// The shipping label size.
	ShippingLabelSize *string `json:"shipping_label_size,omitempty"`
	// The shipping label resolution.
	ShippingLabelResolution *string `json:"shipping_label_resolution,omitempty"`
	// The tracking number associated to this parcel.
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// The tracking status for this parcel, automatically updated in real time by the shipping carrier.
	TrackingStatus *string `json:"tracking_status,omitempty"`
	// Additional information about the tracking status, automatically updated in real time by the shipping carrier.
	TrackingStatusDetail *string `json:"tracking_status_detail,omitempty"`
	// Time at which the parcel's tracking status was last updated.
	TrackingStatusUpdatedAt *string `json:"tracking_status_updated_at,omitempty"`
	// The parcel's full tracking history, automatically updated in real time by the shipping carrier.
	TrackingDetails *string `json:"tracking_details,omitempty"`
	// The weight of the parcel as measured by the carrier in ounces (if available)
	CarrierWeightOz *string `json:"carrier_weight_oz,omitempty"`
	// The name of the person who signed for the parcel (if available)
	SignedBy *string `json:"signed_by,omitempty"`
	// The type of Incoterm (if available).
	Incoterm *string `json:"incoterm,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
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

// NewParcelDataAttributes instantiates a new ParcelDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelDataAttributes() *ParcelDataAttributes {
	this := ParcelDataAttributes{}
	return &this
}

// NewParcelDataAttributesWithDefaults instantiates a new ParcelDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelDataAttributesWithDefaults() *ParcelDataAttributes {
	this := ParcelDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ParcelDataAttributes) SetNumber(v string) {
	o.Number = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *ParcelDataAttributes) SetWeight(v float32) {
	o.Weight = &v
}

// GetUnitOfWeight returns the UnitOfWeight field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetUnitOfWeight() string {
	if o == nil || o.UnitOfWeight == nil {
		var ret string
		return ret
	}
	return *o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetUnitOfWeightOk() (*string, bool) {
	if o == nil || o.UnitOfWeight == nil {
		return nil, false
	}
	return o.UnitOfWeight, true
}

// HasUnitOfWeight returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasUnitOfWeight() bool {
	if o != nil && o.UnitOfWeight != nil {
		return true
	}

	return false
}

// SetUnitOfWeight gets a reference to the given string and assigns it to the UnitOfWeight field.
func (o *ParcelDataAttributes) SetUnitOfWeight(v string) {
	o.UnitOfWeight = &v
}

// GetEelPfc returns the EelPfc field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetEelPfc() string {
	if o == nil || o.EelPfc == nil {
		var ret string
		return ret
	}
	return *o.EelPfc
}

// GetEelPfcOk returns a tuple with the EelPfc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetEelPfcOk() (*string, bool) {
	if o == nil || o.EelPfc == nil {
		return nil, false
	}
	return o.EelPfc, true
}

// HasEelPfc returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasEelPfc() bool {
	if o != nil && o.EelPfc != nil {
		return true
	}

	return false
}

// SetEelPfc gets a reference to the given string and assigns it to the EelPfc field.
func (o *ParcelDataAttributes) SetEelPfc(v string) {
	o.EelPfc = &v
}

// GetContentsType returns the ContentsType field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetContentsType() string {
	if o == nil || o.ContentsType == nil {
		var ret string
		return ret
	}
	return *o.ContentsType
}

// GetContentsTypeOk returns a tuple with the ContentsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetContentsTypeOk() (*string, bool) {
	if o == nil || o.ContentsType == nil {
		return nil, false
	}
	return o.ContentsType, true
}

// HasContentsType returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasContentsType() bool {
	if o != nil && o.ContentsType != nil {
		return true
	}

	return false
}

// SetContentsType gets a reference to the given string and assigns it to the ContentsType field.
func (o *ParcelDataAttributes) SetContentsType(v string) {
	o.ContentsType = &v
}

// GetContentsExplanation returns the ContentsExplanation field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetContentsExplanation() string {
	if o == nil || o.ContentsExplanation == nil {
		var ret string
		return ret
	}
	return *o.ContentsExplanation
}

// GetContentsExplanationOk returns a tuple with the ContentsExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetContentsExplanationOk() (*string, bool) {
	if o == nil || o.ContentsExplanation == nil {
		return nil, false
	}
	return o.ContentsExplanation, true
}

// HasContentsExplanation returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasContentsExplanation() bool {
	if o != nil && o.ContentsExplanation != nil {
		return true
	}

	return false
}

// SetContentsExplanation gets a reference to the given string and assigns it to the ContentsExplanation field.
func (o *ParcelDataAttributes) SetContentsExplanation(v string) {
	o.ContentsExplanation = &v
}

// GetCustomsCertify returns the CustomsCertify field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetCustomsCertify() bool {
	if o == nil || o.CustomsCertify == nil {
		var ret bool
		return ret
	}
	return *o.CustomsCertify
}

// GetCustomsCertifyOk returns a tuple with the CustomsCertify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetCustomsCertifyOk() (*bool, bool) {
	if o == nil || o.CustomsCertify == nil {
		return nil, false
	}
	return o.CustomsCertify, true
}

// HasCustomsCertify returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasCustomsCertify() bool {
	if o != nil && o.CustomsCertify != nil {
		return true
	}

	return false
}

// SetCustomsCertify gets a reference to the given bool and assigns it to the CustomsCertify field.
func (o *ParcelDataAttributes) SetCustomsCertify(v bool) {
	o.CustomsCertify = &v
}

// GetCustomsSigner returns the CustomsSigner field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetCustomsSigner() string {
	if o == nil || o.CustomsSigner == nil {
		var ret string
		return ret
	}
	return *o.CustomsSigner
}

// GetCustomsSignerOk returns a tuple with the CustomsSigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetCustomsSignerOk() (*string, bool) {
	if o == nil || o.CustomsSigner == nil {
		return nil, false
	}
	return o.CustomsSigner, true
}

// HasCustomsSigner returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasCustomsSigner() bool {
	if o != nil && o.CustomsSigner != nil {
		return true
	}

	return false
}

// SetCustomsSigner gets a reference to the given string and assigns it to the CustomsSigner field.
func (o *ParcelDataAttributes) SetCustomsSigner(v string) {
	o.CustomsSigner = &v
}

// GetNonDeliveryOption returns the NonDeliveryOption field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetNonDeliveryOption() string {
	if o == nil || o.NonDeliveryOption == nil {
		var ret string
		return ret
	}
	return *o.NonDeliveryOption
}

// GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetNonDeliveryOptionOk() (*string, bool) {
	if o == nil || o.NonDeliveryOption == nil {
		return nil, false
	}
	return o.NonDeliveryOption, true
}

// HasNonDeliveryOption returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasNonDeliveryOption() bool {
	if o != nil && o.NonDeliveryOption != nil {
		return true
	}

	return false
}

// SetNonDeliveryOption gets a reference to the given string and assigns it to the NonDeliveryOption field.
func (o *ParcelDataAttributes) SetNonDeliveryOption(v string) {
	o.NonDeliveryOption = &v
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetRestrictionType() string {
	if o == nil || o.RestrictionType == nil {
		var ret string
		return ret
	}
	return *o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetRestrictionTypeOk() (*string, bool) {
	if o == nil || o.RestrictionType == nil {
		return nil, false
	}
	return o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasRestrictionType() bool {
	if o != nil && o.RestrictionType != nil {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given string and assigns it to the RestrictionType field.
func (o *ParcelDataAttributes) SetRestrictionType(v string) {
	o.RestrictionType = &v
}

// GetRestrictionComments returns the RestrictionComments field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetRestrictionComments() string {
	if o == nil || o.RestrictionComments == nil {
		var ret string
		return ret
	}
	return *o.RestrictionComments
}

// GetRestrictionCommentsOk returns a tuple with the RestrictionComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetRestrictionCommentsOk() (*string, bool) {
	if o == nil || o.RestrictionComments == nil {
		return nil, false
	}
	return o.RestrictionComments, true
}

// HasRestrictionComments returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasRestrictionComments() bool {
	if o != nil && o.RestrictionComments != nil {
		return true
	}

	return false
}

// SetRestrictionComments gets a reference to the given string and assigns it to the RestrictionComments field.
func (o *ParcelDataAttributes) SetRestrictionComments(v string) {
	o.RestrictionComments = &v
}

// GetCustomsInfoRequired returns the CustomsInfoRequired field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetCustomsInfoRequired() bool {
	if o == nil || o.CustomsInfoRequired == nil {
		var ret bool
		return ret
	}
	return *o.CustomsInfoRequired
}

// GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetCustomsInfoRequiredOk() (*bool, bool) {
	if o == nil || o.CustomsInfoRequired == nil {
		return nil, false
	}
	return o.CustomsInfoRequired, true
}

// HasCustomsInfoRequired returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasCustomsInfoRequired() bool {
	if o != nil && o.CustomsInfoRequired != nil {
		return true
	}

	return false
}

// SetCustomsInfoRequired gets a reference to the given bool and assigns it to the CustomsInfoRequired field.
func (o *ParcelDataAttributes) SetCustomsInfoRequired(v bool) {
	o.CustomsInfoRequired = &v
}

// GetShippingLabelUrl returns the ShippingLabelUrl field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetShippingLabelUrl() string {
	if o == nil || o.ShippingLabelUrl == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelUrl
}

// GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetShippingLabelUrlOk() (*string, bool) {
	if o == nil || o.ShippingLabelUrl == nil {
		return nil, false
	}
	return o.ShippingLabelUrl, true
}

// HasShippingLabelUrl returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasShippingLabelUrl() bool {
	if o != nil && o.ShippingLabelUrl != nil {
		return true
	}

	return false
}

// SetShippingLabelUrl gets a reference to the given string and assigns it to the ShippingLabelUrl field.
func (o *ParcelDataAttributes) SetShippingLabelUrl(v string) {
	o.ShippingLabelUrl = &v
}

// GetShippingLabelFileType returns the ShippingLabelFileType field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetShippingLabelFileType() string {
	if o == nil || o.ShippingLabelFileType == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelFileType
}

// GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetShippingLabelFileTypeOk() (*string, bool) {
	if o == nil || o.ShippingLabelFileType == nil {
		return nil, false
	}
	return o.ShippingLabelFileType, true
}

// HasShippingLabelFileType returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasShippingLabelFileType() bool {
	if o != nil && o.ShippingLabelFileType != nil {
		return true
	}

	return false
}

// SetShippingLabelFileType gets a reference to the given string and assigns it to the ShippingLabelFileType field.
func (o *ParcelDataAttributes) SetShippingLabelFileType(v string) {
	o.ShippingLabelFileType = &v
}

// GetShippingLabelSize returns the ShippingLabelSize field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetShippingLabelSize() string {
	if o == nil || o.ShippingLabelSize == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelSize
}

// GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetShippingLabelSizeOk() (*string, bool) {
	if o == nil || o.ShippingLabelSize == nil {
		return nil, false
	}
	return o.ShippingLabelSize, true
}

// HasShippingLabelSize returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasShippingLabelSize() bool {
	if o != nil && o.ShippingLabelSize != nil {
		return true
	}

	return false
}

// SetShippingLabelSize gets a reference to the given string and assigns it to the ShippingLabelSize field.
func (o *ParcelDataAttributes) SetShippingLabelSize(v string) {
	o.ShippingLabelSize = &v
}

// GetShippingLabelResolution returns the ShippingLabelResolution field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetShippingLabelResolution() string {
	if o == nil || o.ShippingLabelResolution == nil {
		var ret string
		return ret
	}
	return *o.ShippingLabelResolution
}

// GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetShippingLabelResolutionOk() (*string, bool) {
	if o == nil || o.ShippingLabelResolution == nil {
		return nil, false
	}
	return o.ShippingLabelResolution, true
}

// HasShippingLabelResolution returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasShippingLabelResolution() bool {
	if o != nil && o.ShippingLabelResolution != nil {
		return true
	}

	return false
}

// SetShippingLabelResolution gets a reference to the given string and assigns it to the ShippingLabelResolution field.
func (o *ParcelDataAttributes) SetShippingLabelResolution(v string) {
	o.ShippingLabelResolution = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetTrackingNumber() string {
	if o == nil || o.TrackingNumber == nil {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetTrackingNumberOk() (*string, bool) {
	if o == nil || o.TrackingNumber == nil {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasTrackingNumber() bool {
	if o != nil && o.TrackingNumber != nil {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *ParcelDataAttributes) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetTrackingStatus returns the TrackingStatus field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetTrackingStatus() string {
	if o == nil || o.TrackingStatus == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatus
}

// GetTrackingStatusOk returns a tuple with the TrackingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetTrackingStatusOk() (*string, bool) {
	if o == nil || o.TrackingStatus == nil {
		return nil, false
	}
	return o.TrackingStatus, true
}

// HasTrackingStatus returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasTrackingStatus() bool {
	if o != nil && o.TrackingStatus != nil {
		return true
	}

	return false
}

// SetTrackingStatus gets a reference to the given string and assigns it to the TrackingStatus field.
func (o *ParcelDataAttributes) SetTrackingStatus(v string) {
	o.TrackingStatus = &v
}

// GetTrackingStatusDetail returns the TrackingStatusDetail field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetTrackingStatusDetail() string {
	if o == nil || o.TrackingStatusDetail == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatusDetail
}

// GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetTrackingStatusDetailOk() (*string, bool) {
	if o == nil || o.TrackingStatusDetail == nil {
		return nil, false
	}
	return o.TrackingStatusDetail, true
}

// HasTrackingStatusDetail returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasTrackingStatusDetail() bool {
	if o != nil && o.TrackingStatusDetail != nil {
		return true
	}

	return false
}

// SetTrackingStatusDetail gets a reference to the given string and assigns it to the TrackingStatusDetail field.
func (o *ParcelDataAttributes) SetTrackingStatusDetail(v string) {
	o.TrackingStatusDetail = &v
}

// GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetTrackingStatusUpdatedAt() string {
	if o == nil || o.TrackingStatusUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.TrackingStatusUpdatedAt
}

// GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetTrackingStatusUpdatedAtOk() (*string, bool) {
	if o == nil || o.TrackingStatusUpdatedAt == nil {
		return nil, false
	}
	return o.TrackingStatusUpdatedAt, true
}

// HasTrackingStatusUpdatedAt returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasTrackingStatusUpdatedAt() bool {
	if o != nil && o.TrackingStatusUpdatedAt != nil {
		return true
	}

	return false
}

// SetTrackingStatusUpdatedAt gets a reference to the given string and assigns it to the TrackingStatusUpdatedAt field.
func (o *ParcelDataAttributes) SetTrackingStatusUpdatedAt(v string) {
	o.TrackingStatusUpdatedAt = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetTrackingDetails() string {
	if o == nil || o.TrackingDetails == nil {
		var ret string
		return ret
	}
	return *o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetTrackingDetailsOk() (*string, bool) {
	if o == nil || o.TrackingDetails == nil {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasTrackingDetails() bool {
	if o != nil && o.TrackingDetails != nil {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given string and assigns it to the TrackingDetails field.
func (o *ParcelDataAttributes) SetTrackingDetails(v string) {
	o.TrackingDetails = &v
}

// GetCarrierWeightOz returns the CarrierWeightOz field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetCarrierWeightOz() string {
	if o == nil || o.CarrierWeightOz == nil {
		var ret string
		return ret
	}
	return *o.CarrierWeightOz
}

// GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetCarrierWeightOzOk() (*string, bool) {
	if o == nil || o.CarrierWeightOz == nil {
		return nil, false
	}
	return o.CarrierWeightOz, true
}

// HasCarrierWeightOz returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasCarrierWeightOz() bool {
	if o != nil && o.CarrierWeightOz != nil {
		return true
	}

	return false
}

// SetCarrierWeightOz gets a reference to the given string and assigns it to the CarrierWeightOz field.
func (o *ParcelDataAttributes) SetCarrierWeightOz(v string) {
	o.CarrierWeightOz = &v
}

// GetSignedBy returns the SignedBy field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetSignedBy() string {
	if o == nil || o.SignedBy == nil {
		var ret string
		return ret
	}
	return *o.SignedBy
}

// GetSignedByOk returns a tuple with the SignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetSignedByOk() (*string, bool) {
	if o == nil || o.SignedBy == nil {
		return nil, false
	}
	return o.SignedBy, true
}

// HasSignedBy returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasSignedBy() bool {
	if o != nil && o.SignedBy != nil {
		return true
	}

	return false
}

// SetSignedBy gets a reference to the given string and assigns it to the SignedBy field.
func (o *ParcelDataAttributes) SetSignedBy(v string) {
	o.SignedBy = &v
}

// GetIncoterm returns the Incoterm field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetIncoterm() string {
	if o == nil || o.Incoterm == nil {
		var ret string
		return ret
	}
	return *o.Incoterm
}

// GetIncotermOk returns a tuple with the Incoterm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetIncotermOk() (*string, bool) {
	if o == nil || o.Incoterm == nil {
		return nil, false
	}
	return o.Incoterm, true
}

// HasIncoterm returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasIncoterm() bool {
	if o != nil && o.Incoterm != nil {
		return true
	}

	return false
}

// SetIncoterm gets a reference to the given string and assigns it to the Incoterm field.
func (o *ParcelDataAttributes) SetIncoterm(v string) {
	o.Incoterm = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ParcelDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ParcelDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ParcelDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ParcelDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ParcelDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ParcelDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ParcelDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ParcelDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ParcelDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	if o.UnitOfWeight != nil {
		toSerialize["unit_of_weight"] = o.UnitOfWeight
	}
	if o.EelPfc != nil {
		toSerialize["eel_pfc"] = o.EelPfc
	}
	if o.ContentsType != nil {
		toSerialize["contents_type"] = o.ContentsType
	}
	if o.ContentsExplanation != nil {
		toSerialize["contents_explanation"] = o.ContentsExplanation
	}
	if o.CustomsCertify != nil {
		toSerialize["customs_certify"] = o.CustomsCertify
	}
	if o.CustomsSigner != nil {
		toSerialize["customs_signer"] = o.CustomsSigner
	}
	if o.NonDeliveryOption != nil {
		toSerialize["non_delivery_option"] = o.NonDeliveryOption
	}
	if o.RestrictionType != nil {
		toSerialize["restriction_type"] = o.RestrictionType
	}
	if o.RestrictionComments != nil {
		toSerialize["restriction_comments"] = o.RestrictionComments
	}
	if o.CustomsInfoRequired != nil {
		toSerialize["customs_info_required"] = o.CustomsInfoRequired
	}
	if o.ShippingLabelUrl != nil {
		toSerialize["shipping_label_url"] = o.ShippingLabelUrl
	}
	if o.ShippingLabelFileType != nil {
		toSerialize["shipping_label_file_type"] = o.ShippingLabelFileType
	}
	if o.ShippingLabelSize != nil {
		toSerialize["shipping_label_size"] = o.ShippingLabelSize
	}
	if o.ShippingLabelResolution != nil {
		toSerialize["shipping_label_resolution"] = o.ShippingLabelResolution
	}
	if o.TrackingNumber != nil {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if o.TrackingStatus != nil {
		toSerialize["tracking_status"] = o.TrackingStatus
	}
	if o.TrackingStatusDetail != nil {
		toSerialize["tracking_status_detail"] = o.TrackingStatusDetail
	}
	if o.TrackingStatusUpdatedAt != nil {
		toSerialize["tracking_status_updated_at"] = o.TrackingStatusUpdatedAt
	}
	if o.TrackingDetails != nil {
		toSerialize["tracking_details"] = o.TrackingDetails
	}
	if o.CarrierWeightOz != nil {
		toSerialize["carrier_weight_oz"] = o.CarrierWeightOz
	}
	if o.SignedBy != nil {
		toSerialize["signed_by"] = o.SignedBy
	}
	if o.Incoterm != nil {
		toSerialize["incoterm"] = o.Incoterm
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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

type NullableParcelDataAttributes struct {
	value *ParcelDataAttributes
	isSet bool
}

func (v NullableParcelDataAttributes) Get() *ParcelDataAttributes {
	return v.value
}

func (v *NullableParcelDataAttributes) Set(val *ParcelDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelDataAttributes(val *ParcelDataAttributes) *NullableParcelDataAttributes {
	return &NullableParcelDataAttributes{value: val, isSet: true}
}

func (v NullableParcelDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
