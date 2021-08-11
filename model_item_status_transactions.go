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
	"time"
)

// ItemStatusTransactions Information about the last successful and failed transactions update for the Item.
type ItemStatusTransactions struct {
	// ISO 8601 timestamp of the last successful transactions update for the Item. The status will update each time Plaid successfully connects with the institution, regardless of whether any new data is available in the update.
	LastSuccessfulUpdate NullableTime `json:"last_successful_update,omitempty"`
	// ISO 8601 timestamp of the last failed transactions update for the Item. The status will update each time Plaid fails an attempt to connect with the institution, regardless of whether any new data is available in the update.
	LastFailedUpdate NullableTime `json:"last_failed_update,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemStatusTransactions ItemStatusTransactions

// NewItemStatusTransactions instantiates a new ItemStatusTransactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStatusTransactions() *ItemStatusTransactions {
	this := ItemStatusTransactions{}
	return &this
}

// NewItemStatusTransactionsWithDefaults instantiates a new ItemStatusTransactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStatusTransactionsWithDefaults() *ItemStatusTransactions {
	this := ItemStatusTransactions{}
	return &this
}

// GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusTransactions) GetLastSuccessfulUpdate() time.Time {
	if o == nil || o.LastSuccessfulUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSuccessfulUpdate.Get()
}

// GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusTransactions) GetLastSuccessfulUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastSuccessfulUpdate.Get(), o.LastSuccessfulUpdate.IsSet()
}

// HasLastSuccessfulUpdate returns a boolean if a field has been set.
func (o *ItemStatusTransactions) HasLastSuccessfulUpdate() bool {
	if o != nil && o.LastSuccessfulUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastSuccessfulUpdate gets a reference to the given NullableTime and assigns it to the LastSuccessfulUpdate field.
func (o *ItemStatusTransactions) SetLastSuccessfulUpdate(v time.Time) {
	o.LastSuccessfulUpdate.Set(&v)
}
// SetLastSuccessfulUpdateNil sets the value for LastSuccessfulUpdate to be an explicit nil
func (o *ItemStatusTransactions) SetLastSuccessfulUpdateNil() {
	o.LastSuccessfulUpdate.Set(nil)
}

// UnsetLastSuccessfulUpdate ensures that no value is present for LastSuccessfulUpdate, not even an explicit nil
func (o *ItemStatusTransactions) UnsetLastSuccessfulUpdate() {
	o.LastSuccessfulUpdate.Unset()
}

// GetLastFailedUpdate returns the LastFailedUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ItemStatusTransactions) GetLastFailedUpdate() time.Time {
	if o == nil || o.LastFailedUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastFailedUpdate.Get()
}

// GetLastFailedUpdateOk returns a tuple with the LastFailedUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemStatusTransactions) GetLastFailedUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastFailedUpdate.Get(), o.LastFailedUpdate.IsSet()
}

// HasLastFailedUpdate returns a boolean if a field has been set.
func (o *ItemStatusTransactions) HasLastFailedUpdate() bool {
	if o != nil && o.LastFailedUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastFailedUpdate gets a reference to the given NullableTime and assigns it to the LastFailedUpdate field.
func (o *ItemStatusTransactions) SetLastFailedUpdate(v time.Time) {
	o.LastFailedUpdate.Set(&v)
}
// SetLastFailedUpdateNil sets the value for LastFailedUpdate to be an explicit nil
func (o *ItemStatusTransactions) SetLastFailedUpdateNil() {
	o.LastFailedUpdate.Set(nil)
}

// UnsetLastFailedUpdate ensures that no value is present for LastFailedUpdate, not even an explicit nil
func (o *ItemStatusTransactions) UnsetLastFailedUpdate() {
	o.LastFailedUpdate.Unset()
}

func (o ItemStatusTransactions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSuccessfulUpdate.IsSet() {
		toSerialize["last_successful_update"] = o.LastSuccessfulUpdate.Get()
	}
	if o.LastFailedUpdate.IsSet() {
		toSerialize["last_failed_update"] = o.LastFailedUpdate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemStatusTransactions) UnmarshalJSON(bytes []byte) (err error) {
	varItemStatusTransactions := _ItemStatusTransactions{}

	if err = json.Unmarshal(bytes, &varItemStatusTransactions); err == nil {
		*o = ItemStatusTransactions(varItemStatusTransactions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "last_successful_update")
		delete(additionalProperties, "last_failed_update")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemStatusTransactions struct {
	value *ItemStatusTransactions
	isSet bool
}

func (v NullableItemStatusTransactions) Get() *ItemStatusTransactions {
	return v.value
}

func (v *NullableItemStatusTransactions) Set(val *ItemStatusTransactions) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStatusTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStatusTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStatusTransactions(val *ItemStatusTransactions) *NullableItemStatusTransactions {
	return &NullableItemStatusTransactions{value: val, isSet: true}
}

func (v NullableItemStatusTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStatusTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


