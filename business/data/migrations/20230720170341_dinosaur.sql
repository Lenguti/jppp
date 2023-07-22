-- +goose Up
-- +goose StatementBegin
CREATE TABLE dinosaur (
  id uuid NOT NULL,
  cage_id uuid NULL,
  name text,
  species text,
  diet text,
  PRIMARY KEY (id),
  FOREIGN KEY(cage_id) REFERENCES cage(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dinosaur;
-- +goose StatementEnd
