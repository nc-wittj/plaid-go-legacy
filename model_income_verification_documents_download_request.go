/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.20.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// IncomeVerificationDocumentsDownloadRequest IncomeVerificationDocumentsDownloadRequest defines the request schema for `/income/verification/documents/download`.
type IncomeVerificationDocumentsDownloadRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The ID of the verification.
	IncomeVerificationId string `json:"income_verification_id"`
}

// NewIncomeVerificationDocumentsDownloadRequest instantiates a new IncomeVerificationDocumentsDownloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeVerificationDocumentsDownloadRequest(incomeVerificationId string) *IncomeVerificationDocumentsDownloadRequest {
	this := IncomeVerificationDocumentsDownloadRequest{}
	this.IncomeVerificationId = incomeVerificationId
	return &this
}

// NewIncomeVerificationDocumentsDownloadRequestWithDefaults instantiates a new IncomeVerificationDocumentsDownloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeVerificationDocumentsDownloadRequestWithDefaults() *IncomeVerificationDocumentsDownloadRequest {
	this := IncomeVerificationDocumentsDownloadRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *IncomeVerificationDocumentsDownloadRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IncomeVerificationDocumentsDownloadRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IncomeVerificationDocumentsDownloadRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetIncomeVerificationId returns the IncomeVerificationId field value
func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IncomeVerificationId
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value
// and a boolean to check if the value has been set.
func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeVerificationId, true
}

// SetIncomeVerificationId sets field value
func (o *IncomeVerificationDocumentsDownloadRequest) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId = v
}

func (o IncomeVerificationDocumentsDownloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["income_verification_id"] = o.IncomeVerificationId
	}
	return json.Marshal(toSerialize)
}

type NullableIncomeVerificationDocumentsDownloadRequest struct {
	value *IncomeVerificationDocumentsDownloadRequest
	isSet bool
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) Get() *IncomeVerificationDocumentsDownloadRequest {
	return v.value
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) Set(val *IncomeVerificationDocumentsDownloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeVerificationDocumentsDownloadRequest(val *IncomeVerificationDocumentsDownloadRequest) *NullableIncomeVerificationDocumentsDownloadRequest {
	return &NullableIncomeVerificationDocumentsDownloadRequest{value: val, isSet: true}
}

func (v NullableIncomeVerificationDocumentsDownloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeVerificationDocumentsDownloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


