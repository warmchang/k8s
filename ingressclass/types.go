package ingressclass

import "errors"

var (
	ErrInvalidToolsType  = errors.New("type must be string, *networkingv1.IngressClass, networkingv1.IngressClass or runtime.Object")
	ErrInvalidCreateType = errors.New("type must be string, []byte, *networkingv1.IngressClass, networkingv1.IngressClass, runtime.Object, *unstructured.Unstructured, unstructured.Unstructured or map[string]interface{}")
	ErrInvalidUpdateType = ErrInvalidCreateType
	ErrInvalidApplyType  = ErrInvalidCreateType
	ErrInvalidDeleteType = ErrInvalidCreateType
	ErrInvalidGetType    = ErrInvalidCreateType
	ErrInvalidPathType   = errors.New("path data type must be string, []byte, *networkingv1.IngressClass, networkingv1.IngressClass, runtime.Object, *unstructured.Unstructured, unstructured.Unstructured or map[string]interface{}")
)
