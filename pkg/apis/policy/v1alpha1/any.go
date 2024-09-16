package v1alpha1

import (
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/jinzhu/copier"
)

// Any can be any type.
// +k8s:deepcopy-gen=false
type Any struct {
	value any
}

func (in *Any) DeepCopyInto(out *Any) {
	if err := copier.Copy(out, in); err != nil {
		panic("deep copy failed")
	}
}

func (in *Any) DeepCopy() *Any {
	if in == nil {
		return nil
	}
	out := new(Any)
	in.DeepCopyInto(out)
	return out
}

func (a *Any) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.value)
}

func (a *Any) UnmarshalJSON(data []byte) error {
	var v any
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a.value = v
	return nil
}
