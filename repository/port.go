package repository

import (
	"context"
	"database/sql"
	"github.com/jaimera/poc-services/domain/entity"
)

type portRepository struct {
	conn      *sql.DB
	mainTable string
}

func newPortRepository(conn *sql.DB) portRepository {
	return portRepository{
		conn:      conn,
		mainTable: "tb_port",
	}
}

const query = `
		INSERT INTO tb_port
			( slug
			, name
			, city
			, country
			, alias
			, regions
			, coordinates
			, province
			, timezone
			, unlocs
			, code
			)
			VALUES
				( ?
				, ?
				, ?
				, ?
				, ?
				, ?
				, POINT(?, ?)
				, ?
				, ?
				, ?
				, ?
				) as new ON DUPLICATE KEY UPDATE 
				    name = new.name,
					city = new.city,
					country = new.country,
					alias = new.alias,
					regions = new.regions,
					coordinates = new.coordinates,
					province = new.province,
					timezone = new.timezone,
					unlocs = new.unlocs,
					code = new.code
		;
	`

const fetch = `
	select id
	        , slug
			, name
			, city
			, country
			, alias
			, regions
			, province
			, timezone
			, unlocs
			, code from db_poc.tb_port where slug = ?;
`

func (r portRepository) Upsert(context context.Context, port entity.Port) (uint32, error) {

	tx, err := r.conn.Begin()

	ret, err := tx.Exec(query,
		port.Slug,
		port.Name,
		port.City,
		port.Country,
		port.Alias,
		port.Regions,
		port.Coordinates.Latitude,
		port.Coordinates.Longitude,
		port.Province,
		port.Timezone,
		port.Unlocs,
		port.Code,
	)

	if err != nil {
		return 0, err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r portRepository) FetchByCode(context context.Context, code string) (entity.Port, error) {

	row := r.conn.QueryRow(fetch, code)

	port, err := r.parseEntity(row)
	if err != nil {
		return entity.Port{}, err
	}

	return port, nil
}

func (r portRepository) parseEntity(s *sql.Row) (entity.Port, error) {

	var port entity.Port
	err := s.Scan(
		&port.ID,
		&port.Slug,
		&port.Name,
		&port.City,
		&port.Country,
		&port.Alias,
		&port.Regions,
		&port.Province,
		&port.Timezone,
		&port.Unlocs,
		&port.Code,
	)
	if err != nil {
		return entity.Port{}, err
	}

	return port, nil
}
