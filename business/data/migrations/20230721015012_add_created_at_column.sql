-- +goose Up
-- +goose StatementBegin
ALTER TABLE cage
  ADD created_at int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cage
  DROP created_at;
-- +goose StatementEnd
