import os
from datetime import timedelta, datetime

import dotenv
import jwt

dotenv.load_dotenv("../.env")


def generate_token(user_id):
    payload = lambda data, exp: {
        'id': data,
        'exp_time': (datetime.now() + exp).timestamp(),
    }

    print("acces on create", (datetime.now()+timedelta(seconds=30)).timestamp())

    return [
        jwt.encode(payload(user_id, timedelta(seconds=30)), os.getenv('JWT_SECRET'), algorithm='HS256'),
        jwt.encode(payload(user_id, timedelta(days=10)), os.getenv('JWT_SECRET'), algorithm='HS256'),
    ]


def decode_token(token):
    return jwt.decode(token, os.getenv('JWT_SECRET'), algorithms=['HS256'])
