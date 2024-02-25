package handlers

import (
    "encoding/json"
    "net/http"
    
    "1billion/services"
    "1billion/models"
    
    "1billion/utils"
)

// LeadHandler handles HTTP requests related to leads
type LeadHandler struct {
    leadService *services.LeadService
}

// NewLeadHandler creates a new LeadHandler with the provided LeadService
func NewLeadHandler(leadService *services.LeadService) *LeadHandler {
    return &LeadHandler{leadService}
}

// AddLead handles adding new leads
func (h *LeadHandler) AddLead(w http.ResponseWriter, r *http.Request) {
    var lead models.Lead
    if err := json.NewDecoder(r.Body).Decode(&lead); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.leadService.AddLead(&lead); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusCreated, lead)
}

// GetLead handles retrieving lead information by phone number
func (h *LeadHandler) GetLead(w http.ResponseWriter, r *http.Request) {
    phoneNumber := r.URL.Query().Get("phone")
    if phoneNumber == "" {
        utils.RespondWithError(w, http.StatusBadRequest, "Phone number is required")
        return
    }

    lead, err := h.leadService.GetLeadByPhone(phoneNumber)
    if err != nil {
        utils.RespondWithError(w, http.StatusNotFound, "Lead not found")
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, lead)
}

// UpdateLead handles updating lead information
func (h *LeadHandler) UpdateLead(w http.ResponseWriter, r *http.Request) {
    phoneNumber := r.URL.Query().Get("phone")
    if phoneNumber == "" {
        utils.RespondWithError(w, http.StatusBadRequest, "Phone number is required")
        return
    }

    var lead models.Lead
    if err := json.NewDecoder(r.Body).Decode(&lead); err != nil {
        utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.leadService.UpdateLeadByPhone(phoneNumber, &lead); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, lead)
}

// DeleteLead handles deleting lead information
func (h *LeadHandler) DeleteLead(w http.ResponseWriter, r *http.Request) {
    phoneNumber := r.URL.Query().Get("phone")
    if phoneNumber == "" {
        utils.RespondWithError(w, http.StatusBadRequest, "Phone number is required")
        return
    }

    if err := h.leadService.DeleteLeadByPhone(phoneNumber); err != nil {
        utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Lead deleted successfully"})
}
