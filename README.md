# Rafi BE Assessment Service

This is a backend service created as part of the assessment task. It is a simple Go service with the following features:
- Customer Registration
- Account Registration
- Transaction Management (Deposit, Withdrawal)
- Logging and Error Handling
- PostgreSQL Database

## Prerequisites

- Go 1.20 or higher
- Docker (for running PostgreSQL container)
- PostgreSQL database
- Go modules enabled

## Setup and Installation

### Clone the repository
First, clone the repository to your local machine.

```bash
git clone https://github.com/rafialg11/rafi_BE_assesment.git
cd rafi_BE_assesment
```

### Setting up the database

To run the service, you need a PostgreSQL database. The easiest way is to use Docker to run PostgreSQL locally.

**With Docker:**

```bash
docker-compose up -d
```

This will start the PostgreSQL container and create the necessary database.

### Set up `.env` file

Create a `.env` file in the root directory of the project with the following content:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=account_db
DB_PORT=5432
```

### Build the service with Docker

If you want to run the service inside a Docker container, you can use this command:
```
docker-compose up --build
```

## Endpoints

### 1. `POST /api/v1/daftar`

Registers a new customer and creates an associated account.

**Request Body:**

```json
{
    "name": "John Doe",
    "phone": "1234567890",
    "NIK": "1234567890123456"
}
```

**Response:**

```json
{
    "status": 200,
    "message": "SUCCESS",
    "data": {
        "account_number": "1234-2757-000002"
    },
    "error": null
}
```

### 2. `POST /api/v1/tabung`

Deposits an amount into the specified account.

**Request Body:**

```json
{
    "account_number": "1234567890123456",
    "amount": 500
}
```

**Response:**

```json
{
    "status": 200 ,
    "message": "SUCCESS",
    "data": {
        "account_number": "1234567890123456",
        "amount": 500
    },
    "error": null
}
```

### 3. `POST /api/v1/tarik`

Withdraws an amount from the specified account.

**Request Body:**

```json
{
    "account_number": "1234567890123456",
    "amount": 100
}
```

**Response:**

```json
{
    "status": 200 ,
    "message": "SUCCESS",
    "data": {
        "account_number": "1234567890123456",
        "amount": 400
    },
    "error": null
}
```
### 4. `GET /api/v1/account/saldo/{no_rekening}`

Withdraws an amount from the specified account.

**Response:**

```json
{
    "status": 200,
    "message": "SUCCESS",
    "data": {
        "account_number": "1234-0513-000001",
        "amount": 1000000
    },
    "error": null
}
```
## Database Scheme
![schema](https://github.com/user-attachments/assets/f3478f3d-587c-44e9-992e-eda73941bf6e)


## Logging

The service utilizes **Logrus** for structured logging. Logs are stored in the `log/` directory. The log filename is based on the current date (e.g., `2025-02-02.log`).

Log entries are automatically rotated every day using the cron job defined in `InitLogger`.

### Example log format:

```json
{
    "timestamp": "2025-02-02T17:30:00Z",
    "level": "info",
    "message": "Transaction completed successfully",
    "transaction_id": "12345",
    "customer_id": "1",
    "amount": "500"
}
```
