# SkuData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "skus"]
**Attributes** | [**SkuDataAttributes**](SkuDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuDataRelationships**](SkuDataRelationships.md) |  | [optional] 

## Methods

### NewSkuData

`func NewSkuData(type_ string, attributes SkuDataAttributes, ) *SkuData`

NewSkuData instantiates a new SkuData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuDataWithDefaults

`func NewSkuDataWithDefaults() *SkuData`

NewSkuDataWithDefaults instantiates a new SkuData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SkuData) GetAttributes() SkuDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuData) GetAttributesOk() (*SkuDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuData) SetAttributes(v SkuDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuData) GetRelationships() SkuDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuData) GetRelationshipsOk() (*SkuDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuData) SetRelationships(v SkuDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


