-- +goose Up
-- +goose StatementBegin
ALTER TABLE dinosaur
  ADD updated_at int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE dinosaur
  DROP updated_at;
-- +goose StatementEnd
