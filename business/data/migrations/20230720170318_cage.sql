-- +goose Up
-- +goose StatementBegin
CREATE TABLE cage (
  id uuid not null,
  type text,
  capacity int,
  current_capacity int,
  status text,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cage;
-- +goose StatementEnd
