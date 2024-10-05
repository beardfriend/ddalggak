ent-init:
	sh ./script/ent-init.sh

ent-generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent ./internal/schema/ --feature sql/modifier --feature sql/execquery

