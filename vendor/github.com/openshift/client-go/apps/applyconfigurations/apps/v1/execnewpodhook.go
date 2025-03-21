// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// ExecNewPodHookApplyConfiguration represents a declarative configuration of the ExecNewPodHook type for use
// with apply.
type ExecNewPodHookApplyConfiguration struct {
	Command       []string        `json:"command,omitempty"`
	Env           []corev1.EnvVar `json:"env,omitempty"`
	ContainerName *string         `json:"containerName,omitempty"`
	Volumes       []string        `json:"volumes,omitempty"`
}

// ExecNewPodHookApplyConfiguration constructs a declarative configuration of the ExecNewPodHook type for use with
// apply.
func ExecNewPodHook() *ExecNewPodHookApplyConfiguration {
	return &ExecNewPodHookApplyConfiguration{}
}

// WithCommand adds the given value to the Command field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Command field.
func (b *ExecNewPodHookApplyConfiguration) WithCommand(values ...string) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Command = append(b.Command, values[i])
	}
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *ExecNewPodHookApplyConfiguration) WithEnv(values ...corev1.EnvVar) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithContainerName sets the ContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerName field is set to the value of the last call.
func (b *ExecNewPodHookApplyConfiguration) WithContainerName(value string) *ExecNewPodHookApplyConfiguration {
	b.ContainerName = &value
	return b
}

// WithVolumes adds the given value to the Volumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Volumes field.
func (b *ExecNewPodHookApplyConfiguration) WithVolumes(values ...string) *ExecNewPodHookApplyConfiguration {
	for i := range values {
		b.Volumes = append(b.Volumes, values[i])
	}
	return b
}
