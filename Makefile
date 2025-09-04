.PHONY: new_exercise execute_exercise

exercise ?= merged-sorted-array

generate_readme:
	python3 scripts/leetcode.py $(exercise)

execute_exercise:
	go run tasks/$(exercise)/main.go

new_exercise:
	@read -p "Enter exercise name (e.g., two-sum): " exercise_name; \
	if [ -z "$$exercise_name" ]; then \
		echo "Exercise name cannot be empty."; \
		exit 1; \
	fi; \
	sanitized_name=$$(echo "$$exercise_name" | tr ' ' '-' | tr '[:upper:]' '[:lower:]'); \
	exercise_path="tasks/$$sanitized_name"; \
	if [ -d "$$exercise_path" ]; then \
		echo "Error: Directory '$$exercise_path' already exists."; \
		exit 1; \
	fi; \
	echo "Creating new exercise in '$$exercise_path'..."; \
	mkdir -p "$$exercise_path"; \
	(cd "$$exercise_path" && go mod init "leetcode/$$sanitized_name" && echo -e "package main\n\nfunc main() {\n\n}" > main.go && touch README.md); \
	echo "Project '$$sanitized_name' created successfully.";
