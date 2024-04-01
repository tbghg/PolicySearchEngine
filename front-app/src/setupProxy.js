const { createProxyMiddleware } = require('http-proxy-middleware')

module.exports = function (app) {
    app.use(
        createProxyMiddleware('/search', {
            target: 'http://localhost:8080',
            changeOrigin: true
        }),
        createProxyMiddleware('/summary', {
            target: 'http://localhost:8080',
            changeOrigin: true
        }),
        createProxyMiddleware('/api2', {
            target: 'http://localhost:5001',
            changeOrigin: true,
            pathRewrite: { '^/api2': '' }
        }),
    )
}
