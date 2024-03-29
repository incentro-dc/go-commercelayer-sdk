# StockLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETStockLocationsStockLocationId200ResponseDataAttributes**](GETStockLocationsStockLocationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockLocationDataRelationships**](StockLocationDataRelationships.md) |  | [optional] 

## Methods

### NewStockLocationData

`func NewStockLocationData(type_ interface{}, attributes GETStockLocationsStockLocationId200ResponseDataAttributes, ) *StockLocationData`

NewStockLocationData instantiates a new StockLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockLocationDataWithDefaults

`func NewStockLocationDataWithDefaults() *StockLocationData`

NewStockLocationDataWithDefaults instantiates a new StockLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockLocationData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockLocationData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockLocationData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StockLocationData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StockLocationData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StockLocationData) GetAttributes() GETStockLocationsStockLocationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockLocationData) GetAttributesOk() (*GETStockLocationsStockLocationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockLocationData) SetAttributes(v GETStockLocationsStockLocationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockLocationData) GetRelationships() StockLocationDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockLocationData) GetRelationshipsOk() (*StockLocationDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockLocationData) SetRelationships(v StockLocationDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockLocationData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


