/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PullRequestInitParameters struct {

	// Name of the branch serving as the base of the Pull Request.
	// Name of the branch serving as the base of the Pull Request.
	BaseRef *string `json:"baseRef,omitempty" tf:"base_ref,omitempty"`

	// Name of the base repository to retrieve the Pull Requests from.
	// Name of the base repository to retrieve the Pull Requests from.
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/repo/v1alpha1.Repository
	BaseRepository *string `json:"baseRepository,omitempty" tf:"base_repository,omitempty"`

	// Reference to a Repository in repo to populate baseRepository.
	// +kubebuilder:validation:Optional
	BaseRepositoryRef *v1.Reference `json:"baseRepositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate baseRepository.
	// +kubebuilder:validation:Optional
	BaseRepositorySelector *v1.Selector `json:"baseRepositorySelector,omitempty" tf:"-"`

	// Body of the Pull Request.
	// Body of the Pull Request.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Name of the branch serving as the head of the Pull Request.
	// Name of the branch serving as the head of the Pull Request.
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/repo/v1alpha1.Branch
	HeadRef *string `json:"headRef,omitempty" tf:"head_ref,omitempty"`

	// Reference to a Branch in repo to populate headRef.
	// +kubebuilder:validation:Optional
	HeadRefRef *v1.Reference `json:"headRefRef,omitempty" tf:"-"`

	// Selector for a Branch in repo to populate headRef.
	// +kubebuilder:validation:Optional
	HeadRefSelector *v1.Selector `json:"headRefSelector,omitempty" tf:"-"`

	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	// Controls whether the base repository maintainers can modify the Pull Request. Default: 'false'.
	MaintainerCanModify *bool `json:"maintainerCanModify,omitempty" tf:"maintainer_can_modify,omitempty"`

	// Owner of the repository. If not provided, the provider's default owner is used.
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The title of the Pull Request.
	// The title of the Pull Request.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PullRequestObservation struct {

	// Name of the branch serving as the base of the Pull Request.
	// Name of the branch serving as the base of the Pull Request.
	BaseRef *string `json:"baseRef,omitempty" tf:"base_ref,omitempty"`

	// Name of the base repository to retrieve the Pull Requests from.
	// Name of the base repository to retrieve the Pull Requests from.
	BaseRepository *string `json:"baseRepository,omitempty" tf:"base_repository,omitempty"`

	// Head commit SHA of the Pull Request base.
	// Head commit SHA of the Pull Request base.
	BaseSha *string `json:"baseSha,omitempty" tf:"base_sha,omitempty"`

	// Body of the Pull Request.
	// Body of the Pull Request.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Indicates Whether this Pull Request is a draft.
	// Indicates Whether this Pull Request is a draft.
	Draft *bool `json:"draft,omitempty" tf:"draft,omitempty"`

	// Name of the branch serving as the head of the Pull Request.
	// Name of the branch serving as the head of the Pull Request.
	HeadRef *string `json:"headRef,omitempty" tf:"head_ref,omitempty"`

	// Head commit SHA of the Pull Request head.
	// Head commit SHA of the Pull Request head.
	HeadSha *string `json:"headSha,omitempty" tf:"head_sha,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of label names set on the Pull Request.
	// List of names of labels on the PR
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	// Controls whether the base repository maintainers can modify the Pull Request. Default: 'false'.
	MaintainerCanModify *bool `json:"maintainerCanModify,omitempty" tf:"maintainer_can_modify,omitempty"`

	// The number of the Pull Request within the repository.
	// The number of the Pull Request within the repository.
	Number *int64 `json:"number,omitempty" tf:"number,omitempty"`

	// Unix timestamp indicating the Pull Request creation time.
	// Unix timestamp indicating the Pull Request creation time.
	OpenedAt *int64 `json:"openedAt,omitempty" tf:"opened_at,omitempty"`

	// GitHub login of the user who opened the Pull Request.
	// Username of the PR creator
	OpenedBy *string `json:"openedBy,omitempty" tf:"opened_by,omitempty"`

	// Owner of the repository. If not provided, the provider's default owner is used.
	// Owner of the repository. If not provided, the provider's default owner is used.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// the current Pull Request state - can be "open", "closed" or "merged".
	// The current Pull Request state - can be 'open', 'closed' or 'merged'.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The title of the Pull Request.
	// The title of the Pull Request.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`

	// The timestamp of the last Pull Request update.
	// The timestamp of the last Pull Request update.
	UpdatedAt *int64 `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type PullRequestParameters struct {

	// Name of the branch serving as the base of the Pull Request.
	// Name of the branch serving as the base of the Pull Request.
	// +kubebuilder:validation:Optional
	BaseRef *string `json:"baseRef,omitempty" tf:"base_ref,omitempty"`

	// Name of the base repository to retrieve the Pull Requests from.
	// Name of the base repository to retrieve the Pull Requests from.
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/repo/v1alpha1.Repository
	// +kubebuilder:validation:Optional
	BaseRepository *string `json:"baseRepository,omitempty" tf:"base_repository,omitempty"`

	// Reference to a Repository in repo to populate baseRepository.
	// +kubebuilder:validation:Optional
	BaseRepositoryRef *v1.Reference `json:"baseRepositoryRef,omitempty" tf:"-"`

	// Selector for a Repository in repo to populate baseRepository.
	// +kubebuilder:validation:Optional
	BaseRepositorySelector *v1.Selector `json:"baseRepositorySelector,omitempty" tf:"-"`

	// Body of the Pull Request.
	// Body of the Pull Request.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Name of the branch serving as the head of the Pull Request.
	// Name of the branch serving as the head of the Pull Request.
	// +crossplane:generate:reference:type=github.com/coopnorge/provider-github/apis/repo/v1alpha1.Branch
	// +kubebuilder:validation:Optional
	HeadRef *string `json:"headRef,omitempty" tf:"head_ref,omitempty"`

	// Reference to a Branch in repo to populate headRef.
	// +kubebuilder:validation:Optional
	HeadRefRef *v1.Reference `json:"headRefRef,omitempty" tf:"-"`

	// Selector for a Branch in repo to populate headRef.
	// +kubebuilder:validation:Optional
	HeadRefSelector *v1.Selector `json:"headRefSelector,omitempty" tf:"-"`

	// Controls whether the base repository maintainers can modify the Pull Request. Default: false.
	// Controls whether the base repository maintainers can modify the Pull Request. Default: 'false'.
	// +kubebuilder:validation:Optional
	MaintainerCanModify *bool `json:"maintainerCanModify,omitempty" tf:"maintainer_can_modify,omitempty"`

	// Owner of the repository. If not provided, the provider's default owner is used.
	// Owner of the repository. If not provided, the provider's default owner is used.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The title of the Pull Request.
	// The title of the Pull Request.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// PullRequestSpec defines the desired state of PullRequest
type PullRequestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PullRequestParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PullRequestInitParameters `json:"initProvider,omitempty"`
}

// PullRequestStatus defines the observed state of PullRequest.
type PullRequestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PullRequestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PullRequest is the Schema for the PullRequests API. Get information on a single GitHub Pull Request.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type PullRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseRef) || (has(self.initProvider) && has(self.initProvider.baseRef))",message="spec.forProvider.baseRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	Spec   PullRequestSpec   `json:"spec"`
	Status PullRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PullRequestList contains a list of PullRequests
type PullRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PullRequest `json:"items"`
}

// Repository type metadata.
var (
	PullRequest_Kind             = "PullRequest"
	PullRequest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PullRequest_Kind}.String()
	PullRequest_KindAPIVersion   = PullRequest_Kind + "." + CRDGroupVersion.String()
	PullRequest_GroupVersionKind = CRDGroupVersion.WithKind(PullRequest_Kind)
)

func init() {
	SchemeBuilder.Register(&PullRequest{}, &PullRequestList{})
}
