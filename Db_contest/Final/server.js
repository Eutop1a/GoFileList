const http = require('http');
const httpProxy = require('http-proxy');
const finalhandler = require('finalhandler');
const serveStatic = require('serve-static');

const proxyTarget = 'http://localhost:65533';
const proxy = httpProxy.createProxyServer();
const serve = serveStatic('.', { 'index': ['index.html'] });

const server = http.createServer((req, res) => {
    // 设置 CORS 头
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
    res.setHeader('Access-Control-Allow-Headers', 'X-Requested-With,content-type');
    res.setHeader('Access-Control-Allow-Credentials', true);

    serve(req, res, finalhandler(req, res));
});

const port = 65532;
server.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});

proxy.on('error', (err, req, res) => {
    console.error('Proxy error:', err);
    res.writeHead(500, {
        'Content-Type': 'text/plain'
    });
    res.end('Proxy server error');
});
