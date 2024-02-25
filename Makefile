run:
	make gen
	@go run cmd/main.go
gen:
	templ generate
	./tailwindcss  -i ./static/tailwind.css -o ./static/dist/styles.css

templ:
	@templ generate -watch -proxy=http://localhost:3000

tailwind:
	./tailwindcss  -i ./static/tailwind.css -o ./static/dist/styles.css --watch