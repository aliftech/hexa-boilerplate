```
hexa-fiber-gorm/
├── cmd/
│   └── app/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   └── factory/
│   │       └── factory.go
│   │
│   ├── core/
│   │   └── module.go
│   │
│   └── modules/
│       └── user/
│           ├── domain/
│           │   └── user.go                # Entity
│           │
│           ├── port/                      # Interfaces (Ports)
│           │   ├── repository.go          # e.g., UserRepository
│           │   └── usecase.go             # e.g., UserUsecase
│           │
│           ├── usecase/                   # ← BUSINESS LOGIC HERE
│           │   └── user_usecase.go        # Implements UserUsecase port
│           │
│           └── adapter/
│               ├── persistence/
│               │   └── user_repository.go # Implements UserRepository
│               │
│               └── web/
│                   ├── handler.go         # ONLY calls usecase
│                   ├── routes.go          # Route registration
│                   └── dto/               # (Optional) Request/Response DTOs
│                       ├── create_user_request.go
│                       └── user_response.go
│
│           └── bootstrap.go               # Wires deps & registers module
│
├── pkg/
│   └── db/
│       └── db.go
│
├── .env
├── Makefile
└── go.mod
```

- ✅ **internal/modules/** → Each feature (user, auth, session) is a self-contained module.
- ✅ **port/ (not ports/)** → Clean singular name for interface package (common Go style).
- ✅ **adapter/web/handler.go** → Only logic, no routing.
- ✅ **adapter/web/routes.go** → Only path-to-handler mapping.
- ✅ **bootstrap.go** → Uses init() to auto-register the module in the factory.
- ✅ **pkg/db/** → Shared infrastructure (database), used by any module.
- ✅ **internal/app/factory/** → Central place to build all modules with dependencies.