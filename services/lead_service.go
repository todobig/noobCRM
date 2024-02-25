package services

import (
    "1billion/models"
    "1billion/repositories"
)

// LeadService contains methods for lead operations
type LeadService struct {
    leadRepository repositories.LeadRepository
}

// NewLeadService creates a new LeadService with the provided LeadRepository
func NewLeadService(leadRepository repositories.LeadRepository) *LeadService {
    return &LeadService{leadRepository}
}

// AddLead adds a new lead
func (s *LeadService) AddLead(lead *models.Lead) error {
    return s.leadRepository.InsertLead(lead)
}

// GetLeadByPhone retrieves lead information by phone number
func (s *LeadService) GetLeadByPhone(phoneNumber string) (*models.Lead, error) {
    return s.leadRepository.FindLeadByPhone(phoneNumber)
}

// UpdateLeadByPhone updates lead information by phone number
func (s *LeadService) UpdateLeadByPhone(phoneNumber string, lead *models.Lead) error {
    return s.leadRepository.UpdateLeadByPhone(phoneNumber, lead)
}

// DeleteLeadByPhone deletes lead information by phone number
func (s *LeadService) DeleteLeadByPhone(phoneNumber string) error {
    return s.leadRepository.DeleteLeadByPhone(phoneNumber)
}
