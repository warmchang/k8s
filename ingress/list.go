package ingress

import (
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

// List list all ingresses in the k8s cluster, it simply call `ListAll`.
func (h *Handler) List() ([]*networkingv1.Ingress, error) {
	return h.ListAll()
}

// ListByLabel list ingresses by labels.
// Multiple labels separated by comma(",") eg: "name=myapp,role=devops",
// and there is an "And" relationship between multiple labels.
func (h *Handler) ListByLabel(labels string) ([]*networkingv1.Ingress, error) {
	listOptions := h.Options.ListOptions.DeepCopy()
	listOptions.LabelSelector = labels
	ingList, err := h.clientset.NetworkingV1().Ingresses(h.namespace).List(h.ctx, *listOptions)
	if err != nil {
		return nil, err
	}
	return extractList(ingList), nil
}

// ListByField list ingresses by field, work like `kubectl get xxx --field-selector=xxx`.
func (h *Handler) ListByField(field string) ([]*networkingv1.Ingress, error) {
	fieldSelector, err := fields.ParseSelector(field)
	if err != nil {
		return nil, err
	}
	listOptions := h.Options.ListOptions.DeepCopy()
	listOptions.FieldSelector = fieldSelector.String()

	ingList, err := h.clientset.NetworkingV1().Ingresses(h.namespace).List(h.ctx, *listOptions)
	if err != nil {
		return nil, err
	}
	return extractList(ingList), nil
}

// ListByNamespace list all ingresses in the specified namespace.
func (h *Handler) ListByNamespace(namespace string) ([]*networkingv1.Ingress, error) {
	return h.WithNamespace(namespace).ListByLabel("")
}

// ListAll list all ingresses in the k8s cluster.
func (h *Handler) ListAll() ([]*networkingv1.Ingress, error) {
	return h.WithNamespace(metav1.NamespaceAll).ListByLabel("")
}

// extractList
func extractList(ingList *networkingv1.IngressList) []*networkingv1.Ingress {
	var objList []*networkingv1.Ingress
	for i := range ingList.Items {
		objList = append(objList, &ingList.Items[i])
	}
	return objList
}
