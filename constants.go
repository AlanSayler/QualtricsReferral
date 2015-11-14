

package main

import (
//	"sielickin/Qualgo"
//	"database/sql"
//	"fmt"
//	"log"
	_ "github.com/lib/pq"
)

func init () {

}

// TODO:
//	PUT ALL SQL CALLS HERE
 const (
	//edit these out
	schema_name = "research"
	table_name = "tree"
	//keep these in
	CREATE_SCHEMA_BASE = "CREATE SCHEMA ?;"
	CREATE_TABLE_BASE = `
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
	SELECT_IDS_OF_S_BASE = "SELECT RESPONSEID FROM ?.? WHERE STATE = 'S';"
	KILL_OLD_S_BASE = "UPDATE ?.? SET STATE = 'D' WHERE STATE ='S' AND TIME_SENT < now() - interval '7 days';" 
	SELECT_COUPON_CODES_OF_DEAD_BASE = "SELECT COUPON_CODEFROM ?.? WHERE STATE = 'DEAD' AND COUPON_CODE != 'NONE';"
	DELETE_COUPONS_FROM_D_BASE = "UPDATE ?.? SET COUPON_CODE = 'NONE' WHERE STATE = 'DEAD';"
	GET_OLD_Q_BASE = "SELECT RESPONSEID FROM ?.? WHERE STATE = 'Q' AND TIME_SELECTED < now() - interval '1 day';"
	UPDATE_OLD_Q_BASE = "UPDATE ?.? SET STATE = 'S' WHERE STATE = 'Q' and TIME_SELECTED < now() - interval '1 day';"
 )

func main () {

}
