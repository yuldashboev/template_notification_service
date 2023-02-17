package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	pb "genproto/notification_service"

	"gitlab.udevs.io/template/template_notification_service/config"
	"gitlab.udevs.io/template/template_notification_service/pkg/helper"

	// postgres driver
	_ "github.com/lib/pq"
)

type notificationRepo struct {
	db *pgxpool.Pool
}

// NewNotificationRepo constructor
func NewNotificationRepo(db *pgxpool.Pool) *notificationRepo {
	return &notificationRepo{db: db}
}

func (nr *notificationRepo) Create(ctx context.Context, notification *pb.Notification) (string, error) {
	var (
		orderID sql.NullString
	)

	noteID, err := uuid.NewRandom()

	if err != nil {
		return "", err
	}

	orderID = helper.ToNullString(notification.OrderId)

	tx, err := nr.db.Begin(ctx)

	if err != nil {
		return "", err
	}

	_, err = tx.Exec(
		ctx,
		`INSERT INTO notifications(
                          id, 
                          content, 
                          sender_id, 
                          sender_type, 
                          order_id, 
                          shipper_id) 
				VALUES ($1, $2, $3, $4, $5, $6)`,
		noteID,
		notification.Content,
		notification.SenderId,
		notification.SenderType,
		orderID,
		notification.ShipperId,
	)

	if err != nil {
		tx.Rollback(ctx)
		return "", err
	}

	for _, rec := range notification.Receivers {
		id, err := uuid.NewRandom()

		if err != nil {
			tx.Rollback(ctx)
			return "", err
		}

		_, err = tx.Exec(
			ctx,
			`INSERT INTO notif_receivers(id, notif_id, receiver_id, receiver_type, notif_status_id)
			VALUES ($1, $2, $3, $4, $5)`, id, noteID, rec.ReceiverId, rec.ReceiverType, config.NewStatusId)

		if err != nil {
			tx.Rollback(ctx)
			return "", err
		}
	}
	tx.Commit(ctx)
	return noteID.String(), nil
}

func (nr *notificationRepo) ChangeStatus(ctx context.Context, notificationID, receiverID, statusID string) error {
	_, err := nr.db.Exec(ctx, "UPDATE notif_receivers SET notif_status_id=$1, send_at=current_timestamp WHERE notif_id=$2 AND receiver_id=$3",
		statusID, notificationID, receiverID)

	if err != nil {
		return err
	}

	return nil
}

func (nr *notificationRepo) GetByStatus(ctx context.Context, statusID string) ([]*pb.NotificationReceivers, error) {
	var (
		receivers []*pb.NotificationReceivers
	)

	return receivers, nil
}

// GetAllNotifications selects all notifications from the datbase
func (nr *notificationRepo) GetAllUserNotifications(ctx context.Context, receiverID string) (notifications []*pb.Notification, err error) {

	return
}

func (nr *notificationRepo) CreateFromAdmin(ctx context.Context, notification *pb.NotificationFromAdmin) error {
	var (
		toAll bool
	)

	noteID, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	tx, err := nr.db.Begin(ctx)

	if err != nil {
		return err
	}

	if len(notification.Receivers) > 0 {
		toAll = false
	}

	_, err = tx.Exec(
		ctx,
		`INSERT INTO notifications(
                          id, 
                          content, 
						  title,
                          sender_id, 
                          sender_type, 
                          shipper_id,
						  receiver_type,
						  to_all) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		noteID,
		notification.Content,
		notification.Title,
		notification.SenderId,
		"Shipper",
		notification.ShipperId,
		notification.ReceiverType,
		toAll,
	)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, receiverID := range notification.Receivers {
		id, err := uuid.NewRandom()

		if err != nil {
			tx.Rollback(ctx)
			return err
		}

		_, err = tx.Exec(
			ctx,
			`INSERT INTO notif_receivers(id, notif_id, receiver_id, receiver_type, notif_status_id)
			VALUES ($1, $2, $3, $4, $5)`, id, noteID, receiverID, notification.ReceiverType, config.SentStatusId)

		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	tx.Commit(ctx)
	return nil
}

func (n *notificationRepo) GetAllNotificationHistory(ctx context.Context, shipper, receiverType string, limit, page uint64) (*pb.GetAllNotificationHistoryResponse, error) {
	var (
		nots  []*pb.NotificationHistoryResponse
		count uint64
		query string
	)
	offset := (page - 1) * limit
	params := pgx.NamedArgs{
		"limit":         limit,
		"offset":        offset,
		"shipper":       shipper,
		"receiver_type": receiverType,
	}

	// filter = `LIMIT :limit OFFSET :offset`
	// if limit != 0 && offset != 0 {
	// 	filter = `LIMIT :limit OFFSET :offset`
	// }

	if receiverType != "" {
		query = " AND n.receiver_type = @receiver_type "
	}

	rows, err := n.db.Query(
		ctx, `
		SELECT  
			count(n.id) OVER(),
			count(nr.id) as receivers_count,
			n.id,  
			content,
			title,
			n.receiver_type,
			to_all,
			to_char(n.created_at, 'YYYY-MM-DD HH24::MI::SS')
		FROM notifications n
		left join notif_receivers nr on nr.notif_id = n.id
		WHERE shipper_id = @shipper `+query+`
		GROUP BY n.id
		ORDER BY created_at DESC limit @limit offset @offset`,
		params,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			el           pb.NotificationHistoryResponse
			title, rType *string
		)
		if err := rows.Scan(
			&count,
			&el.ReceiversCount,
			&el.Id,
			&el.Content,
			&title,
			&rType,
			&el.ToAll,
			&el.Time,
		); err != nil {
			rows.Close()
			return nil, err
		}
		if title != nil {
			el.Title = *title
		}
		if rType != nil {
			el.ReceiverType = *rType
		}

		nots = append(nots, &el)
	}
	resp := pb.GetAllNotificationHistoryResponse{
		Notifications: nots,
		Count:         uint32(count),
	}
	return &resp, err
}
