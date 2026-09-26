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
	"k8s.io/client-go/rest"
)

type sClient struct {
}

const (
	MockServiceError        = "Mock service didnt work"
	SucceedingServiceName   = "succeed"
	EmptySubsetsServiceName = "empty-subsets"
	FailingServiceName      = "fail"
)

func (s sClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Service, error) {
	if name == FailingServiceName {
		return nil, fmt.Errorf(MockServiceError)
	}
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}, nil
}
func (s sClient) Create(ctx context.Context, ds *v1.Service, opts metav1.CreateOptions) (*v1.Service, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) Delete(ctx context.Context, name string, options metav1.DeleteOptions) error {
	return fmt.Errorf("Not implemented")
}

func (s sClient) DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return fmt.Errorf("Not implemented")
}

func (s sClient) List(ctx context.Context, options metav1.ListOptions) (*v1.ServiceList, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) Update(ctx context.Context, ds *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) UpdateStatus(ctx context.Context, ds *v1.Service, opts metav1.UpdateOptions) (*v1.Service, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) Watch(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) ProxyGet(scheme string, name string, port string, path string, params map[string]string) rest.ResponseWrapper {
	return nil
}

func (s sClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Service, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) Apply(ctx context.Context, service *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Service, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func (s sClient) ApplyStatus(ctx context.Context, service *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Service, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func NewSClient() v1core.ServiceInterface {
	return sClient{}
}
