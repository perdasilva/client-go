// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// LifecycleHooksApplyConfiguration represents a declarative configuration of the LifecycleHooks type for use
// with apply.
type LifecycleHooksApplyConfiguration struct {
	PreDrain     []LifecycleHookApplyConfiguration `json:"preDrain,omitempty"`
	PreTerminate []LifecycleHookApplyConfiguration `json:"preTerminate,omitempty"`
}

// LifecycleHooksApplyConfiguration constructs a declarative configuration of the LifecycleHooks type for use with
// apply.
func LifecycleHooks() *LifecycleHooksApplyConfiguration {
	return &LifecycleHooksApplyConfiguration{}
}

// WithPreDrain adds the given value to the PreDrain field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PreDrain field.
func (b *LifecycleHooksApplyConfiguration) WithPreDrain(values ...*LifecycleHookApplyConfiguration) *LifecycleHooksApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPreDrain")
		}
		b.PreDrain = append(b.PreDrain, *values[i])
	}
	return b
}

// WithPreTerminate adds the given value to the PreTerminate field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PreTerminate field.
func (b *LifecycleHooksApplyConfiguration) WithPreTerminate(values ...*LifecycleHookApplyConfiguration) *LifecycleHooksApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPreTerminate")
		}
		b.PreTerminate = append(b.PreTerminate, *values[i])
	}
	return b
}
