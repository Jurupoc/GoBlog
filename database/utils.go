package database

const DBDefaultPath string = "./app_db"
const DBCreatePostsTable string = `CREATE TABLE IF NOT EXISTS posts (
							id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
							author STRING NOT NULL,
							content STRING,
							createdAt DATETIME
						);`
