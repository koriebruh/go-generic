package trial_new

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

// //
type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

///

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (vp *MyVicePresident) GetName() string {
	return vp.Name
}

func (vp *MyVicePresident) GetVicePresidentName() string {
	return vp.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Jamal", GetName[Manager](&MyManager{Name: "Jamal"}))
	assert.Equal(t, "Jamal", GetName[VicePresident](&MyVicePresident{Name: "Jamal"}))
}
