package schema

import (
	specs "github.com/opencontainers/image-spec/specs-go"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type Subject struct {
	// SubjectType is the location of the Resource that
	// explains client interaction with the Subject.
	//
	// PredicateType is a Docker URL.
	SubjectType string `json:"subjectType,omitempty"`
	// Subject is a Resource used to locate the Subject
	// of the Statement.
	//
	// In practical terms, the Subject is the thing being
	// described by the predicate or the target of the
	// relationship of the Object.
	Noun ResourceAddress `json:"noun,omitempty"`
}

// Predicate describes a triple's predicate
type Predicate struct {
	// PredicateType is the location of the Resource that
	// explains client interaction with the Predicate.
	//
	// PredicateType is a Docker URL.
	PredicateType string `json:"predicateType,omitempty"`
	// Predicate is a Resource used to express the Predicate
	// of the Statement.
	//
	// In practical terms, the Predicate is the description
	// of the Subject or the description of the relationship
	// of the Object to the Subject.
	Noun ResourceAddress `json:"noun,omitempty"`
}

// Object describes a triple's object
type Object struct {
	// ObjectType is the location of the Resource that
	// explains client interaction with the Object.
	//
	// ObjectType is a Docker URL.
	ObjectType string `json:"objectType,omitempty"`
	// Object is a Resource used to locate the Object
	// of the Statement.
	//
	// In practical terms, the Object is the thing
	// establishing the relationship to the Subject
	// that is described by the Predicate.
	Noun ResourceAddress `json:"Noun,omitempty"`
}

// Resource is the address of a resource. It is a
// json.RawMessage so that the ResourceType can
// specify its address format and retrieval
// instructions.
type ResourceAddress map[string]interface{}

// Statement provides `application/vnd.oci.statement.v1+json` mediatype structure when marshalled to JSON.
type Statement struct {
	specs.Versioned

	// MediaType specifies the type of this document data structure e.g. `application/vnd.oci.statement.v1+json`
	MediaType string `json:"mediaType,omitempty"`

	// Subject is the subject of the statement.
	Subject *Subject `json:"subject,omitempty"`

	// Predicate is the predicate of the statement.
	Predicate *Predicate `json:"predicate,omitempty"`

	// Object is the object of the statement.
	Object *Object `json:"object,omitempty"`
}

type StatementRecord struct {
	Statement *Statement `json:"statement,omitempty"`
	Location  *Location  `json:"location,omitempty"`
}

// Location is the
type Location struct {
	Descriptor ispec.Descriptor `json:"descriptor,omitempty"`
	Namespace  string           `json:"namespace,omitempty"`
}
