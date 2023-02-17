package service

import (
	"context"
	"fmt"
	pbc "genproto/courier_service"
	pb "genproto/notification_service"
	pbu "genproto/user_service"

	"gitlab.udevs.io/template/template_notification_service/config"
	"gitlab.udevs.io/template/template_notification_service/grpc/client"
	"gitlab.udevs.io/template/template_notification_service/pkg/helper"
	logger "gitlab.udevs.io/template/template_notification_service/pkg/logger"
	ws "gitlab.udevs.io/template/template_notification_service/socket"
	storage "gitlab.udevs.io/template/template_notification_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/rpc/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

// NotificationService struct
type NotificationService struct {
	cfg      config.Config
	storage  storage.StorageI
	logger   logger.Logger
	client   client.GrpcClientI
	wsClient ws.ClientI
}

// NewNotificationService constructor
func NewNotificationService(cfg config.Config, log logger.Logger, storage storage.StorageI, client client.GrpcClientI, wsClient ws.ClientI) *NotificationService {
	return &NotificationService{
		cfg:      cfg,
		storage:  storage,
		logger:   log,
		client:   client,
		wsClient: wsClient,
	}
}

// Create method
func (ns *NotificationService) Create(ctx context.Context, req *pb.Notification) (*empty.Empty, error) {
	ns.logger.Info("Sending ---", logger.Any("req: ", req))
	for _, rec := range req.Receivers {
		if rec.ReceiverType == config.COURIER {
			ns.logger.Info("GetCourier in notification create", logger.Any("resIDcourier- ", rec.ReceiverId))

			res, err := ns.client.CourierService().GetCourier(ctx, &pbc.GetCourierRequest{
				Id: rec.ReceiverId,
			})
			ns.logger.Info("GetCourier in notification create", logger.Any("res:--- ", res))
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "error while getting courier", logger.Any("courier_id", rec.ReceiverId))
			}

			if res.Courier.FcmToken.GetValue() != "" {
				thing, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
					Platform: gorush.PlatFormAndroid,
					Tokens:   []string{res.Courier.FcmToken.GetValue()},
					Title:    req.Title,
					Sound:    "default",
					Message:  req.Content,
					Data: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"title": {
								Kind: &structpb.Value_StringValue{StringValue: req.Title},
							},
							"body": {
								Kind: &structpb.Value_StringValue{StringValue: req.Content},
							},
							"sound": {
								Kind: &structpb.Value_StringValue{StringValue: "default"},
							},
						},
					},
				})

				ns.logger.Info("Returning gorush SERVICE ----", logger.Any("res: ", thing))

				if err != nil {
					return nil, helper.HandleError(ns.logger, err, "Error while sending notification to courier", logger.Any("receiver", rec))
				}

				ns.logger.Info("Notification sent to couriers", logger.Any("req: ", req))
			}
		} else if rec.ReceiverType == config.BRANCH_USER {
			branchUser, err := ns.client.BranchUserService().Get(ctx, &pbu.BranchUserId{
				Id: rec.ReceiverId,
			})

			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "error while getting branch user", logger.Any("branch_user_id:", rec.ReceiverId))
			}

			err = ns.wsClient.Push(req.Title, req.Content, req.StatusId, rec.ReceiverId)
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "error while sending push with Socket", logger.Any("branch_user_id:", rec.ReceiverId))
			}

			if branchUser.FcmToken.GetValue() != "" {
				if branchUser.PlatformId.GetValue() == config.ANDROID_PLATFORM_ID || branchUser.PlatformId.GetValue() == "" {
					_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
						Platform: gorush.PlatFormAndroid,
						Tokens:   []string{branchUser.FcmToken.GetValue()},
						Data: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"screen": {
									Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
								},
								"sound": {
									Kind: &structpb.Value_StringValue{StringValue: "default"},
								},
								"title": {
									Kind: &structpb.Value_StringValue{StringValue: req.Title},
								},
								"body": {
									Kind: &structpb.Value_StringValue{StringValue: req.Content},
								},
							},
						},
					})
					if err != nil {
						return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", rec))
					}

				} else if branchUser.PlatformId.GetValue() == config.IOS_PLATFORM_ID {
					_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
						Platform: gorush.PlatFormAndroid,
						Tokens:   []string{branchUser.FcmToken.GetValue()},
						Message:  req.Content,
						Title:    req.Title,
						Sound:    "default",
						Data: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"screen": {
									Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
								},
								"sound": {
									Kind: &structpb.Value_StringValue{StringValue: "default"},
								},
							},
						},
					})

					if err != nil {
						return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", rec))
					}
				}

				ns.logger.Info("Notification sent to branch users", logger.Any("req: ", req))
			}
		} else if rec.ReceiverType == config.CLIENT {
			customer, err := ns.client.CustomerService().Get(
				context.Background(),
				&pbu.CustomerId{
					Id: rec.ReceiverId,
				})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while getting customer", req)
			}

			if customer.FcmToken.GetValue() != "" {
				if customer.PlatformId.GetValue() == config.ANDROID_PLATFORM_ID {
					_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
						Platform: gorush.PlatFormAndroid,
						Tokens:   []string{customer.FcmToken.GetValue()},
						Title:    req.Title,
						Sound:    "default",
						Message:  req.Content,
						Data: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"screen": {
									Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
								},
								"sound": {
									Kind: &structpb.Value_StringValue{StringValue: "default"},
								},
								// "title": {
								// 	Kind: &structpb.Value_StringValue{StringValue: req.Title},
								// },
								// "body": {
								// 	Kind: &structpb.Value_StringValue{StringValue: req.Content},
								// },
							},
						},
					})
					if err != nil {
						return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", rec))
					}
				} else if customer.PlatformId.GetValue() == config.IOS_PLATFORM_ID {
					_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
						Platform: gorush.PlatFormAndroid,
						Tokens:   []string{customer.FcmToken.GetValue()},
						Title:    req.Title,
						Sound:    "default",
						Message:  req.Content,
						Data: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"screen": {
									Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
								},
								"sound": {
									Kind: &structpb.Value_StringValue{StringValue: "default"},
								},
							},
						},
					})
					if err != nil {
						return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", rec))
					}
				}
			}

			ns.logger.Info("Notification sent to client", logger.Any("req: ", req))
		}
	}

	return &empty.Empty{}, nil
}

