package helper

import (
	"database/sql"

	"gitlab.udevs.io/template/template_notification_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(log logger.Logger, err error, message string, req interface{}) error {
	if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return status.Error(codes.NotFound, "Not Found")
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return status.Error(codes.Internal, message)
	}
	return nil
}

func ToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

func StringValue(ns sql.NullString) string {
	return ns.String
}
