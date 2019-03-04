GEN_DIR ?= "gen"

.PHONY: check
check: 
	which rsvg-convert 1>/dev/null || (echo "\nInstall librsvg2 (ubuntu: librsvg2-bin) to convert SVGs!\n\n"; exit 1)
	
.PHONY: convert.pngs
convert.pngs: check
	@echo ">> Converting SVGs to PNGs"
	@mkdir -p ./gen/png
	@for f in `ls svg/`; \
		do echo "$${f}..." && rsvg-convert "svg/$${f}" > "./gen/png/`echo $${f} | sed 's/\(.*\)\..*/\1/'`.png"; \
	done;

.PHONY: serve
serve:
	go run src/cmd/serve/main.go .

.PHONY: update.manifest
update.manifest:
	go run src/cmd/manifest/main.go
