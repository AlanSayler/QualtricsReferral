package main

import (
	"sielickin/Qualgo"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

// TODO:
//	PUT ALL SQL CALLS HERE
// const (
//	CREATE_SCHEMA_BASE = "CREATE SCHEMA %v;"
//	CREATE_TABLE_BASE = CREATE TABLE %v.%v (
//	RESPONSEID    TEXT,
//	EMAIL    TEXT,
//	SURVEY   BOOLEAN,
//	STATE    CHAR(1),
//	PARENTID TEXT,
//	TIME_SELECTED TIMESTAMP,
//	TIME_SENT     TIMESTAMP,
//	CHILDRENIDS   TEXT[5],
//	COUPON_CODE   TEXT,
//	SIBLINGIDS    TEXT[5]
//	);
//
// )

func  CREATE_SCHEMA(){
	return fmt.Sprintf(CREATE_SCHEMA_BASE, schema_name );
}  

func CREATE_TABLE (){
	return fmt.Sprintf(CREATE_TABLE_BASE, schema_name, table_name);
}

// States:
// S: Sent
// Q: Queued 
// N: Not Queued
// D: Dead
// C: Completed

// SELECT UserID FROM nodes WHERE STATE=='s'

const (
	FatalLog = "FATAL   : "
	CritLog  = "CRITICAL: "
	WarnLog  = "WARNING : "
	OKLog	 = "NORMAL  : "
)

var logLevel := map[string]int {
	FatalLog = 3,
	CritLog = 2,
	Warnlog = 1,
	OKLog = 0,
}

func log(level int, str string){
		makeConn()

}

func main() {
	// open connection to database
	// https://csl.cs.wisc.edu/services/databases/postgresql-database-service
	conn, err := sql.Open("postgres", "postgres://username:password@postgres.cs.wisc.edu/databasename?sslmode=verify-full")
	defer conn.Close()
	if err != nil {
		log.Fatal(FatalLog + "Could not open connection to the database: " + err)
	}
	// TODO: Register qualtrics info

	// Read from qualtrics, update databases state.
	// -->


	//credits tracks maximum number of survey that can potentially be queued this instance
	credits := 0
	//1.
	//deal with sent surveys that have completed, get their children
	//SQl call: get array of responseIDs of all nodes with state S
	for i, val := range (ArrayOfResponseIDs) {
		//API call: getlegacyresponsedata on ith responseID
		err, rspData := Qualgo.Query(Qualgo.GetLegacyResponseData)
		if(/*getlecagyresponsedata returned successfully*/){
			//set state to C

			//API call: sendThankYouSurvey with this node's coupon code

			//parse each response for email address, concatenate @wisc.edu

			//SQL call: add new rows, set their parent to be the row which
			//holds the ith responseID, set their states to N, their email to
			//the email contained in the above response, 

			//API call to add each of the new nodes to the panel

			//SQL call: take the resultant respondentID, and put that in the
			//respondentID field of the child

			//SQL calls make the children the children of the parent

			//SQL check if any new children have same email address as old
			//children, if so, set new children state to D  

			credit += 3
		}
	}

	//2.
	//Deal with sent surveys that have expired

	//SQL statement: get all nodes whose state is S and whose send date was
	//over 1 week ago, set their state to D, also get their coupons, put them
	//in an array
	credit += length(arrayCoupons)
	for( i:= 0; i< length(arrayCoupons); i++ {
		coupons.push(/*ith coupon in array*/)
		/*set owner of ith coupon to D, set their coupon field to blank*/
	}

	//3.
	//Deal with queued surveys that are ready to send

	//SQL statement: get all nodes whose state is Q and whose queued date is
	//over 1 day ago, set their state to S, keep an array of their response IDs

	for i:= 0; i<length(/*arrayofresponseIDs*/); i++ {
		//API call send survey to each individual in array
	}

	//4.
	//queue up new surveys
	for credit > 0 && !coupons.isEmpty() && /*SQL statement getting all nodes whose state is N does not return empty*/  {
		maxDist float32 = 0
		currDist float32 = 0
		//UUID of longest distance N node
		whichMax string = nil
		//SQL statement: get all nodes of state N
		for (i := 0; i<length(/*nodesArray*/); i++{
			currDist = 0
			maxDist = 0
			//SQL statement: get all nodes of states Q,S or C
			for(j := 0; j<length(/*arrayofQSC*/); j++{
				currDist += dist(/*ithNodesArray, jthArraryofQSC*/)
			}
			if(currDist>maxDist){
				maxDist = currDist
				whichMax = /*ithUUID in array*/
			}
		}
		/*SQL statement: set chosen node's state to Q, record time in queued time */
		/*SQL statement: get chosen node's coupon field*/ = coupons.pop()
		credit--
	}
}

func dist(UUID1 int,UUID2 int){
	//Make SQL statments, getting an array of parents for each node. Compare the two arrays for the lowest common parent. Return sum from bottoms.
}
