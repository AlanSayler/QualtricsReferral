

package main

import (
//	"github.com/sielickin/Qualgo"
//	"database/sql"
//	"fmt"
//	"log"
	_ "github.com/lib/pq"
)

func init () {

}

 const (
	schema_name = "research"
	table_name = "tree"
	createSchema_base = "CREATE SCHEMA ?;"
	createTableBase = `
CREATE TABLE ?.? (
RESPONSEID    TEXT,
EMAIL    TEXT,
SURVEY   BOOLEAN,
STATE    CHAR(1),
PARENTID TEXT,
TIME_SELECTED TIMESTAMP,
TIME_SENT     TIMESTAMP,
CHILDRENIDS   TEXT[5],
COUPON_CODE   TEXT,
SIBLINGIDS    TEXT[5]
);
`
	selectIdsOfSBase = "SELECT RESPONSEID FROM ?.? WHERE STATE = 'S';"

	killOldSBase = "UPDATE ?.? SET STATE = 'D' WHERE STATE ='S' AND TIME_SENT < now() - interval '7 days';" 


	selectCouponCodesOfDeadBase = "SELECT COUPON_CODEFROM ?.? WHERE STATE = 'DEAD' AND COUPON_CODE != 'NONE';"


	deleteCouponsFromDBase = "UPDATE ?.? SET COUPON_CODE = 'NONE' WHERE STATE = 'DEAD';"


	getOldQBase = "SELECT RESPONSEID FROM ?.? WHERE STATE = 'Q' AND TIME_SELECTED < now() - interval '1 day';"


	updateOldQBase = "UPDATE ?.? SET STATE = 'S' WHERE STATE = 'Q' and TIME_SELECTED < now() - interval '1 day';"


        insertEmailBase = "INSERT INTO ?.? (EMAIL, STATE)  VALUES ('?', 'Q');"
 )

func main () {

}
