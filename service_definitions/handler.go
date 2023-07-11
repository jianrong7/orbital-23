package main

import (
	"context"
)

// IDLManagementImpl implements the last service interface defined in the IDL.
type IDLManagementImpl struct{}

// GetThriftFile implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) GetThriftFile(ctx context.Context, serviceName string, serviceVersion string) (resp string, err error) {
	// TODO: Your code here...
	return
}
