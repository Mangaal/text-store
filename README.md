# text-store



## apis

### Create 

 curl -X POST -F "files=@pod-2.yaml" -F "files=@pod.yaml"  http://localhost:80/apis/file

### Update

 curl -X POST -d    http://localhost:80/apis/file/new.txt/oldname.txt

 ### Delete

 curl -X POST -d    http://localhost:80/apis/file/new.txt

 ### Options
curl   http://localhost:80/apis/file/option/d/2

response
{"items":[{"Word":"name:","Frequency":2},{"Word":"-","Frequency":2}],"totalWordCount":23}



 