package admin

import (
    "time"
)

// SupportTicketStatus represents the status of a support ticket
type SupportTicketStatus string

const (
    TicketStatusOpen     SupportTicketStatus = "open"
    TicketStatusPending  SupportTicketStatus = "pending"
    TicketStatusResolved SupportTicketStatus = "resolved"
    TicketStatusClosed   SupportTicketStatus = "closed"
)

// SupportTicketPriority represents the priority of a support ticket
type SupportTicketPriority string

const (
    PriorityLow    SupportTicketPriority = "low"
    PriorityMedium SupportTicketPriority = "medium"
    PriorityHigh   SupportTicketPriority = "high"
    PriorityUrgent SupportTicketPriority = "urgent"
)

// SupportTicket represents a support ticket
type SupportTicket struct {
    ID          uint                 `gorm:"primaryKey" json:"id"`
    UserID      uint                 `gorm:"not null" json:"user_id"`
    Subject     string               `gorm:"not null" json:"subject"`
    Description string               `gorm:"not null" json:"description"`
    Status      SupportTicketStatus  `gorm:"not null;default:'open'" json:"status"`
    Priority    SupportTicketPriority `gorm:"not null;default:'medium'" json:"priority"`
    Category    string               `gorm:"not null" json:"category"`
    AssignedTo  *uint                `json:"assigned_to,omitempty"`
    ResolvedAt  *time.Time           `json:"resolved_at,omitempty"`
    ClosedAt    *time.Time           `json:"closed_at,omitempty"`
    CreatedAt   time.Time            `json:"created_at"`
    UpdatedAt   time.Time            `json:"updated_at"`
}

// SupportMessage represents a message in a support ticket
type SupportMessage struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    TicketID   uint      `gorm:"not null" json:"ticket_id"`
    UserID     uint      `gorm:"not null" json:"user_id"`
    Message    string    `gorm:"not null" json:"message"`
    IsInternal bool      `gorm:"not null;default:false" json:"is_internal"`
    CreatedAt  time.Time `json:"created_at"`
}

// AuditLog represents an audit log entry
type AuditLog struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    UserID      *uint     `json:"user_id,omitempty"`
    Action      string    `gorm:"not null" json:"action"`
    EntityType  string    `gorm:"not null" json:"entity_type"`
    EntityID    uint      `gorm:"not null" json:"entity_id"`
    Description string    `json:"description,omitempty"`
    IPAddress   string    `json:"ip_address,omitempty"`
    UserAgent   string    `json:"user_agent,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
}

// SystemMetric represents a system metric
type SystemMetric struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"not null" json:"name"`
    Value     float64   `gorm:"not null" json:"value"`
    Timestamp time.Time `gorm:"not null" json:"timestamp"`
    CreatedAt time.Time `json:"created_at"`
}