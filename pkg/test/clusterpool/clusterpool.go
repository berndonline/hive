package clusterpool

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	hivev1aws "github.com/openshift/hive/pkg/apis/hive/v1/aws"
	"github.com/openshift/hive/pkg/test/generic"
)

// Option defines a function signature for any function that wants to be passed into Build
type Option func(*hivev1.ClusterPool)

// Build runs each of the functions passed in to generate the object.
func Build(opts ...Option) *hivev1.ClusterPool {
	retval := &hivev1.ClusterPool{}
	for _, o := range opts {
		o(retval)
	}

	return retval
}

type Builder interface {
	Build(opts ...Option) *hivev1.ClusterPool

	Options(opts ...Option) Builder

	GenericOptions(opts ...generic.Option) Builder
}

func BasicBuilder() Builder {
	return &builder{}
}

func FullBuilder(namespace, name string, typer runtime.ObjectTyper) Builder {
	b := &builder{}
	return b.GenericOptions(
		generic.WithTypeMeta(typer),
		generic.WithResourceVersion("1"),
		generic.WithNamespace(namespace),
		generic.WithName(name),
	)
}

type builder struct {
	options []Option
}

func (b *builder) Build(opts ...Option) *hivev1.ClusterPool {
	return Build(append(b.options, opts...)...)
}

func (b *builder) Options(opts ...Option) Builder {
	return &builder{
		options: append(b.options, opts...),
	}
}

func (b *builder) GenericOptions(opts ...generic.Option) Builder {
	options := make([]Option, len(opts))
	for i, o := range opts {
		options[i] = Generic(o)
	}
	return b.Options(options...)
}

// Generic allows common functions applicable to all objects to be used as Options to Build
func Generic(opt generic.Option) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		opt(clusterPool)
	}
}

func ForAWS(credsSecretName, region string) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.Platform.AWS = &hivev1aws.Platform{
			CredentialsSecretRef: corev1.LocalObjectReference{Name: credsSecretName},
			Region:               "region",
		}
	}
}

func WithPullSecret(pullSecretName string) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.PullSecretRef = &corev1.LocalObjectReference{Name: pullSecretName}
	}
}

func WithSize(size int) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.Size = int32(size)
	}
}

func WithClusterDeploymentLabels(labels map[string]string) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.Labels = labels
	}
}

func WithBaseDomain(baseDomain string) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.BaseDomain = baseDomain
	}
}

func WithImageSet(clusterImageSetName string) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		clusterPool.Spec.ImageSetRef = hivev1.ClusterImageSetReference{Name: clusterImageSetName}
	}
}

// WithCondition adds the specified condition to the ClusterPool
func WithCondition(cond hivev1.ClusterPoolCondition) Option {
	return func(clusterPool *hivev1.ClusterPool) {
		for i, c := range clusterPool.Status.Conditions {
			if c.Type == cond.Type {
				clusterPool.Status.Conditions[i] = cond
				return
			}
		}
		clusterPool.Status.Conditions = append(clusterPool.Status.Conditions, cond)
	}
}
