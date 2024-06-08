package gophermongo

import (
	"context"
	"errors"

	"github.com/xfrr/gophersys/internal/gopher"
	"github.com/xfrr/gophersys/pkg/mongoutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongomigrations "github.com/xfrr/gophersys/internal/storage/mongodb/migrations"
)

var (
	ErrExistsFilterRequired = errors.New("at least one filter is required")
)

const (
	mongoDatabaseName   = "gophers"
	mongoCollectionName = "gophers"
)

var _ gopher.Repository = (*MongoDBRepository)(nil)

type MongoDBRepository struct {
	gophers *mongo.Collection
}

func NewRepository(db *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{
		gophers: db.Database(mongoDatabaseName).Collection(mongoCollectionName),
	}
}

func (r *MongoDBRepository) Save(ctx context.Context, g *gopher.Aggregate) error {
	dto := gopherToDTO(*g)
	filter := bson.M{"_id": dto.ID}
	update := bson.M{"$set": dto}
	opts := options.Update().SetUpsert(true)
	_, err := r.gophers.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *MongoDBRepository) Delete(ctx context.Context, id gopher.ID) error {
	filter := bson.M{"_id": id}
	_, err := r.gophers.DeleteOne(ctx, filter)
	return err
}

func (r *MongoDBRepository) Exists(ctx context.Context, fopts ...gopher.Filter) (bool, error) {
	gopherFilter := gopher.Filters{}
	for _, filter := range fopts {
		filter(&gopherFilter)
	}

	filters := makeGopherQueryFilters(gopherFilter)
	if len(filters) == 0 {
		return false, ErrExistsFilterRequired
	}

	var query []bson.D
	query = append(query, makeMatchOrGopherQuery(filters))
	query = append(query, makeCountGopherQuery())

	cursor, err := r.gophers.Aggregate(ctx, mongo.Pipeline(query))
	if err != nil {
		return false, err
	}

	var result struct {
		Count int `bson:"count"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return false, err
		}
	}

	return result.Count > 0, nil
}

func (r *MongoDBRepository) Get(ctx context.Context, fopts ...gopher.Filter) (*gopher.Aggregate, error) {
	gopherFilter := gopher.Filters{}
	for _, filter := range fopts {
		filter(&gopherFilter)
	}

	filters := makeGopherQueryFilters(gopherFilter)
	if len(filters) == 0 {
		return nil, ErrExistsFilterRequired
	}

	var query []bson.D
	query = append(query, makeMatchOrGopherQuery(filters))

	cursor, err := r.gophers.Aggregate(ctx, mongo.Pipeline(query))
	if err != nil {
		return nil, err
	}

	var res gopherDTO
	if cursor.Next(ctx) {
		if err := cursor.Decode(&res); err != nil {
			return nil, err
		}
	}

	return dtoToGopher(res)
}

func (r *MongoDBRepository) Search(ctx context.Context, query gopher.SearchQuery) ([]*gopher.Aggregate, error) {
	var filters []bson.D
	if query.Names != nil {
		filters = append(filters, makeNamesFilter(query.Names))
	}

	if query.Statuses != nil {
		filters = append(filters, makeStatusesFilter(query.Statuses))
	}

	var mongoQuery []bson.D
	if len(filters) > 0 {
		mongoQuery = append(mongoQuery, makeMatchAndGopherQuery(filters))
	}

	cursor, err := r.gophers.Aggregate(ctx, mongo.Pipeline(mongoQuery))
	if err != nil {
		return nil, err
	}

	var gophers []*gopher.Aggregate
	for cursor.Next(ctx) {
		var res gopherDTO
		if err := cursor.Decode(&res); err != nil {
			return nil, err
		}

		g, err := dtoToGopher(res)
		if err != nil {
			return nil, err
		}

		gophers = append(gophers, g)
	}

	return gophers, nil
}

func (r *MongoDBRepository) RunMigrations(ctx context.Context) (err error) {
	migrator := mongoutils.NewMongoDBMigrator(
		r.gophers.Database().Client(),
		mongoCollectionName,
	)

	err = migrator.ApplyMigrations(ctx,
		mongomigrations.NewCreateGophersCollection(r.gophers.Database()),
	)
	if err != nil {
		return err
	}

	return nil
}

func makeCountGopherQuery() bson.D {
	return bson.D{{"$count", "count"}}
}

func makeMatchOrGopherQuery(filters []bson.D) bson.D {
	return bson.D{{"$match", makeOrFilter(filters)}}
}

func makeMatchAndGopherQuery(filters []bson.D) bson.D {
	return bson.D{{"$match", makeAndFilter(filters)}}
}

func makeOrFilter(filters []bson.D) bson.D {
	return bson.D{{"$or", filters}}
}

func makeAndFilter(filters []bson.D) bson.D {
	return bson.D{{"$and", filters}}
}

func makeGopherQueryFilters(filters gopher.Filters) []bson.D {
	var queryFilters []bson.D
	if filters.ID != "" {
		queryFilters = append(queryFilters, bson.D{{"_id", filters.ID}})
	}

	if filters.Username != "" {
		queryFilters = append(queryFilters, bson.D{{"username", filters.Username}})
	}
	return queryFilters
}

func makeNamesFilter(names []string) bson.D {
	var nameFilters []bson.D
	for _, name := range names {
		nameFilters = append(nameFilters, bson.D{{"name", name}})
	}

	return bson.D{{"$or", nameFilters}}
}

func makeStatusesFilter(statuses []gopher.Status) bson.D {
	var statusFilters []bson.D
	for _, status := range statuses {
		statusFilters = append(statusFilters, bson.D{{"status", status.String()}})
	}

	return bson.D{{"$or", statusFilters}}
}
