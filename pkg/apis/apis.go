/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package apis contains Kubernetes API groups.
package apis

import (
	"github.com/ellistarn/etcd-operator/pkg/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/webhook/resourcesemantics"
)

var (
	// Builder includes all types within the apis package
	Builder = runtime.NewSchemeBuilder(v1alpha1.SchemeBuilder.AddToScheme)
	// AddToScheme may be used to add all resources defined in the project to a Scheme
	AddToScheme = Builder.AddToScheme
	// Resources defined in the project
	Resources = map[schema.GroupVersionKind]resourcesemantics.GenericCRD{
		v1alpha1.SchemeGroupVersion.WithKind("Cluster"): v1alpha1.ClusterReference,
	}
)
