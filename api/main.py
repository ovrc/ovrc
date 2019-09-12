from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware
from api.config import CORS_ALLOW_ORIGINS

from .routers import auth, user

app = FastAPI(title="ovrc")

app.add_middleware(
    CORSMiddleware, allow_origins=CORS_ALLOW_ORIGINS, allow_credentials=True
)

app.include_router(auth.router, prefix="/auth", tags=["auth"])
app.include_router(user.router, prefix="/users", tags=["users"])