func (ns *NotificationService) CreateForShipperCustomers(ctx context.Context, req *pb.Notification) (*empty.Empty, error) {
	fcmTokens, err := ns.client.CustomerService().GetFcmTokens(
		context.Background(),
		&pbu.GetFcmTokensRequest{
			ShipperId: req.ShipperId,
		})
	if err != nil {
		ns.logger.Info("Error while getting shippers all customers' fcm tokens", logger.Any("req: ", req))
		return nil, err
	}

	fmt.Printf("%+v", fcmTokens)

	_, err = ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
		Platform: gorush.PlatFormAndroid,
		Tokens:   fcmTokens.AndroidFcmTokens,
		Title:    req.Title,
		Sound:    "default",
		Message:  req.Content,
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"screen": {
					Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
				},
				"sound": {
					Kind: &structpb.Value_StringValue{StringValue: "default"},
				},
				// "title": {
				// 	Kind: &structpb.Value_StringValue{StringValue: req.Title},
				// },
				// "body": {
				// 	Kind: &structpb.Value_StringValue{StringValue: req.Content},
				// },
			},
		},
	})
	if err != nil {
		return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", ""))
	}

	_, err = ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
		Platform: gorush.PlatFormAndroid,
		Tokens:   fcmTokens.IosFcmTokens,
		Message:  req.Content,
		Title:    req.Title,
		Sound:    "default",
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"screen": {
					Kind: &structpb.Value_StringValue{StringValue: req.StatusId},
				},
				"sound": {
					Kind: &structpb.Value_StringValue{StringValue: "default"},
				},
			},
		},
	})
	if err != nil {
		return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", ""))
	}

	ns.logger.Info("Notification sent to client", logger.Any("req: ", req))

	return &empty.Empty{}, nil
}

// ChangeStatus method
func (ns *NotificationService) ChangeStatus(ctx context.Context, req *pb.ChangeStatusRequest) (*empty.Empty, error) {
	err := ns.storage.Notification().ChangeStatus(ctx, req.NotificationId, req.ReceiverId, req.StatusId)

	if err != nil {
		ns.logger.Error("error while updating notification status id", logger.Error(err))
		return nil, err
	}

	return &empty.Empty{}, nil
}

// GetByStatusId method
func (ns *NotificationService) GetByStatusId(statusId string) ([]*pb.NotificationReceivers, error) {
	return nil, nil
}

// GetAllUserNotifications method
func (ns *NotificationService) GetAllUserNotifications(ctx context.Context, in *pb.GetAllUserNotificationsRequest) (*pb.GetAllUserNotificationsResponse, error) {
	return nil, nil
}

