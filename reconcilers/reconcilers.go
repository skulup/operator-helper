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

package reconcilers

import (
	"fmt"
	"github.com/skulup/operator-helper/reconciler"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Configure let the added reconcilers to configure themselves
func Configure(manager ctrl.Manager, reconcilers ...reconciler.Reconciler) error {
	ctx := reconciler.NewContext(manager)
	for _, r := range reconcilers {
		fmt.Printf("configuring the reconciler: %T\n", r)
		if err := r.Configure(ctx); err != nil {
			return err
		}
	}
	return nil
}
