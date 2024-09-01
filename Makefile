# Define the compiler for Go
CC = go
CC_OUT = dist/

# Define the compiler flags
CFLAGS = -ldflags="-s -w"

# Define the color codes for output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput sgr0)

.PHONY: all

all: help

# Check if the target is 'build' and ensure 'file' is provided
ifeq ($(MAKECMDGOALS),build)
  ifeq ($(file),)
    $(error You must provide a source file. Usage: make build file=source.go)
  endif
  OUT_NAME := $(basename $(notdir $(file)))
  OUT = -o $(CC_OUT)
endif

## Build:
build: ## Compile the provided source file into an executable
	@echo "${YELLOW}Building...${RESET}"
	@if [ ! -d "$(CC_OUT)" ]; then \
		mkdir -p "$(CC_OUT)"; \
	fi
	$(CC) build $(CFLAGS) $(OUT) $(file)
	@echo "${GREEN}Build complete!${RESET}"

# Define the list of extensions or filenames to ignore during clean
IGNORE_EXTENSIONS = mod go txt docx png
IGNORE_NAMES = Makefile README.md .gitignore

## Clean:
clean: ## Clean the directory by removing all files that are not in the ignore list
	find . -type f $(foreach ext, $(IGNORE_EXTENSIONS), ! -name '*.$(ext)') \
		$(foreach name, $(IGNORE_NAMES), ! -name '$(name)') -not -path './.git/*' -exec rm -f {} +
	rm -fr $(CC_OUT)

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  $(YELLOW)make$(RESET) $(GREEN)<target>$(RESET) $(GREEN)<options>$(RESET)'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    $(YELLOW)%-20s$(GREEN)%s$(RESET)\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  $(CYAN)%s$(RESET)\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)