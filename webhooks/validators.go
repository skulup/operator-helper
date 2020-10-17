/*
 * Copyright 2020 Skulup Ltd, Open Collaborators
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package webhooks

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Validate is a helper function to validate a CR with a error
// list object to aggregate the error and return the invalid error or nil depending
// if the validator func added at least one error to the passed error list object
func Validate(gvk schema.GroupVersionKind, name string, validatorFuncs ...func(list *ErrorList)) error {
	errList := &ErrorList{list: make([]*field.Error, 0)}
	for _, Func := range validatorFuncs {
		Func(errList)
	}
	if len(errList.list) == 0 {
		return nil
	}
	return errors.NewInvalid(gvk.GroupKind(), name, errList.list)
}

// ErrorList is a wrapper of field.Error
// to enable convenient error adding
type ErrorList struct {
	list []*field.Error
}

// Add adds the specified error the error list array
func (e *ErrorList) Add(err *field.Error) {
	e.list = append(e.list, err)
}
