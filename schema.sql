DROP TABLE IF EXISTS ;
CREATE TABLE IF NOT EXISTS websites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at INTEGER NOT NULL
    site_title TEXT NOT NULL,
    domain_name TEXT NOT NULL,
);
CREATE INDEX idx_articles_on_created_at ON articles (created_at DESC);
INSERT INTO articles (title, body, created_at) VALUES (
    'title of example post',
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
    unixepoch()
);