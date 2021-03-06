# ParcelCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **float32** | The parcel weight, used to automatically calculate the tax rates from the available carrier accounts. | 
**UnitOfWeight** | **string** | Can be one of &#39;gr&#39;, &#39;lb&#39;, or &#39;oz&#39; | 
**EelPfc** | Pointer to **string** | When shipping outside the US, you need to provide either an Exemption and Exclusion Legend (EEL) code or a Proof of Filing Citation (PFC). Which you need is based on the value of the goods being shipped. Value can be one of \&quot;EEL\&quot; o \&quot;PFC\&quot;. | [optional] 
**ContentsType** | Pointer to **string** | The type of item you are sending. Can be one of &#39;merchandise&#39;, &#39;gift&#39;, &#39;documents&#39;, &#39;returned_goods&#39;, &#39;sample&#39;, or &#39;other&#39;. | [optional] 
**ContentsExplanation** | Pointer to **string** | If you specify &#39;other&#39; in the &#39;contents_type&#39; attribute, you must supply a brief description in this attribute. | [optional] 
**CustomsCertify** | Pointer to **bool** | Indicates if the provided information is accurate | [optional] 
**CustomsSigner** | Pointer to **string** | This is the name of the person who is certifying that the information provided on the customs form is accurate. Use a name of the person in your organization who is responsible for this. | [optional] 
**NonDeliveryOption** | Pointer to **string** | In case the shipment cannot be delivered, this option tells the carrier what you want to happen to the parcel. You can pass either &#39;return&#39;, or &#39;abandon&#39;. The value defaults to &#39;return&#39;. If you pass &#39;abandon&#39;, you will not receive the parcel back if it cannot be delivered. | [optional] 
**RestrictionType** | Pointer to **string** | Describes if your parcel requires any special treatment or quarantine when entering the country. Can be one of &#39;none&#39;, &#39;other&#39;, &#39;quarantine&#39;, or &#39;sanitary_phytosanitary_inspection&#39;. | [optional] 
**RestrictionComments** | Pointer to **string** | If you specify &#39;other&#39; in the restriction type, you must supply a brief description of what is required. | [optional] 
**CustomsInfoRequired** | Pointer to **bool** | Indicates if the parcel requires customs info to get the shipping rates. | [optional] 
**TrackingNumber** | Pointer to **string** | The tracking number associated to this parcel. | [optional] 
**TrackingStatus** | Pointer to **string** | The tracking status for this parcel, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusDetail** | Pointer to **string** | Additional information about the tracking status, automatically updated in real time by the shipping carrier. | [optional] 
**TrackingStatusUpdatedAt** | Pointer to **string** | Time at which the parcel&#39;s tracking status was last updated. | [optional] 
**TrackingDetails** | Pointer to **string** | The parcel&#39;s full tracking history, automatically updated in real time by the shipping carrier. | [optional] 
**CarrierWeightOz** | Pointer to **string** | The weight of the parcel as measured by the carrier in ounces (if available) | [optional] 
**SignedBy** | Pointer to **string** | The name of the person who signed for the parcel (if available) | [optional] 
**Incoterm** | Pointer to **string** | The type of Incoterm (if available). | [optional] 
**Reference** | Pointer to **string** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **string** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewParcelCreateDataAttributes

`func NewParcelCreateDataAttributes(weight float32, unitOfWeight string, ) *ParcelCreateDataAttributes`

NewParcelCreateDataAttributes instantiates a new ParcelCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelCreateDataAttributesWithDefaults

`func NewParcelCreateDataAttributesWithDefaults() *ParcelCreateDataAttributes`

