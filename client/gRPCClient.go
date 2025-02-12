package client

import (
	"context"
	protoapi "protoapi/proto"
)

func GetDateTime(ctx context.Context, m protoapi.RandomClient) (*protoapi.DateTime, error) {

	request := &protoapi.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}

func GetPassword(ctx context.Context, m protoapi.RandomClient, seed int64, length int64) (*protoapi.RandomPass, error) {
	request := &protoapi.RequestPass{
		Seed:   seed,
		Length: length,
	}
	return m.GetRandomPass(ctx, request)
}

func GetRandom(ctx context.Context, m protoapi.RandomClient, seed int64, place int64) (*protoapi.RandomInt, error) {
	request := &protoapi.RandomParams{
		Seed:  seed,
		Place: place,
	}

	return m.GetRandom(ctx, request)
}
