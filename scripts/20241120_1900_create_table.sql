CREATE TABLE IF NOT EXISTS global_config ( 
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key TEXT UNIQUE NOT NULL,
			value TEXT NOT NULL,
			created_at TEXT NOT NULL,
			scope TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    gender TEXT,
    birthday TEXT,
    password TEXT NOT NULL,
    created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS articles ( 
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    image_urls TEXT NOT NULL CHECK(json_valid(image_urls)), 
    title TEXT NOT NULL,
    uploaded_time TEXT NOT NULL, 
    tags TEXT, 
    content TEXT, 
    created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username INTEGER NOT NULL,
    event_type TEXT NOT NULL,
    content TEXT NOT NULL CHECK(json_valid(content)), 
    created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_diaries (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TEXT NOT NULL
);
