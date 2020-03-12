#### 1.设置代理

```
'/api/': {
  target: 'http://localhost:8090/',
  changeOrigin: true,
  pathRewrite: { '^/api': '/api' }, // /server/api/currentUser -> /api/currentUser
}
```
