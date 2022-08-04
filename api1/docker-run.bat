rem docker run -d --name gohadrapi1 -p 10001:80 -e API_TO_CALL="http://localhost:10002" gohadrapi1:d.0.1

docker run -d --name gohadrapi1 -p 10001:80 -e API_TO_CALL="http://pokeapi.co/api/v2/pokedex/kanto/" gohadrapi1:d.0.1
