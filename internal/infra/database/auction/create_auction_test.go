package auction

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/internal/entity/auction_entity"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_DB = "mongodb://admin:admin@127.0.0.1:27017/auctions?authSource=admin"

func TestCreateAuctionAndUpdateStatus(t *testing.T) {
	// Configurar o banco de dados MongoDB em memória ou de teste
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB))
	assert.NoError(t, err)
	defer client.Disconnect(context.TODO())

	database := client.Database("testdb")
	collection := database.Collection("auctions")
	repository := NewAuctionRepository(database)

	// Configurar o intervalo de leilão para teste
	os.Setenv("AUCTION_INTERVAL", "2s") // Intervalo curto para teste

	auction, err := auction_entity.CreateAuction(
		"iPhone 15 Pro",
		"Electronics",
		"Brand new iPhone 15 Pro 256GB",
		auction_entity.New,
	)
	assert.Nil(t, err)

	// verifica se é possivel conectar ao banco
	err = client.Ping(context.TODO(), nil)
	assert.NoError(t, err)

	err = repository.CreateAuction(context.TODO(), auction)
	assert.Nil(t, err)

	// Verificar o status inicial no banco de dados
	var result AuctionEntityMongo
	err = collection.FindOne(context.TODO(), bson.M{"_id": auction.Id}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, auction_entity.Active, result.Status)

	// Aguardar o intervalo para que o status seja atualizado
	time.Sleep(3 * time.Second) // Aguarde um pouco mais que o intervalo configurado

	// Verificar o status atualizado no banco de dados
	err = collection.FindOne(context.TODO(), bson.M{"_id": auction.Id}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, auction_entity.Completed, result.Status)
}

func TestCreate10AuctionsAndUpdateStatus(t *testing.T) {
	// Setup MongoDB test connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_DB))
	assert.NoError(t, err)
	defer client.Disconnect(context.TODO())

	database := client.Database("testdb")
	collection := database.Collection("auctions")
	repository := NewAuctionRepository(database)

	// Clear the collection first
	_, err = collection.DeleteMany(context.TODO(), bson.M{})
	assert.NoError(t, err)

	// Configure auction interval for test
	os.Setenv("AUCTION_INTERVAL", "2s") // Short interval for testing

	// Create 10 auctions
	auctions := make([]*auction_entity.Auction, 10)
	for i := 0; i < 10; i++ {
		auction, err := auction_entity.CreateAuction(
			fmt.Sprintf("Test Item %d", i+1),
			"Test Category",
			fmt.Sprintf("Description for test item %d", i+1),
			auction_entity.New,
		)
		assert.Nil(t, err)
		err = repository.CreateAuction(context.TODO(), auction)
		assert.Nil(t, err)
		auctions[i] = auction
	}

	// Verify initial status in the database is Active for all auctions
	for _, auction := range auctions {
		var result AuctionEntityMongo
		err = collection.FindOne(context.TODO(), bson.M{"_id": auction.Id}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, auction_entity.Active, result.Status)
	}

	// Wait for a bit longer than the auction interval to allow status update
	time.Sleep(3 * time.Second)

	// Verify that all auctions have been updated to Completed
	for _, auction := range auctions {
		var result AuctionEntityMongo
		err = collection.FindOne(context.TODO(), bson.M{"_id": auction.Id}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, auction_entity.Completed, result.Status)
	}
}
