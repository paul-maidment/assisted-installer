// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RegistrySourcesApplyConfiguration represents an declarative configuration of the RegistrySources type for use
// with apply.
type RegistrySourcesApplyConfiguration struct {
	InsecureRegistries               []string `json:"insecureRegistries,omitempty"`
	BlockedRegistries                []string `json:"blockedRegistries,omitempty"`
	AllowedRegistries                []string `json:"allowedRegistries,omitempty"`
	ContainerRuntimeSearchRegistries []string `json:"containerRuntimeSearchRegistries,omitempty"`
}

// RegistrySourcesApplyConfiguration constructs an declarative configuration of the RegistrySources type for use with
// apply.
func RegistrySources() *RegistrySourcesApplyConfiguration {
	return &RegistrySourcesApplyConfiguration{}
}

// WithInsecureRegistries adds the given value to the InsecureRegistries field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the InsecureRegistries field.
func (b *RegistrySourcesApplyConfiguration) WithInsecureRegistries(values ...string) *RegistrySourcesApplyConfiguration {
	for i := range values {
		b.InsecureRegistries = append(b.InsecureRegistries, values[i])
	}
	return b
}

// WithBlockedRegistries adds the given value to the BlockedRegistries field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the BlockedRegistries field.
func (b *RegistrySourcesApplyConfiguration) WithBlockedRegistries(values ...string) *RegistrySourcesApplyConfiguration {
	for i := range values {
		b.BlockedRegistries = append(b.BlockedRegistries, values[i])
	}
	return b
}

// WithAllowedRegistries adds the given value to the AllowedRegistries field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AllowedRegistries field.
func (b *RegistrySourcesApplyConfiguration) WithAllowedRegistries(values ...string) *RegistrySourcesApplyConfiguration {
	for i := range values {
		b.AllowedRegistries = append(b.AllowedRegistries, values[i])
	}
	return b
}

// WithContainerRuntimeSearchRegistries adds the given value to the ContainerRuntimeSearchRegistries field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ContainerRuntimeSearchRegistries field.
func (b *RegistrySourcesApplyConfiguration) WithContainerRuntimeSearchRegistries(values ...string) *RegistrySourcesApplyConfiguration {
	for i := range values {
		b.ContainerRuntimeSearchRegistries = append(b.ContainerRuntimeSearchRegistries, values[i])
	}
	return b
}
