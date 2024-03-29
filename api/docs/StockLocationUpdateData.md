# StockLocationUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHStockLocationsStockLocationId200ResponseDataAttributes**](PATCHStockLocationsStockLocationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**MerchantUpdateDataRelationships**](MerchantUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewStockLocationUpdateData

`func NewStockLocationUpdateData(type_ interface{}, id interface{}, attributes PATCHStockLocationsStockLocationId200ResponseDataAttributes, ) *StockLocationUpdateData`

NewStockLocationUpdateData instantiates a new StockLocationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationUpdateDataWithDefaults

`func NewStockLocationUpdateDataWithDefaults() *StockLocationUpdateData`

NewStockLocationUpdateDataWithDefaults instantiates a new StockLocationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLocationUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLocationUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLocationUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockLocationUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockLocationUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *StockLocationUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StockLocationUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StockLocationUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *StockLocationUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StockLocationUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *StockLocationUpdateData) GetAttributes() PATCHStockLocationsStockLocationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLocationUpdateData) GetAttributesOk() (*PATCHStockLocationsStockLocationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLocationUpdateData) SetAttributes(v PATCHStockLocationsStockLocationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLocationUpdateData) GetRelationships() MerchantUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLocationUpdateData) GetRelationshipsOk() (*MerchantUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLocationUpdateData) SetRelationships(v MerchantUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLocationUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


