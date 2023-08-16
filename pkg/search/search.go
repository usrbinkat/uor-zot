package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	"zotregistry.io/zot/ent"
	swhere "zotregistry.io/zot/ent/statement"
	sschema "zotregistry.io/zot/pkg/search/schema" // Only struct definitions. No ent definitions.
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
	ctx := context.Background()
	bytes, err := json.Marshal(descriptor)
	if err != nil {
		return fmt.Errorf("marshalling error: %v", err)
	}
	var mdescriptor map[string]interface{}
	if err := json.Unmarshal(bytes, &mdescriptor); err != nil {
		return fmt.Errorf("unmarshalling error: %v", err)
	}

	// Query existing statements with the same namespace
	existingStatements, err := eclient.Statement.Query().
		Where(swhere.NamespaceEQ(repo)).
		All(ctx)
	if err != nil {
		return fmt.Errorf("error querying statements: %v", err)
	}
	for _, existingStatement := range existingStatements {
		existingStatementJSON, _ := json.Marshal(existingStatement.Statement)
		newStatementJSON, _ := json.Marshal(mdescriptor)
		if string(existingStatementJSON) == string(newStatementJSON) {
			return fmt.Errorf("duplicate statement found for namespace: %s", repo)
		}
	}

	statementCreate := eclient.Statement.Create().SetStatement(mdescriptor).SetNamespace(repo)

	if statement.Object.Noun != nil {
		object, err := eclient.Object.Create().
			SetObject(statement.Object.Noun).
			SetObjectType(statement.Object.ObjectType).
			Save(ctx)
		if err != nil {
			return err
		}
		statementCreate.AddObjects(object)
	}

	if statement.Predicate.Noun != nil {
		predicate, err := eclient.Spredicate.Create().
			SetPredicate(statement.Predicate.Noun).
			SetPredicateType(statement.Predicate.PredicateType).
			Save(ctx)
		if err != nil {
			return err
		}
		statementCreate.AddPredicates(predicate)
	}

	if statement.Subject.Noun != nil {
		subject, err := eclient.Subject.Create().
			SetSubject(statement.Subject.Noun).
			SetSubjectType(statement.Subject.SubjectType).
			Save(ctx)
		if err != nil {
			return err
		}
		statementCreate.AddSubjects(subject)
	}

	_, err = statementCreate.Save(ctx)
	if err != nil {
		return fmt.Errorf("error saving statement: %v", err)
	}

	return nil
}
