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

// SystemQuotasRateLimitRequest struct for SystemQuotasRateLimitRequest
type SystemQuotasRateLimitRequest struct {
	// If set, when a client reaches a rate limit threshold, the client will be prohibited from any further requests until after the 'block_interval' has elapsed.
	BlockInterval *int32 `json:"block_interval,omitempty"`
	// The duration to enforce rate limiting for (default '1s').
	Interval *int32 `json:"interval,omitempty"`
	// Path of the mount or namespace to apply the quota. A blank path configures a global quota. For example namespace1/ adds a quota to a full namespace, namespace1/auth/userpass adds a quota to userpass in namespace1.
	Path *string `json:"path,omitempty"`
	// The maximum number of requests in a given interval to be allowed by the quota rule. The 'rate' must be positive.
	Rate *float32 `json:"rate,omitempty"`
	// Type of the quota rule.
	Type *string `json:"type,omitempty"`
}

// NewSystemQuotasRateLimitRequest instantiates a new SystemQuotasRateLimitRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemQuotasRateLimitRequest() *SystemQuotasRateLimitRequest {
	this := SystemQuotasRateLimitRequest{}
	return &this
}

// NewSystemQuotasRateLimitRequestWithDefaults instantiates a new SystemQuotasRateLimitRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemQuotasRateLimitRequestWithDefaults() *SystemQuotasRateLimitRequest {
	this := SystemQuotasRateLimitRequest{}
	return &this
}

// GetBlockInterval returns the BlockInterval field value if set, zero value otherwise.
func (o *SystemQuotasRateLimitRequest) GetBlockInterval() int32 {
	if o == nil || o.BlockInterval == nil {
		var ret int32
		return ret
	}
	return *o.BlockInterval
}

// GetBlockIntervalOk returns a tuple with the BlockInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemQuotasRateLimitRequest) GetBlockIntervalOk() (*int32, bool) {
	if o == nil || o.BlockInterval == nil {
		return nil, false
	}
	return o.BlockInterval, true
}

// HasBlockInterval returns a boolean if a field has been set.
func (o *SystemQuotasRateLimitRequest) HasBlockInterval() bool {
	if o != nil && o.BlockInterval != nil {
		return true
	}

	return false
}

// SetBlockInterval gets a reference to the given int32 and assigns it to the BlockInterval field.
func (o *SystemQuotasRateLimitRequest) SetBlockInterval(v int32) {
	o.BlockInterval = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *SystemQuotasRateLimitRequest) GetInterval() int32 {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemQuotasRateLimitRequest) GetIntervalOk() (*int32, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *SystemQuotasRateLimitRequest) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *SystemQuotasRateLimitRequest) SetInterval(v int32) {
	o.Interval = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SystemQuotasRateLimitRequest) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemQuotasRateLimitRequest) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SystemQuotasRateLimitRequest) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SystemQuotasRateLimitRequest) SetPath(v string) {
	o.Path = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *SystemQuotasRateLimitRequest) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemQuotasRateLimitRequest) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *SystemQuotasRateLimitRequest) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *SystemQuotasRateLimitRequest) SetRate(v float32) {
	o.Rate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SystemQuotasRateLimitRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemQuotasRateLimitRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SystemQuotasRateLimitRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SystemQuotasRateLimitRequest) SetType(v string) {
	o.Type = &v
}

func (o SystemQuotasRateLimitRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockInterval != nil {
		toSerialize["block_interval"] = o.BlockInterval
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSystemQuotasRateLimitRequest struct {
	value *SystemQuotasRateLimitRequest
	isSet bool
}

func (v NullableSystemQuotasRateLimitRequest) Get() *SystemQuotasRateLimitRequest {
	return v.value
}

func (v *NullableSystemQuotasRateLimitRequest) Set(val *SystemQuotasRateLimitRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemQuotasRateLimitRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemQuotasRateLimitRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemQuotasRateLimitRequest(val *SystemQuotasRateLimitRequest) *NullableSystemQuotasRateLimitRequest {
	return &NullableSystemQuotasRateLimitRequest{value: val, isSet: true}
}

func (v NullableSystemQuotasRateLimitRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemQuotasRateLimitRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


