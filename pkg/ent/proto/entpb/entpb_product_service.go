// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent"
	category "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/category"
	order "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/order"
	product "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/product"
	user "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/user"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// ProductService implements ProductServiceServer
type ProductService struct {
	client *ent.Client
	UnimplementedProductServiceServer
}

// NewProductService returns a new ProductService
func NewProductService(client *ent.Client) *ProductService {
	return &ProductService{
		client: client,
	}
}

// toProtoProduct transforms the ent type to the pb type
func toProtoProduct(e *ent.Product) (*Product, error) {
	v := &Product{}
	description := wrapperspb.String(e.Description)
	v.Description = description
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	name := e.Name
	v.Name = name
	pictureurl := wrapperspb.String(e.PictureURL)
	v.PictureUrl = pictureurl
	price := e.Price
	v.Price = price
	quantity := int32(e.Quantity)
	v.Quantity = quantity
	for _, edg := range e.Edges.Categories {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Categories = append(v.Categories, &Category{
			Id: id,
		})
	}
	for _, edg := range e.Edges.Orders {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Orders = append(v.Orders, &Order{
			Id: id,
		})
	}
	for _, edg := range e.Edges.ShoppingCartOwners {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.ShoppingCartOwners = append(v.ShoppingCartOwners, &User{
			Id: id,
		})
	}
	return v, nil
}

// Create implements ProductServiceServer.Create
func (svc *ProductService) Create(ctx context.Context, req *CreateProductRequest) (*Product, error) {
	product := req.GetProduct()
	m := svc.client.Product.Create()
	if product.GetDescription() != nil {
		productDescription := product.GetDescription().GetValue()
		m.SetDescription(productDescription)
	}
	productName := product.GetName()
	m.SetName(productName)
	if product.GetPictureUrl() != nil {
		productPictureURL := product.GetPictureUrl().GetValue()
		m.SetPictureURL(productPictureURL)
	}
	productPrice := float64(product.GetPrice())
	m.SetPrice(productPrice)
	productQuantity := int(product.GetQuantity())
	m.SetQuantity(productQuantity)
	for _, item := range product.GetCategories() {
		var categories uuid.UUID
		if err := (&categories).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddCategoryIDs(categories)
	}
	for _, item := range product.GetOrders() {
		var orders uuid.UUID
		if err := (&orders).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddOrderIDs(orders)
	}
	for _, item := range product.GetShoppingCartOwners() {
		var shoppingcartowners uuid.UUID
		if err := (&shoppingcartowners).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddShoppingCartOwnerIDs(shoppingcartowners)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoProduct(res)
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

// Get implements ProductServiceServer.Get
func (svc *ProductService) Get(ctx context.Context, req *GetProductRequest) (*Product, error) {
	var (
		err error
		get *ent.Product
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetProductRequest_VIEW_UNSPECIFIED, GetProductRequest_BASIC:
		get, err = svc.client.Product.Get(ctx, id)
	case GetProductRequest_WITH_EDGE_IDS:
		get, err = svc.client.Product.Query().
			Where(product.ID(id)).
			WithCategories(func(query *ent.CategoryQuery) {
				query.Select(category.FieldID)
			}).
			WithOrders(func(query *ent.OrderQuery) {
				query.Select(order.FieldID)
			}).
			WithShoppingCartOwners(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoProduct(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements ProductServiceServer.Update
func (svc *ProductService) Update(ctx context.Context, req *UpdateProductRequest) (*Product, error) {
	product := req.GetProduct()
	var productID uuid.UUID
	if err := (&productID).UnmarshalBinary(product.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Product.UpdateOneID(productID)
	if product.GetDescription() != nil {
		productDescription := product.GetDescription().GetValue()
		m.SetDescription(productDescription)
	}
	productName := product.GetName()
	m.SetName(productName)
	if product.GetPictureUrl() != nil {
		productPictureURL := product.GetPictureUrl().GetValue()
		m.SetPictureURL(productPictureURL)
	}
	productPrice := float64(product.GetPrice())
	m.SetPrice(productPrice)
	productQuantity := int(product.GetQuantity())
	m.SetQuantity(productQuantity)
	for _, item := range product.GetCategories() {
		var categories uuid.UUID
		if err := (&categories).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddCategoryIDs(categories)
	}
	for _, item := range product.GetOrders() {
		var orders uuid.UUID
		if err := (&orders).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddOrderIDs(orders)
	}
	for _, item := range product.GetShoppingCartOwners() {
		var shoppingcartowners uuid.UUID
		if err := (&shoppingcartowners).UnmarshalBinary(item.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
		}
		m.AddShoppingCartOwnerIDs(shoppingcartowners)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoProduct(res)
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

// Delete implements ProductServiceServer.Delete
func (svc *ProductService) Delete(ctx context.Context, req *DeleteProductRequest) (*emptypb.Empty, error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Product.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements ProductServiceServer.List
func (svc *ProductService) List(ctx context.Context, req *ListProductRequest) (*ListProductResponse, error) {
	var (
		err      error
		entList  []*ent.Product
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Product.Query().
		Order(ent.Desc(product.FieldID)).
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
			Where(product.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListProductRequest_VIEW_UNSPECIFIED, ListProductRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListProductRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithCategories(func(query *ent.CategoryQuery) {
				query.Select(category.FieldID)
			}).
			WithOrders(func(query *ent.OrderQuery) {
				query.Select(order.FieldID)
			}).
			WithShoppingCartOwners(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
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
		var pbList []*Product
		for _, entEntity := range entList {
			pbEntity, err := toProtoProduct(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListProductResponse{
			ProductList:   pbList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}