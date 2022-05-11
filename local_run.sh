(go run ./cmd/repository/run_repository.go > repository.log &) && \
(go run ./cmd/profile/run_profile.go > profile.log &) && \
(go run ./cmd/mailbox/run_mailbox.go > mailbox.log &) && \
(go run ./cmd/folder_manager/run_folder_manager.go > folder_manager.log &) && \
(go run ./cmd/auth/run_auth.go > auth.log &) && \
(go run ./cmd/smtp/run_smtp.go > smtp.log &) && \
go run ./cmd/app/main.go