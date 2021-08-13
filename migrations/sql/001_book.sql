-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book (   
  "id"                serial,   
  "title"             VARCHAR(160),
  "author"            VARCHAR(160),    
  "year"              VARCHAR(160),
  PRIMARY KEY ("id")  
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS book;
-- +goose StatementEnd

