# PrebuildDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** |  | 
**CommitInterval** | **int32** |  | 
**Id** | **string** |  | 
**ProjectConfigName** | **string** |  | 
**TriggerFiles** | **[]string** |  | 

## Methods

### NewPrebuildDTO

`func NewPrebuildDTO(branch string, commitInterval int32, id string, projectConfigName string, triggerFiles []string, ) *PrebuildDTO`

NewPrebuildDTO instantiates a new PrebuildDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrebuildDTOWithDefaults

`func NewPrebuildDTOWithDefaults() *PrebuildDTO`

NewPrebuildDTOWithDefaults instantiates a new PrebuildDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *PrebuildDTO) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PrebuildDTO) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PrebuildDTO) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetCommitInterval

`func (o *PrebuildDTO) GetCommitInterval() int32`

GetCommitInterval returns the CommitInterval field if non-nil, zero value otherwise.

### GetCommitIntervalOk

`func (o *PrebuildDTO) GetCommitIntervalOk() (*int32, bool)`

GetCommitIntervalOk returns a tuple with the CommitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitInterval

`func (o *PrebuildDTO) SetCommitInterval(v int32)`

SetCommitInterval sets CommitInterval field to given value.


### GetId

`func (o *PrebuildDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrebuildDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrebuildDTO) SetId(v string)`

SetId sets Id field to given value.


### GetProjectConfigName

`func (o *PrebuildDTO) GetProjectConfigName() string`

GetProjectConfigName returns the ProjectConfigName field if non-nil, zero value otherwise.

### GetProjectConfigNameOk

`func (o *PrebuildDTO) GetProjectConfigNameOk() (*string, bool)`

GetProjectConfigNameOk returns a tuple with the ProjectConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectConfigName

`func (o *PrebuildDTO) SetProjectConfigName(v string)`

SetProjectConfigName sets ProjectConfigName field to given value.


### GetTriggerFiles

`func (o *PrebuildDTO) GetTriggerFiles() []string`

GetTriggerFiles returns the TriggerFiles field if non-nil, zero value otherwise.

### GetTriggerFilesOk

`func (o *PrebuildDTO) GetTriggerFilesOk() (*[]string, bool)`

GetTriggerFilesOk returns a tuple with the TriggerFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerFiles

`func (o *PrebuildDTO) SetTriggerFiles(v []string)`

SetTriggerFiles sets TriggerFiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

