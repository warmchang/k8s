package clusterrolebinding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

// Create creates clusterrolebinding from type string, []byte, *rbacv1.ClusterRoleBinding,
// rbacv1.ClusterRoleBinding, runtime.Object or map[string]interface{}.
func (h *Handler) Create(obj interface{}) (*rbacv1.ClusterRoleBinding, error) {
	switch val := obj.(type) {
	case string:
		return h.CreateFromFile(val)
	case []byte:
		return h.CreateFromBytes(val)
	case *rbacv1.ClusterRoleBinding:
		return h.CreateFromObject(val)
	case rbacv1.ClusterRoleBinding:
		return h.CreateFromObject(&val)
	case runtime.Object:
		return h.CreateFromObject(val)
	case map[string]interface{}:
		return h.CreateFromUnstructured(val)
	default:
		return nil, ERR_TYPE_CREATE
	}
}

// CreateFromFile creates clusterrolebinding from yaml file.
func (h *Handler) CreateFromFile(filename string) (*rbacv1.ClusterRoleBinding, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return h.CreateFromBytes(data)
}

// CreateFromBytes creates clusterrolebinding from bytes.
func (h *Handler) CreateFromBytes(data []byte) (*rbacv1.ClusterRoleBinding, error) {
	crbJson, err := yaml.ToJSON(data)
	if err != nil {
		return nil, err
	}

	crb := &rbacv1.ClusterRoleBinding{}
	err = json.Unmarshal(crbJson, crb)
	if err != nil {
		return nil, err
	}
	return h.createCRB(crb)
}

// CreateFromObject creates clusterrolebinding from runtime.Object.
func (h *Handler) CreateFromObject(obj runtime.Object) (*rbacv1.ClusterRoleBinding, error) {
	crb, ok := obj.(*rbacv1.ClusterRoleBinding)
	if !ok {
		return nil, fmt.Errorf("object is not *rbacv1.ClusterRoleBinding")
	}
	return h.createCRB(crb)
}

// CreateFromUnstructured creates clusterrolebinding from map[string]interface{}.
func (h *Handler) CreateFromUnstructured(u map[string]interface{}) (*rbacv1.ClusterRoleBinding, error) {
	crb := &rbacv1.ClusterRoleBinding{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u, crb)
	if err != nil {
		return nil, err
	}
	return h.createCRB(crb)
}

// createCRB
func (h *Handler) createCRB(crb *rbacv1.ClusterRoleBinding) (*rbacv1.ClusterRoleBinding, error) {
	crb.ResourceVersion = ""
	crb.UID = ""
	return h.clientset.RbacV1().ClusterRoleBindings().Create(h.ctx, crb, h.Options.CreateOptions)
}
