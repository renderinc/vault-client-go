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

// SystemRawRequest struct for SystemRawRequest
type SystemRawRequest struct {
	Compressed *bool `json:"compressed,omitempty"`
	CompressionType *string `json:"compression_type,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewSystemRawRequest instantiates a new SystemRawRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemRawRequest() *SystemRawRequest {
	this := SystemRawRequest{}
	return &this
}

// NewSystemRawRequestWithDefaults instantiates a new SystemRawRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemRawRequestWithDefaults() *SystemRawRequest {
	this := SystemRawRequest{}
	return &this
}

// GetCompressed returns the Compressed field value if set, zero value otherwise.
func (o *SystemRawRequest) GetCompressed() bool {
	if o == nil || o.Compressed == nil {
		var ret bool
		return ret
	}
	return *o.Compressed
}

// GetCompressedOk returns a tuple with the Compressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRawRequest) GetCompressedOk() (*bool, bool) {
	if o == nil || o.Compressed == nil {
		return nil, false
	}
	return o.Compressed, true
}

// HasCompressed returns a boolean if a field has been set.
func (o *SystemRawRequest) HasCompressed() bool {
	if o != nil && o.Compressed != nil {
		return true
	}

	return false
}

// SetCompressed gets a reference to the given bool and assigns it to the Compressed field.
func (o *SystemRawRequest) SetCompressed(v bool) {
	o.Compressed = &v
}

// GetCompressionType returns the CompressionType field value if set, zero value otherwise.
func (o *SystemRawRequest) GetCompressionType() string {
	if o == nil || o.CompressionType == nil {
		var ret string
		return ret
	}
	return *o.CompressionType
}

// GetCompressionTypeOk returns a tuple with the CompressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRawRequest) GetCompressionTypeOk() (*string, bool) {
	if o == nil || o.CompressionType == nil {
		return nil, false
	}
	return o.CompressionType, true
}

// HasCompressionType returns a boolean if a field has been set.
func (o *SystemRawRequest) HasCompressionType() bool {
	if o != nil && o.CompressionType != nil {
		return true
	}

	return false
}

// SetCompressionType gets a reference to the given string and assigns it to the CompressionType field.
func (o *SystemRawRequest) SetCompressionType(v string) {
	o.CompressionType = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *SystemRawRequest) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRawRequest) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *SystemRawRequest) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *SystemRawRequest) SetEncoding(v string) {
	o.Encoding = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SystemRawRequest) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemRawRequest) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SystemRawRequest) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SystemRawRequest) SetValue(v string) {
	o.Value = &v
}

func (o SystemRawRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Compressed != nil {
		toSerialize["compressed"] = o.Compressed
	}
	if o.CompressionType != nil {
		toSerialize["compression_type"] = o.CompressionType
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableSystemRawRequest struct {
	value *SystemRawRequest
	isSet bool
}

func (v NullableSystemRawRequest) Get() *SystemRawRequest {
	return v.value
}

func (v *NullableSystemRawRequest) Set(val *SystemRawRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemRawRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemRawRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemRawRequest(val *SystemRawRequest) *NullableSystemRawRequest {
	return &NullableSystemRawRequest{value: val, isSet: true}
}

func (v NullableSystemRawRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemRawRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


