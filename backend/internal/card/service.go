package card

import (
    "errors"

    "gorm.io/gorm"
)

// Service provides card operations
type Service struct {
    db *gorm.DB
}

// NewService creates a new card service
func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

// CreateCardRecord creates a new card record
func (s *Service) CreateCardRecord(record *CardRecord) error {
    if record.ExternalID == "" || record.Issuer == "" || record.CardholderID == "" {
        return errors.New("external ID, issuer, and cardholder ID are required")
    }

    return s.db.Create(record).Error
}

// GetCardRecordByID retrieves a card record by ID
func (s *Service) GetCardRecordByID(id uint) (*CardRecord, error) {
    var record CardRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetCardRecordByExternalID retrieves a card record by external ID
func (s *Service) GetCardRecordByExternalID(externalID string) (*CardRecord, error) {
    var record CardRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdateCardRecord updates a card record
func (s *Service) UpdateCardRecord(record *CardRecord) error {
    return s.db.Save(record).Error
}

// CreateCardholderRecord creates a new cardholder record
func (s *Service) CreateCardholderRecord(record *CardholderRecord) error {
    if record.ExternalID == "" || record.Issuer == "" || record.Email == "" {
        return errors.New("external ID, issuer, and email are required")
    }

    return s.db.Create(record).Error
}

// GetCardholderRecordByID retrieves a cardholder record by ID
func (s *Service) GetCardholderRecordByID(id uint) (*CardholderRecord, error) {
    var record CardholderRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetCardholderRecordByExternalID retrieves a cardholder record by external ID
func (s *Service) GetCardholderRecordByExternalID(externalID string) (*CardholderRecord, error) {
    var record CardholderRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdateCardholderRecord updates a cardholder record
func (s *Service) UpdateCardholderRecord(record *CardholderRecord) error {
    return s.db.Save(record).Error
}

// CreateTransactionRecord creates a new transaction record
func (s *Service) CreateTransactionRecord(record *TransactionRecord) error {
    if record.ExternalID == "" || record.Issuer == "" || record.CardID == "" {
        return errors.New("external ID, issuer, and card ID are required")
    }

    return s.db.Create(record).Error
}

// GetTransactionRecordByID retrieves a transaction record by ID
func (s *Service) GetTransactionRecordByID(id uint) (*TransactionRecord, error) {
    var record TransactionRecord
    err := s.db.First(&record, id).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// GetTransactionRecordByExternalID retrieves a transaction record by external ID
func (s *Service) GetTransactionRecordByExternalID(externalID string) (*TransactionRecord, error) {
    var record TransactionRecord
    err := s.db.Where("external_id = ?", externalID).First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

// UpdateTransactionRecord updates a transaction record
func (s *Service) UpdateTransactionRecord(record *TransactionRecord) error {
    return s.db.Save(record).Error
}