# LineItemCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTLineItems201ResponseDataAttributes**](POSTLineItems201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**LineItemCreateDataRelationships**](LineItemCreateDataRelationships.md) |  | [optional] 

## Methods

### NewLineItemCreateData

`func NewLineItemCreateData(type_ interface{}, attributes POSTLineItems201ResponseDataAttributes, ) *LineItemCreateData`

NewLineItemCreateData instantiates a new LineItemCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineItemCreateDataWithDefaults

`func NewLineItemCreateDataWithDefaults() *LineItemCreateData`

NewLineItemCreateDataWithDefaults instantiates a new LineItemCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LineItemCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LineItemCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LineItemCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *LineItemCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LineItemCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *LineItemCreateData) GetAttributes() POSTLineItems201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *LineItemCreateData) GetAttributesOk() (*POSTLineItems201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *LineItemCreateData) SetAttributes(v POSTLineItems201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *LineItemCreateData) GetRelationships() LineItemCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *LineItemCreateData) GetRelationshipsOk() (*LineItemCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *LineItemCreateData) SetRelationships(v LineItemCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *LineItemCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


