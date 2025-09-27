package payments

import (
    "errors"

    "gorm.io/gorm"
)

// Service provides payment operations
type Service struct {
    db *gorm.DB
}

// NewService creates a new payment service
func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

// CreatePaymentRecord creates a new payment record
func (s *Service) CreatePaymentRecord(record *PaymentRecord) error {
    if record.ExternalID == "" || record.Processor == "" || record.Amount <= 0 {
        return errors.New("external ID, processor, and amount are required")
    }

    return s.db.Create(record).Error
}

// GetPaymentRecordByID retrieves a payment record by ID
func (s *Service) GetPaymentRecordByID(id uint) (*PaymentRecord, error) {
    var record PaymentRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetPaymentRecordByExternalID retrieves a payment record by external ID
func (s *Service) GetPaymentRecordByExternalID(externalID string) (*PaymentRecord, error) {
    var record PaymentRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdatePaymentRecord updates a payment record
func (s *Service) UpdatePaymentRecord(record *PaymentRecord) error {
    return s.db.Save(record).Error
}

// CreateCustomerRecord creates a new customer record
func (s *Service) CreateCustomerRecord(record *CustomerRecord) error {
    if record.ExternalID == "" || record.Processor == "" || record.Email == "" {
        return errors.New("external ID, processor, and email are required")
    }

    return s.db.Create(record).Error
}

// GetCustomerRecordByID retrieves a customer record by ID
func (s *Service) GetCustomerRecordByID(id uint) (*CustomerRecord, error) {
    var record CustomerRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetCustomerRecordByExternalID retrieves a customer record by external ID
func (s *Service) GetCustomerRecordByExternalID(externalID string) (*CustomerRecord, error) {
    var record CustomerRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdateCustomerRecord updates a customer record
func (s *Service) UpdateCustomerRecord(record *CustomerRecord) error {
    return s.db.Save(record).Error
}

// CreatePaymentMethodRecord creates a new payment method record
func (s *Service) CreatePaymentMethodRecord(record *PaymentMethodRecord) error {
    if record.ExternalID == "" || record.Processor == "" || record.Type == "" {
        return errors.New("external ID, processor, and type are required")
    }

    return s.db.Create(record).Error
}

// GetPaymentMethodRecordByID retrieves a payment method record by ID
func (s *Service) GetPaymentMethodRecordByID(id uint) (*PaymentMethodRecord, error) {
    var record PaymentMethodRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetPaymentMethodRecordByExternalID retrieves a payment method record by external ID
func (s *Service) GetPaymentMethodRecordByExternalID(externalID string) (*PaymentMethodRecord, error) {
    var record PaymentMethodRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdatePaymentMethodRecord updates a payment method record
func (s *Service) UpdatePaymentMethodRecord(record *PaymentMethodRecord) error {
    return s.db.Save(record).Error
}