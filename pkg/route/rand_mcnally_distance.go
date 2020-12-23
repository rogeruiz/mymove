package route

import (
	"fmt"

	"github.com/gobuffalo/pop/v5"

	"github.com/transcom/mymove/pkg/models"
)

func randMcNallyZip3Distance(db *pop.Connection, fromZip3 string, toZip3 string) (int, error) {
	var distance models.Zip3Distance
	if fromZip3 == toZip3 {
		return 0, fmt.Errorf("fromZip3 (%s) cannot be the same as toZip3 (%s)", fromZip3, toZip3)
	} else if fromZip3 > toZip3 {
		db.Where("from_zip3 = ? and to_zip3 = ?", toZip3, fromZip3).First(&distance)
	} else {
		db.Where("from_zip3 = ? and to_zip3 = ?", fromZip3, toZip3).First(&distance)
	}
	return distance.DistanceMiles, nil
}