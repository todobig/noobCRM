package repositories

import (
    "context"
    "errors"

    "1billion/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// MongoDBLeadRepository is a MongoDB implementation of LeadRepository
type MongoDBLeadRepository struct {
    collection *mongo.Collection
}

// NewMongoDBLeadRepository creates a new MongoDBLeadRepository
func NewMongoDBLeadRepository(database *mongo.Database) *MongoDBLeadRepository {
    return &MongoDBLeadRepository{
        collection: database.Collection("leads"),
    }
}

// InsertLead inserts a new lead into the database
func (r *MongoDBLeadRepository) InsertLead(lead *models.Lead) error {
    _, err := r.collection.InsertOne(context.Background(), lead)
    return err
}

// FindLeadByPhone retrieves lead information by phone number
func (r *MongoDBLeadRepository) FindLeadByPhone(phoneNumber string) (*models.Lead, error) {
    var lead models.Lead
    err := r.collection.FindOne(context.Background(), bson.M{"phone": phoneNumber}).Decode(&lead)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, nil // Lead not found
        }
        return nil, err
    }
    return &lead, nil
}

// UpdateLeadByPhone updates lead information by phone number
func (r *MongoDBLeadRepository) UpdateLeadByPhone(phoneNumber string, lead *models.Lead) error {
    filter := bson.M{"phone": phoneNumber}
    update := bson.M{"": lead}
    _, err := r.collection.UpdateOne(context.Background(), filter, update)
    return err
}

// DeleteLeadByPhone deletes lead information by phone number
func (r *MongoDBLeadRepository) DeleteLeadByPhone(phoneNumber string) error {
    filter := bson.M{"phone": phoneNumber}
    _, err := r.collection.DeleteOne(context.Background(), filter)
    return err
}
