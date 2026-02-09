
---

# SME-Compliance-Tracker

**SME-Compliance-Tracker** is a web-based platform designed to help small and medium-sized enterprises (SMEs) in Kenya manage their **tax compliance**, **business licenses**, and **financial reporting**. The app automates **tax reminders**, tracks **expenses** and **income**, stores important **business documents**, and generates **financial reports**, making it easier for business owners to stay on top of their obligations and focus on growing their businesses.

---

## Features

* **Tax Compliance Management**: Keep track of tax payments, due dates, and deadlines.
* **Business Licenses**: Store, track, and manage business licenses and their expiration dates.
* **Financial Tracking**: Monitor income, expenses, and generate basic financial reports.
* **Automated Reminders**: Get reminders about tax deadlines and license renewals.
* **User Authentication**: Secure login for different roles (e.g., Admin, SME Owner).

---

## Technologies

This project uses the following technologies:

* **Frontend**: React.js
* **Backend**: Go (Golang)
* **Database**: PostgreSQL
* **Containerization**: Docker, Docker Compose
* **Version Control**: Git & GitHub

---

## Prerequisites

Before you start, make sure you have the following installed:

* **Docker**: For containerization.
* **Docker Compose**: To easily manage multi-container applications.
* **Node.js** (for frontend): Ensure you have `npm` or `yarn` installed.
* **Go** (for backend): Ensure you have Go installed on your system.
* **Git**: For version control and repository management.

---

## Project Structure

Here’s a breakdown of the project structure:

```
SME-Compliance-Tracker/
│
├── backend/                  # Go-based backend
│   ├── Dockerfile            # Dockerfile for the backend
│   ├── go.mod                # Go modules file
│   ├── go.sum                # Go checksum file
│   ├── main.go               # Main Go application file (API)
│   ├── handlers/             # Contains the route handler files
│   ├── models/               # Defines database models (e.g., User, Compliance)
│   └── utils/                # Utility functions (e.g., for validation, date handling)
│
├── frontend/                 # React-based frontend
│   ├── Dockerfile            # Dockerfile for the frontend
│   ├── package.json          # npm package file for managing dependencies
│   ├── public/               # Public files (index.html, images, etc.)
│   ├── src/                  # Source code for React components
│   └── .env                  # Environment variables for frontend configuration
│
├── database/                 # Database setup
│   ├── init.sql              # SQL file for initializing the database schema
│   └── docker-compose.yml    # Docker Compose configuration for the database and services
│
├── .gitignore                # Git ignore file to exclude unnecessary files
├── .env                      # Main environment file for sensitive data
├── README.md                 # Project overview and documentation
└── docker-compose.yml        # Docker Compose configuration for the entire app
```

---

## Getting Started

Follow these instructions to set up the project locally on your machine.

### 1. Clone the Repository

Start by cloning this repository to your local machine:

```bash
git clone https://github.com/eojuma/SME-Compliance-Tracker.git
cd SME-Compliance-Tracker
```

### 2. Build and Run the Application with Docker Compose

This project uses **Docker Compose** to handle the multiple services (frontend, backend, and database). To start the project, run the following commands:

1. **Build and start the containers:**

   ```bash
   docker-compose up --build
   ```

2. **Access the app:**

   * The **frontend** will be accessible at `http://localhost:3000`.
   * The **backend API** can be accessed at `http://localhost:8080`.
   * **PostgreSQL database** is accessible at `localhost:5432`.

3. **Shut down the containers**:
   To stop the containers after you're done:

   ```bash
   docker-compose down
   ```

---

## Configuration

### Environment Variables

Create a `.env` file in the root of your project and add the following configurations:

```bash
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=sme_compliance
PORT=8080
```

Make sure to **not** commit this file to Git as it contains sensitive information (i.e., database credentials). Add it to `.gitignore` to ensure it's excluded from version control.

### Database Schema

The database schema is initialized via `init.sql` inside the `database/` folder. When the **PostgreSQL** container is first launched, it will run this SQL file to create the necessary tables:

```sql
CREATE TABLE compliance (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  tax_status VARCHAR(50),
  due_date DATE
);

CREATE TABLE business_licenses (
  id SERIAL PRIMARY KEY,
  business_name VARCHAR(100),
  license_number VARCHAR(50),
  expiration_date DATE
);
```

---

## Contributing

We welcome contributions! To contribute to the **SME-Compliance-Tracker** project, follow these steps:

1. **Fork the repository** to your own GitHub account.
2. **Clone** the forked repository to your local machine.
3. Create a new branch for your changes:

   ```bash
   git checkout -b feature-name
   ```
4. **Make your changes** and commit them:

   ```bash
   git commit -m "Description of changes"
   ```
5. **Push** your changes to your fork:

   ```bash
   git push origin feature-name
   ```
6. **Open a pull request** to the main repository.

We will review your changes and, if everything looks good, we will merge them into the main project.

---

## License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---

## Acknowledgements

* [Docker](https://www.docker.com/) for containerization.
* [Go](https://golang.org/) for backend development.
* [React](https://reactjs.org/) for frontend development.
* [PostgreSQL](https://www.postgresql.org/) for database management.
* [GitHub](https://github.com/) for hosting the project.

---

## Troubleshooting

If you run into issues while setting up or running the project, try the following:

1. **Docker not starting?**

   * Ensure Docker and Docker Compose are installed and running.
   * Restart Docker and try again.

2. **Error: Port already in use?**

   * Check if port `3000` or `8080` is being used by another application. You can change the ports in `docker-compose.yml`.

3. **Database connection issues?**

   * Double-check your `.env` file and make sure your database credentials are correct.
   * Verify that the `db` service in Docker Compose is running correctly.

---

### Conclusion

This README provides all the necessary instructions and details for setting up, developing, and contributing to the **SME-Compliance-Tracker** project. The app is designed to automate tax and compliance tracking for SMEs in Kenya, helping them stay on top of their business obligations and simplify reporting.

Feel free to contribute, and good luck building the project!

---
