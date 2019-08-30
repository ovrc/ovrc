from fastapi import APIRouter

from datetime import datetime, timedelta

import jwt
from fastapi import APIRouter, Depends, HTTPException
from fastapi.security import OAuth2PasswordRequestForm
from jwt import PyJWTError
from pydantic import BaseModel
from starlette.responses import Response
from starlette.status import HTTP_401_UNAUTHORIZED

from api import OAuth2PasswordBearerWithCookie
from api.config import JWS_SECRET_KEY, JWS_ALGORITHM, JWS_ACCESS_TOKEN_EXPIRE_MINUTES

router = APIRouter()


class User(BaseModel):
    username: str
    email: str = None
    full_name: str = None
    disabled: bool = None


oauth2_scheme = OAuth2PasswordBearerWithCookie(tokenUrl="/auth/token")


def create_access_token(*, data: dict, expires_delta: timedelta = None):
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.utcnow() + expires_delta
    else:
        expire = datetime.utcnow() + timedelta(minutes=15)
    to_encode.update({"exp": expire})

    encoded_jwt = jwt.encode(to_encode, JWS_SECRET_KEY, algorithm=JWS_ALGORITHM)

    return encoded_jwt


async def get_current_user(token: str = Depends(oauth2_scheme)):
    credentials_exception = HTTPException(
        status_code=HTTP_401_UNAUTHORIZED,
        detail="Could not validate credentials",
        headers={"WWW-Authenticate": "Bearer"},
    )

    try:
        payload = jwt.decode(token, JWS_SECRET_KEY, algorithms=[JWS_ALGORITHM])
        username: str = payload.get("sub")
        if username is None:
            raise credentials_exception
        # token_data = TokenData(username=username)
    except PyJWTError as e:
        print(e)
        raise credentials_exception
    # user = get_user(fake_users_db, username=token_data.username)
    user = User(username="joao", email="joao@joao.com")
    if user is None:
        raise credentials_exception
    return user


async def get_current_active_user(current_user: User = Depends(get_current_user)):
    if current_user.disabled:
        raise HTTPException(status_code=400, detail="Inactive user")
    return current_user


@router.post("/token", tags=["auth"])
async def auth_token(
    response: Response, form_data: OAuth2PasswordRequestForm = Depends()
):
    # auth_user(form_data.username, form_data.password)

    username = form_data.username
    password = form_data.password

    print(username, password)

    # Assume user is authenticated via username and password...
    # Add database checks later.

    access_token_expires = timedelta(minutes=JWS_ACCESS_TOKEN_EXPIRE_MINUTES)

    access_token = create_access_token(
        data={"sub": username}, expires_delta=access_token_expires
    ).decode("utf-8")

    response.set_cookie(
        key="access_token", value=f"Bearer {access_token}", httponly=True, secure=True
    )

    return


@router.get("/user", response_model=User)
async def auth_read_user(current_user: User = Depends(get_current_active_user)):
    return current_user
