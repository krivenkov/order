package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound   = status.Error(codes.NotFound, "not found")
	ErrMultiItems = status.Error(codes.InvalidArgument, "multi items")
)
