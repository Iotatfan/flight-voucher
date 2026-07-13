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

   ## 🗄️ Architecture & System Rules

   To meet the core assessment requirements, the application implements the following business logic rules:

   ### 1. Unique Voucher Restrictions
   * **Frontend Guardrail:** Clicking "Generate Vouchers" first triggers a request to `POST /api/check`. If vouchers exist, generation is blocked, and an error message is displayed.
   * **Backend Guardrail:** Inside `GenerateRandomSeats`, the service performs a second `CheckFlight` query **before** writing any data. If a voucher for the same `(flight_number, flight_date)` pair already exists, it immediately returns `ErrVoucherAlreadyExists` and the handler responds with **HTTP 409 Conflict** — regardless of what the frontend sent. An invalid aircraft type returns `ErrInvalidAircraftType` (**HTTP 400**), and insufficient seats return `ErrFewerThan3SeatsAvailable` (**HTTP 500**).
   * **Database Guardrail:** To prevent race conditions and ensure query safety, a **Composite Unique Key** constraint is placed on `(flight_number, flight_date)` within the SQLite schema. Duplicate generations for the same flight date will fail safely.

   ### 2. Aircraft Seating Layouts
   The backend verifies and randomizes exactly 3 unique, non-repeating seats using the structural rules defined by the aircraft type:
   * **ATR:** Rows 1–18, Seats: `A, C, D, F` (Excludes `B` and `E`).
   * **Airbus 320 / Boeing 737 Max:** Rows 1–32, Seats: `A, B, C, D, E, F`.

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
