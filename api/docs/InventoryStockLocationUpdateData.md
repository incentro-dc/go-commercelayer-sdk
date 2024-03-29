# InventoryStockLocationUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes**](PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryReturnLocationUpdateDataRelationships**](InventoryReturnLocationUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryStockLocationUpdateData

`func NewInventoryStockLocationUpdateData(type_ interface{}, id interface{}, attributes PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes, ) *InventoryStockLocationUpdateData`

NewInventoryStockLocationUpdateData instantiates a new InventoryStockLocationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryStockLocationUpdateDataWithDefaults

`func NewInventoryStockLocationUpdateDataWithDefaults() *InventoryStockLocationUpdateData`

NewInventoryStockLocationUpdateDataWithDefaults instantiates a new InventoryStockLocationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryStockLocationUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryStockLocationUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryStockLocationUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *InventoryStockLocationUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InventoryStockLocationUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *InventoryStockLocationUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryStockLocationUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryStockLocationUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *InventoryStockLocationUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InventoryStockLocationUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *InventoryStockLocationUpdateData) GetAttributes() PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryStockLocationUpdateData) GetAttributesOk() (*PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryStockLocationUpdateData) SetAttributes(v PATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryStockLocationUpdateData) GetRelationships() InventoryReturnLocationUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryStockLocationUpdateData) GetRelationshipsOk() (*InventoryReturnLocationUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryStockLocationUpdateData) SetRelationships(v InventoryReturnLocationUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryStockLocationUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


