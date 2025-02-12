package gserver

import (
	"context"
	"math/rand"
	protoapi "protoapi/proto"
	"time"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *protoapi.RequestDateTime) (*protoapi.DateTime, error) {

	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}
	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {

	_min := int(1)
	_max := int(9e8)
	place := r.GetPlace()
	aux := _min + rand.Intn(_max-_min+1) + _min

	for {
		place--
		if place == 0 {
			break
		}
		aux = _min + rand.Intn(_max-_min+1) + _min
	}
	response := &protoapi.RandomInt{Value: int64(aux)}
	return response, nil

}

func (RandomServer) GetRandomPass(ctx context.Context, r *protoapi.RequestPass) (*protoapi.RandomPass, error) {

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, r.GetLength())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	response := &protoapi.RandomPass{Password: string(b)}

	return response, nil
}
