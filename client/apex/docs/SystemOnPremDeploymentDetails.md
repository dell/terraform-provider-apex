# SystemOnPremDeploymentDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | Pointer to [**SystemDeploymentTypeEnum**](SystemDeploymentTypeEnum.md) |  | [optional] [default to SYSTEMDEPLOYMENTTYPEENUM_ONPREM]
**SiteName** | Pointer to **string** | Name of the site where the system is located. | [optional] 
**Location** | Pointer to **string** | User provided description of where the system is located. | [optional] 
**Country** | Pointer to **string** | Name of the country where the system is located. | [optional] 
**State** | Pointer to **string** | Name of the state where the system is located. | [optional] 
**City** | Pointer to **string** | Name of the city where the system is located. | [optional] 
**StreetAddress1** | Pointer to **string** | Street address 1 of where the system is located. | [optional] 
**StreetAddress2** | Pointer to **string** | Street address 2 of where the system is located. | [optional] 
**ZipCode** | Pointer to **string** | State ZIP code of where the system is located. | [optional] 

## Methods

### NewSystemOnPremDeploymentDetails

`func NewSystemOnPremDeploymentDetails() *SystemOnPremDeploymentDetails`

NewSystemOnPremDeploymentDetails instantiates a new SystemOnPremDeploymentDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemOnPremDeploymentDetailsWithDefaults

`func NewSystemOnPremDeploymentDetailsWithDefaults() *SystemOnPremDeploymentDetails`

NewSystemOnPremDeploymentDetailsWithDefaults instantiates a new SystemOnPremDeploymentDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentType

`func (o *SystemOnPremDeploymentDetails) GetDeploymentType() SystemDeploymentTypeEnum`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *SystemOnPremDeploymentDetails) GetDeploymentTypeOk() (*SystemDeploymentTypeEnum, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *SystemOnPremDeploymentDetails) SetDeploymentType(v SystemDeploymentTypeEnum)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *SystemOnPremDeploymentDetails) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetSiteName

`func (o *SystemOnPremDeploymentDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SystemOnPremDeploymentDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SystemOnPremDeploymentDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SystemOnPremDeploymentDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetLocation

`func (o *SystemOnPremDeploymentDetails) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SystemOnPremDeploymentDetails) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SystemOnPremDeploymentDetails) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SystemOnPremDeploymentDetails) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountry

`func (o *SystemOnPremDeploymentDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SystemOnPremDeploymentDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SystemOnPremDeploymentDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SystemOnPremDeploymentDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetState

`func (o *SystemOnPremDeploymentDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SystemOnPremDeploymentDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SystemOnPremDeploymentDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SystemOnPremDeploymentDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCity

`func (o *SystemOnPremDeploymentDetails) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SystemOnPremDeploymentDetails) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SystemOnPremDeploymentDetails) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SystemOnPremDeploymentDetails) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *SystemOnPremDeploymentDetails) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *SystemOnPremDeploymentDetails) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *SystemOnPremDeploymentDetails) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.

### HasStreetAddress1

`func (o *SystemOnPremDeploymentDetails) HasStreetAddress1() bool`

HasStreetAddress1 returns a boolean if a field has been set.

### GetStreetAddress2

`func (o *SystemOnPremDeploymentDetails) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *SystemOnPremDeploymentDetails) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *SystemOnPremDeploymentDetails) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *SystemOnPremDeploymentDetails) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetZipCode

`func (o *SystemOnPremDeploymentDetails) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *SystemOnPremDeploymentDetails) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *SystemOnPremDeploymentDetails) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *SystemOnPremDeploymentDetails) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


