## Book Management System

### Run
#### Run in local
1. install dependencies: `npm install`
2. run the code: `npm run dev`
3. open browser and enter url `http://127.0.0.1:5173/`

#### Run in Docker
1. build Docker `docker build -t book-management-system .`
2. run Docker `docker run -d --name book-management-system -p 80:80 book-management-system:latest`
3. open browser and enter url `http://127.0.0.1/`

### TODO
* Deploy in cloud platform
* Add Unit Tests
* Fuzzy search