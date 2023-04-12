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

// checks if the POSTParcels201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcels201ResponseDataAttributes{}

// POSTParcels201ResponseDataAttributes struct for POSTParcels201ResponseDataAttributes
type POSTParcels201ResponseDataAttributes struct {
	// The parcel weight, used to automatically calculate the tax rates from the available carrier accounts.
	Weight interface{} `json:"weight"`
	// Can be one of 'gr', 'lb', or 'oz'
	UnitOfWeight interface{} `json:"unit_of_weight"`
	// When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \"EEL\" o \"PFC\".
	EelPfc interface{} `json:"eel_pfc,omitempty"`
	// The type of item you are sending. Can be one of 'merchandise', 'gift', 'documents', 'returned_goods', 'sample', or 'other'.
	ContentsType interface{} `json:"contents_type,omitempty"`
	// If you specify 'other' in the 'contents_type' attribute, you must supply a brief description in this attribute.
	ContentsExplanation interface{} `json:"contents_explanation,omitempty"`
	// Indicates if the provided information is accurate
	CustomsCertify interface{} `json:"customs_certify,omitempty"`
	// This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this.
	CustomsSigner interface{} `json:"customs_signer,omitempty"`
	// In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either 'return', or 'abandon'. The value defaults to 'return'. If you pass 'abandon', you will not receive the parcel back if it cannot be delivered.
	NonDeliveryOption interface{} `json:"non_delivery_option,omitempty"`
	// Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of 'none', 'other', 'quarantine', or 'sanitary_phytosanitary_inspection'.
	RestrictionType interface{} `json:"restriction_type,omitempty"`
	// If you specify 'other' in the restriction type, you must supply a brief description of what is required.
	RestrictionComments interface{} `json:"restriction_comments,omitempty"`
	// Indicates if the parcel requires customs info to get the shipping rates.
	CustomsInfoRequired interface{} `json:"customs_info_required,omitempty"`
	// The shipping label url, ready to be downloaded and printed.
	ShippingLabelUrl interface{} `json:"shipping_label_url,omitempty"`
	// The shipping label file type. One of 'application/pdf', 'application/zpl', 'application/epl2', or 'image/png'.
	ShippingLabelFileType interface{} `json:"shipping_label_file_type,omitempty"`
	// The shipping label size.
	ShippingLabelSize interface{} `json:"shipping_label_size,omitempty"`
	// The shipping label resolution.
	ShippingLabelResolution interface{} `json:"shipping_label_resolution,omitempty"`
	// The tracking number associated to this parcel.
	TrackingNumber interface{} `json:"tracking_number,omitempty"`
	// The tracking status for this parcel, automatically updated in real time by the shipping carrier.
	TrackingStatus interface{} `json:"tracking_status,omitempty"`
	// Additional information about the tracking status, automatically updated in real time by the shipping carrier.
	TrackingStatusDetail interface{} `json:"tracking_status_detail,omitempty"`
	// Time at which the parcel's tracking status was last updated.
	TrackingStatusUpdatedAt interface{} `json:"tracking_status_updated_at,omitempty"`
	// The parcel's full tracking history, automatically updated in real time by the shipping carrier.
	TrackingDetails interface{} `json:"tracking_details,omitempty"`
	// The weight of the parcel as measured by the carrier in ounces (if available)
	CarrierWeightOz interface{} `json:"carrier_weight_oz,omitempty"`
	// The name of the person who signed for the parcel (if available)
	SignedBy interface{} `json:"signed_by,omitempty"`
	// The type of Incoterm (if available).
	Incoterm interface{} `json:"incoterm,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTParcels201ResponseDataAttributes instantiates a new POSTParcels201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcels201ResponseDataAttributes(weight interface{}, unitOfWeight interface{}) *POSTParcels201ResponseDataAttributes {
	this := POSTParcels201ResponseDataAttributes{}
	this.Weight = weight
	this.UnitOfWeight = unitOfWeight
	return &this
}

// NewPOSTParcels201ResponseDataAttributesWithDefaults instantiates a new POSTParcels201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcels201ResponseDataAttributesWithDefaults() *POSTParcels201ResponseDataAttributes {
	this := POSTParcels201ResponseDataAttributes{}
	return &this
}

// GetWeight returns the Weight field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTParcels201ResponseDataAttributes) GetWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *POSTParcels201ResponseDataAttributes) SetWeight(v interface{}) {
	o.Weight = v
}

// GetUnitOfWeight returns the UnitOfWeight field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTParcels201ResponseDataAttributes) GetUnitOfWeight() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.UnitOfWeight
}

// GetUnitOfWeightOk returns a tuple with the UnitOfWeight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetUnitOfWeightOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UnitOfWeight) {
		return nil, false
	}
	return &o.UnitOfWeight, true
}

// SetUnitOfWeight sets field value
func (o *POSTParcels201ResponseDataAttributes) SetUnitOfWeight(v interface{}) {
	o.UnitOfWeight = v
}

// GetEelPfc returns the EelPfc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetEelPfc() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EelPfc
}

// GetEelPfcOk returns a tuple with the EelPfc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetEelPfcOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EelPfc) {
		return nil, false
	}
	return &o.EelPfc, true
}

// HasEelPfc returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasEelPfc() bool {
	if o != nil && IsNil(o.EelPfc) {
		return true
	}

	return false
}

// SetEelPfc gets a reference to the given interface{} and assigns it to the EelPfc field.
func (o *POSTParcels201ResponseDataAttributes) SetEelPfc(v interface{}) {
	o.EelPfc = v
}

// GetContentsType returns the ContentsType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetContentsType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ContentsType
}

// GetContentsTypeOk returns a tuple with the ContentsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetContentsTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ContentsType) {
		return nil, false
	}
	return &o.ContentsType, true
}

// HasContentsType returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasContentsType() bool {
	if o != nil && IsNil(o.ContentsType) {
		return true
	}

	return false
}

// SetContentsType gets a reference to the given interface{} and assigns it to the ContentsType field.
func (o *POSTParcels201ResponseDataAttributes) SetContentsType(v interface{}) {
	o.ContentsType = v
}

// GetContentsExplanation returns the ContentsExplanation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetContentsExplanation() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ContentsExplanation
}

// GetContentsExplanationOk returns a tuple with the ContentsExplanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetContentsExplanationOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ContentsExplanation) {
		return nil, false
	}
	return &o.ContentsExplanation, true
}

// HasContentsExplanation returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasContentsExplanation() bool {
	if o != nil && IsNil(o.ContentsExplanation) {
		return true
	}

	return false
}

// SetContentsExplanation gets a reference to the given interface{} and assigns it to the ContentsExplanation field.
func (o *POSTParcels201ResponseDataAttributes) SetContentsExplanation(v interface{}) {
	o.ContentsExplanation = v
}

// GetCustomsCertify returns the CustomsCertify field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetCustomsCertify() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomsCertify
}

// GetCustomsCertifyOk returns a tuple with the CustomsCertify field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetCustomsCertifyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomsCertify) {
		return nil, false
	}
	return &o.CustomsCertify, true
}

// HasCustomsCertify returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasCustomsCertify() bool {
	if o != nil && IsNil(o.CustomsCertify) {
		return true
	}

	return false
}

// SetCustomsCertify gets a reference to the given interface{} and assigns it to the CustomsCertify field.
func (o *POSTParcels201ResponseDataAttributes) SetCustomsCertify(v interface{}) {
	o.CustomsCertify = v
}

// GetCustomsSigner returns the CustomsSigner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetCustomsSigner() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomsSigner
}

// GetCustomsSignerOk returns a tuple with the CustomsSigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetCustomsSignerOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomsSigner) {
		return nil, false
	}
	return &o.CustomsSigner, true
}

// HasCustomsSigner returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasCustomsSigner() bool {
	if o != nil && IsNil(o.CustomsSigner) {
		return true
	}

	return false
}

// SetCustomsSigner gets a reference to the given interface{} and assigns it to the CustomsSigner field.
func (o *POSTParcels201ResponseDataAttributes) SetCustomsSigner(v interface{}) {
	o.CustomsSigner = v
}

// GetNonDeliveryOption returns the NonDeliveryOption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetNonDeliveryOption() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NonDeliveryOption
}

// GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetNonDeliveryOptionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NonDeliveryOption) {
		return nil, false
	}
	return &o.NonDeliveryOption, true
}

// HasNonDeliveryOption returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasNonDeliveryOption() bool {
	if o != nil && IsNil(o.NonDeliveryOption) {
		return true
	}

	return false
}

// SetNonDeliveryOption gets a reference to the given interface{} and assigns it to the NonDeliveryOption field.
func (o *POSTParcels201ResponseDataAttributes) SetNonDeliveryOption(v interface{}) {
	o.NonDeliveryOption = v
}

// GetRestrictionType returns the RestrictionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetRestrictionType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RestrictionType
}

// GetRestrictionTypeOk returns a tuple with the RestrictionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetRestrictionTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RestrictionType) {
		return nil, false
	}
	return &o.RestrictionType, true
}

// HasRestrictionType returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasRestrictionType() bool {
	if o != nil && IsNil(o.RestrictionType) {
		return true
	}

	return false
}

// SetRestrictionType gets a reference to the given interface{} and assigns it to the RestrictionType field.
func (o *POSTParcels201ResponseDataAttributes) SetRestrictionType(v interface{}) {
	o.RestrictionType = v
}

// GetRestrictionComments returns the RestrictionComments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetRestrictionComments() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RestrictionComments
}

// GetRestrictionCommentsOk returns a tuple with the RestrictionComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetRestrictionCommentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RestrictionComments) {
		return nil, false
	}
	return &o.RestrictionComments, true
}

// HasRestrictionComments returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasRestrictionComments() bool {
	if o != nil && IsNil(o.RestrictionComments) {
		return true
	}

	return false
}

// SetRestrictionComments gets a reference to the given interface{} and assigns it to the RestrictionComments field.
func (o *POSTParcels201ResponseDataAttributes) SetRestrictionComments(v interface{}) {
	o.RestrictionComments = v
}

// GetCustomsInfoRequired returns the CustomsInfoRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetCustomsInfoRequired() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomsInfoRequired
}

// GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetCustomsInfoRequiredOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomsInfoRequired) {
		return nil, false
	}
	return &o.CustomsInfoRequired, true
}

// HasCustomsInfoRequired returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasCustomsInfoRequired() bool {
	if o != nil && IsNil(o.CustomsInfoRequired) {
		return true
	}

	return false
}

// SetCustomsInfoRequired gets a reference to the given interface{} and assigns it to the CustomsInfoRequired field.
func (o *POSTParcels201ResponseDataAttributes) SetCustomsInfoRequired(v interface{}) {
	o.CustomsInfoRequired = v
}

// GetShippingLabelUrl returns the ShippingLabelUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingLabelUrl
}

// GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingLabelUrl) {
		return nil, false
	}
	return &o.ShippingLabelUrl, true
}

// HasShippingLabelUrl returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasShippingLabelUrl() bool {
	if o != nil && IsNil(o.ShippingLabelUrl) {
		return true
	}

	return false
}

// SetShippingLabelUrl gets a reference to the given interface{} and assigns it to the ShippingLabelUrl field.
func (o *POSTParcels201ResponseDataAttributes) SetShippingLabelUrl(v interface{}) {
	o.ShippingLabelUrl = v
}

// GetShippingLabelFileType returns the ShippingLabelFileType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelFileType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingLabelFileType
}

// GetShippingLabelFileTypeOk returns a tuple with the ShippingLabelFileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelFileTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingLabelFileType) {
		return nil, false
	}
	return &o.ShippingLabelFileType, true
}

// HasShippingLabelFileType returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasShippingLabelFileType() bool {
	if o != nil && IsNil(o.ShippingLabelFileType) {
		return true
	}

	return false
}

// SetShippingLabelFileType gets a reference to the given interface{} and assigns it to the ShippingLabelFileType field.
func (o *POSTParcels201ResponseDataAttributes) SetShippingLabelFileType(v interface{}) {
	o.ShippingLabelFileType = v
}

// GetShippingLabelSize returns the ShippingLabelSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelSize() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingLabelSize
}

// GetShippingLabelSizeOk returns a tuple with the ShippingLabelSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelSizeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingLabelSize) {
		return nil, false
	}
	return &o.ShippingLabelSize, true
}

// HasShippingLabelSize returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasShippingLabelSize() bool {
	if o != nil && IsNil(o.ShippingLabelSize) {
		return true
	}

	return false
}

// SetShippingLabelSize gets a reference to the given interface{} and assigns it to the ShippingLabelSize field.
func (o *POSTParcels201ResponseDataAttributes) SetShippingLabelSize(v interface{}) {
	o.ShippingLabelSize = v
}

// GetShippingLabelResolution returns the ShippingLabelResolution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelResolution() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ShippingLabelResolution
}

// GetShippingLabelResolutionOk returns a tuple with the ShippingLabelResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetShippingLabelResolutionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ShippingLabelResolution) {
		return nil, false
	}
	return &o.ShippingLabelResolution, true
}

// HasShippingLabelResolution returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasShippingLabelResolution() bool {
	if o != nil && IsNil(o.ShippingLabelResolution) {
		return true
	}

	return false
}

// SetShippingLabelResolution gets a reference to the given interface{} and assigns it to the ShippingLabelResolution field.
func (o *POSTParcels201ResponseDataAttributes) SetShippingLabelResolution(v interface{}) {
	o.ShippingLabelResolution = v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetTrackingNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetTrackingNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return &o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasTrackingNumber() bool {
	if o != nil && IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given interface{} and assigns it to the TrackingNumber field.
func (o *POSTParcels201ResponseDataAttributes) SetTrackingNumber(v interface{}) {
	o.TrackingNumber = v
}

// GetTrackingStatus returns the TrackingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingStatus
}

// GetTrackingStatusOk returns a tuple with the TrackingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TrackingStatus) {
		return nil, false
	}
	return &o.TrackingStatus, true
}

// HasTrackingStatus returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatus() bool {
	if o != nil && IsNil(o.TrackingStatus) {
		return true
	}

	return false
}

// SetTrackingStatus gets a reference to the given interface{} and assigns it to the TrackingStatus field.
func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatus(v interface{}) {
	o.TrackingStatus = v
}

// GetTrackingStatusDetail returns the TrackingStatusDetail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingStatusDetail
}

// GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusDetailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TrackingStatusDetail) {
		return nil, false
	}
	return &o.TrackingStatusDetail, true
}

// HasTrackingStatusDetail returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatusDetail() bool {
	if o != nil && IsNil(o.TrackingStatusDetail) {
		return true
	}

	return false
}

// SetTrackingStatusDetail gets a reference to the given interface{} and assigns it to the TrackingStatusDetail field.
func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatusDetail(v interface{}) {
	o.TrackingStatusDetail = v
}

// GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingStatusUpdatedAt
}

// GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetTrackingStatusUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TrackingStatusUpdatedAt) {
		return nil, false
	}
	return &o.TrackingStatusUpdatedAt, true
}

// HasTrackingStatusUpdatedAt returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasTrackingStatusUpdatedAt() bool {
	if o != nil && IsNil(o.TrackingStatusUpdatedAt) {
		return true
	}

	return false
}

// SetTrackingStatusUpdatedAt gets a reference to the given interface{} and assigns it to the TrackingStatusUpdatedAt field.
func (o *POSTParcels201ResponseDataAttributes) SetTrackingStatusUpdatedAt(v interface{}) {
	o.TrackingStatusUpdatedAt = v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetTrackingDetails() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetTrackingDetailsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TrackingDetails) {
		return nil, false
	}
	return &o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasTrackingDetails() bool {
	if o != nil && IsNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given interface{} and assigns it to the TrackingDetails field.
func (o *POSTParcels201ResponseDataAttributes) SetTrackingDetails(v interface{}) {
	o.TrackingDetails = v
}

// GetCarrierWeightOz returns the CarrierWeightOz field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetCarrierWeightOz() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CarrierWeightOz
}

// GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetCarrierWeightOzOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CarrierWeightOz) {
		return nil, false
	}
	return &o.CarrierWeightOz, true
}

// HasCarrierWeightOz returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasCarrierWeightOz() bool {
	if o != nil && IsNil(o.CarrierWeightOz) {
		return true
	}

	return false
}

// SetCarrierWeightOz gets a reference to the given interface{} and assigns it to the CarrierWeightOz field.
func (o *POSTParcels201ResponseDataAttributes) SetCarrierWeightOz(v interface{}) {
	o.CarrierWeightOz = v
}

// GetSignedBy returns the SignedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetSignedBy() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SignedBy
}

// GetSignedByOk returns a tuple with the SignedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetSignedByOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SignedBy) {
		return nil, false
	}
	return &o.SignedBy, true
}

// HasSignedBy returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasSignedBy() bool {
	if o != nil && IsNil(o.SignedBy) {
		return true
	}

	return false
}

// SetSignedBy gets a reference to the given interface{} and assigns it to the SignedBy field.
func (o *POSTParcels201ResponseDataAttributes) SetSignedBy(v interface{}) {
	o.SignedBy = v
}

// GetIncoterm returns the Incoterm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetIncoterm() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Incoterm
}

// GetIncotermOk returns a tuple with the Incoterm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetIncotermOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Incoterm) {
		return nil, false
	}
	return &o.Incoterm, true
}

// HasIncoterm returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasIncoterm() bool {
	if o != nil && IsNil(o.Incoterm) {
		return true
	}

	return false
}

// SetIncoterm gets a reference to the given interface{} and assigns it to the Incoterm field.
func (o *POSTParcels201ResponseDataAttributes) SetIncoterm(v interface{}) {
	o.Incoterm = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTParcels201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTParcels201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcels201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcels201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTParcels201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTParcels201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTParcels201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcels201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePOSTParcels201ResponseDataAttributes struct {
	value *POSTParcels201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTParcels201ResponseDataAttributes) Get() *POSTParcels201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTParcels201ResponseDataAttributes) Set(val *POSTParcels201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcels201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcels201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcels201ResponseDataAttributes(val *POSTParcels201ResponseDataAttributes) *NullablePOSTParcels201ResponseDataAttributes {
	return &NullablePOSTParcels201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTParcels201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcels201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
