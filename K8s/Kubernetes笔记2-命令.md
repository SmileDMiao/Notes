## 命令
```shell
# get pods of namesapce
kubectl get pods -n staging

# container sh
kubectl exec -it pods -n namespace sh

# exec sh and attach 
kubectl exec "$(kubectl get pod --namespace=production -l app.kubernetes.io/name=dongapp-server -o jsonpath="{.items[0].metadata.name}")" --namespace production -it -- sh
kubectl attach "$(kubectl get pod --namespace=production -l app.kubernetes.io/name=sports-data -o jsonpath="{.items[0].metadata.name}")" --namespace production
```