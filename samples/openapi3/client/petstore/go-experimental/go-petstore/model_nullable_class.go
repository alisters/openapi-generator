/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
	"encoding/json"
)
// NullableClass struct for NullableClass
type NullableClass struct {
	IntegerProp *int32 `json:"integer_prop,omitempty"`
	isExplicitNullIntegerProp bool `json:"-"`
	NumberProp *float32 `json:"number_prop,omitempty"`
	isExplicitNullNumberProp bool `json:"-"`
	BooleanProp *bool `json:"boolean_prop,omitempty"`
	isExplicitNullBooleanProp bool `json:"-"`
	StringProp *string `json:"string_prop,omitempty"`
	isExplicitNullStringProp bool `json:"-"`
	DateProp *string `json:"date_prop,omitempty"`
	isExplicitNullDateProp bool `json:"-"`
	DatetimeProp *time.Time `json:"datetime_prop,omitempty"`
	isExplicitNullDatetimeProp bool `json:"-"`
	ArrayNullableProp *[]map[string]interface{} `json:"array_nullable_prop,omitempty"`
	isExplicitNullArrayNullableProp bool `json:"-"`
	ArrayAndItemsNullableProp *[]map[string]interface{} `json:"array_and_items_nullable_prop,omitempty"`
	isExplicitNullArrayAndItemsNullableProp bool `json:"-"`
	ArrayItemsNullable *[]map[string]interface{} `json:"array_items_nullable,omitempty"`

	ObjectNullableProp *map[string]map[string]interface{} `json:"object_nullable_prop,omitempty"`
	isExplicitNullObjectNullableProp bool `json:"-"`
	ObjectAndItemsNullableProp *map[string]map[string]interface{} `json:"object_and_items_nullable_prop,omitempty"`
	isExplicitNullObjectAndItemsNullableProp bool `json:"-"`
	ObjectItemsNullable *map[string]map[string]interface{} `json:"object_items_nullable,omitempty"`

}

// GetIntegerProp returns the IntegerProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetIntegerProp() int32 {
	if o == nil || o.IntegerProp == nil {
		var ret int32
		return ret
	}
	return *o.IntegerProp
}

// GetIntegerPropOk returns a tuple with the IntegerProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetIntegerPropOk() (int32, bool) {
	if o == nil || o.IntegerProp == nil {
		var ret int32
		return ret, false
	}
	return *o.IntegerProp, true
}

// HasIntegerProp returns a boolean if a field has been set.
func (o *NullableClass) HasIntegerProp() bool {
	if o != nil && o.IntegerProp != nil {
		return true
	}

	return false
}

// SetIntegerProp gets a reference to the given int32 and assigns it to the IntegerProp field.
func (o *NullableClass) SetIntegerProp(v int32) {
	o.IntegerProp = &v
}

// SetIntegerPropExplicitNull (un)sets IntegerProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The IntegerProp value is set to nil even if false is passed
func (o *NullableClass) SetIntegerPropExplicitNull(b bool) {
	o.IntegerProp = nil
	o.isExplicitNullIntegerProp = b
}
// GetNumberProp returns the NumberProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetNumberProp() float32 {
	if o == nil || o.NumberProp == nil {
		var ret float32
		return ret
	}
	return *o.NumberProp
}

// GetNumberPropOk returns a tuple with the NumberProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetNumberPropOk() (float32, bool) {
	if o == nil || o.NumberProp == nil {
		var ret float32
		return ret, false
	}
	return *o.NumberProp, true
}

// HasNumberProp returns a boolean if a field has been set.
func (o *NullableClass) HasNumberProp() bool {
	if o != nil && o.NumberProp != nil {
		return true
	}

	return false
}

// SetNumberProp gets a reference to the given float32 and assigns it to the NumberProp field.
func (o *NullableClass) SetNumberProp(v float32) {
	o.NumberProp = &v
}

// SetNumberPropExplicitNull (un)sets NumberProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The NumberProp value is set to nil even if false is passed
func (o *NullableClass) SetNumberPropExplicitNull(b bool) {
	o.NumberProp = nil
	o.isExplicitNullNumberProp = b
}
// GetBooleanProp returns the BooleanProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetBooleanProp() bool {
	if o == nil || o.BooleanProp == nil {
		var ret bool
		return ret
	}
	return *o.BooleanProp
}

// GetBooleanPropOk returns a tuple with the BooleanProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetBooleanPropOk() (bool, bool) {
	if o == nil || o.BooleanProp == nil {
		var ret bool
		return ret, false
	}
	return *o.BooleanProp, true
}

// HasBooleanProp returns a boolean if a field has been set.
func (o *NullableClass) HasBooleanProp() bool {
	if o != nil && o.BooleanProp != nil {
		return true
	}

	return false
}

