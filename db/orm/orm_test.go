package orm

import (
	"testing"
	//	"time"
)

/*entity*/
type Chair struct {
	ID int `table:"device_chair" column:"id"`
	//	Time time.Time `column:"time"`
	Day string `column:"day"`
	//	Origin          string    `column:"origin"`
	//	ParentSerial    string    `column:"parent_serial"`
	//	Serial          string
	//	RunningNumber   string  `column:"running_number"`
	//	PointA          float32 `column:"point_a"`
	//	PointB          float32 `column:"point_b"`
	//	PointC          float32 `column:"point_c"`
	//	PointD          float32 `column:"point_d"`
	//	PointE          float32 `column:"point_e"`
	//	PointF          float32 `column:"point_f"`
	//	Voltage         float32 `column:"voltage"`
	//	SeatingPosition uint8   `column:"seating_position"`
}

func TestDB(t *testing.T) {

	chair := new(Chair)
	chair.ID = 12377
	//	chair.SeatingPosition = 10
	//	chair.Day = chair.Time.Format("2006-01-02")
	chair.Day = "asdfasfd"

	Register(&Chair{})

	DropTable(&Chair{})
	CreateTable(&Chair{})
	Save(chair)
}
