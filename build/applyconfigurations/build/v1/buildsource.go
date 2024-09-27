// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/build/v1"
	corev1 "k8s.io/api/core/v1"
)

// BuildSourceApplyConfiguration represents a declarative configuration of the BuildSource type for use
// with apply.
type BuildSourceApplyConfiguration struct {
	Type         *v1.BuildSourceType                      `json:"type,omitempty"`
	Binary       *BinaryBuildSourceApplyConfiguration     `json:"binary,omitempty"`
	Dockerfile   *string                                  `json:"dockerfile,omitempty"`
	Git          *GitBuildSourceApplyConfiguration        `json:"git,omitempty"`
	Images       []ImageSourceApplyConfiguration          `json:"images,omitempty"`
	ContextDir   *string                                  `json:"contextDir,omitempty"`
	SourceSecret *corev1.LocalObjectReference             `json:"sourceSecret,omitempty"`
	Secrets      []SecretBuildSourceApplyConfiguration    `json:"secrets,omitempty"`
	ConfigMaps   []ConfigMapBuildSourceApplyConfiguration `json:"configMaps,omitempty"`
}

// BuildSourceApplyConfiguration constructs a declarative configuration of the BuildSource type for use with
// apply.
func BuildSource() *BuildSourceApplyConfiguration {
	return &BuildSourceApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithType(value v1.BuildSourceType) *BuildSourceApplyConfiguration {
	b.Type = &value
	return b
}

// WithBinary sets the Binary field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Binary field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithBinary(value *BinaryBuildSourceApplyConfiguration) *BuildSourceApplyConfiguration {
	b.Binary = value
	return b
}

// WithDockerfile sets the Dockerfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Dockerfile field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithDockerfile(value string) *BuildSourceApplyConfiguration {
	b.Dockerfile = &value
	return b
}

// WithGit sets the Git field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Git field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithGit(value *GitBuildSourceApplyConfiguration) *BuildSourceApplyConfiguration {
	b.Git = value
	return b
}

// WithImages adds the given value to the Images field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Images field.
func (b *BuildSourceApplyConfiguration) WithImages(values ...*ImageSourceApplyConfiguration) *BuildSourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImages")
		}
		b.Images = append(b.Images, *values[i])
	}
	return b
}

// WithContextDir sets the ContextDir field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContextDir field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithContextDir(value string) *BuildSourceApplyConfiguration {
	b.ContextDir = &value
	return b
}

// WithSourceSecret sets the SourceSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SourceSecret field is set to the value of the last call.
func (b *BuildSourceApplyConfiguration) WithSourceSecret(value corev1.LocalObjectReference) *BuildSourceApplyConfiguration {
	b.SourceSecret = &value
	return b
}

// WithSecrets adds the given value to the Secrets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Secrets field.
func (b *BuildSourceApplyConfiguration) WithSecrets(values ...*SecretBuildSourceApplyConfiguration) *BuildSourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSecrets")
		}
		b.Secrets = append(b.Secrets, *values[i])
	}
	return b
}

// WithConfigMaps adds the given value to the ConfigMaps field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConfigMaps field.
func (b *BuildSourceApplyConfiguration) WithConfigMaps(values ...*ConfigMapBuildSourceApplyConfiguration) *BuildSourceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConfigMaps")
		}
		b.ConfigMaps = append(b.ConfigMaps, *values[i])
	}
	return b
}
