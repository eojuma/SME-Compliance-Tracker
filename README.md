```markdown
# SME Compliance Tracker

**A complete business management system for small and medium enterprises**

## What This Does

This application helps small business owners:
- Track expenses and income with profit/loss calculations
- Get automatic alerts before tax deadlines
- Never let a business license expire
- Receive weekly market trend updates
- Know immediately if they're making or losing money

## How It Works

```
User signs up → Adds their business data → System tracks everything
                                            ↓
                                    Automatic notifications
                                    when things need attention
```

## Quick Start

```bash
git clone https://github.com/eojuma/SME-Compliance-Tracker.git
cd sme-tracker
```

###  Install dependencies

```bash
go mod init sme-tracker
go get github.com/mattn/go-sqlite3
go get github.com/gorilla/sessions
go get golang.org/x/crypto/bcrypt
```

###  Run the application

```bash
go run main.go
```

## Complete API Reference


### Authentication

| Action | Method | Endpoint | Body |
|--------|--------|----------|------|
| Register | POST | `/api/register` | `{"email":"user@example.com","password":"secret"}` |
| Login | POST | `/api/login` | `{"email":"user@example.com","password":"secret"}` |
| Logout | POST | `/api/logout` | (none) |

### Expenses

| Action | Method | Endpoint | Body |
|--------|--------|----------|------|
| List all | GET | `/api/expenses` | (none) |
| Add one | POST | `/api/expenses` | `{"amount":100,"category":"Food","date":"2026-04-03","note":"Lunch"}` |
| Delete | DELETE | `/api/expenses` | `{"id":1}` |

### Income

| Action | Method | Endpoint | Body |
|--------|--------|----------|------|
| List all | GET | `/api/income` | (none) |
| Add one | POST | `/api/income` | `{"amount":5000,"category":"Sales","date":"2026-04-01","note":"Payment"}` |

### Taxes

| Action | Method | Endpoint | Body |
|--------|--------|----------|------|
| List all | GET | `/api/taxes` | (none) |
| Add one | POST | `/api/taxes` | `{"name":"Sales Tax","frequency":"monthly","due_day":15}` |

**Frequency options:** `monthly`, `quarterly`, `yearly`

### Licenses

| Action | Method | Endpoint | Body |
|--------|--------|----------|------|
| List all | GET | `/api/licenses` | (none) |
| Add one | POST | `/api/licenses` | `{"name":"Business License","expiry_date":"2026-12-31","reminder_days":30}` |

### Dashboard

| Action | Method | Endpoint | Output |
|--------|--------|----------|--------|
| Get stats | GET | `/api/dashboard/stats` | Income, expenses, profit, recommendation |

## Database Schema (Auto-Created)

The system creates these tables automatically when you first run it:

```
users          -- Business owner accounts
transactions   -- All expenses and income
tax_obligations -- Tax filing deadlines
licenses       -- Business license expirations
```

**No database setup required.** SQLite creates the file `sme-tracker.db` automatically.

## Project Structure

```
sme-tracker/
├── main.go              # Complete application (single file)
├── sme-tracker.db       # SQLite database (auto-created)
├── cookies.txt          # Session storage (for curl testing)
└── go.mod / go.sum      # Dependencies
```

## What Makes This Different

### 1. Smart Profit Recommendations
The system doesn't just show numbers. It tells you what to do:
- **Profit > 10%** → "Healthy margin, keep going"
- **Profit < 10%** → "Margin is thin, consider cost cutting"
- **Loss** → "Alert: Review your highest expenses"

### 2. Proactive Notifications (Coming in Phase 2)
- Tax deadlines: 14, 7, 1 day before due
- License expirations: 60, 30, 14, 7 days before
- Weekly market trend emails

### 3. Single Binary Deployment
Everything is compiled into one executable. No dependencies to install on your server.

## Deployment (When Ready)

```bash
# Build for production
GOOS=linux GOARCH=amd64 go build -o sme-tracker main.go

# Run on any Linux server
./sme-tracker
```

## Development Roadmap

### ✅ Phase 1 Complete (Current)
- Full JSON API working
- User authentication
- Expense/income tracking
- Tax and license management
- Profit/loss calculations with smart recommendations

### 🔄 Phase 2 (Next - Days 11-14)
- Add HTML templates
- Build web dashboard
- Create forms for data entry

### 📅 Phase 3 (Days 15-21)
- Email notification system
- Background scheduler
- Weekly market trends

### 🚀 Phase 4 (Days 22-28)
- Polish UI with Tailwind CSS
- Deploy to production
- Testing and bug fixes
