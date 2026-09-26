package mocks

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	v1apps "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type dClient struct {
}

const (
	SucceedingDaemonsetName         = "DAEMONSET_SUCCEED"
	FailingDaemonsetName            = "DAEMONSET_FAIL"
	CorrectNamespaceDaemonsetName   = "CORRECT_DAEMONSET_NAMESPACE"
	IncorrectNamespaceDaemonsetName = "INCORRECT_DAEMONSET_NAMESPACE"
	CorrectDaemonsetNamespace       = "CORRECT_DAEMONSET"

	FailingMatchLabelsDaemonsetName  = "DAEMONSET_INCORRECT_MATCH_LABELS"
	NotReadyMatchLabelsDaemonsetName = "DAEMONSET_NOT_READY_MATCH_LABELS"
)

func (d dClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*appsv1.DaemonSet, error) {
	matchLabelName := MockContainerName

	if name == FailingDaemonsetName {
		return nil, fmt.Errorf("Mock daemonset didnt work")
	} else if name == FailingMatchLabelsDaemonsetName {
		matchLabelName = FailingMatchLabel
	} else if name == NotReadyMatchLabelsDaemonsetName {
		matchLabelName = SameHostNotReadyMatchLabel
	}

	ds := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"name": matchLabelName},
			},
		},
	}

	if name == CorrectNamespaceDaemonsetName {
		ds.ObjectMeta.Namespace = CorrectDaemonsetNamespace
	} else if name == IncorrectNamespaceDaemonsetName {
		return nil, fmt.Errorf("Mock daemonset didnt work")
	}

	return ds, nil
}
func (d dClient) Create(ctx context.Context, ds *appsv1.DaemonSet, opts metav1.CreateOptions) (*appsv1.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Delete(ctx context.Context, name string, options metav1.DeleteOptions) error {
	return fmt.Errorf("Not implemented")
}
func (d dClient) List(ctx context.Context, options metav1.ListOptions) (*appsv1.DaemonSetList, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Update(ctx context.Context, ds *appsv1.DaemonSet, opts metav1.UpdateOptions) (*appsv1.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) UpdateStatus(ctx context.Context, ds *appsv1.DaemonSet, opts metav1.UpdateOptions) (*appsv1.DaemonSet, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) DeleteCollection(ctx context.Context, options metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return fmt.Errorf("Not implemented")
}

func (d dClient) Watch(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *appsv1.DaemonSet, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) Apply(ctx context.Context, daemonSet *applyconfigurationsappsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *appsv1.DaemonSet, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func (d dClient) ApplyStatus(ctx context.Context, daemonSet *applyconfigurationsappsv1.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (result *appsv1.DaemonSet, err error) {
	return nil, fmt.Errorf("Not implemented")
}

func NewDSClient() v1apps.DaemonSetInterface {
	return dClient{}
}