func (ns *NotificationService) CreateFromAdmin(ctx context.Context, req *pb.NotificationFromAdmin) (*empty.Empty, error) {
	ns.logger.Info("Sending notification", logger.Any("req: ", req))

	if req.ReceiverType == config.COURIER {
		var tokens []string
		res, err := ns.client.CourierService().GetAllCouriersFmcTokens(ctx, &pbc.GetAllGetAllCouriersFmcTokensRequest{
			ReceiversIds: req.GetReceivers(),
			ShipperId:    req.GetShipperId(),
		})
		if err != nil {
			return nil, helper.HandleError(ns.logger, err, "error while getting couriers tokens", logger.Any("courier_ids", req.Receivers))
		}

		for _, id := range res.GetTokens() {
			if id != "" {
				tokens = append(tokens, id)
			}
		}

		if len(tokens) > 0 {
			_, err = ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
				Platform: gorush.PlatFormAndroid,
				Tokens:   tokens,
				Title:    req.Title,
				Sound:    "default",
				Message:  req.Content,
				Data: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"title": {
							Kind: &structpb.Value_StringValue{StringValue: req.Title},
						},
						"body": {
							Kind: &structpb.Value_StringValue{StringValue: req.Content},
						},
						"sound": {
							Kind: &structpb.Value_StringValue{StringValue: "default"},
						},
					},
				},
			})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while sending notification to courier", logger.Any("receiver", req))
			}

			ns.logger.Info("Notification sent to couriers", logger.Any("tokens: ", tokens))
		}

	} else if req.ReceiverType == config.BRANCH_USER {
		var (
			iosTokens, androidTokens []string
		)
		branchUser, err := ns.client.BranchUserService().GetAllBranchUsersFcmTokens(ctx, &pbu.GetAllBranchUsersFcmTokensRequest{
			ShipperId: req.GetShipperId(),
			UsersIds:  req.GetReceivers(),
		})

		if err != nil {
			return nil, helper.HandleError(ns.logger, err, "error while getting branch user", logger.Any("branch_user_id:", req.Receivers))
		}

		for _, item := range branchUser.GetTokens() {
			if item.GetPlatformId() == config.ANDROID_PLATFORM_ID && item.GetFcmToken() != "" {
				androidTokens = append(androidTokens, item.GetFcmToken())
			} else if item.GetPlatformId() == config.IOS_PLATFORM_ID && item.GetFcmToken() != "" {
				iosTokens = append(iosTokens, item.GetFcmToken())
			}
		}

		if len(androidTokens) > 0 {
			_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
				Platform: gorush.PlatFormAndroid,
				Tokens:   androidTokens,
				Data: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"sound": {
							Kind: &structpb.Value_StringValue{StringValue: "default"},
						},
						"title": {
							Kind: &structpb.Value_StringValue{StringValue: req.Title},
						},
						"body": {
							Kind: &structpb.Value_StringValue{StringValue: req.Content},
						},
					},
				},
			})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("androidTokens", androidTokens))
			}
		}
		if len(iosTokens) > 0 {
			_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
				Platform: gorush.PlatFormAndroid,
				Tokens:   iosTokens,
				Message:  req.Content,
				Title:    req.Title,
				Sound:    "default",
				Data: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"sound": {
							Kind: &structpb.Value_StringValue{StringValue: "default"},
						},
					},
				},
			})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("iosTokens", iosTokens))
			}
		}

		ns.logger.Info("Notification sent to branch users", logger.Any("req: ", req))

	} else if req.ReceiverType == config.CLIENT {
		customer, err := ns.client.CustomerService().GetFcmTokens(
			context.Background(),
			&pbu.GetFcmTokensRequest{
				ShipperId: req.GetShipperId(),
			})
		if err != nil {
			return nil, helper.HandleError(ns.logger, err, "Error while getting customer", req)
		}

		if len(customer.GetAndroidFcmTokens()) > 0 {
			_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
				Platform: gorush.PlatFormAndroid,
				Tokens:   customer.GetAndroidFcmTokens(),
				Title:    req.Title,
				Sound:    "default",
				Message:  req.Content,
				Data: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"sound": {
							Kind: &structpb.Value_StringValue{StringValue: "default"},
						},
					},
				},
			})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("tokens", customer.GetAndroidFcmTokens()))
			}
		}
		if len(customer.GetIosFcmTokens()) > 0 {
			_, err := ns.client.GorushService().Send(ctx, &proto.NotificationRequest{
				Platform: gorush.PlatFormAndroid,
				Tokens:   customer.GetIosFcmTokens(),
				Message:  req.Content,
				Title:    req.Title,
				Sound:    "default",
				Data: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"sound": {
							Kind: &structpb.Value_StringValue{StringValue: "default"},
						},
					},
				},
			})
			if err != nil {
				return nil, helper.HandleError(ns.logger, err, "Error while sending notification", logger.Any("receiver", req))
			}
		}

		ns.logger.Info("Notification sent to client", logger.Any("tokens: ", customer.GetIosFcmTokens()))
	}

	err := ns.storage.Notification().CreateFromAdmin(ctx, req)

	if err != nil {
		return nil, helper.HandleError(ns.logger, err, "Error while saving notifications", logger.Any("receiver", req))
	}

	return &empty.Empty{}, nil
}

// GetAllUserNotifications method
func (ns *NotificationService) GetAllNotificationHistory(ctx context.Context, in *pb.GetAllNotificationHistoryRequest) (*pb.GetAllNotificationHistoryResponse, error) {
	resp, err := ns.storage.Notification().GetAllNotificationHistory(ctx, in.GetShipperId(), in.GetReceiverType(), in.GetLimit(), in.GetPage())
	if err != nil {
		helper.HandleError(ns.logger, err, "error while getting all sms history on sms service", logger.Any("receiver", in))
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return resp, nil
}
