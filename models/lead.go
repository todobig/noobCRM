package models

// Lead represents lead information
type Lead struct {
    FirstName    string `json:"first_name"`
    LastName     string `json:"last_name"`
    Address      string `json:"address"`
    City         string `json:"city"`
    State        string `json:"state"`
    Zip          string `json:"zip"`
    Phone        string `json:"phone"`
    Email        string `json:"email"`
    // Add other lead fields here...
}
