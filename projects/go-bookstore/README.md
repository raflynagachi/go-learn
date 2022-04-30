## Book Management System
Build book management system with requirement as follows:
- Database: MySQL
- Package:
  - GORM
  - json
  - httprouter
- Routes table
  Book model
  | Routes   | Path            | Method |
  | :------- | :-------------- | :----- |
  | Create   | /book           | POST   |
  | Update   | /book/{:bookId} | PUT    |
  | Delete   | /book/{:bookId} | DELETE |
  | FindById | /book/{:bookId} | GET    |
  | FindAll  | /book           | GET    |
