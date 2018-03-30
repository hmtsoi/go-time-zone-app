package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kellydunn/golang-geo"
	"thmlogwork.com/go-time-zone-app/domain"
)

type timezoneDao struct {
	db *sql.DB
}

func NewTimezoneDao() timezoneDao {
	mapper, err := geo.HandleWithSQL()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connecting to db")
	db := mapper.SqlDbConn()
	_, connErr := db.Begin()

	if connErr != nil {
		log.Fatal("Connection to db failed")
	}
	log.Println("Connect to db successfully")

	return timezoneDao{db: db}
}

func (dao timezoneDao) Get(l domain.LatLon) *domain.TimezoneInfo {

	log.Printf("Getting (Latitude, Longitude): (%v, %v)\n", l.Latitude, l.Longitude)
	entity := &domain.TimezoneInfo{}

	query := fmt.Sprintf(
		"SELECT utc_format, tz_name1st, places, dst_places from timezones where ST_Contains(geom, ST_GeomFromText('POINT(%v %v)', 4326))",
		l.Longitude, l.Latitude)
	err := dao.db.QueryRow(query).Scan(&entity.UtcFormat, &entity.TimezoneParameter, &entity.Places, &entity.DstPlaces)

	if err != nil {
		log.Println(err.Error())
		log.Printf("Error querying for %v", l)
	}
	return entity
}
