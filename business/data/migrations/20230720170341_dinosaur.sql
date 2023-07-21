-- +goose Up
-- +goose StatementBegin
CREATE TABLE dinosaur (
  id uuid not null,
  cage_id uuid,
  name text,
  species text,
  diet text,
  PRIMARY KEY (id),
  fk_cage uuid REFERENCES cage(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dinosaur;
-- +goose StatementEnd
