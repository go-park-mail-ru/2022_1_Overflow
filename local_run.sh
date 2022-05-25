(go run ./cmd/repository/run_repository.go &) && \
(go run ./cmd/profile/run_profile.go &) && \
(go run ./cmd/mailbox/run_mailbox.go &) && \
(go run ./cmd/folder_manager/run_folder_manager.go &) && \
(go run ./cmd/auth/run_auth.go > auth.log &) && \
(go run ./cmd/attach/run_attach.go > attach.log &) && \
go run ./cmd/app/main.go