# MongoDB StatefulSet


- Iniciar o minikube

```bash
minikube start

eval $(minikube docker-env)
```

- Rodar o script

```bash
sh deploy.sh
```

- Interarir com MongoDB

```bash
kubectl exec -it mongodb-0 /bin/bash

mongo -u admin -p password --authenticationDatabase admin

# logar como admin
db.auth('admin','password');

# Acessar o banco admin
use admin

# mostrar usuários
show users

# mostrar bancos disponíveis
show dbs

# Criar um novo banco
use teste

# Criar um novo usuário com acesso apenas ao novo banco
db.createUser({user: 'root', pwd: 'root', roles:[{'role': 'readWrite', 'db': 'teste'}]});

# mostrar usuários
show users

# Logar com o novo usuário
mongo -u root -p root --authenticationDatabase teste

# O novo usuário consegue ler e escrever apenas no novo banco
use teste

show collections

db.items.insert({name: 'TESTE'});

db.items.find();
```
## Referências

https://medium.com/@dilipkumar/standalone-mongodb-on-kubernetes-cluster-19e7b5896b27

https://kubernetes.io/docs/concepts/configuration/configmap/

https://kubernetes.io/docs/concepts/services-networking/service/