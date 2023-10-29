

get:
	curl http://localhost:8080/books
	curl http://localhost:8080/books/1

post:
	curl http://localhost:8080/books \
		--include \
		--header "Content-Type: application/json" \
		--request "POST" \
		--data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

put:
	curl http://localhost:8080/books/1 \
		--include \
		--header "Content-Type: application/json" \
		--request "PUT" \
		--data '{"id": "1","price": 100.99}'

delete:
	curl http://localhost:8080/books/4 \
		--include \
		--header "Content-Type: application/json" \
		--request "DELETE"