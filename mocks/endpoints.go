package mocks

import (
	"context"
	"fmt"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscorev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
)

type eClient struct {
}

const (
	MockEndpointError = "Mock endpoint didnt work"
)

func (e eClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Endpoints, error) {
	if name == FailingServiceName {
		return nil, fmt.Errorf(MockEndpointError)
	}

	subsets := []v1.EndpointSubset{}

	if name != EmptySubsetsServiceName {
		subsets = []v1.EndpointSubset{
			{
				Addresses: []v1.EndpointAddress{
					{IP: "127.0.0.1"},
				},
			},
		}
	}

	endpoint := &v1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Subsets:    subsets,
	}

	return endpoint, nil
}
func (e eClient) Create(ctx context.Context, ds *v1.Endpoints, opts metav1.CreateOptions) (*v1.Endpoints, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (e eClient) Delete(ctx context.Context, name string, options metav1.DeleteOptions) error {
	return fmt.Errorf("Not implemented")
}

func (e eClient) DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return fmt.Errorf("Not implemented")
}

func (e eClient) List(ctx context.Context, options metav1.ListOptions) (*v1.EndpointsList, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (e eClient) Update(ctx context.Context, ds *v1.Endpoints, opts metav1.UpdateOptions) (*v1.Endpoints, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (e eClient) Watch(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (e eClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoints, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func (e eClient) Apply(ctx context.Context, endpoints *applyconfigurationscorev1.EndpointsApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Endpoints, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func NewEClient() v1core.EndpointsInterface {
	return eClient{}
}
