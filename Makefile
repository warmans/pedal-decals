DIST_DIR ?= "dist"

.PHONY: check
check: 
	which rsvg-convert 1>/dev/null || (echo "\nInstall librsvg2 (ubuntu: librsvg2-bin) to convert SVGs!\n\n"; exit 1)

.PHONY: convert.pngs
convert.pngs: check
	@echo ">> Converting SVGs to PNGs"
	@for f in `ls dist/svg/`; \
		do echo "$${f}..." && rsvg-convert "${DIST_DIR}/svg/$${f}" > "${DIST_DIR}/png/`echo $${f} | sed 's/\(.*\)\..*/\1/'`.png"; \
	done;

