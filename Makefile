.PHONY: docker-deploy
docker-deploy:
	@ docker build \
		--force-rm \
		--compress \
		--file $(pwd)/deployment/Dockerfile \
		--tag promotion:latest
		.