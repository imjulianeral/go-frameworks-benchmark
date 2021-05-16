# Performance banchmark for GO web frameworks

The tests were performed with [autocannon](https://github.com/mcollina/autocannon)

## [Fiber](https://gofiber.io/)

```bash
npx autocannon -c 125 -m POST -H "Content-Type":"application/json" -i ../file.json http://localhost:3000/json
Running 10s test @ http://localhost:3000/json
125 connections

┌─────────┬──────┬──────┬───────┬───────┬─────────┬──────────┬────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%   │ Avg     │ Stdev    │ Max    │
├─────────┼──────┼──────┼───────┼───────┼─────────┼──────────┼────────┤
│ Latency │ 1 ms │ 3 ms │ 58 ms │ 71 ms │ 12.6 ms │ 17.15 ms │ 188 ms │
└─────────┴──────┴──────┴───────┴───────┴─────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬─────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg     │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Req/Sec   │ 8839   │ 8839   │ 9719   │ 9767   │ 9563.21 │ 300.88  │ 8838   │
├───────────┼────────┼────────┼────────┼────────┼─────────┼─────────┼────────┤
│ Bytes/Sec │ 373 MB │ 373 MB │ 410 MB │ 412 MB │ 404 MB  │ 12.7 MB │ 373 MB │
└───────────┴────────┴────────┴────────┴────────┴─────────┴─────────┴────────┘
Req/Bytes counts sampled once per second.

96k requests in 10.04s, 4.04 GB read

```

## [Mux](https://github.com/gorilla/mux)

```bash
npx autocannon -c 125 -m POST -H "Content-Type":"application/json" -i ../file.json http://localhost:4001/json
Running 10s test @ http://localhost:4001/json
125 connections

┌─────────┬──────┬──────┬────────┬────────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5%  │ 99%    │ Avg      │ Stdev    │ Max    │
├─────────┼──────┼──────┼────────┼────────┼──────────┼──────────┼────────┤
│ Latency │ 1 ms │ 6 ms │ 104 ms │ 132 ms │ 24.58 ms │ 30.39 ms │ 313 ms │
└─────────┴──────┴──────┴────────┴────────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 4579   │ 4579   │ 5035   │ 5135   │ 4976.8 │ 172.37  │ 4577   │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 193 MB │ 193 MB │ 212 MB │ 217 MB │ 210 MB │ 7.27 MB │ 193 MB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

50k requests in 10.04s, 2.1 GB read
```

## [Iris](https://www.iris-go.com/)

```bash
npx autocannon -c 125 -m POST -H "Content-Type":"application/json" -i ../file.json http://localhost:4002/json
Running 10s test @ http://localhost:4002/json
125 connections

┌─────────┬──────┬──────┬────────┬────────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5%  │ 99%    │ Avg      │ Stdev    │ Max    │
├─────────┼──────┼──────┼────────┼────────┼──────────┼──────────┼────────┤
│ Latency │ 2 ms │ 8 ms │ 132 ms │ 165 ms │ 31.34 ms │ 38.08 ms │ 398 ms │
└─────────┴──────┴──────┴────────┴────────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 3529   │ 3529   │ 3977   │ 4043   │ 3919.4 │ 152.36  │ 3529   │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 208 MB │ 208 MB │ 235 MB │ 239 MB │ 231 MB │ 8.97 MB │ 208 MB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

39k requests in 10.04s, 2.31 GB read
```

## [Echo](https://echo.labstack.com/)

```bash
npx autocannon -c 125 -m POST -H "Content-Type":"application/json" -i ../file.json http://localhost:4003/json
Running 10s test @ http://localhost:4003/json
125 connections

┌─────────┬──────┬──────┬───────┬────────┬──────────┬──────────┬────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5% │ 99%    │ Avg      │ Stdev    │ Max    │
├─────────┼──────┼──────┼───────┼────────┼──────────┼──────────┼────────┤
│ Latency │ 1 ms │ 6 ms │ 92 ms │ 113 ms │ 22.16 ms │ 26.65 ms │ 261 ms │
└─────────┴──────┴──────┴───────┴────────┴──────────┴──────────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 5091   │ 5091   │ 5535   │ 5619   │ 5510.4 │ 148.15  │ 5091   │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 215 MB │ 215 MB │ 234 MB │ 237 MB │ 233 MB │ 6.24 MB │ 215 MB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

55k requests in 10.04s, 2.33 GB read
```

## [GIN](https://gin-gonic.com/)

```bash
npx autocannon -c 125 -m POST -H "Content-Type":"application/json" -i ../file.json http://localhost:4004/json
Running 10s test @ http://localhost:4004/json
125 connections

┌─────────┬──────┬──────┬────────┬────────┬──────────┬───────┬────────┐
│ Stat    │ 2.5% │ 50%  │ 97.5%  │ 99%    │ Avg      │ Stdev │ Max    │
├─────────┼──────┼──────┼────────┼────────┼──────────┼───────┼────────┤
│ Latency │ 1 ms │ 6 ms │ 112 ms │ 137 ms │ 27.05 ms │ 33 ms │ 285 ms │
└─────────┴──────┴──────┴────────┴────────┴──────────┴───────┴────────┘
┌───────────┬────────┬────────┬────────┬────────┬────────┬─────────┬────────┐
│ Stat      │ 1%     │ 2.5%   │ 50%    │ 97.5%  │ Avg    │ Stdev   │ Min    │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Req/Sec   │ 4071   │ 4071   │ 4567   │ 4671   │ 4522.1 │ 154.31  │ 4070   │
├───────────┼────────┼────────┼────────┼────────┼────────┼─────────┼────────┤
│ Bytes/Sec │ 172 MB │ 172 MB │ 193 MB │ 197 MB │ 191 MB │ 6.52 MB │ 172 MB │
└───────────┴────────┴────────┴────────┴────────┴────────┴─────────┴────────┘

Req/Bytes counts sampled once per second.

50k requests in 11.03s, 2.1 GB read
```
