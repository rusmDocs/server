import datetime
import os

from auth.dto import auth_pb2_grpc
from auth.dto import auth_pb2
from datetime import datetime
import jwt
import dotenv

import redis

from auth.utils.payload import generate_token, decode_token


class JWT(auth_pb2_grpc.AuthServiceServicer):
    r = redis.Redis(host="localhost", port=6379, db=0) 

    def CreateTokens(self, request, context):
        if request.id:
            access_token, refresh_token = generate_token(request.id)
            self.r.set(refresh_token, request.id)
            return auth_pb2.JWTTokens(
                access_token=access_token,
                refresh_token=refresh_token
            )

    def CheckTokens(self, request, context):
        access, refresh = request.access_token, request.refresh_token
        try:
            access_payload, refresh_payload = decode_token(access), decode_token(refresh) 
        except jwt.InvalidTokenError:
            return auth_pb2.UserTokens(
                access_token="",
                refresh_token="",
                id="",
                status=1  
            )
        if not self.r.get(refresh): 
            return auth_pb2.UserTokens(
                access_token="",
                refresh_token="",
                id="",
                status=2  
            )
        elif datetime.fromtimestamp(refresh_payload['exp']) < datetime.now(): 
            return auth_pb2.UserTokens(
                access_token="",
                refresh_token="",
                id="",
                status=3  
            )
        elif datetime.fromtimestamp(access_payload['exp']) < datetime.now(): 
            self.r.delete(refresh)
            access, refresh = generate_token(refresh_payload['id'])
            self.r.set(refresh, refresh_payload['id'])
            return auth_pb2.UserTokens(
                access_token=access,
                refresh_token=refresh,
                id=refresh_payload['id'],
                status=4  
            )
        else:
            return auth_pb2.UserTokens(
                access_token=access,
                refresh_token=refresh,
                id=access_payload['id'],
                status=0  
            )