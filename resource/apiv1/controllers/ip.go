package controllers

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mohrezfadaei/ipresist/internal/db"
	"github.com/mohrezfadaei/ipresist/utils"
)

var ValidStatuses = map[string]db.IPStatus{
	"active":    db.Active,
	"blocked":   db.Blocked,
	"suspended": db.Suspended,
}

// IPController handles the business logic for IP operations
type IPController struct{}

// GetAll retrieves all IP addresses with optional filters
func (ctrl IPController) GetAll(status, sort string, offset, limit int) ([]db.IP, error) {
	var ips []db.IP
	query := db.DB

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sort != "" {
		if sort[0] == '-' {
			query = query.Order(sort[1:] + " DESC")
		} else {
			query = query.Order(sort + " ASC")
		}
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&ips).Error; err != nil {
		return nil, err
	}
	return ips, nil
}

// GetByID retrieves a specific IP address by its UUID
func (ctrl IPController) GetByID(id uuid.UUID) (*db.IP, error) {
	var ip db.IP
	if err := db.DB.Where("id = ?", id).First(&ip).Error; err != nil {
		return nil, err
	}
	return &ip, nil
}

// Create adds a new IP address to the database
func (ctrl IPController) Create(ip *db.IP) error {
	// Validate status
	if _, ok := ValidStatuses[string(ip.Status)]; !ok {
		return errors.New("invalid status value")
	}
	return db.DB.Create(ip).Error
}

// Update modifies an existing IP address
func (ctrl IPController) Update(id uuid.UUID, data map[string]interface{}) (*db.IP, error) {
	var ip db.IP
	if err := db.DB.Where("id = ?", id).First(&ip).Error; err != nil {
		return nil, err
	}

	if ipAddress, ok := data["ipaddress"].(string); ok {
		ip.IPAddress = ipAddress
	}
	if note, ok := data["note"].(string); ok {
		ip.Note = note
	}
	if status, ok := data["status"].(string); ok {
		if _, valid := ValidStatuses[status]; !valid {
			return nil, errors.New("invalid status value")
		}
		ip.Status = db.IPStatus(status)
	}
	ip.LastUpdatedAt = new(time.Time)
	*ip.LastUpdatedAt = utils.Now()

	if err := db.DB.Save(&ip).Error; err != nil {
		return nil, err
	}
	return &ip, nil
}

// Delete removes an IP address by its UUID
func (ctrl IPController) Delete(id uuid.UUID) error {
	return db.DB.Where("id = ?", id).Delete(&db.IP{}).Error
}
