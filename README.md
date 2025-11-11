```
hexa-fiber-gorm/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── factory/
│   │       └── factory.go                 # Module registry & builder
│   │
│   ├── core/
│   │   └── module.go                      # Module interface (RegisterRoutes)
│   │
│   └── modules/
│       └── user/                          # Each feature = self-contained module
│           ├── domain/
│           │   └── user.go                # Pure domain entity
│           │
│           ├── port/                      # "Ports" = interfaces (contracts)
│           │   └── repository.go          # e.g., UserRepository
│           │
│           └── adapter/
│               ├── persistence/
│               │   └── user_repository.go # GORM implementation
│               │
│               └── web/
│                   ├── handler.go         # HTTP handler logic
│                   ├── routes.go          # Route registration only
│                   └── create_user_dto.go # (Optional) Request DTOs
│
│           └── bootstrap.go               # Self-registers module in factory
│
├── pkg/
│   └── db/
│       └── db.go                          # Shared DB connection (GORM)
│
├── .env                                   # Environment variables
├── Makefile
├── go.mod
└── go.sum
```

- ✅ **internal/modules/** → Each feature (user, auth, session) is a self-contained module.
- ✅ **port/ (not ports/)** → Clean singular name for interface package (common Go style).
- ✅ **adapter/web/handler.go** → Only logic, no routing.
- ✅ **adapter/web/routes.go** → Only path-to-handler mapping.
- ✅ **bootstrap.go** → Uses init() to auto-register the module in the factory.
- ✅ **pkg/db/** → Shared infrastructure (database), used by any module.
- ✅ **internal/app/factory/** → Central place to build all modules with dependencies.