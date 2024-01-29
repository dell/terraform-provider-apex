# StorageSystemDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name to identify the deployed system | 
**CloudOptions** | [**StorageSystemDeploymentRequestCloudOptions**](StorageSystemDeploymentRequestCloudOptions.md) |  | 
**StorageOptions** | [**StorageSystemDeploymentRequestStorageOptions**](StorageSystemDeploymentRequestStorageOptions.md) |  | 
**IsTermsAndConditionsAgreed** | **bool** | By setting this option to true, you, on your behalf of your company, agree the following will apply:  Your evaluation of Dell APEX Navigator for Multicloud is subject to and governed by Dell’s Cloud Service Offering Agreement (https://www.dell.com/learn/us/en/uscorp1/legal_terms-conditions_dellwebpage/csoa-agreement) and  the Dell APEX Navigator for Multicloud Storage Service Offering Description (https://www.dell.com/learn/us/en/uscorp1/apex-services).  Your evaluation of Dell APEX Block Storage for AWS is subject to and governed by Dell’s Software Evaluation Agreement  (https://vault.pactsafe.io/s/3cb1f636-b99f-47ab-ad13-3ce168930a55/legal.html?g&#x3D;38113#contract-s12zpzeii).  Your evaluation of Dell APEX Navigator for Multicloud will be available for 90 days free of charge. At the end of the evaluation, the Dell APEX Navigator for Multicloud Storage evaluation service and associated data will remain accessible, subject to the terms identified in this paragraph, however, certain features will be limited until you order a new subscription. Payment for any required public cloud service provider infrastructure requirements is your company’s responsibility. | 

## Methods

### NewStorageSystemDeploymentRequest

`func NewStorageSystemDeploymentRequest(name string, cloudOptions StorageSystemDeploymentRequestCloudOptions, storageOptions StorageSystemDeploymentRequestStorageOptions, isTermsAndConditionsAgreed bool, ) *StorageSystemDeploymentRequest`

NewStorageSystemDeploymentRequest instantiates a new StorageSystemDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSystemDeploymentRequestWithDefaults

`func NewStorageSystemDeploymentRequestWithDefaults() *StorageSystemDeploymentRequest`

NewStorageSystemDeploymentRequestWithDefaults instantiates a new StorageSystemDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageSystemDeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageSystemDeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageSystemDeploymentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCloudOptions

`func (o *StorageSystemDeploymentRequest) GetCloudOptions() StorageSystemDeploymentRequestCloudOptions`

GetCloudOptions returns the CloudOptions field if non-nil, zero value otherwise.

### GetCloudOptionsOk

`func (o *StorageSystemDeploymentRequest) GetCloudOptionsOk() (*StorageSystemDeploymentRequestCloudOptions, bool)`

GetCloudOptionsOk returns a tuple with the CloudOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudOptions

`func (o *StorageSystemDeploymentRequest) SetCloudOptions(v StorageSystemDeploymentRequestCloudOptions)`

SetCloudOptions sets CloudOptions field to given value.


### GetStorageOptions

`func (o *StorageSystemDeploymentRequest) GetStorageOptions() StorageSystemDeploymentRequestStorageOptions`

GetStorageOptions returns the StorageOptions field if non-nil, zero value otherwise.

### GetStorageOptionsOk

`func (o *StorageSystemDeploymentRequest) GetStorageOptionsOk() (*StorageSystemDeploymentRequestStorageOptions, bool)`

GetStorageOptionsOk returns a tuple with the StorageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageOptions

`func (o *StorageSystemDeploymentRequest) SetStorageOptions(v StorageSystemDeploymentRequestStorageOptions)`

SetStorageOptions sets StorageOptions field to given value.


### GetIsTermsAndConditionsAgreed

`func (o *StorageSystemDeploymentRequest) GetIsTermsAndConditionsAgreed() bool`

GetIsTermsAndConditionsAgreed returns the IsTermsAndConditionsAgreed field if non-nil, zero value otherwise.

### GetIsTermsAndConditionsAgreedOk

`func (o *StorageSystemDeploymentRequest) GetIsTermsAndConditionsAgreedOk() (*bool, bool)`

GetIsTermsAndConditionsAgreedOk returns a tuple with the IsTermsAndConditionsAgreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTermsAndConditionsAgreed

`func (o *StorageSystemDeploymentRequest) SetIsTermsAndConditionsAgreed(v bool)`

SetIsTermsAndConditionsAgreed sets IsTermsAndConditionsAgreed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


