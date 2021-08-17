// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NotebookInstanceSpec defines the desired state of NotebookInstance.
type NotebookInstanceSpec struct {
	// A list of Elastic Inference (EI) instance types to associate with this notebook
	// instance. Currently, only one instance type can be associated with a notebook
	// instance. For more information, see Using Elastic Inference in Amazon SageMaker
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html).
	AcceleratorTypes []*string `json:"acceleratorTypes,omitempty"`
	// An array of up to three Git repositories to associate with the notebook instance.
	// These can be either the names of Git repositories stored as resources in
	// your account, or the URL of Git repositories in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. These repositories are cloned at the same
	// level as the default repository of your notebook instance. For more information,
	// see Associating Git Repositories with Amazon SageMaker Notebook Instances
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	AdditionalCodeRepositories []*string `json:"additionalCodeRepositories,omitempty"`
	// A Git repository to associate with the notebook instance as its default code
	// repository. This can be either the name of a Git repository stored as a resource
	// in your account, or the URL of a Git repository in AWS CodeCommit (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html)
	// or in any other Git repository. When you open a notebook instance, it opens
	// in the directory that contains this repository. For more information, see
	// Associating Git Repositories with Amazon SageMaker Notebook Instances (https://docs.aws.amazon.com/sagemaker/latest/dg/nbi-git-repo.html).
	DefaultCodeRepository *string `json:"defaultCodeRepository,omitempty"`
	// Sets whether Amazon SageMaker provides internet access to the notebook instance.
	// If you set this to Disabled this notebook instance will be able to access
	// resources only in your VPC, and will not be able to connect to Amazon SageMaker
	// training and endpoint services unless your configure a NAT Gateway in your
	// VPC.
	//
	// For more information, see Notebook Instances Are Internet-Enabled by Default
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/appendix-additional-considerations.html#appendix-notebook-and-internet-access).
	// You can set the value of this parameter to Disabled only if you set a value
	// for the SubnetId parameter.
	DirectInternetAccess *string `json:"directInternetAccess,omitempty"`
	// The type of ML compute instance to launch for the notebook instance.
	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType"`
	// The Amazon Resource Name (ARN) of a AWS Key Management Service key that Amazon
	// SageMaker uses to encrypt data on the storage volume attached to your notebook
	// instance. The KMS key you provide must be enabled. For information, see Enabling
	// and Disabling Keys (https://docs.aws.amazon.com/kms/latest/developerguide/enabling-keys.html)
	// in the AWS Key Management Service Developer Guide.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The name of a lifecycle configuration to associate with the notebook instance.
	// For information about lifestyle configurations, see Step 2.1: (Optional)
	// Customize a Notebook Instance (https://docs.aws.amazon.com/sagemaker/latest/dg/notebook-lifecycle-config.html).
	LifecycleConfigName *string `json:"lifecycleConfigName,omitempty"`
	// The name of the new notebook instance.
	// +kubebuilder:validation:Required
	NotebookInstanceName *string `json:"notebookInstanceName"`
	// When you send any requests to AWS resources from the notebook instance, Amazon
	// SageMaker assumes this role to perform tasks on your behalf. You must grant
	// this role necessary permissions so Amazon SageMaker can perform these tasks.
	// The policy must allow the Amazon SageMaker service principal (sagemaker.amazonaws.com)
	// permissions to assume this role. For more information, see Amazon SageMaker
	// Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).
	//
	// To be able to pass this role to Amazon SageMaker, the caller of this API
	// must have the iam:PassRole permission.
	// +kubebuilder:validation:Required
	RoleARN *string `json:"roleARN"`
	// Whether root access is enabled or disabled for users of the notebook instance.
	// The default value is Enabled.
	//
	// Lifecycle configurations need root access to be able to set up a notebook
	// instance. Because of this, lifecycle configurations associated with a notebook
	// instance always run with root access even if you disable root access for
	// users.
	RootAccess *string `json:"rootAccess,omitempty"`
	// The VPC security group IDs, in the form sg-xxxxxxxx. The security groups
	// must be for the same VPC as specified in the subnet.
	SecurityGroupIDs []*string `json:"securityGroupIDs,omitempty"`
	// The ID of the subnet in a VPC to which you would like to have a connectivity
	// from your ML compute instance.
	SubnetID *string `json:"subnetID,omitempty"`
	// An array of key-value pairs. You can use tags to categorize your AWS resources
	// in different ways, for example, by purpose, owner, or environment. For more
	// information, see Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags []*Tag `json:"tags,omitempty"`
	// The size, in GB, of the ML storage volume to attach to the notebook instance.
	// The default value is 5 GB.
	VolumeSizeInGB *int64 `json:"volumeSizeInGB,omitempty"`
}

// NotebookInstanceStatus defines the observed state of NotebookInstance
type NotebookInstanceStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// If status is Failed, the reason it failed.
	FailureReason *string `json:"failureReason,omitempty"`
	// The status of the notebook instance.
	NotebookInstanceStatus *string `json:"notebookInstanceStatus,omitempty"`
	// The URL that you use to connect to the Jupyter notebook that is running in
	// your notebook instance.
	URL *string `json:"url,omitempty"`
	// The URL that you use to connect to the Jupyter notebook that is running in
	// your notebook instance.
	StoppedByControllerMetadata *string `json:"stoppedByControllerMetadata,omitempty"`
}

// NotebookInstance is the Schema for the NotebookInstances API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FAILURE-REASON",type=string,priority=1,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.notebookInstanceStatus`
type NotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotebookInstanceSpec   `json:"spec,omitempty"`
	Status            NotebookInstanceStatus `json:"status,omitempty"`
}

// NotebookInstanceList contains a list of NotebookInstance
// +kubebuilder:object:root=true
type NotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NotebookInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NotebookInstance{}, &NotebookInstanceList{})
}
