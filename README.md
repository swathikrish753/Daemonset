# Daemonset
step 1:
kubectl apply -f deamonset.yaml,   to Create the DaemonSet in the Kubernetes cluster
step 2:
kubectl get daemonset, get a list of DaemonSets in your Kubernetes cluster 
step 3:
kubectl port-forward compute-usage-monitor-fnb6m 8080:8080, any traffic sent to localhost:8080 on your local machine will be forwarded to port 8080 on the specified pod.
step 4:
http://localhost:8080/metrics, hit this api in postman to fetch the cpu and memory utilisation of the host system
