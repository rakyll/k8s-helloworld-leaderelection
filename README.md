```
$ kubectl apply -f ./deploy.yaml

$ kubectl get pods
NAME                             READY   STATUS    RESTARTS   AGE
leaderelection-968bc8fb4-g2pjr   1/1     Running   0          5m19s
leaderelection-968bc8fb4-p5rx6   1/1     Running   0          5m19s
leaderelection-968bc8fb4-ssmfl   1/1     Running   0          5m19s
leaderelection-968bc8fb4-tvrsj   1/1     Running   0          5m19s
leaderelection-968bc8fb4-zq762   1/1     Running   0          5m19s

$ kubectl logs -f leaderelection-968bc8fb4-g2pjr
I0415 20:02:12.342517       1 leaderelection.go:248] attempting to acquire leader lease leaderelection-example/lock...
2022/04/15 20:02:12 New leader is "leaderelection-78c5b56b9-4d8bx"
2022/04/15 20:02:22 New leader is "leaderelection-968bc8fb4-p5rx6"
```
