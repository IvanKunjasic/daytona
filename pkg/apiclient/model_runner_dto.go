/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RunnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunnerDTO{}

// RunnerDTO struct for RunnerDTO
type RunnerDTO struct {
	Alias    string          `json:"alias"`
	Id       string          `json:"id"`
	Metadata *RunnerMetadata `json:"metadata,omitempty"`
	State    string          `json:"state"`
}

type _RunnerDTO RunnerDTO

// NewRunnerDTO instantiates a new RunnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunnerDTO(alias string, id string, state string) *RunnerDTO {
	this := RunnerDTO{}
	this.Alias = alias
	this.Id = id
	this.State = state
	return &this
}

// NewRunnerDTOWithDefaults instantiates a new RunnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunnerDTOWithDefaults() *RunnerDTO {
	this := RunnerDTO{}
	return &this
}

// GetAlias returns the Alias field value
func (o *RunnerDTO) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *RunnerDTO) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *RunnerDTO) SetAlias(v string) {
	o.Alias = v
}

// GetId returns the Id field value
func (o *RunnerDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RunnerDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RunnerDTO) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RunnerDTO) GetMetadata() RunnerMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret RunnerMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerDTO) GetMetadataOk() (*RunnerMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RunnerDTO) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RunnerMetadata and assigns it to the Metadata field.
func (o *RunnerDTO) SetMetadata(v RunnerMetadata) {
	o.Metadata = &v
}

// GetState returns the State field value
func (o *RunnerDTO) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *RunnerDTO) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *RunnerDTO) SetState(v string) {
	o.State = v
}

func (o RunnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunnerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *RunnerDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"id",
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRunnerDTO := _RunnerDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRunnerDTO)

	if err != nil {
		return err
	}

	*o = RunnerDTO(varRunnerDTO)

	return err
}

type NullableRunnerDTO struct {
	value *RunnerDTO
	isSet bool
}

func (v NullableRunnerDTO) Get() *RunnerDTO {
	return v.value
}

func (v *NullableRunnerDTO) Set(val *RunnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRunnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRunnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunnerDTO(val *RunnerDTO) *NullableRunnerDTO {
	return &NullableRunnerDTO{value: val, isSet: true}
}

func (v NullableRunnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
