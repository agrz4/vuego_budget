BudgetApp: Budget and Expense Management Application
BudgetApp is a simple full-stack web application designed to help users manage their budgets and track their expenses. This application provides user authentication features, a summary dashboard, CRUD (Create, Read, Update, Delete) capabilities for budgets and expenses, and simple visualizations.

🎯 Features
User Authentication:

Registration: Create a new user account.

Login: Log in to an existing account using username and password.

Uses JWT (JSON Web Token) for session authentication.

User Dashboard:

Displays a summary of total budget, total expenses, and remaining budget.

Visualization of expenses per category in a pie chart.

Budget Management:

Add new budgets with category, amount, and description.

View a list of all created budgets.

Edit existing budget details.

Delete budgets.

Expense Management:

Add new expenses with category, amount, description, and date.

View a list of all expenses.

Filter & Search: Filter expenses by category, date range, and search by description or category.

Edit existing expense details.

Delete expenses.

Form Validation: Input validation on both frontend and backend to ensure data integrity.

Responsive: Simple and responsive design, suitable for both desktop and mobile views.

🧱 Stack and Technologies
Frontend
Vue.js 3: Progressive JavaScript framework for building user interfaces.

Composition API: For more organized and reusable component logic.

Vue Router: For client-side routing.

Axios: Promise-based HTTP client for making requests to the backend API.

TailwindCSS: A utility-first CSS framework for rapid and responsive styling.

Chart.js: For data visualization (pie chart).

Lodash: For utilities like debounce on search inputs.

Backend
REST API Development:

Developed a robust RESTful API using Golang with the Gin Web Framework, facilitating seamless communication between client applications and server-side resources.

Implemented user login and registration systems with secure authentication using JWT (JSON Web Token).

Ensured efficient data storage, retrieval, and management by utilizing PostgreSQL as the primary database, integrated via GORM (ORM).

Database Migration: Automated schema management with GORM's AutoMigrate, defining user, budget, and expense tables with relevant columns.

Implemented secure password hashing using golang.org/x/crypto/bcrypt to protect user data and ensure the security of sensitive information.

📁 Project Folder Structure
.
├── backend/                  # Golang Project
│   ├── main.go               # Backend application entry point
│   ├── models/               # Database model definitions (User, Budget, Expense)
│   ├── controllers/          # Handler logic for APIs (Auth, Budget, Expense)
│   ├── middleware/           # Custom middleware (JWT authentication)
│   ├── database/             # Database connection configuration
│   ├── utils/                # Utilities (JWT generation/parsing)
│   ├── go.mod                # Go modules and dependencies
│   └── go.sum
├── frontend/                 # Vue.js Project
│   ├── public/               # Static assets
│   ├── src/                  # Vue.js source code
│   │   ├── assets/           # Other static assets (e.g., images)
│   │   ├── components/       # Reusable Vue components (e.g., Chart.vue)
│   │   ├── router/           # Vue Router configuration
│   │   ├── services/         # API services for backend interaction
│   │   ├── views/            # Application views/pages
│   │   │   ├── Auth/         # Login and Register pages
│   │   │   ├── Budget/       # Budget list and form pages
│   │   │   ├── Expense/      # Expense list and form pages
│   │   │   └── Dashboard.vue # Dashboard page
│   │   ├── App.vue           # Main application component
│   │   └── main.js           # Vue.js application entry point
│   ├── index.html            # Main HTML file
│   ├── package.json          # Node.js dependencies
│   ├── postcss.config.js     # PostCSS configuration (for TailwindCSS)
│   ├── tailwind.config.js    # TailwindCSS configuration
│   └── vite.config.js        # Vite configuration
└── .env                      # Environment variables for configuration



🚀 How to Run the Application
To run the application, you will execute the frontend and backend components manually.

1. Manual Prerequisites
Install Node.js and npm/Yarn:

Download from https://nodejs.org/ (LTS version recommended).

Verify installation: node -v and npm -v (or yarn -v).

Install Golang:

Download from https://golang.org/dl/.

