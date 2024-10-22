package database

import "github.com/stretchr/testify/mock"

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) Connect() interface{} {
	ret := m.Called()

	r0 := ret.Get(0)
	if r0 == nil {
		return nil
	}
	return r0
}