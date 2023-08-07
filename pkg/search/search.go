package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"zotregistry.io/zot/ent"

	_ "github.com/mattn/go-sqlite3"
	godigest "github.com/opencontainers/go-digest"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"zotregistry.io/zot/pkg/search/schema"
	storageTypes "zotregistry.io/zot/pkg/storage/types"
)

func InitDatabase() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return nil, err
	}
	fmt.Println("sqlite database initialized")
	return client, nil
}

func AddStatement(imgStore storageTypes.ImageStore, repo string, digest godigest.Digest, eclient *ent.Client) error {
	fmt.Println("AddStatement called")
	sdescriptor, err := imgStore.GetStatementDescriptor(repo, digest)
	if err != nil {
		fmt.Println("No statement found")
		return err
	}
	var uDescriptor specs.Descriptor
	if err := json.Unmarshal(sdescriptor, &uDescriptor); err != nil {
		fmt.Println("error unmarshalling statement descriptor")
		return err
	}
	fmt.Printf("unmarshalled descriptor: %v\n", uDescriptor)

	statement, err := imgStore.GetBlobContent(repo, digest)
	if err != nil {
		fmt.Println("error getting blob content")
		return err
	}
	fmt.Printf("statementfound: %v\n", string(statement))
	uStatement := schema.Statement{}
	if err := json.Unmarshal(statement, &uStatement); err != nil {
		fmt.Println("error unmarshalling statement")
		return err
	}
	fmt.Printf("unmarshalled statement: %v\n", uStatement)

	record := schema.StatementRecord{}
	fmt.Println("empty record formed")
	record.Statement = &uStatement
	fmt.Println("statement added to formed record")
	record.Location = &schema.Location{}
	record.Location.Descriptor = uDescriptor
	fmt.Println("location added to formed record")
	record.Location.Namespace = repo
	fmt.Println("namespace added to formed record")

	recordbytes, err := json.Marshal(record)
	if err != nil {
		fmt.Println("error marshalling record")
		return err
	}
	fmt.Printf("record to be written: %v\n", string(recordbytes))
	ctx := context.Background()

	indexed, err := eclient.StatementIndex.
		Create().
		SetPredicate(*record.Statement.Predicate).
		SetSubject(*record.Statement.Subject).
		SetObject(*record.Statement.Object).
		SetStatement(*record.Location).
		Save(ctx)
	if err != nil {
		fmt.Println("error saving statement")
		return err
	}
	fmt.Printf("u: %v\n", indexed)
	return nil
}

func DeleteStatement() {

}

func QueryStatement() {

}
