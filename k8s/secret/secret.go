package secret

import (
	"github.com/wireltd/operator-pkg/util"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// New creates a new configMap
func New(namespace, name string, data map[string][]byte) *v12.Secret {
	return &v12.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}
}

// NewPassword creates a new password of length len
func NewPassword(len int) (string, error) {
	return util.RandomString(len)
}
