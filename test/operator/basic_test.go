package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/nats-io/nats-operator/pkg/spec"
	"github.com/nats-io/nats-streaming-operator/internal/operator"
	stanv1alpha1 "github.com/nats-io/nats-streaming-operator/pkg/apis/streaming/v1alpha1"
	stancrdclient "github.com/nats-io/nats-streaming-operator/pkg/client/v1alpha1"
	k8scrdclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8sclient "k8s.io/client-go/kubernetes/typed/core/v1"
	k8srestapi "k8s.io/client-go/rest"
	k8sclientcmd "k8s.io/client-go/tools/clientcmd"
)

func TestCreateCluster(t *testing.T) {
	kc, err := newKubeClients()
	if err != nil {
		t.Fatal(err)
	}
	controller := operator.NewController(nil)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go controller.Run(ctx)

	// Create cluster, wait for pods then cancel the test.
	cluster := &stanv1alpha1.NatsStreamingCluster{
		TypeMeta: k8smetav1.TypeMeta{
			Kind:       stanv1alpha1.CRDResourceKind,
			APIVersion: stanv1alpha1.SchemeGroupVersion.String(),
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Name:      "stan-cluster-basic-test",
			Namespace: "default",
		},
		Spec: spec.ClusterSpec{
			Size:        3,
			NatsService: "nats-cluster-stan-test",
		},
	}
	kc.stan.Create(cluster)

	select {
	case <-ctx.Done():
		if ctx.Err() != nil && ctx.Err() != context.Canceled {
			t.Errorf("Unexpected error: %v", ctx.Err())
		}
	}
}

type clients struct {
	core k8sclient.CoreV1Interface
	crd  k8scrdclient.Interface
	stan stancrdclient.Interface
}

func newKubeClients() (*clients, error) {
	var err error
	var cfg *k8srestapi.Config
	if kubeconfig := os.Getenv("KUBERNETES_CONFIG_FILE"); kubeconfig != "" {
		cfg, err = k8sclientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("KUBERNETES_CONFIG_FILE env variable must be set")
	}
	kc, err := k8sclient.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	kcrdc, err := k8scrdclient.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	ncr, err := stancrdclient.NewForConfig(cfg)
	if err != nil {
		return err
	}

	return &clients{
		core: kc,
		crd:  kcrdc,
		stan: ncr,
	}, nil
}
