// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent"
	picture "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/picture"
	product "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/product"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// PictureService implements PictureServiceServer
type PictureService struct {
	client *ent.Client
	UnimplementedPictureServiceServer
}

// NewPictureService returns a new PictureService
func NewPictureService(client *ent.Client) *PictureService {
	return &PictureService{
		client: client,
	}
}

// toProtoPicture transforms the ent type to the pb type
func toProtoPicture(e *ent.Picture) (*Picture, error) {
	v := &Picture{}
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	pictureurl := e.PictureURL
	v.PictureUrl = pictureurl
	if edg := e.Edges.Product; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Product = &Product{
			Id: id,
		}
	}
	return v, nil
}

// Create implements PictureServiceServer.Create
func (svc *PictureService) Create(ctx context.Context, req *CreatePictureRequest) (*Picture, error) {
	picture := req.GetPicture()
	m := svc.client.Picture.Create()
	picturePictureURL := picture.GetPictureUrl()
	m.SetPictureURL(picturePictureURL)
	var pictureProduct uuid.UUID
	if err := (&pictureProduct).UnmarshalBinary(picture.GetProduct().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetProductID(pictureProduct)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPicture(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements PictureServiceServer.Get
func (svc *PictureService) Get(ctx context.Context, req *GetPictureRequest) (*Picture, error) {
	var (
		err error
		get *ent.Picture
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetPictureRequest_VIEW_UNSPECIFIED, GetPictureRequest_BASIC:
		get, err = svc.client.Picture.Get(ctx, id)
	case GetPictureRequest_WITH_EDGE_IDS:
		get, err = svc.client.Picture.Query().
			Where(picture.ID(id)).
			WithProduct(func(query *ent.ProductQuery) {
				query.Select(product.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoPicture(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements PictureServiceServer.Update
func (svc *PictureService) Update(ctx context.Context, req *UpdatePictureRequest) (*Picture, error) {
	picture := req.GetPicture()
	var pictureID uuid.UUID
	if err := (&pictureID).UnmarshalBinary(picture.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Picture.UpdateOneID(pictureID)
	picturePictureURL := picture.GetPictureUrl()
	m.SetPictureURL(picturePictureURL)
	var pictureProduct uuid.UUID
	if err := (&pictureProduct).UnmarshalBinary(picture.GetProduct().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetProductID(pictureProduct)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPicture(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements PictureServiceServer.Delete
func (svc *PictureService) Delete(ctx context.Context, req *DeletePictureRequest) (*emptypb.Empty, error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Picture.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements PictureServiceServer.List
func (svc *PictureService) List(ctx context.Context, req *ListPictureRequest) (*ListPictureResponse, error) {
	var (
		err      error
		entList  []*ent.Picture
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Picture.Query().
		Order(ent.Desc(picture.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken, err := uuid.ParseBytes(bytes)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		listQuery = listQuery.
			Where(picture.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListPictureRequest_VIEW_UNSPECIFIED, ListPictureRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListPictureRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithProduct(func(query *ent.ProductQuery) {
				query.Select(product.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		var pbList []*Picture
		for _, entEntity := range entList {
			pbEntity, err := toProtoPicture(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListPictureResponse{
			PictureList:   pbList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}
