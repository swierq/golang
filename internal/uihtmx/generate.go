package uihtmx

//go:generate go run github.com/a-h/templ/cmd/templ@v0.3.833 generate

//go:generate tailwindcss -c ./tailwind.config.js -i ./ui/assets/css/input.css -o ./ui/assets/css/tailwind.css --minify
