# ParcelCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "parcels"]
**Attributes** | [**ParcelCreateDataAttributes**](ParcelCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**ParcelCreateDataRelationships**](ParcelCreateDataRelationships.md) |  | [optional] 

## Methods

### NewParcelCreateData

`func NewParcelCreateData(type_ string, attributes ParcelCreateDataAttributes, ) *ParcelCreateData`

NewParcelCreateData instantiates a new ParcelCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelCreateDataWithDefaults

`func NewParcelCreateDataWithDefaults() *ParcelCreateData`

NewParcelCreateDataWithDefaults instantiates a new ParcelCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ParcelCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParcelCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParcelCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *ParcelCreateData) GetAttributes() ParcelCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ParcelCreateData) GetAttributesOk() (*ParcelCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ParcelCreateData) SetAttributes(v ParcelCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ParcelCreateData) GetRelationships() ParcelCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ParcelCreateData) GetRelationshipsOk() (*ParcelCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ParcelCreateData) SetRelationships(v ParcelCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ParcelCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


