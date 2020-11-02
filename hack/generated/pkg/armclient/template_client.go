/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package armclient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
)

// TODO: Naming?
type Applier interface {
	CreateDeployment(ctx context.Context, deployment *Deployment) (*Deployment, error)
	DeleteDeployment(ctx context.Context, deploymentId string) error
	GetDeployment(ctx context.Context, deploymentId string) (*Deployment, error)
	NewResourceGroupDeployment(resourceGroup string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment
	NewSubscriptionDeployment(location string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment

	// TODO: These functions take an empty status and fill it out with the response from Azure (rather than as
	// TODO: the return type. I don't love that pattern but don't have a better one either.
	BeginDeleteResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error
	GetResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error
	HeadResource(ctx context.Context, id string, apiVersion string) (bool, error)
}

type AzureTemplateClient struct {
	RawClient      *Client
	Logger         logr.Logger
	SubscriptionID string
}

type Template struct {
	Schema         string            `json:"$schema,omitempty"`
	ContentVersion string            `json:"contentVersion,omitempty"`
	Parameters     interface{}       `json:"parameters,omitempty"`
	Variables      interface{}       `json:"variables,omitempty"`
	Resources      []interface{}     `json:"resources,omitempty"`
	Outputs        map[string]Output `json:"outputs,omitempty"`
}

type Output struct {
	Condition string `json:"condition,omitempty"`
	Type      string `json:"type,omitempty"`
	Value     string `json:"value,omitempty"`
}

/*
	TemplateResourceObjectOutput represents the structure output from a deployment template for a given resource when
	requesting a 'Full' representation. The structure for a resource group is as follows:
		{
		  "apiVersion": "2018-05-01",
		  "location": "westus2",
		  "properties": {
			"provisioningState": "Succeeded"
		  },
		  "subscriptionId": "guid",
		  "scope": "",
		  "resourceId": "Microsoft.Resources/resourceGroups/foo",
		  "referenceApiVersion": "2018-05-01",
		  "condition": true,
		  "isConditionTrue": true,
		  "isTemplateResource": false,
		  "isAction": false,
		  "provisioningOperation": "Read"
		}
*/
type TemplateResourceObjectOutput struct {
	APIVersion            string      `json:"apiVersion,omitempty"`
	Location              string      `json:"location,omitempty"`
	Properties            interface{} `json:"properties,omitempty"`
	SubscriptionID        string      `json:"subscriptionId,omitempty"`
	Scope                 string      `json:"scope,omitempty"`
	ID                    string      `json:"id,omitempty"`
	ResourceID            string      `json:"resourceId,omitempty"`
	ReferenceAPIVersion   string      `json:"referenceApiVersion,omitempty"`
	Condition             *bool       `json:"condition,omitempty"`
	IsCondition           *bool       `json:"isConditionTrue,omitempty"`
	IsTemplateResource    *bool       `json:"isTemplateResource,omitempty"`
	IsAction              *bool       `json:"isAction,omitempty"`
	ProvisioningOperation string      `json:"provisioningOperation,omitempty"`
}

type TemplateOutput struct {
	Type  string                       `json:"type,omitempty"`
	Value TemplateResourceObjectOutput `json:"value,omitempty"`
}

type RetryConfig struct {
	Attempts   int
	Backoff    time.Duration
	MaxBackoff time.Duration
}

type ClientConfig struct {
	Logger  logr.Logger
	Retries *RetryConfig
}

type AzureTemplateClientOption func(config *ClientConfig) *ClientConfig

var _ Applier = &AzureTemplateClient{}

func WithLogger(logger logr.Logger) func(*ClientConfig) *ClientConfig {
	return func(cfg *ClientConfig) *ClientConfig {
		cfg.Logger = logger
		return cfg
	}
}

func WithRetries(retries *RetryConfig) func(*ClientConfig) *ClientConfig {
	return func(cfg *ClientConfig) *ClientConfig {
		cfg.Retries = retries
		return cfg
	}
}

func WithDefaultRetries() func(*ClientConfig) *ClientConfig {
	return WithRetries(&RetryConfig{
		Attempts:   5,
		Backoff:    2 * time.Second,
		MaxBackoff: 30 * time.Second,
	})
}

func AuthorizerFromEnvironment() (autorest.Authorizer, error) {
	envSettings, err := auth.GetSettingsFromEnvironment()
	if err != nil {
		return nil, err
	}

	authorizer, err := envSettings.GetAuthorizer()
	if err != nil {
		return nil, err
	}

	return authorizer, nil
}

func NewAzureTemplateClient(authorizer autorest.Authorizer, opts ...AzureTemplateClientOption) (*AzureTemplateClient, error) {
	cfg := &ClientConfig{
		Logger: ctrl.Log.WithName("azure_template_client"),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	subID := os.Getenv(auth.SubscriptionID)
	if subID == "" {
		return nil, errors.Errorf("env var %q was not set", auth.SubscriptionID)
	}

	rawClient := NewClient(authorizer)

	if cfg.Retries != nil {
		rawClient = rawClient.WithExponentialRetries(
			cfg.Retries.Attempts,
			cfg.Retries.Backoff,
			cfg.Retries.MaxBackoff)
	}

	return &AzureTemplateClient{
		RawClient:      rawClient,
		Logger:         cfg.Logger,
		SubscriptionID: subID,
	}, nil
}

func (atc *AzureTemplateClient) GetResource(ctx context.Context, id string, apiVersion string, status genruntime.ArmResourceStatus) error {
	if id == "" {
		return errors.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version=%s", id, apiVersion)
	err := atc.RawClient.GetResource(ctx, path, &status) // TODO: is this right?
	return err
}

// CreateDeployment deploys a resource to Azure via a deployment template
func (atc *AzureTemplateClient) CreateDeployment(ctx context.Context, deployment *Deployment) (*Deployment, error) {
	return atc.RawClient.PutDeployment(ctx, deployment)
}

// DeleteDeployment deletes a deployment. If the deployment doesn't exist it does not return an error
func (atc *AzureTemplateClient) DeleteDeployment(ctx context.Context, deploymentId string) error {
	err := atc.RawClient.DeleteResource(ctx, idWithAPIVersion(deploymentId), nil)

	// NotFound is a success
	if IsNotFound(err) {
		return nil
	}

	return err
}

func (atc *AzureTemplateClient) GetDeployment(ctx context.Context, deploymentId string) (*Deployment, error) {
	var deployment Deployment
	if err := atc.RawClient.GetResource(ctx, idWithAPIVersion(deploymentId), &deployment); err != nil {
		return &deployment, err
	}
	return &deployment, nil
}

func createResourceIdTemplate(resourceSpec genruntime.ArmResourceSpec) map[string]Output {
	resourceName := resourceSpec.GetName()
	names := strings.Split(resourceName, "/")
	formattedNames := make([]string, len(names))
	for i, name := range names {
		formattedNames[i] = fmt.Sprintf("'%s'", name)
	}

	resourceIdTemplateFunction := fmt.Sprintf("resourceId('%s', %s)", resourceSpec.GetType(), strings.Join(formattedNames, ", "))
	result := map[string]Output{
		"resourceId": {
			Type:  "string",
			Value: fmt.Sprintf("[%s]", resourceIdTemplateFunction),
		},
	}

	return result
}

func (atc *AzureTemplateClient) NewResourceGroupDeployment(resourceGroup string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment {
	deployment := NewResourceGroupDeployment(atc.SubscriptionID, resourceGroup, deploymentName, resourceSpec)
	deployment.Properties.Template.Outputs = createResourceIdTemplate(resourceSpec)
	return deployment
}

func (atc *AzureTemplateClient) NewSubscriptionDeployment(location string, deploymentName string, resourceSpec genruntime.ArmResourceSpec) *Deployment {
	deployment := NewSubscriptionDeployment(atc.SubscriptionID, location, deploymentName, resourceSpec)
	deployment.Properties.Template.Outputs = createResourceIdTemplate(resourceSpec)
	return deployment
}

func (atc *AzureTemplateClient) BeginDeleteResource(
	ctx context.Context,
	id string,
	apiVersion string,
	status genruntime.ArmResourceStatus) error {

	if id == "" {
		return errors.Errorf("resource ID cannot be empty")
	}

	path := fmt.Sprintf("%s?api-version=%s", id, apiVersion)
	if err := atc.RawClient.DeleteResource(ctx, path, &status); err != nil {
		return errors.Wrapf(err, "failed deleting %s", id)
	}

	return nil
}

// HeadResource checks to see if the resource exists
//
// Note: this doesn't actually use HTTP HEAD as Azure Resource Manager does not uniformly implement HEAD for all
// all resources. Also, ARM returns a 400 rather than 405 when requesting HEAD for a resource which the Resource
// Provider does not implement HEAD. For these reasons, we use an HTTP GET
func (atc *AzureTemplateClient) HeadResource(ctx context.Context, id string, apiVersion string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("resource ID cannot be empty")
	}

	idAndAPIVersion := id + fmt.Sprintf("?api-version=%s", apiVersion)
	ignored := struct{}{}
	err := atc.RawClient.GetResource(ctx, idAndAPIVersion, &ignored)
	switch {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func MakeArmResourceId(subscriptionId string, segments ...string) (string, error) {
	// There should be an even number of segments
	if len(segments)%2 != 0 {
		return "", errors.Errorf("expected even number of ARM resource ID segments, got: %d", len(segments))
	}

	start := "/subscriptions/" + subscriptionId
	remaining := strings.Join(segments, "/")

	return start + "/" + remaining, nil
}
