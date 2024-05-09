/your-project-name
├── cmd
│ ├── web4
│ │ └── main.go # Entry point for the web application
├── internal
│ ├── cms
│ │ ├── handler.go # CMS handlers
│ │ └── service.go # CMS business logic
│ ├── blog
│ │ ├── handler.go # Blog page handlers
│ │ └── service.go # Blog business logic
│ └── user
│ ├── handler.go # User management handlers
│ └── service.go # User management business logic
├── pkg
│ ├── api
│ │ └── api.go # API handlers
│ └── utils
│ └── utils.go # Utility functions
├── ui
│ ├── html
│ │ ├── layout.html # Base layout
│ │ └── index.html # Home page template
│ └── static
│ ├── css
│ │ └── app.css # Compiled Tailwind CSS
│ └── js
│ └── script.js # JavaScript files
├── configs
│ └── config.toml # Configuration file
├── scripts
│ └── deploy.sh # Deployment script
├── migrations
│ └── 001_init_schema.sql # SQL database migrations
├── tests
│ └── api
│ └── api_test.go # API tests
├── vendor # Vendor dependencies
├── go.mod # Go module dependencies
├── go.sum # Go module sums
└── .gitignore # Specifies intentionally untracked files to ignore
