from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse, PlainTextResponse
from fastapi.templating import Jinja2Templates
import os

app = FastAPI()
templates = Jinja2Templates(directory="templates")
backend_id = os.getenv("BACKEND_ID", "unknown")

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

@app.get("/health")
def health_check_root():
    return {"status": "ok"}

# Print registered routes for debugging
for route in app.routes:
    print(route.path, route.methods)