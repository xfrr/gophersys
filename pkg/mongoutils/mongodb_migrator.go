package mongoutils

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMigrationsCollectionName = "gophers_migrations"
)

var (
	// ErrMigrationsCollectionNotFound error returned when the migrations collection is not found
	ErrMigrationsCollectionNotFound = fmt.Errorf("mongodb migrations collection not found")
)

// migrationRecord struct to store migration information in the database
type migrationRecord struct {
	Name      string    `bson:"name"`
	AppliedAt time.Time `bson:"applied_at"`
}

func newMigrationRecord(name string) migrationRecord {
	return migrationRecord{
		Name:      name,
		AppliedAt: time.Now(),
	}
}

// MongoMigrator struct that implements the Migrator interface
type MongoMigrator struct {
	database       *mongo.Database
	migrationsColl *mongo.Collection
}

// NewMongoDBMigrator creates a new MongoMigrator
func NewMongoDBMigrator(client *mongo.Client, dbName string) *MongoMigrator {
	database := client.Database(dbName)
	migrationsCollection := database.Collection(defaultMigrationsCollectionName)
	return &MongoMigrator{
		database:       database,
		migrationsColl: migrationsCollection,
	}
}

// ApplyMigrations applies all migrations that haven't been applied yet
func (m *MongoMigrator) ApplyMigrations(ctx context.Context, migrations ...Migration) error {
	err := m.ensureCollection(ctx)
	if err != nil {
		return err
	}

	err = m.applyMigrations(ctx, migrations...)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoMigrator) ensureCollection(ctx context.Context) error {
	collNames, err := m.database.ListCollectionNames(ctx, map[string]string{})
	if err != nil {
		return err
	}

	var exists bool
	for _, name := range collNames {
		if name == defaultMigrationsCollectionName {
			exists = true
			break
		}
	}

	if !exists {
		return m.createMigrationsCollection(ctx)
	}

	return nil
}

func (m *MongoMigrator) createMigrationsCollection(ctx context.Context) error {
	err := m.database.CreateCollection(ctx, defaultMigrationsCollectionName)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoMigrator) getMigrationRecords(ctx context.Context) ([]migrationRecord, error) {
	cursor, err := m.migrationsColl.Find(ctx, bson.D{}, options.Find().SetSort(bson.D{{"applied_at", 1}}))
	if err != nil {
		return nil, err
	}

	var records []migrationRecord
	if err = cursor.All(ctx, &records); err != nil {
		return nil, err
	}

	return records, nil
}

func (m *MongoMigrator) applyMigrations(ctx context.Context, migrations ...Migration) error {
	for _, migration := range migrations {
		isApplied, err := m.isMigrationApplied(ctx, migration)
		if err != nil {
			return err
		}

		if isApplied {
			continue
		}

		if err := m.applyMigration(ctx, migration); err != nil {
			return err
		}
	}

	return nil
}

func (m *MongoMigrator) applyMigration(ctx context.Context, migration Migration) error {
	if err := migration.Up(ctx); err != nil {
		return err
	}

	_, err := m.migrationsColl.InsertOne(ctx, newMigrationRecord(migration.Name()))
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoMigrator) isMigrationApplied(ctx context.Context, migration Migration) (isApplied bool, err error) {
	record, err := m.getMigrationRecord(ctx, migration.Name())
	if err != nil {
		return false, err
	}

	isApplied = record != nil && !record.AppliedAt.IsZero()
	return isApplied, nil
}

func (m *MongoMigrator) getMigrationRecord(ctx context.Context, name string) (*migrationRecord, error) {
	filter := bson.D{{"name", name}}
	var record migrationRecord
	err := m.migrationsColl.FindOne(ctx, filter).Decode(&record)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &record, nil
}
