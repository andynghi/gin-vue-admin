## server project structure

```shell
├── api
│   └── v1
├── config
├── core
├── docs
├── global
├── initialize
│   └── internal
├── middleware
├── model
│   ├── request
│   └── response
├── packfile
├── resource
│   ├── excel
│   ├── page
│   └── template
├── router
├── service
├── source
└── utils
    ├── timer
    └── upload
```

| Folder | Description | Description |
| ------------ | ----------------------- | --------------------------- |
| `api` | api layer | api layer |
| `--v1` | v1 version interface | v1 version interface |
| `config` | Configuration package | Configuration structure corresponding to config.yaml |
| `core` | Core file | Initialization of core components (zap, viper, server) |
| `docs` | swagger document directory | swagger document directory |
| `global` | global object | global object |
| `initialize` | Initialization | Initialization of router, redis, gorm, validator, timer |
| `--internal` | Initialize internal functions | gorm's longger customization, the functions in this folder can only be called by the `initialize` layer |
| `middleware` | Middleware layer | Used to store `gin` middleware code |
| `model` | Model layer | Model corresponding data table |
| `--request` | Input parameter structure | Receive data sent from the front end to the back end. |
| `--response` | Parameter structure | Data structure returned to the front end |
| `packfile` | Static file packaging | Static file packaging |
| `resource` | Static resource folder | Responsible for storing static files |
| `--excel` | Excel import and export default path | Excel import and export default path |
| `--page` | Form builder | Form builder packaged dist |
| `--template` | Template | Template folder, which stores the templates of the code generator |
| `router` | routing layer | routing layer |
| `service` | service layer | stores business logic issues |
| `source` | source layer | function to store initialization data |
| `utils` | toolkit | tool function encapsulation |
| `--timer` | timer | timer interface encapsulation |
| `--upload` | oss | oss interface encapsulation |

