# MobilityMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the member (e.g. PowerFlex volume identifier) | 
**Name** | **string** | Name of the member (e.g. name of the volume) | 
**Size** | **string** | Size of the member (e.g. volume size in bytes) | 

## Methods

### NewMobilityMember

`func NewMobilityMember(id string, name string, size string, ) *MobilityMember`

NewMobilityMember instantiates a new MobilityMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobilityMemberWithDefaults

`func NewMobilityMemberWithDefaults() *MobilityMember`

NewMobilityMemberWithDefaults instantiates a new MobilityMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MobilityMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MobilityMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MobilityMember) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MobilityMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobilityMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobilityMember) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *MobilityMember) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MobilityMember) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MobilityMember) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