NewParcelCreateDataAttributesWithDefaults instantiates a new ParcelCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *ParcelCreateDataAttributes) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ParcelCreateDataAttributes) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ParcelCreateDataAttributes) SetWeight(v float32)`

SetWeight sets Weight field to given value.


### GetUnitOfWeight

`func (o *ParcelCreateDataAttributes) GetUnitOfWeight() string`

GetUnitOfWeight returns the UnitOfWeight field if non-nil, zero value otherwise.

### GetUnitOfWeightOk

`func (o *ParcelCreateDataAttributes) GetUnitOfWeightOk() (*string, bool)`

GetUnitOfWeightOk returns a tuple with the UnitOfWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfWeight

`func (o *ParcelCreateDataAttributes) SetUnitOfWeight(v string)`

SetUnitOfWeight sets UnitOfWeight field to given value.


### GetEelPfc

`func (o *ParcelCreateDataAttributes) GetEelPfc() string`

GetEelPfc returns the EelPfc field if non-nil, zero value otherwise.

### GetEelPfcOk

`func (o *ParcelCreateDataAttributes) GetEelPfcOk() (*string, bool)`

GetEelPfcOk returns a tuple with the EelPfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEelPfc

`func (o *ParcelCreateDataAttributes) SetEelPfc(v string)`

SetEelPfc sets EelPfc field to given value.

### HasEelPfc

`func (o *ParcelCreateDataAttributes) HasEelPfc() bool`

HasEelPfc returns a boolean if a field has been set.

### GetContentsType

`func (o *ParcelCreateDataAttributes) GetContentsType() string`

GetContentsType returns the ContentsType field if non-nil, zero value otherwise.

### GetContentsTypeOk

`func (o *ParcelCreateDataAttributes) GetContentsTypeOk() (*string, bool)`

GetContentsTypeOk returns a tuple with the ContentsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsType

`func (o *ParcelCreateDataAttributes) SetContentsType(v string)`

SetContentsType sets ContentsType field to given value.

### HasContentsType

`func (o *ParcelCreateDataAttributes) HasContentsType() bool`

HasContentsType returns a boolean if a field has been set.

### GetContentsExplanation

`func (o *ParcelCreateDataAttributes) GetContentsExplanation() string`

GetContentsExplanation returns the ContentsExplanation field if non-nil, zero value otherwise.

### GetContentsExplanationOk

`func (o *ParcelCreateDataAttributes) GetContentsExplanationOk() (*string, bool)`

GetContentsExplanationOk returns a tuple with the ContentsExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentsExplanation

`func (o *ParcelCreateDataAttributes) SetContentsExplanation(v string)`

SetContentsExplanation sets ContentsExplanation field to given value.

### HasContentsExplanation

`func (o *ParcelCreateDataAttributes) HasContentsExplanation() bool`

HasContentsExplanation returns a boolean if a field has been set.

### GetCustomsCertify

`func (o *ParcelCreateDataAttributes) GetCustomsCertify() bool`

GetCustomsCertify returns the CustomsCertify field if non-nil, zero value otherwise.

### GetCustomsCertifyOk

`func (o *ParcelCreateDataAttributes) GetCustomsCertifyOk() (*bool, bool)`

GetCustomsCertifyOk returns a tuple with the CustomsCertify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsCertify

`func (o *ParcelCreateDataAttributes) SetCustomsCertify(v bool)`

SetCustomsCertify sets CustomsCertify field to given value.

### HasCustomsCertify

`func (o *ParcelCreateDataAttributes) HasCustomsCertify() bool`

HasCustomsCertify returns a boolean if a field has been set.

### GetCustomsSigner

`func (o *ParcelCreateDataAttributes) GetCustomsSigner() string`

GetCustomsSigner returns the CustomsSigner field if non-nil, zero value otherwise.

### GetCustomsSignerOk

`func (o *ParcelCreateDataAttributes) GetCustomsSignerOk() (*string, bool)`

GetCustomsSignerOk returns a tuple with the CustomsSigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsSigner

`func (o *ParcelCreateDataAttributes) SetCustomsSigner(v string)`

SetCustomsSigner sets CustomsSigner field to given value.

### HasCustomsSigner

`func (o *ParcelCreateDataAttributes) HasCustomsSigner() bool`

HasCustomsSigner returns a boolean if a field has been set.

### GetNonDeliveryOption

`func (o *ParcelCreateDataAttributes) GetNonDeliveryOption() string`

GetNonDeliveryOption returns the NonDeliveryOption field if non-nil, zero value otherwise.

### GetNonDeliveryOptionOk

`func (o *ParcelCreateDataAttributes) GetNonDeliveryOptionOk() (*string, bool)`

GetNonDeliveryOptionOk returns a tuple with the NonDeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeliveryOption

`func (o *ParcelCreateDataAttributes) SetNonDeliveryOption(v string)`

SetNonDeliveryOption sets NonDeliveryOption field to given value.

### HasNonDeliveryOption

`func (o *ParcelCreateDataAttributes) HasNonDeliveryOption() bool`

HasNonDeliveryOption returns a boolean if a field has been set.

### GetRestrictionType

`func (o *ParcelCreateDataAttributes) GetRestrictionType() string`

GetRestrictionType returns the RestrictionType field if non-nil, zero value otherwise.

### GetRestrictionTypeOk

`func (o *ParcelCreateDataAttributes) GetRestrictionTypeOk() (*string, bool)`

GetRestrictionTypeOk returns a tuple with the RestrictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionType

`func (o *ParcelCreateDataAttributes) SetRestrictionType(v string)`

SetRestrictionType sets RestrictionType field to given value.

### HasRestrictionType

`func (o *ParcelCreateDataAttributes) HasRestrictionType() bool`

HasRestrictionType returns a boolean if a field has been set.

### GetRestrictionComments

`func (o *ParcelCreateDataAttributes) GetRestrictionComments() string`

GetRestrictionComments returns the RestrictionComments field if non-nil, zero value otherwise.

### GetRestrictionCommentsOk

`func (o *ParcelCreateDataAttributes) GetRestrictionCommentsOk() (*string, bool)`

GetRestrictionCommentsOk returns a tuple with the RestrictionComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionComments

`func (o *ParcelCreateDataAttributes) SetRestrictionComments(v string)`

SetRestrictionComments sets RestrictionComments field to given value.

### HasRestrictionComments

`func (o *ParcelCreateDataAttributes) HasRestrictionComments() bool`

HasRestrictionComments returns a boolean if a field has been set.

### GetCustomsInfoRequired

`func (o *ParcelCreateDataAttributes) GetCustomsInfoRequired() bool`

GetCustomsInfoRequired returns the CustomsInfoRequired field if non-nil, zero value otherwise.

### GetCustomsInfoRequiredOk

`func (o *ParcelCreateDataAttributes) GetCustomsInfoRequiredOk() (*bool, bool)`

GetCustomsInfoRequiredOk returns a tuple with the CustomsInfoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsInfoRequired

`func (o *ParcelCreateDataAttributes) SetCustomsInfoRequired(v bool)`

SetCustomsInfoRequired sets CustomsInfoRequired field to given value.

### HasCustomsInfoRequired

`func (o *ParcelCreateDataAttributes) HasCustomsInfoRequired() bool`

HasCustomsInfoRequired returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ParcelCreateDataAttributes) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ParcelCreateDataAttributes) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ParcelCreateDataAttributes) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ParcelCreateDataAttributes) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingStatus

`func (o *ParcelCreateDataAttributes) GetTrackingStatus() string`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *ParcelCreateDataAttributes) GetTrackingStatusOk() (*string, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *ParcelCreateDataAttributes) SetTrackingStatus(v string)`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *ParcelCreateDataAttributes) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### GetTrackingStatusDetail

