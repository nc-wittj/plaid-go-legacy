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

// AccountAssetsAllOf struct for AccountAssetsAllOf
type AccountAssetsAllOf struct {
	// The duration of transaction history available for this Item, typically defined as the time since the date of the earliest transaction in that account. Only returned by Assets endpoints.
	DaysAvailable float32 `json:"days_available"`
	// Transaction history associated with the account. Only returned by Assets endpoints. Transaction history returned by endpoints such as `/transactions/get` or `/investments/transactions/get` will be returned in the top-level `transactions` field instead.
	Transactions []AssetReportTransaction `json:"transactions"`
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array.
	Owners []Owner `json:"owners"`
	// Calculated data about the historical balances on the account. Only returned by Assets endpoints.
	HistoricalBalances []HistoricalBalance `json:"historical_balances"`
}

// NewAccountAssetsAllOf instantiates a new AccountAssetsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAssetsAllOf(daysAvailable float32, transactions []AssetReportTransaction, owners []Owner, historicalBalances []HistoricalBalance) *AccountAssetsAllOf {
	this := AccountAssetsAllOf{}
	this.DaysAvailable = daysAvailable
	this.Transactions = transactions
	this.Owners = owners
	this.HistoricalBalances = historicalBalances
	return &this
}

// NewAccountAssetsAllOfWithDefaults instantiates a new AccountAssetsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAssetsAllOfWithDefaults() *AccountAssetsAllOf {
	this := AccountAssetsAllOf{}
	return &this
}

// GetDaysAvailable returns the DaysAvailable field value
func (o *AccountAssetsAllOf) GetDaysAvailable() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DaysAvailable
}

// GetDaysAvailableOk returns a tuple with the DaysAvailable field value
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetDaysAvailableOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysAvailable, true
}

// SetDaysAvailable sets field value
func (o *AccountAssetsAllOf) SetDaysAvailable(v float32) {
	o.DaysAvailable = v
}

// GetTransactions returns the Transactions field value
func (o *AccountAssetsAllOf) GetTransactions() []AssetReportTransaction {
	if o == nil {
		var ret []AssetReportTransaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetTransactionsOk() (*[]AssetReportTransaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *AccountAssetsAllOf) SetTransactions(v []AssetReportTransaction) {
	o.Transactions = v
}

// GetOwners returns the Owners field value
func (o *AccountAssetsAllOf) GetOwners() []Owner {
	if o == nil {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *AccountAssetsAllOf) SetOwners(v []Owner) {
	o.Owners = v
}

// GetHistoricalBalances returns the HistoricalBalances field value
func (o *AccountAssetsAllOf) GetHistoricalBalances() []HistoricalBalance {
	if o == nil {
		var ret []HistoricalBalance
		return ret
	}

	return o.HistoricalBalances
}

// GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field value
// and a boolean to check if the value has been set.
func (o *AccountAssetsAllOf) GetHistoricalBalancesOk() (*[]HistoricalBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HistoricalBalances, true
}

// SetHistoricalBalances sets field value
func (o *AccountAssetsAllOf) SetHistoricalBalances(v []HistoricalBalance) {
	o.HistoricalBalances = v
}

func (o AccountAssetsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days_available"] = o.DaysAvailable
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	if true {
		toSerialize["owners"] = o.Owners
	}
	if true {
		toSerialize["historical_balances"] = o.HistoricalBalances
	}
	return json.Marshal(toSerialize)
}

type NullableAccountAssetsAllOf struct {
	value *AccountAssetsAllOf
	isSet bool
}

func (v NullableAccountAssetsAllOf) Get() *AccountAssetsAllOf {
	return v.value
}

func (v *NullableAccountAssetsAllOf) Set(val *AccountAssetsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAssetsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAssetsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAssetsAllOf(val *AccountAssetsAllOf) *NullableAccountAssetsAllOf {
	return &NullableAccountAssetsAllOf{value: val, isSet: true}
}

func (v NullableAccountAssetsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAssetsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


