package v1alpha1

import (
	"github.com/kyverno/kyverno-json/pkg/core/assertion"
	"github.com/kyverno/kyverno-json/pkg/core/templating"
	"k8s.io/apimachinery/pkg/util/json"
)

// +k8s:deepcopy-gen=false
// +kubebuilder:validation:XPreserveUnknownFields
// +kubebuilder:validation:Type:=""
// AssertionTree represents an assertion tree.
type AssertionTree struct {
	_tree any
}

func NewAssertionTree(value any) AssertionTree {
	return AssertionTree{
		_tree: value,
	}
}

func (t *AssertionTree) Assertion(compiler templating.Compiler) (assertion.Assertion, error) {
	if t._tree == nil {
		return nil, nil
	}
	return assertion.Parse(t._tree, compiler)
}

func (a *AssertionTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(a._tree)
}

func (a *AssertionTree) UnmarshalJSON(data []byte) error {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a._tree = v
	return nil
}

func (in *AssertionTree) DeepCopyInto(out *AssertionTree) {
	out._tree = deepCopy(in._tree)
}