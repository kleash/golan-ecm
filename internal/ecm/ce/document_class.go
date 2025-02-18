package ce

// DocumentClass represents a category of a document
type DocumentClass interface {

	// GetWorkspaceId returns the Id of the workspace that owns this DocumentClass
	GetWorkspaceId() string

	// SetWorkspaceId assigns the workspace owner of this DocumentClass
	SetWorkspaceId(objectId string) error

	// GetName returns the name of the DocumentClass.
	// e.g. PolicyDocument, MedicalDocument_XRay, MedicalDocument_Cardio
	GetName() string

	// GetLabel returns the human-readable name of the DocumentClass
	// e.g. Policy Document, Medical Document / XRay, Medical Document / Cardio
	GetLabel() string

	// GetDescription returns the description of the DocumentClass
	// e.g. An Insurance Policy Document
	GetDescription() string

	SetPropertyFields(propertyFields []PropertyField)
	GetPropertyFields() []PropertyField

	Object
}