Follow the installation instructions for your operating system.

Verify installation: go version.

PostgreSQL Installation and Configuration:

Install PostgreSQL on your operating system (Windows, macOS, Linux). Installation guides can be found on the official PostgreSQL website: https://www.postgresql.org/download/.

After installation, create a new database user and database. Example commands (for Linux/macOS, or use pgAdmin on Windows):

sudo -u postgres psql
CREATE USER user_budget WITH PASSWORD 'password_budget';
CREATE DATABASE budget_db OWNER user_budget;
GRANT ALL PRIVILEGES ON DATABASE budget_db TO user_budget;
\q

Ensure the credentials (user_budget, password_budget, budget_db) match what you will use in your .env file.

2. .env File Configuration
Update the .env file in your project root folder.

# PostgreSQL Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=user_budget
DB_PASSWORD=password_budget
DB_NAME=budget_db

# JWT Configuration
JWT_SECRET=supersecretkeyyangsangatpanjangdanaman

# Golang Backend Configuration
PORT=8080
GIN_MODE=debug
FRONTEND_DOMAIN=http://localhost:5173

# Vue.js Frontend Configuration (accessed by Vite)
VITE_API_BASE_URL=http://localhost:8080/api

3. Running the Golang Backend
Navigate to the Backend Directory:

cd backend

Download Golang Dependencies:

go mod tidy

Run the Backend Application:

go run main.go

The backend server will run on http://localhost:8080.

4. Running the Vue.js Frontend
Navigate to the Frontend Directory:
Open a new terminal and navigate to the frontend/ folder:

cd frontend

Install Frontend Dependencies:

npm install
npm install chart.js lodash # Make sure these are also installed

or

yarn
yarn add chart.js lodash

Run the Frontend Development Server:

npm run dev

or

yarn dev

The frontend application will run on http://localhost:5173.

⚙️ API Endpoints (Golang Backend)
All API endpoints start with /api.

Authentication
POST /api/register - Register a new user.

Request Body: { "username": "string", "password": "string" }

POST /api/login - Log in a user and get a JWT token.

Request Body: { "username": "string", "password": "string" }

Response: { "message": "Login successful", "token": "jwt_token_string" }

Budgets (Requires JWT Authentication)
POST /api/budgets - Create a new budget.

Request Body: { "category": "string", "amount": float64, "description": "string" }

GET /api/budgets - Get all user budgets.

GET /api/budgets/:id - Get a budget by ID.

PUT /api/budgets/:id - Update a budget by ID.

Request Body: { "category": "string", "amount": float64, "description": "string" }

DELETE /api/budgets/:id - Delete a budget by ID.

Expenses (Requires JWT Authentication)
POST /api/expenses - Create a new expense.

Request Body: { "category": "string", "amount": float64, "description": "string", "date": "ISO_8601_date_string" }

GET /api/expenses - Get all user expenses.

Query Params: category, start_date (YYYY-MM-DD), end_date (YYYY-MM-DD), search (description/category).

GET /api/expenses/:id - Get an expense by ID.

PUT /api/expenses/:id - Update an expense by ID.

Request Body: { "category": "string", "amount": float64, "description": "string", "date": "ISO_8601_date_string" }

DELETE /api/expenses/:id - Delete an expense by ID.

Dashboard (Requires JWT Authentication)
GET /api/dashboard - Get a summary of total budget, total expenses, remaining budget, and expenses by category.

⚠️ Important Notes
JWT Security: Ensure that JWT_SECRET in your .env file is a very strong, long, and random string. Never expose it in a public repository.

Database Migration: The Golang backend will automatically perform database migrations (create tables) when it runs for the first time if the tables do not exist.

CORS: The CORS configuration in the backend currently only allows http://localhost:5173. Adjust it if you are running the frontend on a different domain/port.

Validation: Basic validation has been implemented. For production applications, consider more advanced validation and more detailed error handling.

Error Handling: This application has basic error handling that will display error messages to the user. For production, you might want to implement a more comprehensive logging system.

Enjoy using BudgetApp!