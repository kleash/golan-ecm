package service

import (
	"cjavellana.me/ecm/golan/internal/ecm/ce"
	"cjavellana.me/ecm/golan/internal/ecm/pb"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"strings"
	"time"
)

type ObjectStoreService struct {
	ObjectStore ce.ObjectStore

	// Required for future compatibility see https://github.com/grpc/grpc-go/blob/master/cmd/protoc-gen-go-grpc/README.md
	pb.UnimplementedContentEngineServer
}

func (s *ObjectStoreService) CreateWorkspace(_ context.Context, in *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceResponse, error) {

	log.Infof("received create workspace request: %s", in.WorkspaceName)

	now := time.Now()

	w := s.ObjectStore.NewWorkspace(in.WorkspaceName, in.Description)
	w.SetCreatedBy("UserFromAuthToken")
	w.SetDateCreated(&now)

	w, _ = s.ObjectStore.SaveWorkspace(w)

	return &pb.CreateWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}

func (s *ObjectStoreService) GetWorkspace(
	_ context.Context,
	in *pb.GetWorkspaceRequest) (*pb.GetWorkspaceResponse, error) {

	log.Infof("received get workspace query: %s", in.Query)

	w, err := s.ObjectStore.GetWorkspaceByObjectId(in.GetQuery())
	if err != nil {
		return nil, err
	}

	return &pb.GetWorkspaceResponse{
		ObjectId: w.ObjectId(),
	}, nil
}

func (s *ObjectStoreService) CreateDocumentClass(
	_ context.Context,
	in *pb.CreateDocumentClassRequest) (*pb.CreateDocumentClassResponse, error) {

	docClass := s.ObjectStore.NewDocumentClass(in.Name, in.Label, in.Description)

	err := docClass.SetWorkspaceId(in.WorkspaceId)
	if err != nil {
		// an error may be returned when the given workspace id
		// cannot converted to a valid mongodb id
		return nil, err
	}

	var propertyFields []ce.PropertyField
	for _, propertyField := range in.PropertyFields {
		fieldType := ce.FieldType(strings.ToLower(propertyField.FieldType.String()))

		p := s.ObjectStore.NewPropertyField(
			propertyField.Name,
			propertyField.Label,
			fieldType,
			propertyField.Description,
		)
		propertyFields = append(propertyFields, p)
	}

	docClass.SetPropertyFields(propertyFields)

	docClass, err = s.ObjectStore.SaveDocumentClass(docClass)
	if err != nil {
		return nil, err
	}

	return &pb.CreateDocumentClassResponse{
		ObjectId:    docClass.ObjectId(),
		Name:        in.Name,
		Label:       in.Label,
		Description: in.Description,
	}, nil
}
