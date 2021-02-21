.PHONY: docker-deploy
docker-deploy:
	@ docker build --force-rm --compress --file $(pwd)deployment\app\Dockerfile --tag bitcodr/promotion:latest .