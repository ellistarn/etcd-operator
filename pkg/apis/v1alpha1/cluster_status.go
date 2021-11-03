/*
Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
*/

package v1alpha1

import (
	"knative.dev/pkg/apis"
)

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
	// +optional
	Conditions apis.Conditions `json:"conditions,omitempty"`
}

func (p *Cluster) StatusConditions() apis.ConditionManager {
	return apis.NewLivingConditionSet().Manage(p)
}

func (p *Cluster) GetConditions() apis.Conditions {
	return p.Status.Conditions
}

func (p *Cluster) SetConditions(conditions apis.Conditions) {
	p.Status.Conditions = conditions
}
