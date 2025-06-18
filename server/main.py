from fastapi import FastAPI, Request, HTTPException
from fastapi.responses import HTMLResponse, PlainTextResponse
from fastapi.templating import Jinja2Templates
import os
import asyncpg

app = FastAPI()
templates = Jinja2Templates(directory="templates")
backend_id = os.getenv("BACKEND_ID", "unknown")
DATABASE_URL = os.getenv("DATABASE_URL", "postgresql://postgresUser:password@localhost/testdb")
pool = None

@app.on_event("startup")
async def startup():
    global pool
    pool = await asyncpg.create_pool(DATABASE_URL)

@app.on_event("shutdown")
async def shutdown():
    global pool
    if pool:
        await pool.close()

@app.get("/")
def read_root_json():
    return f"Hello from FastAPI backend {backend_id}"

@app.get("/text", response_class=PlainTextResponse)
def read_text():
    return f"Hello from FastAPI backend {backend_id}"

@app.get("/page", response_class=HTMLResponse)
async def read_page(request: Request):
    return templates.TemplateResponse("index.html", {"request": request, "backend_id": backend_id})

@app.get("/text/health")
def health_check_text():
    return {"status": "ok"}

@app.get("/page/health")
def health_check_page():
    return {"status": "ok"}

@app.get("/users/health")
def health_check_page():
    return {"status": "ok"}

@app.get("/health")
def health_check_root():
    return {"status": "ok"}

@app.get("/users")
async def get_users():
    try:
        async with pool.acquire() as connection:
            rows = await connection.fetch("SELECT * FROM users")
            return [{"id": row["id"], "name": row["name"], "email": row["email"]} for row in rows]
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

# Print registered routes for debugging
for route in app.routes:
    print(route.path, route.methods)