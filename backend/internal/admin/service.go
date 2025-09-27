package admin

import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// AdminService handles admin operations
type AdminService struct {
    db *gorm.DB
}

// NewAdminService creates a new admin service
func NewAdminService(db *gorm.DB) *AdminService {
    return &AdminService{
        db: db,
    }
}

// InitializeDB initializes the database tables for admin service
func (s *AdminService) InitializeDB() error {
    return s.db.AutoMigrate(&SupportTicket{}, &SupportMessage{}, &AuditLog{}, &SystemMetric{})
}

// CreateSupportTicket creates a new support ticket
func (s *AdminService) CreateSupportTicket(ctx context.Context, ticket SupportTicket) (*SupportTicket, error) {
    ticket.CreatedAt = time.Now()
    ticket.UpdatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(&ticket).Error; err != nil {
        return nil, fmt.Errorf("failed to create support ticket: %w", err)
    }
    
    return &ticket, nil
}

// GetSupportTicket retrieves a support ticket by ID
func (s *AdminService) GetSupportTicket(ctx context.Context, ticketID uint) (*SupportTicket, error) {
    var ticket SupportTicket
    if err := s.db.WithContext(ctx).First(&ticket, "id = ?", ticketID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, fmt.Errorf("support ticket not found")
        }
        return nil, fmt.Errorf("failed to get support ticket: %w", err)
    }
    
    return &ticket, nil
}

// UpdateSupportTicket updates a support ticket
func (s *AdminService) UpdateSupportTicket(ctx context.Context, ticketID uint, updates map[string]interface{}) error {
    updates["updated_at"] = time.Now()
    return s.db.WithContext(ctx).Model(&SupportTicket{}).Where("id = ?", ticketID).Updates(updates).Error
}

// GetSupportTicketsByUser retrieves support tickets for a user
func (s *AdminService) GetSupportTicketsByUser(ctx context.Context, userID uint, limit, offset int) ([]SupportTicket, error) {
    var tickets []SupportTicket
    if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&tickets).Error; err != nil {
        return nil, fmt.Errorf("failed to get support tickets: %w", err)
    }
    
    return tickets, nil
}

// GetSupportTicketsByStatus retrieves support tickets by status
func (s *AdminService) GetSupportTicketsByStatus(ctx context.Context, status SupportTicketStatus, limit, offset int) ([]SupportTicket, error) {
    var tickets []SupportTicket
    if err := s.db.WithContext(ctx).Where("status = ?", status).Order("created_at DESC").Limit(limit).Offset(offset).Find(&tickets).Error; err != nil {
        return nil, fmt.Errorf("failed to get support tickets: %w", err)
    }
    
    return tickets, nil
}

// AddSupportMessage adds a message to a support ticket
func (s *AdminService) AddSupportMessage(ctx context.Context, message SupportMessage) (*SupportMessage, error) {
    message.CreatedAt = time.Now()
    
    if err := s.db.WithContext(ctx).Create(&message).Error; err != nil {
        return nil, fmt.Errorf("failed to add support message: %w", err)
    }
    
    return &message, nil
}

// GetSupportMessages retrieves messages for a support ticket
func (s *AdminService) GetSupportMessages(ctx context.Context, ticketID uint) ([]SupportMessage, error) {
    var messages []SupportMessage
    if err := s.db.WithContext(ctx).Where("ticket_id = ?", ticketID).Order("created_at ASC").Find(&messages).Error; err != nil {
        return nil, fmt.Errorf("failed to get support messages: %w", err)
    }
    
    return messages, nil
}

// CreateAuditLog creates a new audit log entry
func (s *AdminService) CreateAuditLog(ctx context.Context, log AuditLog) error {
    log.CreatedAt = time.Now()
    
    return s.db.WithContext(ctx).Create(&log).Error
}

// GetAuditLogs retrieves audit logs
func (s *AdminService) GetAuditLogs(ctx context.Context, limit, offset int) ([]AuditLog, error) {
    var logs []AuditLog
    if err := s.db.WithContext(ctx).Order("created_at DESC").Limit(limit).Offset(offset).Find(&logs).Error; err != nil {
        return nil, fmt.Errorf("failed to get audit logs: %w", err)
    }
    
    return logs, nil
}

// GetAuditLogsByUser retrieves audit logs for a user
func (s *AdminService) GetAuditLogsByUser(ctx context.Context, userID uint, limit, offset int) ([]AuditLog, error) {
    var logs []AuditLog
    if err := s.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Offset(offset).Find(&logs).Error; err != nil {
        return nil, fmt.Errorf("failed to get audit logs: %w", err)
    }
    
    return logs, nil
}

// CreateSystemMetric creates a new system metric
func (s *AdminService) CreateSystemMetric(ctx context.Context, metric SystemMetric) error {
    metric.CreatedAt = time.Now()
    
    return s.db.WithContext(ctx).Create(&metric).Error
}

// GetSystemMetrics retrieves system metrics
func (s *AdminService) GetSystemMetrics(ctx context.Context, limit, offset int) ([]SystemMetric, error) {
    var metrics []SystemMetric
    if err := s.db.WithContext(ctx).Order("timestamp DESC").Limit(limit).Offset(offset).Find(&metrics).Error; err != nil {
        return nil, fmt.Errorf("failed to get system metrics: %w", err)
    }
    
    return metrics, nil
}
