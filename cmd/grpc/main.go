package main

import (
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent"
	"github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/ent/proto/entpb"
	pkg_grpc "github.com/Pirika-Pirilala-Popolina-Peperuto/database-system/pkg/grpc"
	"google.golang.org/grpc"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			ent.NewPostgresClientWithLC,
			fx.Annotate(
				pkg_grpc.NewServerWithLC,
				fx.As(new(grpc.ServiceRegistrar)),
			),
			fx.Annotate(
				entpb.NewCategoryService,
				fx.As(new(entpb.CategoryServiceServer)),
			),
			fx.Annotate(
				entpb.NewDiscountService,
				fx.As(new(entpb.DiscountServiceServer)),
			),
			fx.Annotate(
				entpb.NewOrderService,
				fx.As(new(entpb.OrderServiceServer)),
			),
			fx.Annotate(
				entpb.NewProductService,
				fx.As(new(entpb.ProductServiceServer)),
			),
			fx.Annotate(
				entpb.NewUserService,
				fx.As(new(entpb.UserServiceServer)),
			),
		),
		fx.Invoke(
			entpb.RegisterCategoryServiceServer,
			entpb.RegisterDiscountServiceServer,
			entpb.RegisterOrderServiceServer,
			entpb.RegisterProductServiceServer,
			entpb.RegisterUserServiceServer,
		),
	).Run()
}
