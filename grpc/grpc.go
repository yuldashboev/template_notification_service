package grpc

import (
	"genproto/notification_service"
	"net/http"

	"gitlab.udevs.io/template/template_notification_service/config"
	"gitlab.udevs.io/template/template_notification_service/grpc/client"
	"gitlab.udevs.io/template/template_notification_service/grpc/service"
	"gitlab.udevs.io/template/template_notification_service/pkg/logger"
	"gitlab.udevs.io/template/template_notification_service/socket"

	//ws "gitlab.udevs.io/template/template_notification_service/socket"
	"gitlab.udevs.io/template/template_notification_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.Logger, storage storage.StorageI, svcs client.GrpcClientI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	//websocket
	hub := socket.NewHub()
	wsClient := socket.NewClient(hub)

	http.HandleFunc("/v1/notification/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(hub, w, r)
	})
	http.HandleFunc("/", socket.ServeHome)

	notification_service.RegisterNotificationServiceServer(grpcServer, service.NewNotificationService(cfg, log, storage, svcs, wsClient))

	go func() {
		log.Info("http: server running", logger.String("port", cfg.HTTPPort))

		err := http.ListenAndServe(cfg.HTTPPort, nil)
		if err != nil {
			log.Fatal("Error while listenning to http server: %v", logger.Error(err))
		}
	}()

	go hub.Run()

	reflection.Register(grpcServer)
	return grpcServer
}
