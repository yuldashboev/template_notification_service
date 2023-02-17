package storage

import (
	"context"
	pb "genproto/notification_service"
)

type StorageI interface {
	CloseDB()
	Notification() NotificationStorageI
}

type NotificationStorageI interface {
	Create(ctx context.Context, notification *pb.Notification) (string, error)
	ChangeStatus(ctx context.Context, notificationID, receiverID, statusID string) error
	GetByStatus(ctx context.Context, statusID string) ([]*pb.NotificationReceivers, error)
	GetAllUserNotifications(ctx context.Context, receiverID string) ([]*pb.Notification, error)
	CreateFromAdmin(ctx context.Context, notification *pb.NotificationFromAdmin) error
	GetAllNotificationHistory(ctx context.Context, shipper, receiver_type string, limit, page uint64) (*pb.GetAllNotificationHistoryResponse, error)
}
