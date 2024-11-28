# linux
Shell script for Process Management :

1. The script will list all running processes.
2. It will accept a process ID (PID) and display detailed information about the process : PID, user, group, memory and CPU consumed by the process, etc
3. Allow the user to terminate a process by PID.
4. Handle errors, such as invalid PID or insufficient permissions, etc
5. Handle all possible error messages for each of the above operations and print them on the screen.

# Docker  
Create a simple web application that displays information : ID, Name, Age, Gender, etc of multiple people (Use any frontend - backend tool).
The website should have an option / form to add, delete and edit records. At run time, I should be able to add, edit and delete records and the same should reflect on the web application.
Use any database (mysql, postgres, mongodb,etc) to save these records.
Write docker files to make this web application work end to end. Everything should be containerised (frontend, backend, database, etc).
Don't use ready docker images like mysql, nodejs, php, mongodb, etc. Create your Dockerfile from scratch. Try to make docker images as lightweight as possible.
Configure the volume for the database such that it will store the data outside the container so that data is not lost when the database container gets stopped / killed.
Write a docker-compose file to include all these containers. Configure the file in such a way that the database container gets started before the web app container.

# Kubernetes 
Create a simple web application that displays information : ID, Name, Age, Gender, etc of multiple people (Use any frontend - backend tool). The website should have an option / form to add, delete and edit records. At run time, I should be able to add, edit and delete records and the same should reflect on the web application.
Use one of these databases - mysql, postgres, mongodb to save these records. Don't use any other database apart from the ones mentioned above. Those who had used SQLite will have to create the docker image again using one of the 2 mentioned above.
Write docker files to make this web application work end to end. Everything should be containerised (frontend, backend, database, etc). Each component should have a separate docker image. Those who had merged 2 docker containers into 1 last time, should separate it out now.
Don't use ready docker images like mysql, nodejs, php, mongodb, etc. Create your Dockerfile from scratch. Try to make docker images as lightweight as possible.
Setup a local kubernetes cluster using minikube or kubeadm. It should have 3 nodes - one master and 2 worker nodes. Configure the nodes such that user applications don't get deployed on the master node. 
Create a deployment for the database using the above Docker image. Configure the imagePullPolicy to use your local image. Create PV and PVC to store the data. Create a Secret / ConfigMap to store the database credentials : username, password, DB name, etc. Create a Service for the database to expose it. Configure your deployment spec to use these secret and service.
Create a deployment for the frontend using the frontend docker image that you created above. Configure the imagePullPolicy to use your local image. This deployment should have 3 replicas. Expose the frontend using a Service of type NodePort or LoadBalancer. Configure the deployment to use Rolling update.
Configure the deployments so that they are able to communicate with each other.
For the database PV, the data should be retained even if the pod is deleted / restarted.
Try scaling the frontend application, restarting / deleting the pods, upgrading the pods while you are inserting data from the UI. See that the data is preserved inside the database during the above operations.
