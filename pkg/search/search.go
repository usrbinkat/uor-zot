package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	storageTypes "zotregistry.io/zot/pkg/storage/types"

	_ "github.com/mattn/go-sqlite3"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	"zotregistry.io/zot/ent"
	swhere "zotregistry.io/zot/ent/statement"
	sschema "zotregistry.io/zot/pkg/search/schema" // Only struct definitions. No ent definitions.
)

func InitDatabase(is storageTypes.ImageStore) (*ent.Client, error) {
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
	fmt.Printf("preparing to write statement: %v\n", statement)

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
			fmt.Printf("existing statement: %v\n", existingStatement.Statement)
			fmt.Printf("new statement: %v\n", mdescriptor)
			fmt.Printf("duplicate statement found for namespace: %s", repo)
			return nil
		}
	}

	statementCreate := eclient.Statement.Create().SetStatement(mdescriptor).SetNamespace(repo)
	fmt.Printf("preparing to write statement: %v\n", statement)
	if statement.Object != nil && statement.Object.Noun != nil {
		object, err := eclient.Object.Create().
			SetObject(statement.Object.Noun).
			SetObjectType(statement.Object.ObjectType).
			Save(ctx)
		if err != nil {
			return err
		}
		statementCreate.AddObjects(object)
	}

	if statement.Predicate != nil && statement.Predicate.Noun != nil {
		predicate, err := eclient.Spredicate.Create().
			SetPredicate(statement.Predicate.Noun).
			SetPredicateType(statement.Predicate.PredicateType).
			Save(ctx)
		if err != nil {
			return err
		}
		statementCreate.AddPredicates(predicate)
	}

	if statement.Subject != nil && statement.Subject.Noun != nil {
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

func Manifest2Statement(manifest ispec.Manifest) (sschema.Statement, error) {
	var statement sschema.Statement
	fmt.Println("Manifest2Statement called")

	// Handle the config object
	bConfig, err := json.Marshal(manifest.Config)
	if err != nil {
		return statement, fmt.Errorf("error marshalling config: %v", err)
	}
	fmt.Println("config marshalled")
	mConfig := make(map[string]interface{})
	if err := json.Unmarshal(bConfig, &mConfig); err != nil {
		return statement, fmt.Errorf("error unmarshalling config: %v", err)
	}
	fmt.Println("config unmarshalled")
	if len(mConfig) != 0 {
		statement.Object = &sschema.Object{
			ObjectType: manifest.Config.MediaType,
			Noun:       mConfig,
		}

		fmt.Printf("config is: %v\n", statement.Object)
	} else {
		statement.Object = nil
		fmt.Println("config is nil")
	}

	mLayers := make(map[string]interface{})
	for i, layer := range manifest.Layers {
		bLayer, err := json.Marshal(layer)
		if err != nil {
			return statement, fmt.Errorf("error marshalling layer: %v", err)
		}
		var layerMap map[string]interface{}
		if err := json.Unmarshal(bLayer, &layerMap); err != nil {
			return statement, fmt.Errorf("error unmarshalling layer: %v", err)
		}
		mLayers[fmt.Sprintf("layer%d", i)] = layerMap
	}
	statement.Subject = &sschema.Subject{
		SubjectType: manifest.MediaType,
		Noun:        mLayers,
	}

	fmt.Printf("layers are: %+v\n", statement.Subject)

	cManifest := ispec.Manifest{}
	cManifest = manifest
	cManifest.Layers = nil
	cManifest.Config = ispec.Descriptor{}
	bManifest, err := json.Marshal(cManifest)
	if err != nil {
		return statement, fmt.Errorf("error marshalling manifest: %v", err)
	}
	mManifest := make(map[string]interface{})
	if err := json.Unmarshal(bManifest, &mManifest); err != nil {
		return statement, fmt.Errorf("error unmarshalling manifest: %v", err)
	}
	statement.Predicate = &sschema.Predicate{
		Noun:          mManifest,
		PredicateType: manifest.MediaType,
	}

	fmt.Printf("statement: %+v\n", statement)

	return statement, nil
}
