# AvalaraAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "avalara_accounts"]
**Attributes** | [**AvalaraAccountDataAttributes**](AvalaraAccountDataAttributes.md) |  | 
**Relationships** | Pointer to [**AvalaraAccountDataRelationships**](AvalaraAccountDataRelationships.md) |  | [optional] 

## Methods

### NewAvalaraAccountData

`func NewAvalaraAccountData(type_ string, attributes AvalaraAccountDataAttributes, ) *AvalaraAccountData`

NewAvalaraAccountData instantiates a new AvalaraAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvalaraAccountDataWithDefaults

`func NewAvalaraAccountDataWithDefaults() *AvalaraAccountData`

NewAvalaraAccountDataWithDefaults instantiates a new AvalaraAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AvalaraAccountData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AvalaraAccountData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AvalaraAccountData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AvalaraAccountData) GetAttributes() AvalaraAccountDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AvalaraAccountData) GetAttributesOk() (*AvalaraAccountDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AvalaraAccountData) SetAttributes(v AvalaraAccountDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AvalaraAccountData) GetRelationships() AvalaraAccountDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AvalaraAccountData) GetRelationshipsOk() (*AvalaraAccountDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AvalaraAccountData) SetRelationships(v AvalaraAccountDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AvalaraAccountData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


