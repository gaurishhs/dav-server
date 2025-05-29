css-dev:
	@exec_command=$$(command -v bunx || command -v npx); \
	if [ -z "$$exec_command" ]; then \
		echo "Error: please install bunx or npx to run this command."; \
		exit 1; \
	fi; \
	$$exec_command @tailwindcss/cli -i ./internal/web/styles/app.css -o ./internal/web/assets/main.css --watch
