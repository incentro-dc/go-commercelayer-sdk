# PATCHPackagesPackageId200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The resource&#39;s id | [optional] 
**Type** | Pointer to **string** | The resource&#39;s type | [optional] [default to "packages"]
**Links** | Pointer to [**GETAddresses200ResponseDataInnerLinks**](GETAddresses200ResponseDataInnerLinks.md) |  | [optional] 
**Attributes** | Pointer to [**PATCHPackagesPackageId200ResponseDataAttributes**](PATCHPackagesPackageId200ResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**PATCHPackagesPackageId200ResponseDataRelationships**](PATCHPackagesPackageId200ResponseDataRelationships.md) |  | [optional] 

## Methods

### NewPATCHPackagesPackageId200ResponseData

`func NewPATCHPackagesPackageId200ResponseData() *PATCHPackagesPackageId200ResponseData`

NewPATCHPackagesPackageId200ResponseData instantiates a new PATCHPackagesPackageId200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHPackagesPackageId200ResponseDataWithDefaults

`func NewPATCHPackagesPackageId200ResponseDataWithDefaults() *PATCHPackagesPackageId200ResponseData`

NewPATCHPackagesPackageId200ResponseDataWithDefaults instantiates a new PATCHPackagesPackageId200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PATCHPackagesPackageId200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PATCHPackagesPackageId200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PATCHPackagesPackageId200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PATCHPackagesPackageId200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PATCHPackagesPackageId200ResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PATCHPackagesPackageId200ResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PATCHPackagesPackageId200ResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PATCHPackagesPackageId200ResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *PATCHPackagesPackageId200ResponseData) GetLinks() GETAddresses200ResponseDataInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PATCHPackagesPackageId200ResponseData) GetLinksOk() (*GETAddresses200ResponseDataInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PATCHPackagesPackageId200ResponseData) SetLinks(v GETAddresses200ResponseDataInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PATCHPackagesPackageId200ResponseData) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAttributes

`func (o *PATCHPackagesPackageId200ResponseData) GetAttributes() PATCHPackagesPackageId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PATCHPackagesPackageId200ResponseData) GetAttributesOk() (*PATCHPackagesPackageId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PATCHPackagesPackageId200ResponseData) SetAttributes(v PATCHPackagesPackageId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PATCHPackagesPackageId200ResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *PATCHPackagesPackageId200ResponseData) GetRelationships() PATCHPackagesPackageId200ResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PATCHPackagesPackageId200ResponseData) GetRelationshipsOk() (*PATCHPackagesPackageId200ResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PATCHPackagesPackageId200ResponseData) SetRelationships(v PATCHPackagesPackageId200ResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PATCHPackagesPackageId200ResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

