# GETParcels200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "parcels"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**GETParcels200ResponseDataInnerAttributes**](GETParcels200ResponseDataInnerAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GETParcels200ResponseDataInnerRelationships**](GETParcels200ResponseDataInnerRelationships.md) |  | [optional] 

## Methods

### NewGETParcels200ResponseDataInner

`func NewGETParcels200ResponseDataInner() *GETParcels200ResponseDataInner`

NewGETParcels200ResponseDataInner instantiates a new GETParcels200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETParcels200ResponseDataInnerWithDefaults

`func NewGETParcels200ResponseDataInnerWithDefaults() *GETParcels200ResponseDataInner`

NewGETParcels200ResponseDataInnerWithDefaults instantiates a new GETParcels200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GETParcels200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GETParcels200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GETParcels200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GETParcels200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GETParcels200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETParcels200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETParcels200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GETParcels200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *GETParcels200ResponseDataInner) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GETParcels200ResponseDataInner) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GETParcels200ResponseDataInner) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GETParcels200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *GETParcels200ResponseDataInner) GetAttributes() GETParcels200ResponseDataInnerAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GETParcels200ResponseDataInner) GetAttributesOk() (*GETParcels200ResponseDataInnerAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GETParcels200ResponseDataInner) SetAttributes(v GETParcels200ResponseDataInnerAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GETParcels200ResponseDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GETParcels200ResponseDataInner) GetRelationships() GETParcels200ResponseDataInnerRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GETParcels200ResponseDataInner) GetRelationshipsOk() (*GETParcels200ResponseDataInnerRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GETParcels200ResponseDataInner) SetRelationships(v GETParcels200ResponseDataInnerRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GETParcels200ResponseDataInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

