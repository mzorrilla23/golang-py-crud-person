package services

import (
	"context"
	"errors"

	"example.com/sarang-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonServiceImpl struct {
	personcollection *mongo.Collection
	ctx              context.Context
}

func NewPersonService(personcollection *mongo.Collection, ctx context.Context) PersonService {
	return &PersonServiceImpl{
		personcollection: personcollection,
		ctx:              ctx,
	}
}

func (p *PersonServiceImpl) CreatePerson(person *models.Person) error {
	_, err := p.personcollection.InsertOne(p.ctx, person)
	return err
}

func (p *PersonServiceImpl) GetPerson(id *int) (*models.Person, error) {
	var person *models.Person
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := p.personcollection.FindOne(p.ctx, query).Decode(&person)
	return person, err
}

func (p *PersonServiceImpl) GetAll() ([]*models.Person, error) {
	var persons []*models.Person
	cursor, err := p.personcollection.Find(p.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(p.ctx) {
		var person models.Person
		err := cursor.Decode(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, &person)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(p.ctx)
	if len(persons) == 0 {
		return nil, errors.New("documents not found")
	}
	return persons, nil
}

func (p *PersonServiceImpl) UpdatePerson(person *models.Person) error {
	filter := bson.D{bson.E{Key: "id", Value: person.Id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: person.Id}, bson.E{Key: "first_name", Value: person.FirstName}, bson.E{Key: "last_name", Value: person.LastName}}}}
	result, _ := p.personcollection.UpdateOne(p.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("error !!!!")
	}
	return nil
}

func (p *PersonServiceImpl) DeletePerson(id *int) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := p.personcollection.DeleteOne(p.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("error !!!!")
	}
	return nil
}
