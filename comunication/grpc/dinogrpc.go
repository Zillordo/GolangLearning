package dinogrpc

import (
	"../../database"
	model "../../models"
	"context"
	"fmt"
)

type GrpcServer struct {
	dbHandler database.DinoDBHandler
}

func NewDinoGrpcServer(dbType database.DbType, connection string) (*GrpcServer, error) {
	handler, err := database.GetDatabaseHandler(dbType, connection)
	if err != nil {
		return nil, fmt.Errorf("could not create a database handler object, err: %v", err)
	}
	return &GrpcServer{
		dbHandler: handler,
	}, nil
}

func (server *GrpcServer) GetAnimal(_ context.Context, r *Request) (*Animal, error) {
	animal, err := server.dbHandler.GetDynoByNickname(r.GetNickname())
	return convertToDinoGRPCAnimal(animal), err
}

func (server *GrpcServer) GetAllAnimals(_ *Request, stream DinoService_GetAllAnimalsServer) error {
	animals, err := server.dbHandler.GetDynos()
	if err != nil {
		return err
	}
	for _, animal := range animals {
		a := convertToDinoGRPCAnimal(animal)
		if err := stream.Send(a); err != nil {
			return err
		}
	}
	return nil
}

func convertToDinoGRPCAnimal(animal model.Animal) *Animal {
	return &Animal{
		Id:         int32(animal.ID),
		AnimalType: animal.AnimalType,
		Nickname:   animal.Nickname,
		Zone:       int32(animal.Zone),
		Age:        int32(animal.Age),
	}
}
