mkdir logs

docker stop poder-judicial-backend && docker rm poder-judicial-backend

docker image rm poder-judicial-backend

docker build -t poder-judicial-backend .

docker run -d \
--restart always \
--name poder-judicial-backend \
--net=luis \
-p 8890:80 \
-v $(pwd)/logs:/etc/app \
poder-judicial-backend