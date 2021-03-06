# AvalaraAccountUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "avalara_accounts"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AvalaraAccountUpdateDataAttributes**](AvalaraAccountUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountCreateDataRelationships**](AvalaraAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAvalaraAccountUpdateData

`func NewAvalaraAccountUpdateData(type_ string, id string, attributes AvalaraAccountUpdateDataAttributes, ) *AvalaraAccountUpdateData`

NewAvalaraAccountUpdateData instantiates a new AvalaraAccountUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvalaraAccountUpdateDataWithDefaults

`func NewAvalaraAccountUpdateDataWithDefaults() *AvalaraAccountUpdateData`

NewAvalaraAccountUpdateDataWithDefaults instantiates a new AvalaraAccountUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AvalaraAccountUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AvalaraAccountUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AvalaraAccountUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AvalaraAccountUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvalaraAccountUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvalaraAccountUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AvalaraAccountUpdateData) GetAttributes() AvalaraAccountUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AvalaraAccountUpdateData) GetAttributesOk() (*AvalaraAccountUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AvalaraAccountUpdateData) SetAttributes(v AvalaraAccountUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AvalaraAccountUpdateData) GetRelationships() AvalaraAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AvalaraAccountUpdateData) GetRelationshipsOk() (*AvalaraAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AvalaraAccountUpdateData) SetRelationships(v AvalaraAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AvalaraAccountUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


