-- +goose Up
-- +goose StatementBegin
ALTER TABLE dinosaur
  ADD created_at int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE dinosaur
  DROP created_at;
-- +goose StatementEnd