`func (o *ParcelCreateDataAttributes) GetTrackingStatusDetail() string`

GetTrackingStatusDetail returns the TrackingStatusDetail field if non-nil, zero value otherwise.

### GetTrackingStatusDetailOk

`func (o *ParcelCreateDataAttributes) GetTrackingStatusDetailOk() (*string, bool)`

GetTrackingStatusDetailOk returns a tuple with the TrackingStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusDetail

`func (o *ParcelCreateDataAttributes) SetTrackingStatusDetail(v string)`

SetTrackingStatusDetail sets TrackingStatusDetail field to given value.

### HasTrackingStatusDetail

`func (o *ParcelCreateDataAttributes) HasTrackingStatusDetail() bool`

HasTrackingStatusDetail returns a boolean if a field has been set.

### GetTrackingStatusUpdatedAt

`func (o *ParcelCreateDataAttributes) GetTrackingStatusUpdatedAt() string`

GetTrackingStatusUpdatedAt returns the TrackingStatusUpdatedAt field if non-nil, zero value otherwise.

### GetTrackingStatusUpdatedAtOk

`func (o *ParcelCreateDataAttributes) GetTrackingStatusUpdatedAtOk() (*string, bool)`

GetTrackingStatusUpdatedAtOk returns a tuple with the TrackingStatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatusUpdatedAt

