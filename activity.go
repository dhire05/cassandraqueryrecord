package cassandraqueryrecord

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/gocql/gocql"	
)

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-CassandraQueryRecord")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	clusterIP := context.GetInput("ClusterIP").(string)
	keySpace := context.GetInput("Keyspace").(string)
	tableName := context.GetInput("TableName").(string)
	selectElements := context.GetInput("Select").(string)
	whereClause := context.GetInput("Where").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says [%s] to [%s] with table [%s]", clusterIP, keySpace, tableName)
	log.Debugf("Flogo is about to select [%s] from table [%s].[%s] where [%s] on cluster [%s]", selectElements, keySpace, tableName, whereClause, clusterIP)
	
	fmt.Println("The Flogo engine says "+clusterIP+" to "+keySpace+" with table "+tableName)
	fmt.Println("Flogo is about to select "+selectElements+" from table "+keySpace+"."+tableName+" where "+whereClause+" on cluster "+clusterIP)

	// Provide the cassandra cluster instance here.
	cluster := gocql.NewCluster(clusterIP)

	// gocql requires the keyspace to be provided before the session is created.
	// In future there might be provisions to do this later.
	cluster.Keyspace = keySpace

	
	session, err := cluster.CreateSession()

	if err != nil {
		log.Debugf("Could not connect to cassandra cluster: ", err)
		fmt.Println("Could not connect to cassandra cluster: ", err)
	}
	log.Debugf("Session Created Sucessfully")
	log.Debugf("Session : ", session)
	log.Debugf("Cluster : ", clusterIP)
	log.Debugf("Keyspace : ", keySpace)
	log.Debugf("Session Timeout : ", cluster.Timeout)

	log.Debugf("Next Step is Query Execution")
	
	//fmt.Println("Session Created Sucessfully")
	//fmt.Println("Session : ", session)
	//fmt.Println("Cluster : ", clusterIP)
	//fmt.Println("Keyspace : ", keySpace)
	//fmt.Println("Session Timeout : ", cluster.Timeout)

	fmt.Println("Next Step is Query Execution")

	queryString := "SELECT " + selectElements + " FROM " + tableName
	if whereClause != "" {
		queryString += " where " + whereClause
	}
	log.Debugf("Query string: [%s]", queryString)
	fmt.Println("Query string: ", queryString)

	iter := session.Query(queryString).Iter()
	log.Debugf("number of columns: %v", len(iter.Columns()))
	var result []map[string]interface{}

	for i := 0; i < iter.NumRows(); i++ {
		row := make(map[string]interface{})
		if !iter.MapScan(row) {
			log.Debug("Error Select")
			iter.Close()
		}
		result = append(result, row)
		for _, column := range iter.Columns() {
			log.Debugf("Record [%v], Field [%v], value [%v]", i, column.Name, row[column.Name])
		}
	}

	context.SetOutput("result", result)

	fmt.Println(result)
	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
