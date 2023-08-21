package biz

import (
	"context"

	v1 "kratos-realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// 业务逻辑的组装层
// 有操作数据库的方法

// biz 领域模型层 调data的接口

type ArticleRepo interface{}

type CommentRepo interface{}

type TagRepo interface{}

// SocialUseCase is a realworld usecase.
type SocialUseCase struct {
	ar  ArticleRepo
	cr  CommentRepo
	tr  TagRepo
	log *log.Helper
}

func NewSocialUseCase(ar ArticleRepo, cr CommentRepo, tr TagRepo, log *log.Helper) *SocialUseCase {
	return &SocialUseCase{ar: ar, cr: cr, tr: tr, log: log}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *SocialUseCase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)

	return nil, nil
}

// 类似于领域驱动设计（DDD）中的 "repo"（repository），它们的作用是在业务逻辑层（"biz"）和数据层之间建立抽象接口。这种抽象接口有助于将业务逻辑与数据存储解耦，并遵循依赖倒置原则，

// 依赖倒置原则（Dependency Inversion Principle，简称DIP）是面向对象编程中的一个重要原则，它是SOLID原则之一，用于指导如何组织代码以减少耦合性和提高灵活性。依赖倒置原则有以下核心思想：
//
//1. **高层模块不应该依赖于底层模块，两者都应该依赖于抽象。**
//2. **抽象不应该依赖于具体实现，具体实现应该依赖于抽象。**
//
//这意味着：
//
//- 代码中的高层模块（通常是应用程序的业务逻辑）不应该直接依赖于低层模块（通常是数据存储、数据库等）的具体实现细节。相反，它们应该依赖于抽象接口或抽象类。
//- 抽象接口或抽象类不应该依赖于它们的具体实现类。具体实现类应该依赖于抽象。
//
//这个原则的目标是实现松散耦合（Loose Coupling），从而使代码更容易维护、扩展和测试。它也支持多态性，使得可以在不改变高层模块的情况下轻松替换底层模块的具体实现。
//
//依赖倒置原则的一些关键概念和技术包括：
//
//- **抽象类和接口**：通过定义抽象的基类或接口，高层模块可以依赖于这些抽象，而不是具体实现。
//
//- **依赖注入**：通过将依赖项传递给高层模块，而不是在高层模块内部创建依赖，可以实现依赖倒置。这通常通过构造函数注入、方法注入或属性注入来实现。
//
//- **反转控制容器**：一些框架和容器（如Spring框架）提供了反转控制容器，可以自动管理依赖项的创建和注入。
//
//总之，依赖倒置原则是面向对象编程中的一个基本原则，有助于创建更灵活、可维护和可扩展的代码。它提倡将高层模块与低层模块之间的依赖关系反转，从而提高代码的可测试性和可重用性。
