--创建用户表
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    gender VARCHAR(10),
    email VARCHAR(100),
    age INTEGER,
    created TIMESTAMP DEFAULT now()
);
INSERT INTO users (name, username, password) VALUES ('PanWJ' ,'admin', '123456');


--创建数字身份表
CREATE TABLE IF NOT EXISTS dids (
    id SERIAL PRIMARY KEY,
    did_id TEXT UNIQUE,
    public_key_id INT REFERENCES public_keys(id),
    status TEXT NOT NULL,
    auth TEXT NOT NULL,
    created TIMESTAMP DEFAULT now(),
    updated TIMESTAMP DEFAULT now()
);
--创建私秘钥
CREATE TABLE IF NOT EXISTS public_keys (
    id SERIAL PRIMARY KEY,
    public_key TEXT UNIQUE
);

--创建发证方表
CREATE TABLE IF NOT EXISTS issuers (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    did TEXT NOT NULL,
    website TEXT NOT NULL,
    endpoint TEXT NOT NULL,
    short_description TEXT NOT NULL,
    long_description TEXT NOT NULL,
    service_type TEXT NOT NULL,
    request_data JSONB NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);
INSERT INTO issuers (uuid, did, website, endpoint, short_description, long_description, service_type, request_data, create_time, update_time)
VALUES ('407ab47f-1f7e-4a8c-86b0-e5c12dddf89d', 'did:iss:4455718552da01afccf74dadd86d475f', 'https://localhost:8081/home', '127.0.0.1', '数字认证声明', '分布式数字身份系统认证声明', 'RealNameAuthentication', '{"basicData": ["Name", "WEB"]}', '2023-3-14T13:33:05+08:00', '2023-3-14T13:33:05+08:00');

--创建应用中心表
CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    appdid VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    urls TEXT NOT NULL,
    issdid TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

--创建智能合约表
CREATE TABLE contracts (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  type TEXT NOT NULL,
  content TEXT NOT NULL,
  author TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);


