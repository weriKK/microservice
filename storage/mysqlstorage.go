package storage

import (
	"database/sql"
	"log"

	// blank import is needed by database/sql
	_ "github.com/go-sql-driver/mysql"

	"github.com/weriKK/microservice"
)

// MysqlStorageService ...
type MysqlStorageService struct {
	db *sql.DB
}

type subscriptionRecord struct {
	AfID                           string `json:"afID"`
	ReferenceID                    int    `json:"referenceID"`
	EventType                      string `json:"eventType"`
	NotificationDestinationAddress string `json:"notifDestAddr"`
	MaxReports                     int    `json:"maxReports"`
	SentReports                    int    `json:"sentReports"`
}

// NewMysqlStorageService ...
func NewMysqlStorageService() (microservice.SubscriptionStorage, error) {

	db, err := sql.Open("mysql", "root:trustno1@tcp(192.168.99.100:3306)/montesubscriptiondb")
	if err != nil {
		return nil, err
	}

	return &MysqlStorageService{
		db: db,
	}, nil
}

// Save ...
func (m *MysqlStorageService) Save(afID string, mes microservice.MonitoringEventSubscription) (referenceID int, err error) {
	log.Printf("storage.Save: %+v, %+v", afID, mes)

	stmt, err := m.db.Prepare("INSERT INTO `subscription` (`afID`, `eventType`, `notifDestAddr`, `maxReports`, `sentReports`) VALUES(?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(afID, mes.EventType, mes.NotificationDestinationAddress, mes.MaxReports, 0)
	if err != nil {
		return 0, err
	}

	// We have a compound primary key(afID, referenceID), but only one field of it can be auto increment,
	// and LastInsertId() returns this value (in our case referenceID)
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	referenceID = int(lastID)

	return referenceID, nil
}

// Get ...
func (m *MysqlStorageService) Get(afID string, referenceID int) (microservice.MonitoringEventSubscription, error) {
	log.Printf("storage.Get: %+v, %+v", afID, referenceID)

	query := "SELECT * FROM `subscription` WHERE afID=? AND referenceID=? LIMIT 1"
	row := m.db.QueryRow(query, afID, referenceID)

	var record subscriptionRecord
	err := row.Scan(
		&record.AfID,
		&record.ReferenceID,
		&record.EventType,
		&record.NotificationDestinationAddress,
		&record.MaxReports,
		&record.SentReports,
	)

	if err != nil {
		return microservice.MonitoringEventSubscription{}, err
	}

	return microservice.MonitoringEventSubscription{
		EventType:                      record.EventType,
		NotificationDestinationAddress: record.NotificationDestinationAddress,
		MaxReports:                     record.MaxReports,
	}, nil
}

// GetAll ...
func (m *MysqlStorageService) GetAll(afID string) (map[int]microservice.MonitoringEventSubscription, error) {
	log.Printf("storage.GetAll: %+v", afID)

	query := "SELECT * FROM `subscription` WHERE afID=?"
	rows, err := m.db.Query(query, afID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subs := map[int]microservice.MonitoringEventSubscription{}

	for rows.Next() {
		var record subscriptionRecord
		err = rows.Scan(
			&record.AfID,
			&record.ReferenceID,
			&record.EventType,
			&record.NotificationDestinationAddress,
			&record.MaxReports,
			&record.SentReports,
		)

		if err != nil {
			return nil, err
		}

		subs[record.ReferenceID] = microservice.MonitoringEventSubscription{
			EventType:                      record.EventType,
			NotificationDestinationAddress: record.NotificationDestinationAddress,
			MaxReports:                     record.MaxReports,
		}
	}

	return subs, nil
}

// Delete ...
func (m *MysqlStorageService) Delete(afID string, referenceID int) error {
	log.Printf("storage.Delete: %+v, %+v", afID, referenceID)

	query := "DELETE FROM `subscription` WHERE afID=? AND referenceID=?"
	_, err := m.db.Query(query, afID, referenceID)
	return err
}

// IncrementSentReportCount ...
func (m *MysqlStorageService) IncrementSentReportCount(afID string, referenceID int) error {
	log.Printf("storage.IncrementSentReportCount: %+v, %+v", afID, referenceID)

	query := "UPDATE `subscription` SET `sentReports` = `sentReports` + 1 WHERE afID=? AND referenceID=?"
	_, err := m.db.Query(query, afID, referenceID)
	return err
}

// Close ...
func (m *MysqlStorageService) Close() error {
	return m.db.Close()
}
