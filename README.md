---
title: CassandraDB Query Activity
---

# CassandraDB Query Activity
This activity allows you to Query a record from particular table from the CassandraDB server

## Installation
### Flogo CLI
```bash
flogo install github.com/dhire05/cassandraqueryrecord
```

## Schema
Inputs and Outputs:

```json
{   
  "inputs":[
    {
      "name": "ClusterIP",
      "type": "string",
	  "required": true      
    },
	{
      "name": "Keyspace",
      "type": "string",
      "required": true
    },
	{
      "name": "TableName",
      "type": "string",
      "required": true
    },
	{
      "name": "Select",
      "type": "string",
      "required": true
    },
	{
      "name": "Where",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  ]
 }
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| ClusterIP      | True     | The CassandraDB cluster instance |         
| Keyspace       | True     | The name of the Keyspace
| TableName      | True     | The name of table to delete record
| Select         | True		| The Select Element to select all or single column form table
| Where          | True     | The where clause or condition |

## Example
The below example is to delete record from CassandraDB

```json
{
  "id": "CassandraDB_1",
  "name": "CassandraDB connector",
  "description": "Delete record from CassandraDB",
  "activity": {
    "ref": "github.com/dhire05/cassandraqueryrecord",
    "input": {
      "ClusterIP": "127.0.0.1",
      "Keyspace": "sample",
      "TableName": "employee",
	  "Select": "*",
      "Where": "empid = 104"      
    }
  }
}
```