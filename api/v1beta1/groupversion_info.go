// Package v1beta1 contains API Schema definitions for the keycloak v1beta1 API group
// +kubebuilder:object:generate=true
// +groupName=keycloak.hostzero.com
package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "keycloak.hostzero.com", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// Builder builds a new Scheme for mapping go types to Kubernetes GroupVersionKinds.
//
// This is a local replacement for the deprecated
// sigs.k8s.io/controller-runtime/pkg/scheme.Builder, kept here so that this api
// package depends only on k8s.io/apimachinery. The behaviour is intentionally
// identical to the upstream Builder.
type Builder struct {
	GroupVersion schema.GroupVersion
	runtime.SchemeBuilder
}

// Register adds the given objects to the SchemeBuilder so they can be added
// to a Scheme. Returns the Builder for chaining.
func (b *Builder) Register(objects ...runtime.Object) *Builder {
	b.SchemeBuilder.Register(func(scheme *runtime.Scheme) error {
		for _, obj := range objects {
			scheme.AddKnownTypes(b.GroupVersion, obj)
		}
		metav1.AddToGroupVersion(scheme, b.GroupVersion)
		return nil
	})
	return b
}

// AddToScheme adds all the registered types to the given Scheme.
func (b *Builder) AddToScheme(s *runtime.Scheme) error {
	return b.SchemeBuilder.AddToScheme(s)
}
