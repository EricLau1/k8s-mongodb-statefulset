# Golang e MongoDB com Kubernetes no Minikube 

- Fazer o build do app

```bash
sh build.sh
```

- Aplicar o deployment

```bash
kubectl apply -f deployment.yaml
```

- Gerar um EXTERNAL-IP

```bash
    # Executar este comando em um terminal separado
    minikube tunnel

    # vizualizar o EXTERNAL-IP
    kubectl get svc
```

- Testar com as rotas:

```bash
# retorna dados do request
curl http://<EXTERNAL-IP>:9000

# faz um Ping no MongoDB
curl http://<EXTERNAL-IP>:9000/ping
```

## Referências: 

- https://minikube.sigs.k8s.io/docs/handbook/accessing/