build_docker: 
	docker build -t app .

run_docker:
	docker run -p 4000:4000 --rm app