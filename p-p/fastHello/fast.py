from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Dict

# 定义用户模型
class User(BaseModel):
    id: str
    name: str
    email: str

# 模拟的用户数据存储
users: Dict[str, User] = {
    "1": User(id="1", name="Py1", email="alice@example.com"),
    "2": User(id="2", name="Py2", email="bob@example.com"),
}

# 创建 FastAPI 应用
app = FastAPI()

# 获取所有用户
@app.get("/users", response_model=list[User])
async def get_users():
    return list(users.values())

# 根据 ID 获取单个用户
@app.get("/users/{user_id}", response_model=User)
async def get_user(user_id: str):
    if user_id not in users:
        raise HTTPException(status_code=404, detail="User not found")
    return users[user_id]

# 添加新用户
@app.post("/users", response_model=User, status_code=201)
async def create_user(user: User):
    if user.id in users:
        raise HTTPException(status_code=400, detail="User already exists")
    users[user.id] = user
    return user

# 更新用户信息
@app.put("/users/{user_id}", response_model=User)
async def update_user(user_id: str, user: User):
    if user_id not in users:
        raise HTTPException(status_code=404, detail="User not found")
    users[user_id] = user
    return user

# 删除用户
@app.delete("/users/{user_id}", response_model=dict)
async def delete_user(user_id: str):
    if user_id not in users:
        raise HTTPException(status_code=404, detail="User not found")
    del users[user_id]
    return {"message": "User deleted"}