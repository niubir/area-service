docker build -t niubir/area-service:test .

docker run -p 10011:10011 -p 10012:10012 -v /root/area-service/config:/config -v /root/area-service/data/:/data -v /root/area-service/logs/:/logs --name=area-service-1 -d niubir/area-service:test

docker push niubir/area-service:test

docker inspect --format='{{json .Mounts}}' area-service-1