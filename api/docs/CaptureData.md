# CaptureData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "captures"]
**Attributes** | [**CaptureDataAttributes**](CaptureDataAttributes.md) |  | 
**Relationships** | Pointer to [**CaptureDataRelationships**](CaptureDataRelationships.md) |  | [optional] 

## Methods

### NewCaptureData

`func NewCaptureData(type_ string, attributes CaptureDataAttributes, ) *CaptureData`

NewCaptureData instantiates a new CaptureData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureDataWithDefaults

`func NewCaptureDataWithDefaults() *CaptureData`

NewCaptureDataWithDefaults instantiates a new CaptureData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CaptureData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CaptureData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CaptureData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CaptureData) GetAttributes() CaptureDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CaptureData) GetAttributesOk() (*CaptureDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CaptureData) SetAttributes(v CaptureDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CaptureData) GetRelationships() CaptureDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CaptureData) GetRelationshipsOk() (*CaptureDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CaptureData) SetRelationships(v CaptureDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CaptureData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


