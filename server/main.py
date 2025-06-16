from fastapi import FastAPI, Request   
from fastapi.responses import HTMLResponse, PlainTextResponse
from fastapi.templating import Jinja2Templates
import os

app = FastAPI()

templates = Jinja2Templates(directory="templates")

backend_id = os.getenv("BACKEND_ID", "unknown")

@app.get("/")
def read_root():
    return {"message": f"Hello from FastAPI backend {backend_id}"}

@app.get("/text", response_class=PlainTextResponse)
def read_root():
    return f"Hello from FastAPI backend {backend_id}"

@app.get("/page", response_class=HTMLResponse)
async def read_page(request: Request):
    return templates.TemplateResponse("index.html", {"request": request, "backend_id": backend_id})

@app.get("/health")
def health_check():
    return {"status": "ok"}