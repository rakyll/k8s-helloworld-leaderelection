ecr = public.ecr.aws/q1p8v8z2
repo = leaderelection-example
region = us-east-1

publish:
	KO_DOCKER_REPO=$(ecr)/$(repo) ko publish --bare ./

login:
	aws ecr-public get-login-password --region $(region) | docker login --username AWS --password-stdin $(ecr)