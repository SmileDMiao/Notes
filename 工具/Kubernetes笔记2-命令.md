## 命令
```shell
# get pods of namesapce
kubectl get pods -n staging

# container sh
kubectl exec -it pods -n namespace sh

# exec sh and attach 
# production
kubectl exec "$(kubectl get pod --namespace=production -l app.kubernetes.io/name=dongapp-server -o jsonpath="{.items[0].metadata.name}")" --namespace production -it -- sh
kubectl exec "$(kubectl get pod --namespace=production -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace production -it -- sh
kubectl attach "$(kubectl get pod --namespace=production -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace production

# Staging
kubectl exec "$(kubectl get pod --namespace=staging -l app.kubernetes.io/name=dongapp-server -o jsonpath="{.items[0].metadata.name}")" --namespace staging -it -- sh
kubectl exec "$(kubectl get pod --namespace=staging -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace staging -it -- sh
kubectl attach "$(kubectl get pod --namespace=staging -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace staging

# Canary
kubectl exec "$(kubectl get pod --namespace=canary -l app.kubernetes.io/name=dongapp-server -o jsonpath="{.items[0].metadata.name}")" --namespace canary -it -- sh
kubectl exec "$(kubectl get pod --namespace=canary -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace canary -it -- sh
kubectl attach "$(kubectl get pod --namespace=canary -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace canary
```