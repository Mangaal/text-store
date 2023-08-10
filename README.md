# Store

Store  project is a GoLang server application developed using the Gin web framework. It provides flexible API for storing, managing, and performing operations on uploaded files. 
This project provides endpoints to upload documents, removes uploaded documents, lists uploaded documents, and perform basic operations on these documents.


# Feature 
```
 Document Upload: Upload and store  various types of documents for storage in the server.  

 Document Listing: Allows to list of all uploaded documents.

 Count World: Able to count the total no. of words uploaded.

 Delete Document: Can Delete uploaded documents.

 Most Frequent Word: Retrieves the most frequent word used across all uploaded documents.
```

 #### All of these features can be easily utilized either through the CLI tools [store](https://github.com/Mangaal/store-cli#readme)  or by directly interacting with the exposed API.

# Geting Started

 1. Clone the repository:
      ```
       git clone https://github.com/Mangaal/text-store.git

       cd text-store

 2. Install on kubernetes.

    The manifest/deployment.yaml contain kubernetes kind deployment ,namespace and service.
    Ensure the values in the manifest are suitable for your environment. You may need to modify values like  port numbers, resource limits, etc
    ```
       kubectl apply -f manifest/deployment.yaml
    ```

  3. Testing

        ```
           curl http://serviceip:port/

# API Endpoints       
       

### Update or Upload Document. 

Endpoint: POST "apis/file"

Response:  200 OK on successful upload.

Curl Example:

```
      curl -X POST -F "files=@pod-2.yaml" -F "files=@pod.yaml"  http://localhost:80/apis/file

     response:
         {"message":"Files uploaded successfully"}
```


### Get list of Uploaded Document   

Endpoint: GET "apis/files"

Response:  200 OK on successful upload along with the list of the files.

Curl Example:

```
      curl http://localhost:80/apis/files

      response:
       {
          "files": [
            "apple.txt",
            "pod.yaml",
            "pod-2.yaml"
          ]
       }
```


 ### Delete

Endpoint: DELETE "apis/file"

Response:  200 OK on successful Delete.

Curl Example:

```
     curl -X DELETE -d  '{"files": ["new.yaml","new1.txt"] }'  http://localhost:80/apis/file

     response:
         {"message":"Files Deleted successfully"}
```

 ### Update File Name

Endpoint: POST "apis/file/:newname/:oldname"


Response:  200 OK on successful Update.

Curl Example:

```
     curl -X POST  http://localhost:80/apis/file/pod.yaml/apple.txt

     response:
         {"message":"Files Updated successfully"}
```


### Get Total Word count and Frequently Used Words


Endpoint: POST "apis/file/:sort/:limit"

```
sort: order on which the data will be sorted. a = ascending , d = descending
limit: retrieve the top  most frequent words. 
```
Response:  200 OK on successful Update.

Curl Example:

```
     curl -X POST   http://localhost:80apis/file/option/a/10

     response:

        {
             "items": [
               {
                 "Word": "rm",
                 "Frequency": 1
               },
               {
                 "Word": "on",
                 "Frequency": 1
               },
               {
                 "Word": "controller:",
                 "Frequency": 1
               },
               {
                 "Word": "shell",
                 "Frequency": 1
               },
               {
                 "Word": "command:",
                 "Frequency": 1
               },
               {
                 "Word": "same",
                 "Frequency": 1
               },
               {
                 "Word": "parallscssc",
                 "Frequency": 1
               },
               {
                 "Word": "absent.",
                 "Frequency": 1
               },
               {
                 "Word": "metadata:",
                 "Frequency": 1
               },
               {
                 "Word": "List",
                 "Frequency": 1
               }
             ],
             "totalWordCount": 192
        }
```

## Latest Docker Image release

```
 mangaaldochub/store-api:v1.0
```




Other Section
Development
k8s

 kubectl port-forward -n store-api  service/store-app-service 8080:80  



 
