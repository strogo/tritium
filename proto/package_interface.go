package proto

import(
	"slugtest/protoface"
	)

func (m *Type) IGetName() string {
	return m.GetString()
}

func (m *Type) IGetImplements() int32 {
	return m.GetImplements()
}

func (m *Package) IGetName() string {
	return m.GetName()
}

func (m *Package) IGetFunctions() []protoface.Function {
	return m.GetFunctions()
}

func (m *Package) IGetTypes() []protoface.Type {
	return m.GetTypes()
}

func (m *Package) IGetDependencies() []string {
	return m.GetDependencies()
}

func (m *Package) IGetPath() string {
	return m.GetPath()
}