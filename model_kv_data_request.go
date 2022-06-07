/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vault

import (
	"encoding/json"
)

// KvDataRequest struct for KvDataRequest
type KvDataRequest struct {
	// The contents of the data map will be stored and returned on read.
	Data map[string]interface{} `json:"data,omitempty"`
	// Options for writing a KV entry. Set the \"cas\" value to use a Check-And-Set operation. If not set the write will be allowed. If set to 0 a write will only be allowed if the key doesn’t exist. If the index is non-zero the write will only be allowed if the key’s current version matches the version specified in the cas parameter.
	Options map[string]interface{} `json:"options,omitempty"`
	// If provided during a read, the value at the version number will be returned
	Version *int32 `json:"version,omitempty"`
}

// NewKvDataRequest instantiates a new KvDataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvDataRequest() *KvDataRequest {
	this := KvDataRequest{}
	return &this
}

// NewKvDataRequestWithDefaults instantiates a new KvDataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvDataRequestWithDefaults() *KvDataRequest {
	this := KvDataRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KvDataRequest) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvDataRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KvDataRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *KvDataRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *KvDataRequest) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvDataRequest) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *KvDataRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *KvDataRequest) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KvDataRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvDataRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KvDataRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *KvDataRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o KvDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableKvDataRequest struct {
	value *KvDataRequest
	isSet bool
}

func (v NullableKvDataRequest) Get() *KvDataRequest {
	return v.value
}

func (v *NullableKvDataRequest) Set(val *KvDataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableKvDataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableKvDataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvDataRequest(val *KvDataRequest) *NullableKvDataRequest {
	return &NullableKvDataRequest{value: val, isSet: true}
}

func (v NullableKvDataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvDataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