`func (o *ParcelCreateDataAttributes) SetTrackingStatusUpdatedAt(v string)`

SetTrackingStatusUpdatedAt sets TrackingStatusUpdatedAt field to given value.

### HasTrackingStatusUpdatedAt

`func (o *ParcelCreateDataAttributes) HasTrackingStatusUpdatedAt() bool`

HasTrackingStatusUpdatedAt returns a boolean if a field has been set.

### GetTrackingDetails

`func (o *ParcelCreateDataAttributes) GetTrackingDetails() string`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *ParcelCreateDataAttributes) GetTrackingDetailsOk() (*string, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *ParcelCreateDataAttributes) SetTrackingDetails(v string)`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *ParcelCreateDataAttributes) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### GetCarrierWeightOz

`func (o *ParcelCreateDataAttributes) GetCarrierWeightOz() string`

GetCarrierWeightOz returns the CarrierWeightOz field if non-nil, zero value otherwise.

### GetCarrierWeightOzOk

`func (o *ParcelCreateDataAttributes) GetCarrierWeightOzOk() (*string, bool)`

GetCarrierWeightOzOk returns a tuple with the CarrierWeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierWeightOz

`func (o *ParcelCreateDataAttributes) SetCarrierWeightOz(v string)`

SetCarrierWeightOz sets CarrierWeightOz field to given value.

### HasCarrierWeightOz

`func (o *ParcelCreateDataAttributes) HasCarrierWeightOz() bool`

HasCarrierWeightOz returns a boolean if a field has been set.

### GetSignedBy

`func (o *ParcelCreateDataAttributes) GetSignedBy() string`

GetSignedBy returns the SignedBy field if non-nil, zero value otherwise.

### GetSignedByOk

`func (o *ParcelCreateDataAttributes) GetSignedByOk() (*string, bool)`

GetSignedByOk returns a tuple with the SignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedBy

`func (o *ParcelCreateDataAttributes) SetSignedBy(v string)`

SetSignedBy sets SignedBy field to given value.

### HasSignedBy

`func (o *ParcelCreateDataAttributes) HasSignedBy() bool`

HasSignedBy returns a boolean if a field has been set.

### GetIncoterm

`func (o *ParcelCreateDataAttributes) GetIncoterm() string`

GetIncoterm returns the Incoterm field if non-nil, zero value otherwise.

### GetIncotermOk

`func (o *ParcelCreateDataAttributes) GetIncotermOk() (*string, bool)`

GetIncotermOk returns a tuple with the Incoterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncoterm

`func (o *ParcelCreateDataAttributes) SetIncoterm(v string)`

SetIncoterm sets Incoterm field to given value.

### HasIncoterm

`func (o *ParcelCreateDataAttributes) HasIncoterm() bool`

HasIncoterm returns a boolean if a field has been set.

### GetReference

`func (o *ParcelCreateDataAttributes) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ParcelCreateDataAttributes) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ParcelCreateDataAttributes) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ParcelCreateDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceOrigin

`func (o *ParcelCreateDataAttributes) GetReferenceOrigin() string`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *ParcelCreateDataAttributes) GetReferenceOriginOk() (*string, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *ParcelCreateDataAttributes) SetReferenceOrigin(v string)`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *ParcelCreateDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### GetMetadata

`func (o *ParcelCreateDataAttributes) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ParcelCreateDataAttributes) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ParcelCreateDataAttributes) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ParcelCreateDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


