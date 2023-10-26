CREATE TABLE docs (
    id SERIAL NOT NULL UNIQUE PRIMARY KEY,
    title VARCHAR(64),
    category VARCHAR(64),
    description TEXT,
    doc_file VARCHAR(64),
    img_name VARCHAR(64) DEFAULT 'default_image.png',
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

INSERT INTO docs (title, category, doc_file) VALUES 
    ('My first story', 'books', 'Book one'),
    ('My second story', 'books', 'Book two');