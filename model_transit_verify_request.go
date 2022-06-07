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

// TransitVerifyRequest struct for TransitVerifyRequest
type TransitVerifyRequest struct {
	// Deprecated: use \"hash_algorithm\" instead.
	Algorithm *string `json:"algorithm,omitempty"`
	// Base64 encoded context for key derivation. Required if key derivation is enabled; currently only available with ed25519 keys.
	Context *string `json:"context,omitempty"`
	// Hash algorithm to use (POST body parameter). Valid values are: * sha1 * sha2-224 * sha2-256 * sha2-384 * sha2-512 * sha3-224 * sha3-256 * sha3-384 * sha3-512 Defaults to \"sha2-256\". Not valid for all key types.
	HashAlgorithm *string `json:"hash_algorithm,omitempty"`
	// The HMAC, including vault header/key version
	Hmac *string `json:"hmac,omitempty"`
	// The base64-encoded input data to verify
	Input *string `json:"input,omitempty"`
	// The method by which to unmarshal the signature when verifying. The default is 'asn1' which is used by openssl and X.509; can also be set to 'jws' which is used for JWT signatures in which case the signature is also expected to be url-safe base64 encoding instead of standard base64 encoding. Currently only valid for ECDSA P-256 key types\".
	MarshalingAlgorithm *string `json:"marshaling_algorithm,omitempty"`
	// Set to 'true' when the input is already hashed. If the key type is 'rsa-2048', 'rsa-3072' or 'rsa-4096', then the algorithm used to hash the input should be indicated by the 'algorithm' parameter.
	Prehashed *bool `json:"prehashed,omitempty"`
	// The signature, including vault header/key version
	Signature *string `json:"signature,omitempty"`
	// The signature algorithm to use for signature verification. Currently only applies to RSA key types. Options are 'pss' or 'pkcs1v15'. Defaults to 'pss'
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`
	// Hash algorithm to use (POST URL parameter)
	Urlalgorithm *string `json:"urlalgorithm,omitempty"`
}

// NewTransitVerifyRequest instantiates a new TransitVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransitVerifyRequest() *TransitVerifyRequest {
	this := TransitVerifyRequest{}
	var algorithm string = "sha2-256"
	this.Algorithm = &algorithm
	var hashAlgorithm string = "sha2-256"
	this.HashAlgorithm = &hashAlgorithm
	var marshalingAlgorithm string = "asn1"
	this.MarshalingAlgorithm = &marshalingAlgorithm
	return &this
}

// NewTransitVerifyRequestWithDefaults instantiates a new TransitVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransitVerifyRequestWithDefaults() *TransitVerifyRequest {
	this := TransitVerifyRequest{}
	var algorithm string = "sha2-256"
	this.Algorithm = &algorithm
	var hashAlgorithm string = "sha2-256"
	this.HashAlgorithm = &hashAlgorithm
	var marshalingAlgorithm string = "asn1"
	this.MarshalingAlgorithm = &marshalingAlgorithm
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *TransitVerifyRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *TransitVerifyRequest) SetContext(v string) {
	o.Context = &v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetHashAlgorithm() string {
	if o == nil || o.HashAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.HashAlgorithm
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetHashAlgorithmOk() (*string, bool) {
	if o == nil || o.HashAlgorithm == nil {
		return nil, false
	}
	return o.HashAlgorithm, true
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm != nil {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given string and assigns it to the HashAlgorithm field.
func (o *TransitVerifyRequest) SetHashAlgorithm(v string) {
	o.HashAlgorithm = &v
}

// GetHmac returns the Hmac field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetHmac() string {
	if o == nil || o.Hmac == nil {
		var ret string
		return ret
	}
	return *o.Hmac
}

// GetHmacOk returns a tuple with the Hmac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetHmacOk() (*string, bool) {
	if o == nil || o.Hmac == nil {
		return nil, false
	}
	return o.Hmac, true
}

// HasHmac returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasHmac() bool {
	if o != nil && o.Hmac != nil {
		return true
	}

	return false
}

// SetHmac gets a reference to the given string and assigns it to the Hmac field.
func (o *TransitVerifyRequest) SetHmac(v string) {
	o.Hmac = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetInput() string {
	if o == nil || o.Input == nil {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetInputOk() (*string, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *TransitVerifyRequest) SetInput(v string) {
	o.Input = &v
}

// GetMarshalingAlgorithm returns the MarshalingAlgorithm field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetMarshalingAlgorithm() string {
	if o == nil || o.MarshalingAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.MarshalingAlgorithm
}

// GetMarshalingAlgorithmOk returns a tuple with the MarshalingAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetMarshalingAlgorithmOk() (*string, bool) {
	if o == nil || o.MarshalingAlgorithm == nil {
		return nil, false
	}
	return o.MarshalingAlgorithm, true
}

// HasMarshalingAlgorithm returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasMarshalingAlgorithm() bool {
	if o != nil && o.MarshalingAlgorithm != nil {
		return true
	}

	return false
}

// SetMarshalingAlgorithm gets a reference to the given string and assigns it to the MarshalingAlgorithm field.
func (o *TransitVerifyRequest) SetMarshalingAlgorithm(v string) {
	o.MarshalingAlgorithm = &v
}

// GetPrehashed returns the Prehashed field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetPrehashed() bool {
	if o == nil || o.Prehashed == nil {
		var ret bool
		return ret
	}
	return *o.Prehashed
}

// GetPrehashedOk returns a tuple with the Prehashed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetPrehashedOk() (*bool, bool) {
	if o == nil || o.Prehashed == nil {
		return nil, false
	}
	return o.Prehashed, true
}

// HasPrehashed returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasPrehashed() bool {
	if o != nil && o.Prehashed != nil {
		return true
	}

	return false
}

// SetPrehashed gets a reference to the given bool and assigns it to the Prehashed field.
func (o *TransitVerifyRequest) SetPrehashed(v bool) {
	o.Prehashed = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *TransitVerifyRequest) SetSignature(v string) {
	o.Signature = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *TransitVerifyRequest) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetUrlalgorithm returns the Urlalgorithm field value if set, zero value otherwise.
func (o *TransitVerifyRequest) GetUrlalgorithm() string {
	if o == nil || o.Urlalgorithm == nil {
		var ret string
		return ret
	}
	return *o.Urlalgorithm
}

// GetUrlalgorithmOk returns a tuple with the Urlalgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransitVerifyRequest) GetUrlalgorithmOk() (*string, bool) {
	if o == nil || o.Urlalgorithm == nil {
		return nil, false
	}
	return o.Urlalgorithm, true
}

// HasUrlalgorithm returns a boolean if a field has been set.
func (o *TransitVerifyRequest) HasUrlalgorithm() bool {
	if o != nil && o.Urlalgorithm != nil {
		return true
	}

	return false
}

// SetUrlalgorithm gets a reference to the given string and assigns it to the Urlalgorithm field.
func (o *TransitVerifyRequest) SetUrlalgorithm(v string) {
	o.Urlalgorithm = &v
}

func (o TransitVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.HashAlgorithm != nil {
		toSerialize["hash_algorithm"] = o.HashAlgorithm
	}
	if o.Hmac != nil {
		toSerialize["hmac"] = o.Hmac
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.MarshalingAlgorithm != nil {
		toSerialize["marshaling_algorithm"] = o.MarshalingAlgorithm
	}
	if o.Prehashed != nil {
		toSerialize["prehashed"] = o.Prehashed
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.Urlalgorithm != nil {
		toSerialize["urlalgorithm"] = o.Urlalgorithm
	}
	return json.Marshal(toSerialize)
}

type NullableTransitVerifyRequest struct {
	value *TransitVerifyRequest
	isSet bool
}

func (v NullableTransitVerifyRequest) Get() *TransitVerifyRequest {
	return v.value
}

func (v *NullableTransitVerifyRequest) Set(val *TransitVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransitVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransitVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransitVerifyRequest(val *TransitVerifyRequest) *NullableTransitVerifyRequest {
	return &NullableTransitVerifyRequest{value: val, isSet: true}
}

func (v NullableTransitVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransitVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


