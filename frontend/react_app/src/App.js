import React, { useState, useEffect } from "react";
import axios from "axios";
import { Users } from "./components/Users";
import { Pagination } from "./components/Pagination";
import "./App.css";

function App() {
  const [users, setUsers] = useState([]);
  const [loading, setloading] = useState(false);
  const [currPage, setCurrPage] = useState(1);
  const [usersPerPage] = useState(10);

  useEffect(() => {
    const fetchUsers = async () => {
      setloading(true);
      const res = await axios.get("http://localhost:8080/users");
      setUsers(res.data);
      setloading(false);
    };
    fetchUsers();
  }, []);

  const idxLast = currPage * usersPerPage;
  const idxFirst = idxLast - usersPerPage;
  const currUsers = users.slice(idxFirst, idxLast);

  const paginate = (number) => {
    setCurrPage(number);
  };

  return (
    <div className="container mt-5">
      <h1 className="text-primary mb-3">Users</h1>
      <Users users={currUsers} loading={loading} />
      <Pagination
        elementsPerPage={usersPerPage}
        totalElements={users.length}
        paginate={paginate}
      />
    </div>
  );
}

export default App;
