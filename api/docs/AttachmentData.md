# AttachmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "attachments"]
**Attributes** | [**AttachmentDataAttributes**](AttachmentDataAttributes.md) |  | 
**Relationships** | Pointer to [**AttachmentDataRelationships**](AttachmentDataRelationships.md) |  | [optional] 

## Methods

### NewAttachmentData

`func NewAttachmentData(type_ string, attributes AttachmentDataAttributes, ) *AttachmentData`

NewAttachmentData instantiates a new AttachmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentDataWithDefaults

`func NewAttachmentDataWithDefaults() *AttachmentData`

NewAttachmentDataWithDefaults instantiates a new AttachmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AttachmentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachmentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachmentData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AttachmentData) GetAttributes() AttachmentDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AttachmentData) GetAttributesOk() (*AttachmentDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AttachmentData) SetAttributes(v AttachmentDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AttachmentData) GetRelationships() AttachmentDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AttachmentData) GetRelationshipsOk() (*AttachmentDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AttachmentData) SetRelationships(v AttachmentDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AttachmentData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


