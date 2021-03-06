# StockItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "stock_items"]
**Attributes** | [**StockItemCreateDataAttributes**](StockItemCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**StockItemCreateDataRelationships**](StockItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStockItemCreateData

`func NewStockItemCreateData(type_ string, attributes StockItemCreateDataAttributes, ) *StockItemCreateData`

NewStockItemCreateData instantiates a new StockItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemCreateDataWithDefaults

`func NewStockItemCreateDataWithDefaults() *StockItemCreateData`

NewStockItemCreateDataWithDefaults instantiates a new StockItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StockItemCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StockItemCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StockItemCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *StockItemCreateData) GetAttributes() StockItemCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StockItemCreateData) GetAttributesOk() (*StockItemCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StockItemCreateData) SetAttributes(v StockItemCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StockItemCreateData) GetRelationships() StockItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StockItemCreateData) GetRelationshipsOk() (*StockItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StockItemCreateData) SetRelationships(v StockItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StockItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


