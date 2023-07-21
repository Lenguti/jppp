-- +goose Up
-- +goose StatementBegin
ALTER TABLE cage
  ADD updated_at int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE cage
  DROP updated_at;
-- +goose StatementEnd
