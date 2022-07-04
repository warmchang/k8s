package configmap

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ListByLabel list configmaps by labels.
func (h *Handler) ListByLabel(labels string) (*corev1.ConfigMapList, error) {
	listOptions := h.Options.ListOptions.DeepCopy()
	listOptions.LabelSelector = labels
	return h.clientset.CoreV1().ConfigMaps(h.namespace).List(h.ctx, *listOptions)
}

// List list configmaps by labels, alias to "ListByLabel".
func (h *Handler) List(labels string) (*corev1.ConfigMapList, error) {
	return h.ListByLabel(labels)
}

// ListByNamespace list configmaps by namespace.
func (h *Handler) ListByNamespace(namespace string) (*corev1.ConfigMapList, error) {
	return h.WithNamespace(namespace).ListByLabel("")
}

// ListAll list all configmaps in the k8s cluster.
func (h *Handler) ListAll() (*corev1.ConfigMapList, error) {
	return h.WithNamespace(metav1.NamespaceAll).ListByLabel("")
}
