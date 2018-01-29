package persistence

import (
	"log"
	"thmlogwork.com/go-time-zone-app/domain"
	"database/sql"
	"github.com/kellydunn/golang-geo"
	"fmt"
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
	db := mapper.SqlDbConn();
	_, connErr := db.Begin()

	if connErr != nil {
		log.Fatal("Connection to db failed")
	}
	log.Println("Connect to db successfully")

	return timezoneDao{db: db}
}

func (dao timezoneDao) Get(l domain.LatLon) *domain.TimezoneInfoResponse {
	log.Printf("Getting (Latitude, Longitude): (%v, %v)\n", l.Latitude, l.Longitude)
	var utc_format string
	var tz_name1st string
	var places string
	var dst_places string
	query := fmt.Sprintf(
		"SELECT name, utc_format, places, dst_places from timezones where ST_Contains(geom, ST_GeomFromText('POINT(%v %v)', 4326))",
		l.Longitude, l.Latitude)
	err := dao.db.QueryRow(query).Scan(&utc_format, &tz_name1st, &places, &dst_places)

	if err != nil {
		log.Println(err.Error())
		log.Printf("Error querying for %v", l)
	}

	return domain.NewTimezoneInfoResponse(places, dst_places, utc_format, tz_name1st)
}