// SetBooleanProp gets a reference to the given bool and assigns it to the BooleanProp field.
func (o *NullableClass) SetBooleanProp(v bool) {
	o.BooleanProp = &v
}

// SetBooleanPropExplicitNull (un)sets BooleanProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The BooleanProp value is set to nil even if false is passed
func (o *NullableClass) SetBooleanPropExplicitNull(b bool) {
	o.BooleanProp = nil
	o.isExplicitNullBooleanProp = b
}
// GetStringProp returns the StringProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetStringProp() string {
	if o == nil || o.StringProp == nil {
		var ret string
		return ret
	}
	return *o.StringProp
}

// GetStringPropOk returns a tuple with the StringProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetStringPropOk() (string, bool) {
	if o == nil || o.StringProp == nil {
		var ret string
		return ret, false
	}
	return *o.StringProp, true
}

// HasStringProp returns a boolean if a field has been set.
func (o *NullableClass) HasStringProp() bool {
	if o != nil && o.StringProp != nil {
		return true
	}

	return false
}

// SetStringProp gets a reference to the given string and assigns it to the StringProp field.
func (o *NullableClass) SetStringProp(v string) {
	o.StringProp = &v
}

// SetStringPropExplicitNull (un)sets StringProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The StringProp value is set to nil even if false is passed
func (o *NullableClass) SetStringPropExplicitNull(b bool) {
	o.StringProp = nil
	o.isExplicitNullStringProp = b
}
// GetDateProp returns the DateProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetDateProp() string {
	if o == nil || o.DateProp == nil {
		var ret string
		return ret
	}
	return *o.DateProp
}

// GetDatePropOk returns a tuple with the DateProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetDatePropOk() (string, bool) {
	if o == nil || o.DateProp == nil {
		var ret string
		return ret, false
	}
	return *o.DateProp, true
}

// HasDateProp returns a boolean if a field has been set.
func (o *NullableClass) HasDateProp() bool {
	if o != nil && o.DateProp != nil {
		return true
	}

	return false
}

// SetDateProp gets a reference to the given string and assigns it to the DateProp field.
func (o *NullableClass) SetDateProp(v string) {
	o.DateProp = &v
}

// SetDatePropExplicitNull (un)sets DateProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DateProp value is set to nil even if false is passed
func (o *NullableClass) SetDatePropExplicitNull(b bool) {
	o.DateProp = nil
	o.isExplicitNullDateProp = b
}
// GetDatetimeProp returns the DatetimeProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetDatetimeProp() time.Time {
	if o == nil || o.DatetimeProp == nil {
		var ret time.Time
		return ret
	}
	return *o.DatetimeProp
}

// GetDatetimePropOk returns a tuple with the DatetimeProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetDatetimePropOk() (time.Time, bool) {
	if o == nil || o.DatetimeProp == nil {
		var ret time.Time
		return ret, false
	}
	return *o.DatetimeProp, true
}

// HasDatetimeProp returns a boolean if a field has been set.
func (o *NullableClass) HasDatetimeProp() bool {
	if o != nil && o.DatetimeProp != nil {
		return true
	}

	return false
}

// SetDatetimeProp gets a reference to the given time.Time and assigns it to the DatetimeProp field.
func (o *NullableClass) SetDatetimeProp(v time.Time) {
	o.DatetimeProp = &v
}

// SetDatetimePropExplicitNull (un)sets DatetimeProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The DatetimeProp value is set to nil even if false is passed
func (o *NullableClass) SetDatetimePropExplicitNull(b bool) {
	o.DatetimeProp = nil
	o.isExplicitNullDatetimeProp = b
}
// GetArrayNullableProp returns the ArrayNullableProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetArrayNullableProp() []map[string]interface{} {
	if o == nil || o.ArrayNullableProp == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ArrayNullableProp
}

// GetArrayNullablePropOk returns a tuple with the ArrayNullableProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetArrayNullablePropOk() ([]map[string]interface{}, bool) {
	if o == nil || o.ArrayNullableProp == nil {
		var ret []map[string]interface{}
		return ret, false
	}
	return *o.ArrayNullableProp, true
}

// HasArrayNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasArrayNullableProp() bool {
	if o != nil && o.ArrayNullableProp != nil {
		return true
	}

	return false
}

// SetArrayNullableProp gets a reference to the given []map[string]interface{} and assigns it to the ArrayNullableProp field.
func (o *NullableClass) SetArrayNullableProp(v []map[string]interface{}) {
	o.ArrayNullableProp = &v
}

