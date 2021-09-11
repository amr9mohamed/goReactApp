import React, { useState, useEffect } from "react";
import axios from "axios";
import { Users } from "./components/Users";
import "./App.css";

function App() {
  const [users, setUsers] = useState([]);
  const [loading, setloading] = useState(false);
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

  return (
    <div className="container mt-5">
      <h1 className="text-primary mb-3">Users</h1>
      <Users users={users} loading={loading} usersPerPage={usersPerPage} />
    </div>
  );
}

export default App;
