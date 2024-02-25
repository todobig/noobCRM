package repositories

import "1billion/models"

// LeadRepository defines methods for lead data access
type LeadRepository interface {
    InsertLead(lead *models.Lead) error
    FindLeadByPhone(phoneNumber string) (*models.Lead, error)
    UpdateLeadByPhone(phoneNumber string, lead *models.Lead) error
    DeleteLeadByPhone(phoneNumber string) error
}
