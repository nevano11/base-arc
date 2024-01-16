package repository

import (
	"awesomeProject/internal/core-module/config"
	"awesomeProject/internal/entity"
	"context"
	"errors"
)

type TtBandRepository struct {
	//connection *tarantool.Connection
}

func NewTtBandRepository(ctx context.Context, config config.TarantoolConfig) (*TtBandRepository, error) {
	//connection, err := tarantool.Connect(ctx, tarantool.NetDialer{
	//	Address:              config.Address,
	//	User:                 config.User,
	//	Password:             config.Password,
	//	RequiredProtocolInfo: tarantool.ProtocolInfo{},
	//}, tarantool.Opts{})
	//if err != nil {
	//	return nil, err
	//}
	return &TtBandRepository{
		//connection: connection,
	}, nil
}

func (r *TtBandRepository) GetBandById(id int) (entity.Band, error) {
	/*request := tarantool.NewSelectRequest("bands").
		Index("primary").
		Offset(0).
		Limit(1).
		Iterator(tarantool.IterEq).
		Key([]interface{}{id})
	response, err := r.connection.Do(request).Get()
	if err != nil {
		return entity.Band{}, err
	}

	if len(response.Data) > 0 {
		record := response.Data[0].([]interface{})
		return entity.Band{
			Id:       record[0].(int8),
			BandName: record[1].(string),
			Year:     record[2].(uint16),
		}, nil
	}
	for i, v := range response.Data {
		fmt.Println(i, v)
	}*/
	return entity.Band{}, errors.New("Not found")
}
