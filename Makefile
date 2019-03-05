GEN_DIR ?= "gen"

.PHONY: check
check: 
	@which rsvg-convert 1>/dev/null || (echo "\nInstall librsvg2 (ubuntu: librsvg2-bin) to convert SVGs!\n\n"; exit 1)
	
.PHONY: convert.pngs
convert.pngs: check
	@echo ">> Converting SVGs to PNGs"
	@mkdir -p ./gen/png
	@for f in `ls svg/`; \
		do echo "$${f}..." && rsvg-convert "svg/$${f}" > "./gen/png/`echo $${f} | sed 's/\(.*\)\..*/\1/'`.png"; \
	done;

.PHONY: gallery.serve
gallery.serve:
	@echo "OPEN http://localhost:8080/pedal-decals/gallery/"
	@go run tools/cmd/serve/main.go ../

.PHONY: gallery.update
gallery.update:
	@go run tools/cmd/manifest/main.go
