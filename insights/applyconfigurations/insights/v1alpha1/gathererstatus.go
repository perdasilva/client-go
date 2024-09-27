// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// GathererStatusApplyConfiguration represents a declarative configuration of the GathererStatus type for use
// with apply.
type GathererStatusApplyConfiguration struct {
	Conditions         []v1.ConditionApplyConfiguration `json:"conditions,omitempty"`
	Name               *string                          `json:"name,omitempty"`
	LastGatherDuration *metav1.Duration                 `json:"lastGatherDuration,omitempty"`
}

// GathererStatusApplyConfiguration constructs a declarative configuration of the GathererStatus type for use with
// apply.
func GathererStatus() *GathererStatusApplyConfiguration {
	return &GathererStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *GathererStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *GathererStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *GathererStatusApplyConfiguration) WithName(value string) *GathererStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithLastGatherDuration sets the LastGatherDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastGatherDuration field is set to the value of the last call.
func (b *GathererStatusApplyConfiguration) WithLastGatherDuration(value metav1.Duration) *GathererStatusApplyConfiguration {
	b.LastGatherDuration = &value
	return b
}
