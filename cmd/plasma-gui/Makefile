run:
	go run main.go

js:
	gopherjs build -m -o main.js main.go

html:
	echo '<!DOCTYPE html>\n<script src="main.js"></script>' > index.html

deploy: js
	scp main.js data.gopher.se:/var/www/data.gopher.se/js/plasma/