// SetArrayNullablePropExplicitNull (un)sets ArrayNullableProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ArrayNullableProp value is set to nil even if false is passed
func (o *NullableClass) SetArrayNullablePropExplicitNull(b bool) {
	o.ArrayNullableProp = nil
	o.isExplicitNullArrayNullableProp = b
}
// GetArrayAndItemsNullableProp returns the ArrayAndItemsNullableProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetArrayAndItemsNullableProp() []map[string]interface{} {
	if o == nil || o.ArrayAndItemsNullableProp == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ArrayAndItemsNullableProp
}

// GetArrayAndItemsNullablePropOk returns a tuple with the ArrayAndItemsNullableProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetArrayAndItemsNullablePropOk() ([]map[string]interface{}, bool) {
	if o == nil || o.ArrayAndItemsNullableProp == nil {
		var ret []map[string]interface{}
		return ret, false
	}
	return *o.ArrayAndItemsNullableProp, true
}

// HasArrayAndItemsNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasArrayAndItemsNullableProp() bool {
	if o != nil && o.ArrayAndItemsNullableProp != nil {
		return true
	}

	return false
}

// SetArrayAndItemsNullableProp gets a reference to the given []map[string]interface{} and assigns it to the ArrayAndItemsNullableProp field.
func (o *NullableClass) SetArrayAndItemsNullableProp(v []map[string]interface{}) {
	o.ArrayAndItemsNullableProp = &v
}

// SetArrayAndItemsNullablePropExplicitNull (un)sets ArrayAndItemsNullableProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ArrayAndItemsNullableProp value is set to nil even if false is passed
func (o *NullableClass) SetArrayAndItemsNullablePropExplicitNull(b bool) {
	o.ArrayAndItemsNullableProp = nil
	o.isExplicitNullArrayAndItemsNullableProp = b
}
// GetArrayItemsNullable returns the ArrayItemsNullable field if non-nil, zero value otherwise.
func (o *NullableClass) GetArrayItemsNullable() []map[string]interface{} {
	if o == nil || o.ArrayItemsNullable == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.ArrayItemsNullable
}

// GetArrayItemsNullableOk returns a tuple with the ArrayItemsNullable field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetArrayItemsNullableOk() ([]map[string]interface{}, bool) {
	if o == nil || o.ArrayItemsNullable == nil {
		var ret []map[string]interface{}
		return ret, false
	}
	return *o.ArrayItemsNullable, true
}

// HasArrayItemsNullable returns a boolean if a field has been set.
func (o *NullableClass) HasArrayItemsNullable() bool {
	if o != nil && o.ArrayItemsNullable != nil {
		return true
	}

	return false
}

// SetArrayItemsNullable gets a reference to the given []map[string]interface{} and assigns it to the ArrayItemsNullable field.
func (o *NullableClass) SetArrayItemsNullable(v []map[string]interface{}) {
	o.ArrayItemsNullable = &v
}

// GetObjectNullableProp returns the ObjectNullableProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetObjectNullableProp() map[string]map[string]interface{} {
	if o == nil || o.ObjectNullableProp == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.ObjectNullableProp
}

// GetObjectNullablePropOk returns a tuple with the ObjectNullableProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetObjectNullablePropOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectNullableProp == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.ObjectNullableProp, true
}

// HasObjectNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasObjectNullableProp() bool {
	if o != nil && o.ObjectNullableProp != nil {
		return true
	}

	return false
}

// SetObjectNullableProp gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectNullableProp field.
func (o *NullableClass) SetObjectNullableProp(v map[string]map[string]interface{}) {
	o.ObjectNullableProp = &v
}

// SetObjectNullablePropExplicitNull (un)sets ObjectNullableProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ObjectNullableProp value is set to nil even if false is passed
func (o *NullableClass) SetObjectNullablePropExplicitNull(b bool) {
	o.ObjectNullableProp = nil
	o.isExplicitNullObjectNullableProp = b
}
// GetObjectAndItemsNullableProp returns the ObjectAndItemsNullableProp field if non-nil, zero value otherwise.
func (o *NullableClass) GetObjectAndItemsNullableProp() map[string]map[string]interface{} {
	if o == nil || o.ObjectAndItemsNullableProp == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.ObjectAndItemsNullableProp
}

// GetObjectAndItemsNullablePropOk returns a tuple with the ObjectAndItemsNullableProp field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetObjectAndItemsNullablePropOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectAndItemsNullableProp == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.ObjectAndItemsNullableProp, true
}

// HasObjectAndItemsNullableProp returns a boolean if a field has been set.
func (o *NullableClass) HasObjectAndItemsNullableProp() bool {
	if o != nil && o.ObjectAndItemsNullableProp != nil {
		return true
	}

	return false
}

