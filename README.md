# fastbin: a text sharing application

https://fastbin.lab.divyam.dev

A text sharing application made with microservice architecture in golang.
Services:
- API Server (Read and Write)
- Custom Key Generator Service
- Database Service

To-do:
- Add data deletion service which removes data after certain time.
- Make key generator service faster and scalable with bloom filter and redis caching.