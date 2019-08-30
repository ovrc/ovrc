from fastapi import Depends, FastAPI, Header, HTTPException
from fastapi.security import OAuth2PasswordBearer

from .routers import users, auth

app = FastAPI(
    title="ovrc"
)

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/auth/token")


async def get_token_header(x_token: str = Header(...)):
    if x_token != "fake-super-secret-token":
        raise HTTPException(status_code=400, detail="X-Token header invalid")


app.include_router(
    users.router,
    prefix="/users",
    tags=["users"],
    dependencies=[Depends(oauth2_scheme)],
    responses={404: {"description": "Not found"}},
)

app.include_router(
    auth.router,
    prefix="/auth",
    tags=["auth"]
)