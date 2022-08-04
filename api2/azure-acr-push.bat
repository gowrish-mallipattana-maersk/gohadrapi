docker rmi sreacr1.azurecr.io/gohadrapi2:d.0.1

docker login sreacr1.azurecr.io
docker tag gohadrapi2:d.0.1 sreacr1.azurecr.io/gohadrapi2:d.0.1
docker push sreacr1.azurecr.io/gohadrapi2:d.0.1
