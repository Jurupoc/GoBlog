package shared

var DB_DEAFULT_PATH string = "./app_db"
var DB_CREATE_TABLE_POSTS string = `CREATE TABLE IF NOT EXISTS posts (
							id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
							author STRING NOT NULL,
							content STRING,
							createdAt DATETIME
						);`
