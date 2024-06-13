CREATE TABLE IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    runtime INTEGER NOT NULL,
    genres TEXT NOT NULL,               -- JSON field
    version INTEGER NOT NULL DEFAULT 1,

    CHECK (runtime >= 0),
    CHECK (year BETWEEN 1888 AND strftime('%Y', current_timestamp)),
    CHECK (json_array_length(genres) BETWEEN 1 and 5)
);
