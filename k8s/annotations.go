/*
Copyright 2021 Absa Group Limited

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/
// Package k8s provides extensions for k8s apimachinery/pkg/apis
package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MergeAnnotations adds or updates annotations from defaultSource to defaultTarget
func MergeAnnotations(target *metav1.ObjectMeta, source *metav1.ObjectMeta) {
	if target.Annotations == nil {
		target.Annotations = make(map[string]string)
	}
	if source.Annotations == nil {
		source.Annotations = make(map[string]string)
	}
	for k, v := range source.Annotations {
		if target.Annotations[k] != v {
			target.Annotations[k] = v
		}
	}
}

// ContainsAnnotations checks if defaultTarget contains all annotations of defaultSource.
func ContainsAnnotations(target *metav1.ObjectMeta, source *metav1.ObjectMeta) bool {
	if source.Annotations == nil {
		return true
	}
	if target.Annotations == nil {
		return false
	}
	for k, v := range source.Annotations {
		if target.Annotations[k] != v {
			return false
		}
	}
	return true
}
