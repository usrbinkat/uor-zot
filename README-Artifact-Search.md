# Artifact Search

This version of Zot has been modified to support enhanced Artifact search and management capabilities. It is a proof of concept. Expect the API and data model to change as we experiment with different approaches here. 

## Overview

Artifacts pushed to the registry with the statement mediatype ("application/vnd.uor.statement.v1+json") are indexed and made searchable with graphql.

Statements are written as semantic triples. Each element of a Statement triple contains a Resource Type and the resource being described. The Resource Type contains an OCI Reference to a Resource type definition which contains the schema and interface information of the referenced resource in the statement. A resource requires a type definition for a client to understand how to interact with it. 

## Jumping into the POC
*Must have oras.land client installed and in your path*
1. Build it: `make build binary`
2. Run it: `./artifact-test.sh`
3. Open it: http://localhost:8081/poc
4. Query it:
```graphql
{
  predicatesq(jsonPath: "test", jsonValue: "test predicate") {
    predicate
    predicatetype
    statement {
      subjects{
        subject
        subjecttype
      }
      objects {
        object
        objecttype
      }
    }
  }
}
```
# Understanding the POC
## Statements
Statements are written as OCI artifacts to the registry. When a manifest is uploaded to an Registry with the mediatype "application/vnd.oci.statement.v1+json" it is indexed and made searchable via graphql. A client can also search for a statement's blob address when they query the registry's graphql endpoint.

Example:
This file contains a statement example: [statement.json](./statement.json)

1. Let's push this to the registry with ORAS. Remember to set the mediatype of statement.json to "application/vnd.oci.statement.v1+json" on the command line when you publish the artifact:
```bash
oras push 127.0.0.1:8080/hello/test:v1 statement.json:application/vnd.uor.statement.v1+json --plain-http --verbose
```
2. Query the graphql endpoint for the statement. Sometimes you'll want to search by predicate, but other times, you may want to search by the object or subject of a statement. In this case, we'll search by the predicate:
```graphql
{
  predicatesq(jsonPath: "test", jsonValue: "test predicate") {
    predicate
    predicatetype
    statement {
      subjects{
        subject
      }
      objects {
        object
        objecttype
      }
    }
  }
}
```
jsonPath used in the above predicatesq query is the jsonpath of the statement. In this case, the jsonpath is "test", but we can use dot notation to access nested json objects. For example, if we had a statement with the jsonpath "test.nested", we could query it with "test.nested".

View the Statement schema here: [Statement Schema](./pkg/search/schema/statement_index.go)

## Additional Information
## Queries
This POC is the result of hacking entgo, graphql, and sqlite into Zot. The keen eye might notice that the database query mechanism is not ideal. A possibly more preferable approach might be to retrieve a referenced resource type definition from the registry's blob store and use the discovered schema to extend the database schema prior to indexing the statement. This will normalize the query of every part of the statement and result in a much more efficient query.

## Resource Types

Resource types are not yet implemented in the POC. When they are, they convey a resource's schema and also provide a reference to an interface or "driver" with which a client would interact with the resource by. 

As an example, if a client discovered a statement that referenced a resource on blockchain, it would retrieve the resource's type definition, that would have a reference to an interface for interacting with the resource. The client would then retrieve the interface and use it to interact with the resource. Depending on the situation, the interface could be used for CRUD operations or interacting with the resource in other ways, such as running the resource (if it were executable).

## Authorization
This server's graphql responses do not yet respect the registry's namespace authorization. Do not store anything in this POC that could represent a confidentiality concern.

## OCI Artifacts and Container Images
OCI Artifacts/Images are not yet indexed in this POC. The POC will need to index these objects in order to provide reverse lookups from content address to named reference. This can be accomplished by considering an artifact/image as a triple expression that can be stored within a statement. If we consider an artifact/image as a triple, then we can assert the manifest config as the triple's object, the manifest as the triple's predicate, and the manifest's layers as the subject. Every OCI artifact/image that is published to the registry would be indexed and searchable by manifest or layer digest via graphql. 