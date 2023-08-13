package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	"zotregistry.io/zot/ent"

	_ "github.com/mattn/go-sqlite3"
	sschema "zotregistry.io/zot/pkg/search/schema"
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

func AddStatement(statement sschema.Statement, repo string, descriptor ispec.Descriptor, eclient *ent.Client) error {
	fmt.Println("AddStatement called")

	ctx := context.Background()

	object, err := eclient.Object.
		Create().
		SetObject(statement.Object.Noun).
		SetObjectType(statement.Object.ObjectType).
		Save(ctx)
	if err != nil {
		fmt.Println("error saving statement")
		return err
	}
	fmt.Printf("u: %v\n", object)

	predicate, err := eclient.Spredicate.
		Create().
		SetPredicate(statement.Predicate.Noun).
		SetPredicateType(statement.Predicate.PredicateType).
		Save(ctx)
	if err != nil {
		fmt.Println("error saving statement")
		return err
	}
	fmt.Printf("u: %v\n", predicate)

	subject, err := eclient.Subject.
		Create().
		SetSubject(statement.Subject.Noun).
		SetSubjectType(statement.Subject.SubjectType).
		Save(ctx)
	if err != nil {
		fmt.Println("error saving statement")
		return err
	}
	fmt.Printf("u: %v\n", subject)

	bytes, err := json.Marshal(descriptor)
	if err != nil {
		fmt.Printf("Marshalling error: %v", err)
	}
	var mdescriptor map[string]interface{}
	if err := json.Unmarshal(bytes, &mdescriptor); err != nil {
		fmt.Printf("unmarshalling error: %v", err)
	}
	indexed, err := eclient.Statement.
		Create().
		AddSubjects(subject).
		AddObjects(object).
		AddPredicates(predicate).
		SetStatement(mdescriptor).
		SetNamespace(repo).
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
