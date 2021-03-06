# SkuCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "skus"]
**Attributes** | [**SkuCreateDataAttributes**](SkuCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuCreateDataRelationships**](SkuCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuCreateData

`func NewSkuCreateData(type_ string, attributes SkuCreateDataAttributes, ) *SkuCreateData`

NewSkuCreateData instantiates a new SkuCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuCreateDataWithDefaults

`func NewSkuCreateDataWithDefaults() *SkuCreateData`

NewSkuCreateDataWithDefaults instantiates a new SkuCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuCreateData) GetAttributes() SkuCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuCreateData) GetAttributesOk() (*SkuCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuCreateData) SetAttributes(v SkuCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuCreateData) GetRelationships() SkuCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuCreateData) GetRelationshipsOk() (*SkuCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuCreateData) SetRelationships(v SkuCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


