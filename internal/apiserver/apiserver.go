package apiserver

import (
	"context"
	"encoding/json"
	"log"
	"wbl0/model"

	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
)

type Api struct {
	DB     *pgx.Conn
	cache  map[int]model.Order
	config *Config
}

func New(config *Config) *Api {
	return &Api{
		config: config,
	}
}

func (a *Api) CacheRecovery() error {
	query := `SELECT id, data FROM orders`
	rows, err := a.DB.Query(context.Background(), query)
	if err != nil {
		return err
	}

	var order model.Order
	var id int
	a.cache = make(map[int]model.Order)

	for rows.Next() {
		err = rows.Scan(&id, &order)
		if err != nil {
			return err
		}
		a.Set(id, order)
	}

	return nil
}

func (a *Api) Set(key int, val model.Order) {
	a.cache[key] = val
}

func (a *Api) Get(key int) (model.Order, bool) {

	value, ok := a.cache[key]
	if ok {
		return value, ok
	}
	return model.Order{}, ok
}

func (a *Api) ParseAndSaveMsg(msg *stan.Msg) {
	var order model.Order

	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		log.Println(err)
		return
	}

	err = order.Validate()
	if err != nil {
		log.Println(err)
		return
	}

	var id int
	query := `INSERT INTO orders (data) VALUES ($1) RETURNING id`

	err = a.DB.QueryRow(context.Background(), query, order).Scan(&id)
	if err != nil {
		log.Println(err)
		return
	}
	a.Set(id, order)
}
