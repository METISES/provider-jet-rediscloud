/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AlertObservation struct {
}

type AlertParameters struct {

	// Alert name
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Alert value
	// +kubebuilder:validation:Required
	Value *float64 `json:"value" tf:"value,omitempty"`
}

type AllowlistObservation struct {
}

type AllowlistParameters struct {

	// Set of CIDR ranges that are allowed to access the databases associated with this subscription
	// +kubebuilder:validation:Optional
	Cidrs []*string `json:"cidrs,omitempty" tf:"cidrs,omitempty"`

	// Set of security groups that are allowed to access the databases associated with this subscription
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type CloudProviderObservation struct {
}

type CloudProviderParameters struct {

	// Cloud account identifier. Default: Redis Labs internal cloud account (using Cloud Account Id = 1 implies using Redis Labs internal cloud account). Note that a GCP subscription can be created only with Redis Labs internal cloud account
	// +kubebuilder:validation:Optional
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// The cloud provider to use with the subscription, (either `AWS` or `GCP`)
	// +kubebuilder:validation:Optional
	Provider *string `json:"provider,omitempty" tf:"provider,omitempty"`

	// Cloud networking details, per region (single region or multiple regions for Active-Active cluster only)
	// +kubebuilder:validation:Required
	Region []RegionParameters `json:"region" tf:"region,omitempty"`
}

type DatabaseObservation struct {
	DBID *float64 `json:"dbId,omitempty" tf:"db_id,omitempty"`

	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`
}

type DatabaseParameters struct {

	// Set of alerts to enable on the database
	// +kubebuilder:validation:Optional
	Alert []AlertParameters `json:"alert,omitempty" tf:"alert,omitempty"`

	// Relevant only to ram-and-flash clusters. Estimated average size (measured in bytes) of the items stored in the database
	// +kubebuilder:validation:Optional
	AverageItemSizeInBytes *float64 `json:"averageItemSizeInBytes,omitempty" tf:"average_item_size_in_bytes,omitempty"`

	// SSL certificate to authenticate user connections
	// +kubebuilder:validation:Optional
	ClientSSLCertificate *string `json:"clientSslCertificate,omitempty" tf:"client_ssl_certificate,omitempty"`

	// Rate of database data persistence (in persistent storage)
	// +kubebuilder:validation:Optional
	DataPersistence *string `json:"dataPersistence,omitempty" tf:"data_persistence,omitempty"`

	// Use TLS for authentication
	// +kubebuilder:validation:Optional
	EnableTLS *bool `json:"enableTls,omitempty" tf:"enable_tls,omitempty"`

	// Should use the external endpoint for open-source (OSS) Cluster API
	// +kubebuilder:validation:Optional
	ExternalEndpointForOssClusterAPI *bool `json:"externalEndpointForOssClusterApi,omitempty" tf:"external_endpoint_for_oss_cluster_api,omitempty"`

	// List of regular expression rules to shard the database by. See the documentation on clustering for more information on the hashing policy - https://docs.redislabs.com/latest/rc/concepts/clustering/
	// +kubebuilder:validation:Optional
	HashingPolicy []*string `json:"hashingPolicy,omitempty" tf:"hashing_policy,omitempty"`

	// Maximum memory usage for this specific database
	// +kubebuilder:validation:Required
	MemoryLimitInGb *float64 `json:"memoryLimitInGb" tf:"memory_limit_in_gb,omitempty"`

	// A module object
	// +kubebuilder:validation:Optional
	Module []ModuleParameters `json:"module,omitempty" tf:"module,omitempty"`

	// A meaningful name to identify the database
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Password used to access the database
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Path that will be used to store database backup files
	// +kubebuilder:validation:Optional
	PeriodicBackupPath *string `json:"periodicBackupPath,omitempty" tf:"periodic_backup_path,omitempty"`

	// The protocol that will be used to access the database, (either ‘redis’ or 'memcached’)
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Set of Redis database URIs, in the format `redis://user:password@host:port`, that this database will be a replica of. If the URI provided is Redis Labs Cloud instance, only host and port should be provided
	// +kubebuilder:validation:Optional
	ReplicaOf []*string `json:"replicaOf,omitempty" tf:"replica_of,omitempty"`

	// Databases replication
	// +kubebuilder:validation:Optional
	Replication *bool `json:"replication,omitempty" tf:"replication,omitempty"`

	// Set of CIDR addresses to allow access to the database
	// +kubebuilder:validation:Optional
	SourceIps []*string `json:"sourceIps,omitempty" tf:"source_ips,omitempty"`

	// Support Redis open-source (OSS) Cluster API
	// +kubebuilder:validation:Optional
	SupportOssClusterAPI *bool `json:"supportOssClusterApi,omitempty" tf:"support_oss_cluster_api,omitempty"`

	// Throughput measurement method, (either ‘number-of-shards’ or ‘operations-per-second’)
	// +kubebuilder:validation:Required
	ThroughputMeasurementBy *string `json:"throughputMeasurementBy" tf:"throughput_measurement_by,omitempty"`

	// Throughput value (as applies to selected measurement method)
	// +kubebuilder:validation:Required
	ThroughputMeasurementValue *float64 `json:"throughputMeasurementValue" tf:"throughput_measurement_value,omitempty"`
}

type ModuleObservation struct {
}

type ModuleParameters struct {

	// Name of the module to enable
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type NetworksObservation struct {
}

type NetworksParameters struct {

	// +kubebuilder:validation:Required
	NetworkingDeploymentCidr *string `json:"networkingDeploymentCidr" tf:"networking_deployment_cidr,omitempty"`

	// +kubebuilder:validation:Required
	NetworkingSubnetID *string `json:"networkingSubnetId" tf:"networking_subnet_id,omitempty"`

	// +kubebuilder:validation:Required
	NetworkingVPCID *string `json:"networkingVpcId" tf:"networking_vpc_id,omitempty"`
}

type RegionObservation struct {
	Networks []NetworksObservation `json:"networks,omitempty" tf:"networks,omitempty"`
}

type RegionParameters struct {

	// Support deployment on multiple availability zones within the selected region
	// +kubebuilder:validation:Optional
	MultipleAvailabilityZones *bool `json:"multipleAvailabilityZones,omitempty" tf:"multiple_availability_zones,omitempty"`

	// Deployment CIDR mask
	// +kubebuilder:validation:Required
	NetworkingDeploymentCidr *string `json:"networkingDeploymentCidr" tf:"networking_deployment_cidr,omitempty"`

	// Either an existing VPC Id (already exists in the specific region) or create a new VPC (if no VPC is specified)
	// +kubebuilder:validation:Optional
	NetworkingVPCID *string `json:"networkingVpcId,omitempty" tf:"networking_vpc_id,omitempty"`

	// List of availability zones used
	// +kubebuilder:validation:Required
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones" tf:"preferred_availability_zones,omitempty"`

	// Deployment region as defined by cloud provider
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`
}

type SubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubscriptionParameters struct {

	// An allowlist object
	// +kubebuilder:validation:Optional
	Allowlist []AllowlistParameters `json:"allowlist,omitempty" tf:"allowlist,omitempty"`

	// A cloud provider object
	// +kubebuilder:validation:Required
	CloudProvider []CloudProviderParameters `json:"cloudProvider" tf:"cloud_provider,omitempty"`

	// A database object
	// +kubebuilder:validation:Required
	Database []DatabaseParameters `json:"database" tf:"database,omitempty"`

	// Memory storage preference: either ‘ram’ or a combination of 'ram-and-flash’
	// +kubebuilder:validation:Optional
	MemoryStorage *string `json:"memoryStorage,omitempty" tf:"memory_storage,omitempty"`

	// A valid payment method pre-defined in the current account
	// +kubebuilder:validation:Optional
	PaymentMethodID *string `json:"paymentMethodId,omitempty" tf:"payment_method_id,omitempty"`

	// Encrypt data stored in persistent storage. Required for a GCP subscription
	// +kubebuilder:validation:Optional
	PersistentStorageEncryption *bool `json:"persistentStorageEncryption,omitempty" tf:"persistent_storage_encryption,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rediscloudjet}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
