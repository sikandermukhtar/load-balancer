from fastapi import FastAPI
import os

app = FastAPI()

backend_id = os.getenv("BACKEND_ID", "unknown")

@app.get("/")
def read_root():
    return {"message": f"Hello from FastAPI backend {backend_id}"}

@app.get("/health")
def health_check():
    return {"status": "ok"}