package main

import (
	"context"
)

// IDLManagementImpl implements the last service interface defined in the IDL.
type IDLManagementImpl struct{}

// CheckVersion implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) CheckVersion(ctx context.Context) (resp string, err error) {
	// TODO: Your code here...
	return
}

// GetServiceThriftFileName implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) GetServiceThriftFileName(ctx context.Context, serviceName string) (resp string, err error) {
	// TODO: Your code here...
	return
}

// GetThriftFile implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) GetThriftFile(ctx context.Context, serviceName string) (resp string, err error) {
	// TODO: Your code here...
	return
}
