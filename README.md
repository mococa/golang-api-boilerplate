# Golang API Boilerplate

Golang API boilerplate using:
  - Gorm v2
  - S.O.L.I.D principles
  - Service pattern
  - Gorilla Mux
  - Govalidator
  - Postgres

It's a simple api. The database entities are:
  - User
  - Book
  - Borrowed Book
 
 As one user can borrow many books
 
 ### Routes
 <table>
  <th>
    <tr>
      <td>Method</td>
      <td>Endpoint</td>
      <td>Body</td>
      <td>Query</td>
      <td>Description</td>
    </tr>
  </th>
  
  <tr>
    <td>GET</td>
    <td>/user</td>
    <td>none</td>
    <td>none</td>
    <td>Lists all users</td>
  </tr>
  
  <tr>
    <td>POST</td>
    <td>/user</td>
    <td>{ email:string, password: string, name: string, picture: string }</td>
    <td>none</td>
    <td>Creates a user</td>
  </tr>
  
  <tr>
    <td>POST</td>
    <td>/book</td>
    <td>{ name:string, author: string, release_year: number }</td>
    <td>none</td>
    <td>Creates a book</td>
  </tr>
  
   <tr>
    <td>GET</td>
    <td>/book/{userId}/borrowed</td>
    <td>none</td>
    <td>none</td>
    <td>Lists all books borrowed by user with given ID</td>
  </tr>
  
   <tr>
    <td>POST</td>
    <td>/book/{bookId}/borrow</td>
    <td>{ user_id: string }</td>
    <td>none</td>
    <td>Makes user borrow book</td>
  </tr>
 </table>
