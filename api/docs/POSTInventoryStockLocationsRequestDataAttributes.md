# POSTInventoryStockLocationsRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **interface{}** | The stock location priority within the associated invetory model. | 
**OnHold** | Pointer to **interface{}** | Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTInventoryStockLocationsRequestDataAttributes

`func NewPOSTInventoryStockLocationsRequestDataAttributes(priority interface{}, ) *POSTInventoryStockLocationsRequestDataAttributes`

NewPOSTInventoryStockLocationsRequestDataAttributes instantiates a new POSTInventoryStockLocationsRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTInventoryStockLocationsRequestDataAttributesWithDefaults

`func NewPOSTInventoryStockLocationsRequestDataAttributesWithDefaults() *POSTInventoryStockLocationsRequestDataAttributes`

NewPOSTInventoryStockLocationsRequestDataAttributesWithDefaults instantiates a new POSTInventoryStockLocationsRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *POSTInventoryStockLocationsRequestDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetOnHold

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetOnHold() interface{}`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetOnHoldOk() (*interface{}, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetOnHold(v interface{})`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *POSTInventoryStockLocationsRequestDataAttributes) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### SetOnHoldNil

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetOnHoldNil(b bool)`

 SetOnHoldNil sets the value for OnHold to be an explicit nil

### UnsetOnHold
`func (o *POSTInventoryStockLocationsRequestDataAttributes) UnsetOnHold()`

UnsetOnHold ensures that no value is present for OnHold, not even an explicit nil
### GetReference

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTInventoryStockLocationsRequestDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTInventoryStockLocationsRequestDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTInventoryStockLocationsRequestDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTInventoryStockLocationsRequestDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTInventoryStockLocationsRequestDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTInventoryStockLocationsRequestDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTInventoryStockLocationsRequestDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTInventoryStockLocationsRequestDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

