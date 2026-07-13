   # ✈️ Flight Voucher

   A full-stack flight voucher assignment app with **Go** backend (Gin + SQLite) and **React + TypeScript** frontend (Vite + Tailwind CSS).

   ---

   ## 📁 Project Structure

   ```
   flight-voucher/
   ├── backend/          # Go REST API (Gin framework)
   │   ├── config/
   │   ├── internal/
   │   ├── config.yaml
   │   ├── main.go
   │   └── Dockerfile
   ├── frontend/         # React + TypeScript (Vite)
   │   ├── src/
   │   ├── .env
   │   ├── package.json
   │   └── Dockerfile
   └── docker-compose.yml
   ```

   ---

   ## 🧰 Prerequisites

   ### Backend
   | Tool | Version | Download |
   |------|---------|----------|
   | **Go** | 1.23+ | [go.dev/dl](https://go.dev/dl/) |

   ### Frontend
   | Tool | Version | Download |
   |------|---------|----------|
   | **Node.js** | 18+ (LTS recommended) | [nodejs.org](https://nodejs.org/) |

   ---

   ## 📦 Installing Dependencies

   ### Backend

   ```bash
   cd backend
   go mod download
   ```

   ### Frontend

   ```bash
   cd frontend
   npm install
   ```

   ---

   ## 🚀 Running the Application

   ### Backend

   Backend runs on **port 8080** by default (configurable in [`backend/config.yaml`](./backend/config.yaml)).

   ```bash
   cd backend
   go run main.go
   ```

   API will be available at `http://localhost:8080`.

   ### Frontend

   Frontend runs on **port 5173** by default.

   1. Ensure the `.env` file is present in the `frontend/` directory:
      ```env
      VITE_API_BASE_URL=http://localhost:8080
      ```

   2. Start the dev server:
      ```bash
      cd frontend
      npm run dev
      ```

   App will be available at `http://localhost:5173`.

   ---

   ## 🐳 Docker Instructions

   Docker Compose orchestrates both services with a single command.

   Make sure **`docker`** and **`docker compose`** available.

   ### Build & Start All Services

   Run this from the **project root** (where `docker-compose.yml` located):

   ```bash
   docker compose up --build
   ```

   | Service  | URL                     |
   |----------|-------------------------|
   | Backend  | http://localhost:8080   |
   | Frontend | http://localhost:5173   |

   ### Stop All Services

   ```bash
   docker compose down
   ```

   ### Stop and Remove Volumes (including SQLite data)

   ```bash
   docker compose down -v
   ```

   ### Run in Detached Mode (Background)

   ```bash
   docker compose up --build -d
   ```

   > **Note:** SQLite data is persisted in a named Docker volume (`sqlite_data`) so it survives container restarts.

   ---

   ## 🔌 API Endpoints

   | Method | Path         | Description              |
   |--------|--------------|--------------------------|
   | POST   | `/api/check` | Check flight voucher     |
   | POST   | `/api/generate` | Generate random voucher seats |

   ---

   ## 🛠️ Tech Stack

   | Layer    | Technology                         |
   |----------|------------------------------------|
   | Backend  | Go, Gin, GORM, SQLite, Viper       |
   | Frontend | React, TypeScript, Vite, Tailwind CSS |
   | Database | SQLite   |
   | DevOps   | Docker, Docker Compose             |
