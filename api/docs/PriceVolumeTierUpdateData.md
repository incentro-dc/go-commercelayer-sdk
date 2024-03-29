# PriceVolumeTierUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes**](PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PriceFrequencyTierUpdateDataRelationships**](PriceFrequencyTierUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPriceVolumeTierUpdateData

`func NewPriceVolumeTierUpdateData(type_ interface{}, id interface{}, attributes PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes, ) *PriceVolumeTierUpdateData`

NewPriceVolumeTierUpdateData instantiates a new PriceVolumeTierUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceVolumeTierUpdateDataWithDefaults

`func NewPriceVolumeTierUpdateDataWithDefaults() *PriceVolumeTierUpdateData`

NewPriceVolumeTierUpdateDataWithDefaults instantiates a new PriceVolumeTierUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PriceVolumeTierUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PriceVolumeTierUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PriceVolumeTierUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PriceVolumeTierUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PriceVolumeTierUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PriceVolumeTierUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PriceVolumeTierUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PriceVolumeTierUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PriceVolumeTierUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PriceVolumeTierUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PriceVolumeTierUpdateData) GetAttributes() PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PriceVolumeTierUpdateData) GetAttributesOk() (*PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PriceVolumeTierUpdateData) SetAttributes(v PATCHPriceVolumeTiersPriceVolumeTierId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PriceVolumeTierUpdateData) GetRelationships() PriceFrequencyTierUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PriceVolumeTierUpdateData) GetRelationshipsOk() (*PriceFrequencyTierUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PriceVolumeTierUpdateData) SetRelationships(v PriceFrequencyTierUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PriceVolumeTierUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


