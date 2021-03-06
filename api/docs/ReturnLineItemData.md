# ReturnLineItemData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "return_line_items"]
**Attributes** | [**ReturnLineItemDataAttributes**](ReturnLineItemDataAttributes.md) |  | 
**Relationships** | Pointer to [**ReturnLineItemDataRelationships**](ReturnLineItemDataRelationships.md) |  | [optional] 

## Methods

### NewReturnLineItemData

`func NewReturnLineItemData(type_ string, attributes ReturnLineItemDataAttributes, ) *ReturnLineItemData`

NewReturnLineItemData instantiates a new ReturnLineItemData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnLineItemDataWithDefaults

`func NewReturnLineItemDataWithDefaults() *ReturnLineItemData`

NewReturnLineItemDataWithDefaults instantiates a new ReturnLineItemData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReturnLineItemData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReturnLineItemData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReturnLineItemData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ReturnLineItemData) GetAttributes() ReturnLineItemDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ReturnLineItemData) GetAttributesOk() (*ReturnLineItemDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ReturnLineItemData) SetAttributes(v ReturnLineItemDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ReturnLineItemData) GetRelationships() ReturnLineItemDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ReturnLineItemData) GetRelationshipsOk() (*ReturnLineItemDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ReturnLineItemData) SetRelationships(v ReturnLineItemDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ReturnLineItemData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


