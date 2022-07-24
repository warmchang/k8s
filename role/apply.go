package role

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
)

// Apply applies role from type string, []byte, *rbacv1.Role,
// rbacv1.Role, runtime.Object or map[string]interface{}.
func (h *Handler) Apply(obj interface{}) (*rbacv1.Role, error) {
	switch val := obj.(type) {
	case string:
		return h.ApplyFromFile(val)
	case []byte:
		return h.ApplyFromBytes(val)
	case *rbacv1.Role:
		return h.ApplyFromObject(val)
	case rbacv1.Role:
		return h.ApplyFromObject(&val)
	case runtime.Object:
		return h.ApplyFromObject(val)
	case map[string]interface{}:
		return h.ApplyFromUnstructured(val)
	default:
		return nil, ERR_TYPE_APPLY
	}
}

// ApplyFromFile applies role from yaml file.
func (h *Handler) ApplyFromFile(filename string) (role *rbacv1.Role, err error) {
	role, err = h.CreateFromFile(filename)
	if k8serrors.IsAlreadyExists(err) { // if role already exist, update it.
		role, err = h.UpdateFromFile(filename)
	}
	return
}

// ApplyFromBytes pply role from bytes.
func (h *Handler) ApplyFromBytes(data []byte) (role *rbacv1.Role, err error) {
	role, err = h.CreateFromBytes(data)
	if k8serrors.IsAlreadyExists(err) {
		role, err = h.UpdateFromBytes(data)
	}
	return
}

// ApplyFromObject applies role from runtime.Object.
func (h *Handler) ApplyFromObject(obj runtime.Object) (*rbacv1.Role, error) {
	role, ok := obj.(*rbacv1.Role)
	if !ok {
		return nil, fmt.Errorf("object is not *rbacv1.Role")
	}
	return h.applyRole(role)
}

// ApplyFromUnstructured applies role from map[string]interface{}.
func (h *Handler) ApplyFromUnstructured(u map[string]interface{}) (*rbacv1.Role, error) {
	role := &rbacv1.Role{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u, role)
	if err != nil {
		return nil, err
	}
	return h.applyRole(role)
}

// applyRole
func (h *Handler) applyRole(role *rbacv1.Role) (*rbacv1.Role, error) {
	_, err := h.createRole(role)
	if k8serrors.IsAlreadyExists(err) {
		return h.updateRole(role)
	}
	return role, err
}
