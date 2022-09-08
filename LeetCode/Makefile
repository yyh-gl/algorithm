.DEFAULT_GOAL := help
.PHONY: help
help: ## helpを表示
	@echo '  see: https://github.com/yyh-gl/algorithm'
	@echo ''
	@grep -E '^[%/0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'
	@echo ''

.PHONY: setup
setup: ## setup
	docker build . -t algorithm

.PHONY: login
login: ## login
	@docker run --rm -it \
	  -v `pwd`:/go/src/github.com/yyh-gl/Algorithm/LeetCode \
	  algorithm \
	  /bin/sh

.PHONY: new
new: ## create new template for LeetCode's problem
	@docker run --rm -i \
	  -v `pwd`:/go/src/github.com/yyh-gl/Algorithm/LeetCode \
	  algorithm \
	  leetgode generate ${N}

# TODO 動かない
.PHONY: test
test: ## test code for LeetCode's problem on online
	@docker run --rm -i \
	  --env-file=.env \
	  -v `pwd`:/go/src/github.com/yyh-gl/Algorithm/LeetCode \
	  algorithm \
	  leetgode test ${N}

# TODO 動かない
.PHONY: submit
submit: ## submit code for LeetCode's problem
	@docker run --rm -i \
	  -v `pwd`:/go/src/github.com/yyh-gl/Algorithm/LeetCode \
	  algorithm \
	  leetgode exec ${N}
