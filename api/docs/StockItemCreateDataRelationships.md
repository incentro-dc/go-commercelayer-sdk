# StockItemCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockLocation** | [**DeliveryLeadTimeCreateDataRelationshipsStockLocation**](DeliveryLeadTimeCreateDataRelationshipsStockLocation.md) |  | 
**Sku** | Pointer to [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | [optional] 

## Methods

### NewStockItemCreateDataRelationships

`func NewStockItemCreateDataRelationships(stockLocation DeliveryLeadTimeCreateDataRelationshipsStockLocation, ) *StockItemCreateDataRelationships`

NewStockItemCreateDataRelationships instantiates a new StockItemCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStockItemCreateDataRelationshipsWithDefaults

`func NewStockItemCreateDataRelationshipsWithDefaults() *StockItemCreateDataRelationships`

NewStockItemCreateDataRelationshipsWithDefaults instantiates a new StockItemCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStockLocation

`func (o *StockItemCreateDataRelationships) GetStockLocation() DeliveryLeadTimeCreateDataRelationshipsStockLocation`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *StockItemCreateDataRelationships) GetStockLocationOk() (*DeliveryLeadTimeCreateDataRelationshipsStockLocation, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *StockItemCreateDataRelationships) SetStockLocation(v DeliveryLeadTimeCreateDataRelationshipsStockLocation)`

SetStockLocation sets StockLocation field to given value.


### GetSku

`func (o *StockItemCreateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *StockItemCreateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *StockItemCreateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *StockItemCreateDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


