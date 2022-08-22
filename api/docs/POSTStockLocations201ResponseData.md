# POSTStockLocations201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "stock_locations"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTStockLocations201ResponseDataAttributes**](POSTStockLocations201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**POSTMerchants201ResponseDataRelationships**](POSTMerchants201ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTStockLocations201ResponseData

`func NewPOSTStockLocations201ResponseData() *POSTStockLocations201ResponseData`

NewPOSTStockLocations201ResponseData instantiates a new POSTStockLocations201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTStockLocations201ResponseDataWithDefaults

`func NewPOSTStockLocations201ResponseDataWithDefaults() *POSTStockLocations201ResponseData`

NewPOSTStockLocations201ResponseDataWithDefaults instantiates a new POSTStockLocations201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTStockLocations201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTStockLocations201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTStockLocations201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTStockLocations201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTStockLocations201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTStockLocations201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTStockLocations201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTStockLocations201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTStockLocations201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTStockLocations201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTStockLocations201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTStockLocations201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTStockLocations201ResponseData) GetAttributes() POSTStockLocations201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTStockLocations201ResponseData) GetAttributesOk() (*POSTStockLocations201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTStockLocations201ResponseData) SetAttributes(v POSTStockLocations201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTStockLocations201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTStockLocations201ResponseData) GetRelationships() POSTMerchants201ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTStockLocations201ResponseData) GetRelationshipsOk() (*POSTMerchants201ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTStockLocations201ResponseData) SetRelationships(v POSTMerchants201ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTStockLocations201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

