from fastapi import FastAPI

from .routers import auth

app = FastAPI(title="ovrc")


app.include_router(auth.router, prefix="/auth", tags=["auth"])
