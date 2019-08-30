from fastapi import FastAPI
from starlette.middleware.cors import CORSMiddleware

from .routers import auth

app = FastAPI(title="ovrc")

# Cors, allow connections from the frontend and anywhere.
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


app.include_router(auth.router, prefix="/auth", tags=["auth"])
