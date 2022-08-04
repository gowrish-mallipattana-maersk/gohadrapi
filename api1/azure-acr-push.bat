docker rmi sreacr1.azurecr.io/gohadrapi1:d.0.1

docker login sreacr1.azurecr.io
docker tag gohadrapi1:d.0.1 sreacr1.azurecr.io/gohadrapi1:d.0.1
docker push sreacr1.azurecr.io/gohadrapi1:d.0.1
