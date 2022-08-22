# POSTExternalGateways201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "external_gateways"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**POSTExternalGateways201ResponseDataAttributes**](POSTExternalGateways201ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPOSTExternalGateways201ResponseData

`func NewPOSTExternalGateways201ResponseData() *POSTExternalGateways201ResponseData`

NewPOSTExternalGateways201ResponseData instantiates a new POSTExternalGateways201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTExternalGateways201ResponseDataWithDefaults

`func NewPOSTExternalGateways201ResponseDataWithDefaults() *POSTExternalGateways201ResponseData`

NewPOSTExternalGateways201ResponseDataWithDefaults instantiates a new POSTExternalGateways201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *POSTExternalGateways201ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *POSTExternalGateways201ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *POSTExternalGateways201ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *POSTExternalGateways201ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *POSTExternalGateways201ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTExternalGateways201ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTExternalGateways201ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *POSTExternalGateways201ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *POSTExternalGateways201ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *POSTExternalGateways201ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *POSTExternalGateways201ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *POSTExternalGateways201ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *POSTExternalGateways201ResponseData) GetAttributes() POSTExternalGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTExternalGateways201ResponseData) GetAttributesOk() (*POSTExternalGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTExternalGateways201ResponseData) SetAttributes(v POSTExternalGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *POSTExternalGateways201ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *POSTExternalGateways201ResponseData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTExternalGateways201ResponseData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTExternalGateways201ResponseData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTExternalGateways201ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

