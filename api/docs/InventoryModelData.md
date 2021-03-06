# InventoryModelData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "inventory_models"]
**Attributes** | [**InventoryModelDataAttributes**](InventoryModelDataAttributes.md) |  | 
**Relationships** | Pointer to [**InventoryModelDataRelationships**](InventoryModelDataRelationships.md) |  | [optional] 

## Methods

### NewInventoryModelData

`func NewInventoryModelData(type_ string, attributes InventoryModelDataAttributes, ) *InventoryModelData`

NewInventoryModelData instantiates a new InventoryModelData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryModelDataWithDefaults

`func NewInventoryModelDataWithDefaults() *InventoryModelData`

NewInventoryModelDataWithDefaults instantiates a new InventoryModelData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InventoryModelData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InventoryModelData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InventoryModelData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *InventoryModelData) GetAttributes() InventoryModelDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryModelData) GetAttributesOk() (*InventoryModelDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InventoryModelData) SetAttributes(v InventoryModelDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *InventoryModelData) GetRelationships() InventoryModelDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InventoryModelData) GetRelationshipsOk() (*InventoryModelDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InventoryModelData) SetRelationships(v InventoryModelDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InventoryModelData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