// SetObjectAndItemsNullableProp gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectAndItemsNullableProp field.
func (o *NullableClass) SetObjectAndItemsNullableProp(v map[string]map[string]interface{}) {
	o.ObjectAndItemsNullableProp = &v
}

// SetObjectAndItemsNullablePropExplicitNull (un)sets ObjectAndItemsNullableProp to be considered as explicit "null" value
// when serializing to JSON (pass true as argument to set this, false to unset)
// The ObjectAndItemsNullableProp value is set to nil even if false is passed
func (o *NullableClass) SetObjectAndItemsNullablePropExplicitNull(b bool) {
	o.ObjectAndItemsNullableProp = nil
	o.isExplicitNullObjectAndItemsNullableProp = b
}
// GetObjectItemsNullable returns the ObjectItemsNullable field if non-nil, zero value otherwise.
func (o *NullableClass) GetObjectItemsNullable() map[string]map[string]interface{} {
	if o == nil || o.ObjectItemsNullable == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.ObjectItemsNullable
}

// GetObjectItemsNullableOk returns a tuple with the ObjectItemsNullable field if it's non-nil, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NullableClass) GetObjectItemsNullableOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.ObjectItemsNullable == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.ObjectItemsNullable, true
}

// HasObjectItemsNullable returns a boolean if a field has been set.
func (o *NullableClass) HasObjectItemsNullable() bool {
	if o != nil && o.ObjectItemsNullable != nil {
		return true
	}

	return false
}

// SetObjectItemsNullable gets a reference to the given map[string]map[string]interface{} and assigns it to the ObjectItemsNullable field.
func (o *NullableClass) SetObjectItemsNullable(v map[string]map[string]interface{}) {
	o.ObjectItemsNullable = &v
}


// MarshalJSON returns the JSON representation of the model.
func (o NullableClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntegerProp == nil {
		if o.isExplicitNullIntegerProp {
			toSerialize["integer_prop"] = o.IntegerProp
		}
	} else {
		toSerialize["integer_prop"] = o.IntegerProp
	}
	if o.NumberProp == nil {
		if o.isExplicitNullNumberProp {
			toSerialize["number_prop"] = o.NumberProp
		}
	} else {
		toSerialize["number_prop"] = o.NumberProp
	}
	if o.BooleanProp == nil {
		if o.isExplicitNullBooleanProp {
			toSerialize["boolean_prop"] = o.BooleanProp
		}
	} else {
		toSerialize["boolean_prop"] = o.BooleanProp
	}
	if o.StringProp == nil {
		if o.isExplicitNullStringProp {
			toSerialize["string_prop"] = o.StringProp
		}
	} else {
		toSerialize["string_prop"] = o.StringProp
	}
	if o.DateProp == nil {
		if o.isExplicitNullDateProp {
			toSerialize["date_prop"] = o.DateProp
		}
	} else {
		toSerialize["date_prop"] = o.DateProp
	}
	if o.DatetimeProp == nil {
		if o.isExplicitNullDatetimeProp {
			toSerialize["datetime_prop"] = o.DatetimeProp
		}
	} else {
		toSerialize["datetime_prop"] = o.DatetimeProp
	}
	if o.ArrayNullableProp == nil {
		if o.isExplicitNullArrayNullableProp {
			toSerialize["array_nullable_prop"] = o.ArrayNullableProp
		}
	} else {
		toSerialize["array_nullable_prop"] = o.ArrayNullableProp
	}
	if o.ArrayAndItemsNullableProp == nil {
		if o.isExplicitNullArrayAndItemsNullableProp {
			toSerialize["array_and_items_nullable_prop"] = o.ArrayAndItemsNullableProp
		}
	} else {
		toSerialize["array_and_items_nullable_prop"] = o.ArrayAndItemsNullableProp
	}
	if o.ArrayItemsNullable != nil {
		toSerialize["array_items_nullable"] = o.ArrayItemsNullable
	}
	if o.ObjectNullableProp == nil {
		if o.isExplicitNullObjectNullableProp {
			toSerialize["object_nullable_prop"] = o.ObjectNullableProp
		}
	} else {
		toSerialize["object_nullable_prop"] = o.ObjectNullableProp
	}
	if o.ObjectAndItemsNullableProp == nil {
		if o.isExplicitNullObjectAndItemsNullableProp {
			toSerialize["object_and_items_nullable_prop"] = o.ObjectAndItemsNullableProp
		}
	} else {
		toSerialize["object_and_items_nullable_prop"] = o.ObjectAndItemsNullableProp
	}
	if o.ObjectItemsNullable != nil {
		toSerialize["object_items_nullable"] = o.ObjectItemsNullable
	}
	return json.Marshal(toSerialize)
}


